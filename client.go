package cclient

import (
	"net/http"

	"golang.org/x/net/proxy"

	utls "github.com/refraction-networking/utls"
)

func NewClient(clientHello utls.ClientHelloID, proxyUrl ...string) (http.Client, error) {
	if len(proxyUrl) > 0 && len(proxyUrl[0]) > 0 {
		dialer, err := NewConnectDialer(proxyUrl[0])
		if err != nil {
			return http.Client{}, err
		}
		return http.Client{
			Transport: NewRoundTripper(clientHello, dialer),
		}, nil
	} else {
		return http.Client{
			Transport: NewRoundTripper(clientHello, proxy.Direct),
		}, nil
	}
}
