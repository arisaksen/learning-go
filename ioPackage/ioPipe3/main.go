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

	// Use io.TeeReader to transform data on the fly
	transformer := io.TeeReader(reader, uppercaseWriter(writer))
	_, err := io.Copy(writer, transformer)
	if err != nil {
		log.Fatalf("Error in transformData: %v", err)
	}
}

// Helper function to wrap a writer and apply uppercase transformation
func uppercaseWriter(writer io.Writer) io.Writer {
	return &upperWriter{Writer: writer}
}

type upperWriter struct {
	io.Writer
}

func (u *upperWriter) Write(p []byte) (int, error) {
	upper := strings.ToUpper(string(p))
	return u.Writer.Write([]byte(upper))
}

// Reads data from the second PipeReader
func readData(reader *io.PipeReader) {
	defer reader.Close() // Ensure the reader is closed

	// Copy data from reader to stdout
	_, err := io.Copy(io.Discard, reader) // For demonstration, replace io.Discard with another Writer
	if err != nil {
		log.Fatalf("Error reading in readData: %v", err)
	}
	fmt.Println("Data transformation complete")
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
