package protos

import (
	"encoding/base64"
	"fmt"

	"github.com/0xzer/snapper/data/methods"
	"google.golang.org/protobuf/proto"
)

func DecodeGRPCProtoMessage(data []byte, message proto.Message) error {
	protoBytes, err := methods.ReadResponseFrame(data)
	if err != nil {
		return err
	}

	return DecodeProtoMessage(protoBytes, message)
}

func EncodeProtoMessage(message proto.Message) ([]byte, error) {
	data, err := proto.Marshal(message)
	if err != nil {
		return nil, fmt.Errorf("failed to encode proto message: %v", err)
	}
	return data, nil
}

func DecodeProtoMessage(data []byte, message proto.Message) error {
	err := proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("failed to decode proto message: %v", err)
	}
	return nil
}

// protobytes, b64string, err
func EncodeProtoMessageB64(message proto.Message) ([]byte, string, error) {
	protoBytes, protoErr := EncodeProtoMessage(message)
	if protoErr != nil {
		return nil, "", protoErr
	}

	b64Encoded := base64.StdEncoding.EncodeToString(protoBytes)

	return protoBytes, b64Encoded, nil
}