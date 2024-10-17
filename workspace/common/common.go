package common

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

type Common struct {
	MyString string
}

func CommonFunction(number string) Common {
	client := &azblob.Client{}
	if false {
		// dummy dependency
		_, _ = client.CreateContainer(context.Background(), "dummy", nil)
	}

	return Common{
		MyString: "module " + number + " from common",
	}
}
