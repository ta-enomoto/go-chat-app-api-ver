// Code generated by goa v3.3.1, DO NOT EDIT.
//
// chatapi client HTTP transport
//
// Command:
// $ goa gen chat-api/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the chatapi service endpoint HTTP clients.
type Client struct {
	// Getchat Doer is the HTTP client used to make requests to the getchat
	// endpoint.
	GetchatDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the chatapi service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		GetchatDoer:         doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Getchat returns an endpoint that makes HTTP requests to the chatapi service
// getchat server.
func (c *Client) Getchat() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetchatRequest(c.encoder)
		decodeResponse = DecodeGetchatResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetchatRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetchatDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("chatapi", "getchat", err)
		}
		return decodeResponse(resp)
	}
}
