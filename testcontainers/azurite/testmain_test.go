package azurite

import (
	"bytes"
	"context"
	"fmt"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/testcontainers/testcontainers-go/modules/azurite"
	"go.uber.org/zap"
)

const (
	azuriteContainerImg = "mcr.microsoft.com/azure-storage/azurite:3.31.0"
	packetDataContainer = "packetdata"
	blobNameFormat      = "data-%d.bin"
)

var (
	azblobClient *azblob.Client
)

type packetData struct {
	id   int
	data byte
}

func TestMain(m *testing.M) {
	zapConfig := zap.NewDevelopmentConfig()
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger := zap.Must(zapConfig.Build())
	zap.ReplaceGlobals(logger)

	ctx := context.Background()
	var err error
	azuriteContainer, err := azurite.Run(ctx, azuriteContainerImg)
	if err != nil {
		zap.L().Fatal("failed to start container: %s", zap.Error(err))
	}
	blobPort, err := azuriteContainer.MappedPort(ctx, azurite.BlobPort)
	if err != nil {
		zap.L().Fatal("Failed to get mapped blob port", zap.Error(err))
	}
	connectionString := fmt.Sprintf(
		"DefaultEndpointsProtocol=http;AccountName=%s;AccountKey=%s;BlobEndpoint=http://localhost:%s/%s;",
		azurite.AccountName, azurite.AccountKey, blobPort.Port(), azurite.AccountName,
	)
	azblobClient, err = azblob.NewClientFromConnectionString(connectionString, nil)
	if err != nil {
		zap.L().Fatal("NewClientFromConnectionString", zap.Error(err))
	}
	if _, err = azblobClient.CreateContainer(context.Background(), packetDataContainer, nil); err != nil {
		zap.L().Fatal("CreateContainer", zap.Error(err))
	}

	for i := 0; i < 9; i++ {
		blobName := fmt.Sprintf(blobNameFormat, i)
		buf := new(bytes.Buffer)
		buf.WriteByte(0)
		if _, err = azblobClient.UploadStream(context.Background(), packetDataContainer, blobName, buf, nil); err != nil {
			zap.L().Fatal("UploadStream", zap.Error(err))
		}
		buf.Reset()
	}

	exitVal := m.Run()

	if err = azuriteContainer.Terminate(context.Background()); err != nil {
		logger.Fatal("Failed to terminate azuriteContainer", zap.Error(err))
	}
	os.Exit(exitVal)
}
