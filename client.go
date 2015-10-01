package avalara

import (
	"fmt"
	"net/http"

	"gopkg.in/jmcvetta/napping.v2"
)

const PostalEndpoint = "https://taxrates.api.avalara.com:443/postal"

type Client struct {
	key        string
	httpClient *napping.Session
}

func New(key string) *Client {
	header := http.Header{}
	header.Add("Authorization", fmt.Sprintf("AvalaraApiKey %s", key))
	c := &napping.Session{
		Header: &header,
	}
	return &Client{
		key:        key,
		httpClient: c,
	}
}
