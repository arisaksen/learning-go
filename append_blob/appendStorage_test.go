package storage

import (
	"bytes"
	"context"
	"testing"
)

// requires test main with testcontainer
func TestNormal(t *testing.T) {
	blobContainerName := "individualpackets"
	var filenames []string
	_ = storage.client.GetAllBinFileNamesInContainer(context.Background(), blobContainerName, &filenames)

	chunkBuffer := new(bytes.Buffer)
	for _, filename := range filenames {
		_ = storage.client.DownloadStreamToBuffer(context.Background(), storage.downStreamContainer, filename, chunkBuffer)
	}

	result := chunkBuffer.Len()
	expected := 8192 * 9

	if result != expected {
		t.Errorf("Result is incorrect, got: %d, expected: %d.", result, expected)
	}
}

// Test to compare two different methods
func BenchmarkBench(b *testing.B) {

	b.Run("DownloadBufferToBuffer", func(b *testing.B) {
		blobContainerName := "individualpackets"
		var filenames []string
		_ = storage.client.GetAllBinFileNamesInContainer(context.Background(), blobContainerName, &filenames)

		chunkBuffer := new(bytes.Buffer)
		for i := 0; i < b.N; i++ {
			for _, filename := range filenames {
				_ = storage.client.DownloadBufferToBuffer(context.Background(), storage.downStreamContainer, filename, chunkBuffer)
			}
		}
		b.ReportAllocs()
	})

	b.Run("DownloadStreamToBuffer", func(b *testing.B) {
		blobContainerName := "individualpackets"
		var filenames []string
		_ = storage.client.GetAllBinFileNamesInContainer(context.Background(), blobContainerName, &filenames)

		chunkBuffer := new(bytes.Buffer)
		for i := 0; i < b.N; i++ {
			for _, filename := range filenames {
				_ = storage.client.DownloadStreamToBuffer(context.Background(), storage.downStreamContainer, filename, chunkBuffer)
			}
		}
		b.ReportAllocs()
	})
}
