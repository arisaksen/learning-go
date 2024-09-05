package main

import (
	"bytes"
	"encoding/binary"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

const testInputFiles = "input"

func TestGenerateTestData(t *testing.T) {
	packets := make([]UdpPacket, 10)
	for i, _ := range packets {
		packets[i] = UdpPacket{
			Header:  byte(i + 1),
			ID:      byte(i + 1),
			Payload: byte(i + 1),
		}

		buf := new(bytes.Buffer)
		name := filepath.Join(testInputFiles, "udp-packet-"+strconv.Itoa(i+1)+".bin")
		_ = binary.Write(buf, binary.LittleEndian, packets[i])
		err := os.WriteFile(name, buf.Bytes(), fs.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}
