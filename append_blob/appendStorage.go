package storage

import (
	"microsar-processing-line/storage"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/appendblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
)

//type AppendBlobManagers struct {
//	appendBlobClients map[string]AppendBlobClient
//}

type AppendBlobManager struct {
	client             *appendblob.Client
	timerStart         time.Duration
	ticker             *time.Ticker
	firstPacketForBlob bool
}

func NewContainerClient(client *storage.FileClient, containerName string) *container.Client {
	serviceClient := client.client.ServiceClient()
	containerClient := serviceClient.NewContainerClient(containerName)

	return containerClient
}

func NewAppendBlobManager(containerClient *container.Client, blobName string, timeDelta time.Duration) *AppendBlobManager {
	newAppendBlobManager := AppendBlobManager{
		client:             containerClient.NewAppendBlobClient(blobName),
		timerStart:         timeDelta,
		ticker:             time.NewTicker(timeDelta),
		firstPacketForBlob: true,
	}

	return &newAppendBlobManager
}
