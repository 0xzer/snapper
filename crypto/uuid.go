package crypto

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/0xzer/snapper/snaperror"
	"github.com/google/uuid"
)

func EncodeUUID(uuid string) ([]byte, error) {
	uuid = strings.ReplaceAll(uuid, "-", "")
	if len(uuid) != 32 {
		return nil, &snaperror.InvalidUUIDError{ErrorMessage: "uuid must be of length 32 to encode"}
	}

	bytes, err := hex.DecodeString(uuid)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func DecodeUUID(bytes []byte) (string, error) {
	if len(bytes) != 16 {
		return "", &snaperror.InvalidUUIDError{ErrorMessage: "encoded uuid must be of length 16 to decode"}
	}

	uuid := hex.EncodeToString(bytes)
	uuid = fmt.Sprintf("%s-%s-%s-%s-%s", uuid[:8], uuid[8:12], uuid[12:16], uuid[16:20], uuid[20:])

	return uuid, nil
}

func GenerateBitUUID() (*big.Int, error) {
	uuidString := uuid.New().String()
	significantBits := uuidString[0:8] + uuidString[9:13] + uuidString[14:18]

	result := new(big.Int)
	result, success := result.SetString(significantBits, 16)
	if !success {
		return nil, fmt.Errorf("failed to convert hex string to big.Int")
	}

	return result, nil
}