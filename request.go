package snapper

import (
	"bytes"
	bin "encoding/binary"
	"io"
	"net/http"

	"github.com/0xzer/snapper/types"
)

type Request struct {
	client *Client
	http *http.Client
}

func (r *Request) MakeRequest(url string, method string, headers http.Header, payload []byte, requestType types.RequestType) (*http.Response, []byte, error) {
	if requestType == types.GRPC {
		grpcHeader, err := writeGRPCHeader(payload)
		if err != nil {
			return nil, nil, err
		}
		payload = append(grpcHeader.Bytes(), payload...)
	}

	newRequest, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, nil, err
	}
	newRequest.Header = headers

	response, err := r.http.Do(newRequest)
	if err !=  nil {
		return nil, nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}

	return response, responseBody, nil
}

func writeGRPCHeader(message []byte) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	buf.WriteByte(0)
	err := bin.Write(buf, bin.BigEndian, int32(len(message)))
	if err != nil {
		return nil, err
	}
	return buf, nil
}