package httpc

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

// @author valor.

type PackedReq struct {

	// if nil, use global.defaultHttpClient
	Client *http.Client

	Ctx context.Context

	URL string

	Method string

	Header map[string]string

	// if nil, use PublisherNoBody
	ReqBodyPublisher ReqBodyPublisher

	// if nil, use RespBodyNoHandle
	RespBodyHandler RespBodyHandler
}

func (r PackedReq) getBlankHttpRequest() (req *http.Request, err error) {
	if r.Ctx == nil {
		req, err = http.NewRequest("", "", nil)
	} else {
		req, err = http.NewRequestWithContext(r.Ctx, "", "", nil)
	}
	return
}

func (r PackedReq) Send() (interface{}, error) {
	req, err := r.getBlankHttpRequest()
	if err != nil {
		return nil, err
	}

	// set URL
	u, err := url.Parse(r.URL)
	if err != nil {
		return nil, err
	}
	req.URL = u

	// set method
	req.Method = r.Method

	// request body
	if r.ReqBodyPublisher == nil {
		r.ReqBodyPublisher = PublisherNoBody{}
	}
	body := r.ReqBodyPublisher.Subscribe()
	if body.Content != nil {
		rc, ok := body.Content.(io.ReadCloser)
		if !ok {
			rc = io.NopCloser(body.Content)
		}
		req.Body = rc

		req.ContentLength = body.Length
		req.Header.Set("Content-Type", body.Type)
	}

	// set user agent
	req.Header.Set("User-Agent", "httpc v"+version+" @github.com:valord577")

	// set header
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			req.Header.Set(k, v)
		}
	}

	// set HTTP client
	if r.Client == nil {
		r.Client = defaultHttpClient
	}

	// do HTTP request
	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	// response body
	if r.RespBodyHandler == nil {
		r.RespBodyHandler = RespBodyNoHandle{}
	}
	resBody, err := r.RespBodyHandler.Apply(resp.Body)
	if err != nil {
		return nil, err
	}
	return resBody, nil
}
