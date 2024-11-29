package ioAzureUnfinished

import (
	"bytes"
	"fmt"
	"go.uber.org/zap"
	"io"
	"path/filepath"
	"strings"
)

func (as *AzureStorage) GetDataStream(ctx context.Context, containerName string, filename string) (*io.ReadCloser, error) {
	downloadStream, err := as.client.DownloadStream(ctx, containerName, filename, nil)
	if err != nil {
		zap.L().Error("Error DownloadStream",
			zap.Error(err),
			zap.String("container", containerName),
			zap.String("filename", filename),
		)
		downloadStream.Body.Close()

		return nil, err
	}

	zap.L().Debug("GetDataStream successfully",
		zap.String("containerName", containerName),
		zap.String("filename", filename),
	)

	return &downloadStream.Body, nil
}

// io.Copy will wait for pipeWriter to Close
func TestAzureClientGetDataStream(t *testing.T) {
	printMemStats("Before:")

	streamReader, err := azureClient.GetDataStream(context.Background(), upStreamContainer, uploadBlobName)
	if err != nil {
		streamReader.Close()
	}

	buf := new(bytes.Buffer)
	// Because Copy is defined to read from reader until EOF, it does
	// not treat an EOF from pipeReader as an error to be reported.
	_, err = io.Copy(buf, streamReader)
	if err != nil {
		t.Fatal(err)
	}
	streamReader.Close()

	expectedLen := numberOfTestPackets * bytesPerPacket
	if buf.Len() != expectedLen {
		t.Error("actual:", buf.Len())
		t.Error("expected:", expectedLen)
	}

	readBytes := buf.Bytes()
	packetCounter := 0
	for i := 0; i < expectedLen; i += bytesPerPacket {
		streamId := model.ReadStreamId(readBytes[i : i+bytesPerPacket])
		if err != nil {
			t.Fatal(err)
		}
		if int(streamId) != packetCounter {
			t.Error("Expected streamId", packetCounter)
			t.Error("Actual streamId", streamId)
		}
		packetCounter++
	}

	printMemStats("After:")
}

func (ds *SplitterService) SplitStream(packetStream io.ReadCloser, fooPipeWriter, confPipeWriter, barPipeWriter *io.PipeWriter) error {
	buf := make([]byte, model.BytesPerPacket)
	for {
		n, err := packetStream.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			zap.L().Error("Error reading from packetStream", zap.Error(err))
			return err
		}

		packet := buf[:n]
		streamId := model.ReadStreamId(packet)

		writer, err := getWriterForStreamId(streamId, fooPipeWriter, confPipeWriter, barPipeWriter)
		if err != nil {
			zap.L().Error("Invalid streamId encountered", zap.Uint8("streamId", streamId), zap.Binary("packet", packet))
			return err
		}

		if _, err := writer.Write(packet); err != nil {
			zap.L().Error("Error writing packet to stream", zap.Error(err), zap.Uint8("streamId", streamId), zap.Binary("packet", packet))
			return err
		}
	}
	return nil
}

func getWriterForStreamId(streamId uint8, fooPipeWriter, confPipeWriter, barPipeWriter *io.PipeWriter) (*io.PipeWriter, error) {
	switch {
	case streamId >= fooDataMinStreamID && streamId <= fooDataMaxStreamID:
		return fooPipeWriter, nil
	case streamId >= confDataMinStreamID && streamId <= confDataMaxStreamID:
		return confPipeWriter, nil
	case streamId <= barDataMaxStreamID:
		return barPipeWriter, nil
	default:
		return nil, fmt.Errorf("invalid streamId: %d", streamId)
	}
}

func (s *StorageService) ProcessChunkAndUploadStreaming(ctx context.Context, env *env.Env, blobUrl string) error {
	//getDataReader, getDataWriter := io.Pipe()

	packetStream, err := s.storage.GetDataStream(ctx, env.StorageEnv.IncomingDataContainer, blobUrl)
	if err != nil {
		zap.L().Error("Error GetDataStream",
			zap.Error(err),
			zap.String("container", env.StorageEnv.IncomingDataContainer),
			zap.String("filename", blobUrl),
		)
		packetStream.Close()
	}

	blobName := strings.TrimSuffix(blobUrl, filepath.Ext(binaryFileExt))
	fooDataReader, fooDataWriter := io.Pipe()
	confDataReader, confDataWriter := io.Pipe()
	barDataReader, barDataWriter := io.Pipe()
	go func() (err error) {
		defer fooDataWriter.Close()
		defer confDataWriter.Close()
		defer barDataWriter.Close()

		fooDataWriter, confDataWriter, barDataWriter, err = s.splitterService.SplitStream(packetStream, fooDataWriter, confDataWriter, barDataWriter)
		if err != nil {
			zap.L().Error("Error SplitChunkDataStream",
				zap.Error(err),
				zap.String("container", env.StorageEnv.IncomingDataContainer),
				zap.String("filename", blobUrl),
			)
			fooDataWriter.CloseWithError(err)
			confDataWriter.CloseWithError(err)
			barDataWriter.CloseWithError(err)

			return err
		}

		return nil
	}()

	// todo("store datastream foo")
	// todo("publish foo")

	// todo("store datastream conf")
	// todo("publish conf")

	// todo("store datastream bar")

	// todo("store state")
	// todo("RECONSTRUCT JOB: [Datatake, burstId, streamID] -> where to find barData ref, lastPacket received?, numberOfPackets total,
	// Do any foo data or conf data info need to be in reconstruct job?
	// Can this be published to temp queue? Or is this problems with the logic? should it be stored in database and posted by the splitter or maybe checked periodically by another app?
	// Then picket off by another listener that creates "READY RECONSTRUCT JOB"

	// todo("RECONSTRUCT publish") this should be done when LAST_PACKET && NUMBER_OF_PACKETS_TOTAL==packets_received_counted) if not -> then TTL (publish after example 10 sek).
	// Can this be tested?

	// POST to new bar JOB QUEUE after reconstruct done

	return nil
}
