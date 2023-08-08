package methods

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)


func GetTimestamp() int64 {
	return time.Now().UnixMilli()
}

func GetTimestampString() string {
	return strconv.Itoa(int(time.Now().UnixMilli()))
}

func SliceToString(s []string) (string, error) {
	jsonSlice, jsonErr := json.Marshal(s)
	return string(jsonSlice), jsonErr
}

func GenerateKey(length int) ([]byte, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return nil, fmt.Errorf("failed to generate key: %w", err)
	}
	return key, nil
}