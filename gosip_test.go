package gosip

import (
	"io"
	"log"
	"net"
	"net/http"
)

type AnonymousCnfg struct {
	SiteURL string `json:"siteUrl"` // SPSite or SPWeb URL, which is the context target for the API calls
}

// ReadConfig : reads private config with auth options
func (c *AnonymousCnfg) ReadConfig(privateFile string) error {
	return nil
}

// WriteConfig : writes private config with auth options
func (c *AnonymousCnfg) WriteConfig(privateFile string) error {
	return nil
}

// GetAuth : authenticates, receives access token
func (c *AnonymousCnfg) GetAuth() (string, error) {
	return "", nil
}

// GetSiteURL : gets siteURL
func (c *AnonymousCnfg) GetSiteURL() string {
	return c.SiteURL
}

// GetStrategy : gets auth strategy name
func (c *AnonymousCnfg) GetStrategy() string {
	return "anonymous"
}

// SetAuth : authenticate request
func (c *AnonymousCnfg) SetAuth(req *http.Request, httpClient *SPClient) error {
	return nil
}

// Fake server bootstrap helper

func startFakeServer(addr string, handler http.Handler) (io.Closer, error) {
	srv := &http.Server{Addr: addr, Handler: handler}
	if addr == "" {
		addr = ":8989"
	}

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	go func() {
		err := srv.Serve(listener.(*net.TCPListener))
		if err != nil {
			log.Println("HTTP Server Error - ", err)
		}
	}()

	return listener, nil
}
