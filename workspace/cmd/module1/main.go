package main

import (
	"common"
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

type Game struct{}

func main() {
	testImport := common.CommonFunction("1")
	log.Println(testImport)

	client := &azblob.Client{}

	if false {
		// dummy dependency
		_, _ = client.CreateContainer(context.Background(), "dummy", nil)
	}

	println("Hello " + testImport.MyString)
}
