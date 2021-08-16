package httpc

import "io"

// @author valor.

// RespBodyHandler processing response body
type RespBodyHandler interface {
	// Apply the handle of response body
	Apply(body io.ReadCloser) (interface{}, error)
}

// ------ no handle ------

// RespBodyNoHandle is the implement of RespBodyHandler
//   -> No read body
type RespBodyNoHandle struct{}

// Apply the handle
func (r RespBodyNoHandle) Apply(_ io.ReadCloser) (interface{}, error) {
	return nil, nil
}

// ------ read body as bytes ------

// RespBodyAsByteArray is the implement of RespBodyHandler
//   -> Read body as byte array
type RespBodyAsByteArray struct{}

// Apply the handle
func (r RespBodyAsByteArray) Apply(body io.ReadCloser) (interface{}, error) {
	bs, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return bs, nil
}
