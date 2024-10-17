package main

import (
	"common"
	"log"
)

type Game struct{}

func main() {
	testImport := common.CommonFunction("2")
	log.Println(testImport)

	println("Hello " + testImport.MyString)
}
