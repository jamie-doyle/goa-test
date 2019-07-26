// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// TableInfo client
//
// Command:
// $ goa gen github.com/jamie-doyle/goa-test/design

package tableinfo

import (
	"context"

	goa "goa.design/goa"
)

// Client is the "TableInfo" service client.
type Client struct {
	GetInfoEndpoint goa.Endpoint
}

// NewClient initializes a "TableInfo" service client given the endpoints.
func NewClient(getInfo goa.Endpoint) *Client {
	return &Client{
		GetInfoEndpoint: getInfo,
	}
}

// GetInfo calls the "get_info" endpoint of the "TableInfo" service.
func (c *Client) GetInfo(ctx context.Context, p *GetInfoPayload) (res *Tablepayload, err error) {
	var ires interface{}
	ires, err = c.GetInfoEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Tablepayload), nil
}
