package httpc

import (
	"fmt"
	"net/http"
	"testing"
)

// @author valor.

func TestHttpc000(t *testing.T) {
	c := PackedReq{
		URL:              "https://www.google.com",
		Method:           http.MethodGet,
		ReqBodyPublisher: PublisherNoBody{},
		RespBodyHandler:  RespBodyAsByteArray{},
	}

	bs, err := c.Send()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", bs)
}

func TestHttpc001(t *testing.T) {
	c := PackedReq{
		URL:    "https://www.google.com",
		Method: http.MethodPost,
		ReqBodyPublisher: PublisherRawBytesBody{
			Body: nil,
			Type: RawTypeText,
		},
		RespBodyHandler: RespBodyAsByteArray{},
	}

	bs, err := c.Send()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", bs)
}

func TestHttpc002(t *testing.T) {
	c := PackedReq{
		URL:    "https://www.google.com",
		Method: http.MethodPost,
		ReqBodyPublisher: PublisherRawStringBody{
			Body: "",
			Type: RawTypeText,
		},
		RespBodyHandler: RespBodyAsByteArray{},
	}

	bs, err := c.Send()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", bs)
}
