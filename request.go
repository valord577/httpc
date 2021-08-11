package httpc

import (
	"bytes"
	"io"
)

// @author valor.

type ReqBody struct {

	// Request body
	Content io.Reader

	// Content-Length
	Length int64

	// Content-type
	Type string
}

type ReqBodyPublisher interface {
	Subscribe() ReqBody
}

// ------ empty body ------

type PublisherNoBody struct{}

func (e PublisherNoBody) Subscribe() ReqBody {
	return ReqBody{}
}

// ------ raw body ------

type RawType int

const (
	RawTypeText RawType = iota
	RawTypeHtml
	RawTypeJson
	RawTypeXml

	RawTypeUrlEncodedForm
)

var RawTypesMap = map[RawType]string{
	RawTypeText: "text/plain; charset=utf-8",
	RawTypeHtml: "text/html; charset=utf-8",
	RawTypeJson: "application/json; charset=utf-8",
	RawTypeXml:  "application/xml; charset=utf-8",

	RawTypeUrlEncodedForm: "application/x-www-form-urlencoded; charset=utf-8",
}

type PublisherRawBytesBody struct {
	Body []byte
	Type RawType
}

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

type PublisherRawStringBody struct {
	Body string
	Type RawType
}

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
