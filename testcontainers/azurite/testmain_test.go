package storage

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
	"testing"

	pdhtparser "github.com/KSAT-Government-Programs/microsar-processing-line/pkg/pdhtparser_2_1"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/testcontainers/testcontainers-go/modules/azurite"
	"go.uber.org/zap"
)

const (
	downStreamContainer             = "individualpackets"
	upStreamContainer               = "upload"
	upstreamContainerCreateOrUpdate = "reconstructed-sardata"
	bytesPerPacket                  = 8192
)

var (
	client *AzureFileClient
	logger *zap.Logger
)

// Setup for tests
func TestMain(m *testing.M) {
	// Initialize zap logger
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	logger, _ = zap.Config.Build(config)

	// Before tests start up azuriteContainer
	ctx := context.Background()
	var err error
	azuriteContainer, err := azurite.Run(ctx,
		"mcr.microsoft.com/azure-storage/azurite:3.31.0",
	)
	if err != nil {
		logger.Fatal("failed to start container: %s", zap.Error(err))
	}

	// localhost:mappedPort->azuriteContainer:10000
	blobPort, err := azuriteContainer.MappedPort(ctx, azurite.BlobPort)
	if err != nil {
		logger.Fatal("Failed to get mapped blob port", zap.Error(err))
	}
	connectionString := fmt.Sprintf(
		"DefaultEndpointsProtocol=http;AccountName=%s;AccountKey=%s;BlobEndpoint=http://localhost:%s/%s;",
		azurite.AccountName, azurite.AccountKey, blobPort.Port(), azurite.AccountName,
	)

	// Create fileClient
	client, err = NewStorageClient(connectionString, azurite.AccountName, bytesPerPacket, logger)
	if err != nil {
		logger.Fatal("Failed to create storage client", zap.Error(err))
	}

	// Pre upload bin files to Azurite
	UploadTestData(connectionString)

	exitVal := m.Run()

	// After tests stop container
	if err = azuriteContainer.Terminate(context.Background()); err != nil {
		logger.Fatal("Failed to terminate azuriteContainer", zap.Error(err))
	}

	os.Exit(exitVal)
}

func UploadTestData(connectionString string) {
	azblobClient, err := azblob.NewClientFromConnectionString(connectionString, nil)
	if err != nil {
		logger.Fatal("Failed to create azblobClient to put test data", zap.Error(err))
	}
	_, err = azblobClient.CreateContainer(context.Background(), upStreamContainer, nil)
	if err != nil {
		logger.Info("Failed to create azurite storage container", zap.Error(err))
	}
	_, err = azblobClient.CreateContainer(context.Background(), downStreamContainer, nil)
	if err != nil {
		logger.Info("Failed to create azurite storage container", zap.Error(err))
	}
	_, err = azblobClient.CreateContainer(context.Background(), upstreamContainerCreateOrUpdate, nil)
	if err != nil {
		logger.Info("Failed to create azurite storage container", zap.Error(err))
	}
	data := pdhtparser.AncillaryData{
		PacketVersion:         1,
		SatelliteId:           0,
		PayloadId:             0,
		StreamId:              8,
		DatatakeId:            20240925,
		DatatakePacketCounter: 2,
		StreamPacketCounter:   1073741824,

		AncillaryDataVersion: 0,
		Unused:               [27]byte{},
		Checksum:             36444269,

		GNSS1PPSInternalClock:             13312500000,
		GNSSDataInternalClock:             13317187499,
		OrbitDataTimestampWeeks:           2371,
		OrbitDataTimestampSeconds:         50415,
		OrbitDataTimestampNanoSeconds:     0,
		GNSSPositionXaxis:                 110965.27,
		GNSSPositionYaxis:                 -350494.1,
		GNSSPositionZaxis:                 6.9535995e+06,
		GNSSVelocityXaxis:                 3031.4634,
		GNSSVelocityYaxis:                 6899.7085,
		GNSSVelocityZaxis:                 296.47083,
		OrbitDataStatus:                   [4]byte{1, 0, 2, 0},
		StarTrackerDataInternalClock:      0,
		AttitudeDataTimestampSeconds:      3,
		AttitudeDataTimestampMicroSeconds: 463973111,
		StarTrackerAltitudeQ1:             1749995997,
		StarTrackerAltitudeQ2:             0,
		StarTrackerAltitudeQ3:             43016757,
		StarTrackerAltitudeQ4:             13988794,
		AltitudeDataStatus:                136,
	}
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.LittleEndian, data)
	if err != nil {
		logger.Fatal("Failed to write binary to buffer", zap.Error(err))
	}
	for i := 0; i < 9; i++ {
		blobName := "data-" + strconv.Itoa(i) + ".bin"
		_, err := azblobClient.UploadBuffer(context.Background(), downStreamContainer, blobName, buf.Bytes(), nil)
		if err != nil {
			logger.Fatal("Failed to uploadBuffer", zap.Error(err))
		}
	}
}
