// Code generated by goa v3.4.3, DO NOT EDIT.
//
// chatapi HTTP client encoders and decoders
//
// Command:
// $ goa gen chat-api/design

package client

import (
	"bytes"
	chatapi "chat-api/gen/chatapi"
	chatapiviews "chat-api/gen/chatapi/views"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildGetchatRequest instantiates a HTTP request object with method and path
// set to call the "chatapi" service "getchat" endpoint
func (c *Client) BuildGetchatRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id int
	)
	{
		p, ok := v.(*chatapi.GetchatPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("chatapi", "getchat", "*chatapi.GetchatPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetchatChatapiPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("chatapi", "getchat", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetchatRequest returns an encoder for requests sent to the chatapi
// getchat server.
func EncodeGetchatRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*chatapi.GetchatPayload)
		if !ok {
			return goahttp.ErrInvalidType("chatapi", "getchat", "*chatapi.GetchatPayload", v)
		}
		{
			head := p.Key
			req.Header.Set("Authorization", head)
		}
		return nil
	}
}

// DecodeGetchatResponse returns a decoder for responses returned by the
// chatapi getchat endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeGetchatResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetchatResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chatapi", "getchat", err)
			}
			p := NewGetchatGoaChatCollectionOK(body)
			view := "default"
			vres := chatapiviews.GoaChatCollection{Projected: p, View: view}
			if err = chatapiviews.ValidateGoaChatCollection(vres); err != nil {
				return nil, goahttp.ErrValidationError("chatapi", "getchat", err)
			}
			res := chatapi.NewGoaChatCollection(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("chatapi", "getchat", resp.StatusCode, string(body))
		}
	}
}

// BuildPostchatRequest instantiates a HTTP request object with method and path
// set to call the "chatapi" service "postchat" endpoint
func (c *Client) BuildPostchatRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: PostchatChatapiPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("chatapi", "postchat", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodePostchatRequest returns an encoder for requests sent to the chatapi
// postchat server.
func EncodePostchatRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*chatapi.PostchatPayload)
		if !ok {
			return goahttp.ErrInvalidType("chatapi", "postchat", "*chatapi.PostchatPayload", v)
		}
		{
			head := p.Key
			req.Header.Set("Authorization", head)
		}
		body := NewPostchatRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("chatapi", "postchat", err)
		}
		return nil
	}
}

// DecodePostchatResponse returns a decoder for responses returned by the
// chatapi postchat endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodePostchatResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body bool
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chatapi", "postchat", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("chatapi", "postchat", resp.StatusCode, string(body))
		}
	}
}

// unmarshalGoaChatResponseToChatapiviewsGoaChatView builds a value of type
// *chatapiviews.GoaChatView from a value of type *GoaChatResponse.
func unmarshalGoaChatResponseToChatapiviewsGoaChatView(v *GoaChatResponse) *chatapiviews.GoaChatView {
	res := &chatapiviews.GoaChatView{
		ID:       v.ID,
		UserID:   v.UserID,
		RoomName: v.RoomName,
		Member:   v.Member,
		Chat:     v.Chat,
		PostDt:   v.PostDt,
	}

	return res
}
