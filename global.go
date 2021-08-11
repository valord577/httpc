package httpc

import "net/http"

// @author valor.

var defaultHttpClient = http.DefaultClient

func SetGlobalHttpClient(client *http.Client) {
	defaultHttpClient = client
}
