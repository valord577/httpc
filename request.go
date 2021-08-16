package httpc

import (
	"bytes"
	"io"
)

// @author valor.

// ReqBody need by stdlib HTTP client
type ReqBody struct {

	// Request body
	Content io.Reader

	// Content-Length
	Length int64

	// Content-type
	Type string
}

// ReqBodyPublisher processing request body
type ReqBodyPublisher interface {
	Subscribe() ReqBody
}

// ------ empty body ------

// PublisherNoBody is the implement of ReqBodyPublisher
//   -> No request body
type PublisherNoBody struct{}

// Subscribe the empty request body
func (e PublisherNoBody) Subscribe() ReqBody {
	return ReqBody{}
}

// ------ raw body ------

// RawType HTTP content-type enum
type RawType int

const (
	// RawTypeText "text/plain; charset=utf-8"
	RawTypeText RawType = iota
	// RawTypeHtml "text/html; charset=utf-8"
	RawTypeHtml
	// RawTypeJson "application/json; charset=utf-8"
	RawTypeJson
	// RawTypeXml "application/xml; charset=utf-8"
	RawTypeXml

	// RawTypeUrlEncodedForm "application/x-www-form-urlencoded; charset=utf-8"
	RawTypeUrlEncodedForm
)

// RawTypesMap HTTP content-type map
var RawTypesMap = map[RawType]string{
	RawTypeText: "text/plain; charset=utf-8",
	RawTypeHtml: "text/html; charset=utf-8",
	RawTypeJson: "application/json; charset=utf-8",
	RawTypeXml:  "application/xml; charset=utf-8",

	RawTypeUrlEncodedForm: "application/x-www-form-urlencoded; charset=utf-8",
}

// PublisherRawBytesBody is the implement of ReqBodyPublisher
//   -> Byte array request body
type PublisherRawBytesBody struct {
	Body []byte
	Type RawType
}

// Subscribe the byte array request body
func (raw PublisherRawBytesBody) Subscribe() ReqBody {

	length := len(raw.Body)

	buf := bytes.Buffer{}
	buf.Grow(length)
	buf.Write(raw.Body)

	return ReqBody{
		Content: &buf,
		Length:  int64(length),
		Type:    RawTypesMap[raw.Type],
	}
}

// PublisherRawStringBody is the implement of ReqBodyPublisher
//   -> String request body
type PublisherRawStringBody struct {
	Body string
	Type RawType
}

// Subscribe the string request body
func (raw PublisherRawStringBody) Subscribe() ReqBody {

	length := len(raw.Body)

	buf := bytes.Buffer{}
	buf.Grow(length)
	buf.WriteString(raw.Body)

	return ReqBody{
		Content: &buf,
		Length:  int64(length),
		Type:    RawTypesMap[raw.Type],
	}
}
