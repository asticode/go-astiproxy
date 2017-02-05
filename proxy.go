package astiproxy

import (
	"net/http"
	"net/url"
)

// New returns an *http.Client with a proxy
func New(c Configuration) (cl *http.Client, err error) {
	cl = &http.Client{}
	if c.Addr != "" {
		var pu *url.URL
		if pu, err = url.Parse(c.Addr); err != nil {
			return
		}
		cl.Transport = &http.Transport{Proxy: http.ProxyURL(pu)}
	}
	return
}
