package main

import (
	"bytes"
	"encoding/binary"
)

type UdpPacket struct {
	Header  byte
	ID      byte
	Payload byte
}

func Parse(input []byte) (*UdpPacket, error) {
	udpPacket := &UdpPacket{}
	reader := bytes.NewReader(input)
	err := binary.Read(reader, binary.BigEndian, udpPacket)
	if err != nil {
		return nil, err
	}
	return udpPacket, nil
}
