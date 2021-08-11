package httpc

import (
	"fmt"
	"net/http"
	"testing"
)

// @author valor.

func TestHttpc(t *testing.T) {
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
