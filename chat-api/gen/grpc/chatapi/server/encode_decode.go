// Code generated by goa v3.3.1, DO NOT EDIT.
//
// chatapi gRPC server encoders and decoders
//
// Command:
// $ goa gen chat-api/design

package server

import (
	chatapi "chat-api/gen/chatapi"
	chatapiviews "chat-api/gen/chatapi/views"
	chatapipb "chat-api/gen/grpc/chatapi/pb"
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc/metadata"
)

// EncodeGetchatResponse encodes responses from the "chatapi" service "getchat"
// endpoint.
func EncodeGetchatResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	vres, ok := v.(chatapiviews.GoaChatCollection)
	if !ok {
		return nil, goagrpc.ErrInvalidType("chatapi", "getchat", "chatapiviews.GoaChatCollection", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewGoaChatCollection(result)
	return resp, nil
}

// DecodeGetchatRequest decodes requests sent to "chatapi" service "getchat"
// endpoint.
func DecodeGetchatRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *chatapipb.GetchatRequest
		ok      bool
	)
	{
		if message, ok = v.(*chatapipb.GetchatRequest); !ok {
			return nil, goagrpc.ErrInvalidType("chatapi", "getchat", "*chatapipb.GetchatRequest", v)
		}
	}
	var payload *chatapi.GetchatPayload
	{
		payload = NewGetchatPayload(message)
	}
	return payload, nil
}
