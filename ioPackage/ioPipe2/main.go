package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

// Writes data to the first PipeWriter
func writeData(writer *io.PipeWriter) {
	defer writer.Close() // Ensure the writer closes

	chunks := []string{"Hello, ", "world!", " This is a test."}
	for _, chunk := range chunks {
		_, err := writer.Write([]byte(chunk))
		if err != nil {
			log.Fatalf("Failed to write to pipe: %v", err)
		}
	}
}

// Transforms data by reading from the first PipeReader and writing to the second PipeWriter
func transformData(reader *io.PipeReader, writer *io.PipeWriter) {
	defer reader.Close() // Close reader when done
	defer writer.Close() // Close writer when done

	buf := make([]byte, 16)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Error reading in transform: %v", err)
		}

		// Apply transformation (convert to uppercase, for example)
		transformedData := strings.ToUpper(string(buf[:n]))
		_, err = writer.Write([]byte(transformedData))
		if err != nil {
			log.Fatalf("Error writing in transform: %v", err)
		}
	}
}

// Reads data from the second PipeReader
func readData(reader *io.PipeReader) {
	defer reader.Close() // Ensure the reader is closed

	buf := make([]byte, 16)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Error reading in readData: %v", err)
		}
		fmt.Print(string(buf[:n])) // Print transformed data
	}
}

func main() {
	// Create two pipes for each stage of the pipeline
	firstReader, firstWriter := io.Pipe()
	secondReader, secondWriter := io.Pipe()

	// Start the writer in a goroutine
	go writeData(firstWriter)

	// Start the transformer in a goroutine, passing firstReader and secondWriter
	go transformData(firstReader, secondWriter)

	// Finally, read from the second reader in the main function
	readData(secondReader)
}
