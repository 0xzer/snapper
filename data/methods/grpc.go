package methods

import (
	"encoding/binary"
	"errors"
)

const chunkHeaderLen = 5

// this does not return the trailing chunk header (grpc-message & grpc-status) it completely ignores it and only parses the protobuf message bytes
func ReadResponseFrame(incoming []byte) ([]byte, error) {
	offset := 0
	length := len(incoming)

	if length < chunkHeaderLen {
		return nil, errors.New("insufficient data for chunk header")
	}

	headerBytes := incoming[offset : offset+chunkHeaderLen]
	dataLength := int(binary.BigEndian.Uint32(headerBytes[1:]))

	if offset+chunkHeaderLen+dataLength > length {
		return nil, errors.New("insufficient data for chunk")
	}

	return incoming[offset+chunkHeaderLen : offset+chunkHeaderLen+dataLength], nil
}