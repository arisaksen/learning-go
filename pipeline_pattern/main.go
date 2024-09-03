package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/fs"
	"os"
	"path"
)

const (
	inputFiles = "pipeline_pattern/input"
	outFiles   = "pipeline_pattern/output"
)

type Job struct {
	InputPath string
	Packet    UdpPacket
	OutPath   string
}

func loadData(dir string, outDir string) <-chan Job {
	out := make(chan Job)
	fileNames, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	go func() {
		for _, file := range fileNames {
			inPath := path.Join(dir, file.Name())
			content, _ := os.ReadFile(inPath)
			outPath := path.Join(outDir, file.Name())
			payload, _ := Parse(content)
			job := Job{
				InputPath: inPath,
				Packet:    *payload,
				OutPath:   outPath,
			}

			out <- job
		}
		close(out)
	}()
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

func saveData(input <-chan Job) <-chan bool {
	out := make(chan bool)
	go func() {
		for job := range input {
			buf := new(bytes.Buffer)
			_ = binary.Write(buf, binary.LittleEndian, job.Packet)
			_ = os.WriteFile(job.OutPath, buf.Bytes(), fs.ModePerm)
			out <- true
		}
		close(out)
	}()
	return out
}

func main() {

	channel1 := loadData(inputFiles, outFiles)
	channel2 := transform1(channel1)
	channel3 := transform2(channel2)
	writeResults := saveData(channel3)

	for success := range writeResults {
		if success {
			fmt.Println("Success!. See output dir!")
		} else {
			fmt.Println("Failed!")
		}
	}

}

// generate bin files
//func main() {
//	packets := make([]UdpPacket, 10)
//	for i, _ := range packets {
//		packets[i] = UdpPacket{
//			Header:  byte(i + 1),
//			ID:      byte(i + 1),
//			Payload: byte(i + 1),
//		}
//
//		buf := new(bytes.Buffer)
//		name := filepath.Join(inputFiles, "udp-packet-"+strconv.Itoa(i+1)+".bin")
//		_ = binary.Write(buf, binary.LittleEndian, packets[i])
//		err := os.WriteFile(name, buf.Bytes(), fs.ModePerm)
//		if err != nil {
//			panic(err)
//		}
//
//	}
//}
