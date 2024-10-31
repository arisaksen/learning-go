package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/fs"
	"os"
	"path"
	"time"
)

type Job struct {
	InputPath string
	Packet    UdpPacket
	OutPath   string
}

func loadData(input string, output string) <-chan Job {
	out := make(chan Job)
	time.Sleep(1 * time.Second)

	fileNames, err := os.ReadDir(input)
	if err != nil {
		panic(err)
	}

	go func() {
		defer close(out)

		for _, file := range fileNames {
			inPath := path.Join(input, file.Name())
			content, _ := os.ReadFile(inPath)
			outPath := path.Join(output, file.Name())
			payload, _ := Parse(content)
			job := Job{
				InputPath: inPath,
				Packet:    *payload,
				OutPath:   outPath,
			}

			out <- job
		}
	}()

	return out
}

func transform1(input <-chan Job) <-chan Job {
	out := make(chan Job)

	go func() {
		defer close(out)

		for job := range input {
			job.Packet.Payload = job.Packet.Payload + 1
			out <- job
		}
	}()

	return out
}

func transform2(input <-chan Job) <-chan Job {
	out := make(chan Job)

	go func() {
		defer close(out)

		for job := range input {
			job.Packet.Payload = job.Packet.Payload + 1
			out <- job
		}
	}()

	return out
}

func saveData(input <-chan Job) <-chan bool {
	out := make(chan bool)

	go func() {
		defer close(out)

		for job := range input {
			buf := new(bytes.Buffer)
			_ = binary.Write(buf, binary.LittleEndian, job.Packet)
			_ = os.WriteFile(job.OutPath, buf.Bytes(), fs.ModePerm)
			out <- true
		}

	}()

	return out
}

func printResults(input <-chan bool) {
	for result := range input {
		if result {
			fmt.Println("Success!. See output dir!")
		} else {
			fmt.Println("Failed!")
		}
	}
}
