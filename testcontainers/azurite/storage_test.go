package azurite

import (
	"context"
	"fmt"
	"runtime"
	"testing"
)

func printMemStats(title string) {
	fmt.Println(title)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc = %v MB\n", bToMb(m.Alloc))
	fmt.Printf("TotalAlloc = %v MB\n", bToMb(m.TotalAlloc))
	fmt.Printf("Sys = %v MB\n", bToMb(m.Sys))
	fmt.Printf("NumGC = %v MB\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1000 / 1000 // SI unit decimal form
	//return b / 1024 / 1024 // binary form. will
}

func TestAzuriteBuffer(t *testing.T) {
	ctx := context.Background()
	blobName := fmt.Sprintf(blobNameFormat, 1)

	printMemStats("Before:")

	blobClient := azblobClient.ServiceClient().NewContainerClient(packetDataContainer).NewBlobClient(blobName)
	prop, _ := blobClient.GetProperties(ctx, nil)
	buff := make([]byte, *prop.ContentLength)
	_, _ = azblobClient.DownloadBuffer(ctx, packetDataContainer, blobName, buff, nil)

	printMemStats("After:")
}

func TestAzuriteStream(t *testing.T) {
	ctx := context.Background()
	blobName := fmt.Sprintf(blobNameFormat, 1)

	printMemStats("Before:")

	_, _ = azblobClient.DownloadStream(ctx, packetDataContainer, blobName, nil)

	printMemStats("After:")
}

func BenchmarkAzurite(b *testing.B) {
	ctx := context.Background()
	blobName := fmt.Sprintf(blobNameFormat, 1)

	b.Run("DownloadBuffer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			blobClient := azblobClient.ServiceClient().NewContainerClient(packetDataContainer).NewBlobClient(blobName)
			prop, err := blobClient.GetProperties(ctx, nil)
			if err != nil {
				b.Fatal(err)
			}
			buff := make([]byte, *prop.ContentLength)
			_, err = azblobClient.DownloadBuffer(ctx, packetDataContainer, blobName, buff, nil)
			if err != nil {
				b.Fatal(err)
			}
		}

		b.ReportAllocs()
	})

	b.Run("DownloadStream", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = azblobClient.DownloadStream(ctx, packetDataContainer, blobName, nil)
		}
		b.ReportAllocs()
	})

}
