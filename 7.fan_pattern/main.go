package main

import (
	"log"
	"os"
)

const (
	inputFiles = "7.fan_pattern/input"
	outFiles   = "7.fan_pattern/output"
)

func main() {
	dirEntries, err := os.ReadDir(inputFiles)
	if err != nil {
		panic(err)
	}

	// loadData - fan out
	channel1 := loadData(inputFiles, dirEntries, outFiles) // onwards. pipeline pattern
	channel2 := transform1(channel1)                       // pipeline pattern
	channel3 := transform2(channel2)                       // pipeline pattern
	channel4 := transform2(channel3)                       // pipeline pattern

	// collectJobs - fan in
	jobs := collectJobs(channel4, len(dirEntries))
	err = saveData(jobs)
	if err != nil {
		log.Fatal("Error when writing jobs to files")
	}

}
