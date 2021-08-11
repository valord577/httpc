package httpc

import "io"

// @author valor.

type RespBodyHandler interface {
	Apply(body io.ReadCloser) (interface{}, error)
}

// ------ no handle ------

type RespBodyNoHandle struct {}

func (r RespBodyNoHandle) Apply(_ io.ReadCloser) (interface{}, error) {
	return nil, nil
}

// ------ read body as bytes ------

type RespBodyAsByteArray struct {}

func (r RespBodyAsByteArray) Apply(body io.ReadCloser) (interface{}, error) {
	bs, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

