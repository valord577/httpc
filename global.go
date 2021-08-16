package httpc

import "net/http"

// @author valor.

var defaultHttpClient = http.DefaultClient

// SetGlobalHttpClient set default http client
func SetGlobalHttpClient(client *http.Client) {
	defaultHttpClient = client
}
