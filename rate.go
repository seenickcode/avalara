package avalara

import "github.com/jmcvetta/napping"

type TotalRate struct {
	Amount float64 `json:"totalRate"`
	Rates  []Rate  `json:"rates"`
}

type Rate struct {
	Rate float64 `json:"rate"`
	Name string  `json:"name"`
	Type string  `json:"type"`
}

func (c *Client) GetPostalCodeRate(country string, postal string) (*TotalRate, error) {
	r := &TotalRate{}
	p := &napping.Params{
		"country": country,
		"postal":  postal,
	}
	_, err := c.httpClient.Get(PostalEndpoint, p, r, nil)
	return r, err
}
