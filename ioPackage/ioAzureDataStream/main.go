package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"io"
	"log"
)

type AzureStorage struct {
	client azblob.Client
}

func NewAzureStorage() *AzureStorage {
	// Initialize your Azure Blob client here (mocked in this example)
	return &AzureStorage{client: azblob.Client{}}
}

func (as *AzureStorage) ProcessAndStoreBlob(ctx context.Context, containerName, sourceBlobName, destBlobName string) error {
	// 1. Set up a reader to download the data from Azure Blob
	downloadStream, err := as.client.DownloadStream(ctx, containerName, sourceBlobName, nil)
	if err != nil {
		return fmt.Errorf("failed to download blob: %w", err)
	}
	defer downloadStream.Body.Close()

	// 2. Set up the pipes for streaming data through transformation
	transformReader, transformWriter := io.Pipe()

	// 3. Start the transformation in a goroutine
	go func() {
		defer transformWriter.Close()

		if err := transformData(downloadStream.Body, transformWriter); err != nil {
			log.Printf("error in transformation: %v", err)
		}
	}()

	// 4. Upload the transformed data to a new blob
	uploadOptions := &azblob.UploadStreamOptions{
		BlockSize:   8 * 1024 * 1024, // Block size for the upload
		Concurrency: 5,               // Number of concurrent goroutines
	}
	_, err = as.client.UploadStream(ctx, containerName, destBlobName, transformReader, uploadOptions)
	if err != nil {
		return fmt.Errorf("failed to upload transformed blob: %w", err)
	}
	return nil
}

func transformData(input io.Reader, output io.Writer) error {
	buf := make([]byte, 1024)
	for {
		n, err := input.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		// Example transformation: Convert data to uppercase
		transformedData := bytes.ToUpper(buf[:n])

		_, err = output.Write(transformedData)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	ctx := context.Background()
	azureStorage := NewAzureStorage()

	err := azureStorage.ProcessAndStoreBlob(ctx, "container-name", "sourceBlob.bin", "destBlob.bin")
	if err != nil {
		log.Fatalf("error processing and storing blob: %v", err)
	}
	fmt.Println("Data streamed, transformed, and stored successfully.")
}
