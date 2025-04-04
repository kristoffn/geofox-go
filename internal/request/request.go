package request

import (
	"context"
	"net/url"
)

type Request struct {
	BaseURL *url.URL
}

func NewRequest(ctx context.Context, method string, endpoint string) (*Request, error) {
	r := &Request{
		BaseURL: &url.URL{
			Scheme: "https",
			Host:   "gti.geofox.de",
			Path:   "gti" + "/" + endpoint,
		},
	}
	return r, nil
}

func (r *Request) Execute() error {
	return nil
}
