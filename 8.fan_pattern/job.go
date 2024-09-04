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

func loadData(input string, dirEntries []os.DirEntry, output string) <-chan Job {
	out := make(chan Job)

	time.Sleep(2 * time.Second)
	for _, file := range dirEntries {
		fileName := file.Name()
		go func() {
			inPath := path.Join(input, fileName)
			content, _ := os.ReadFile(inPath)
			outPath := path.Join(output, fileName)
			payload, _ := Parse(content)
			job := Job{
				InputPath: inPath,
				Packet:    *payload,
				OutPath:   outPath,
			}

			out <- job
		}()
	}
	return out
}

func transform1(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input {
			job.Packet.Payload = job.Packet.Payload + 1
			out <- job
		}
		close(out)
	}()
	return out
}

func transform2(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input {
			job.Packet.Payload = job.Packet.Payload + 1
			out <- job
		}
		close(out)
	}()
	return out
}

func collectJobs(input <-chan Job, size int) []Job {
	var jobs []Job
	for i := 0; i < size; i++ {
		//read job from input channel
		job := <-input

		jobs = append(jobs, job)
	}

	return jobs
}

func saveData(jobs []Job) error {
	for _, job := range jobs {
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, job.Packet)
		if err != nil {
			return err
		}

		fmt.Printf("Writing to file: %s\n", job.OutPath)
		err = os.WriteFile(job.OutPath, buf.Bytes(), fs.ModePerm)
		if err != nil {
			return err
		}

	}
	return nil
}
