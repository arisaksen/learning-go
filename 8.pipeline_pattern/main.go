package main

const (
	inputFiles = "8.pipeline_pattern/input"
	outFiles   = "8.pipeline_pattern/output"
)

func main() {

	channel1 := loadData(inputFiles, outFiles)
	channel2 := transform1(channel1)
	channel3 := transform2(channel2)
	writeResults := saveData(channel3)
	printResults(writeResults)

}
