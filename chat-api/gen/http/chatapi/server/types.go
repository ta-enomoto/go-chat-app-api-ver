// Code generated by goa v3.4.3, DO NOT EDIT.
//
// chatapi HTTP server types
//
// Command:
// $ goa gen chat-api/design

package server

import (
	chatapi "chat-api/gen/chatapi"
	chatapiviews "chat-api/gen/chatapi/views"

	goa "goa.design/goa/v3/pkg"
)

// PostchatRequestBody is the type of the "chatapi" service "postchat" endpoint
// HTTP request body.
type PostchatRequestBody struct {
	// room id
	ID *string `form:"Id,omitempty" json:"Id,omitempty" xml:"Id,omitempty"`
	// user id
	UserID *string `form:"UserId,omitempty" json:"UserId,omitempty" xml:"UserId,omitempty"`
	// room name
	RoomName *string `form:"RoomName,omitempty" json:"RoomName,omitempty" xml:"RoomName,omitempty"`
	// member
	Member *string `form:"Member,omitempty" json:"Member,omitempty" xml:"Member,omitempty"`
	// chat
	Chat   *string `form:"Chat,omitempty" json:"Chat,omitempty" xml:"Chat,omitempty"`
	PostDt *string `form:"PostDt,omitempty" json:"PostDt,omitempty" xml:"PostDt,omitempty"`
	// cookie
	Cookie *string `form:"Cookie,omitempty" json:"Cookie,omitempty" xml:"Cookie,omitempty"`
}

// GoaChatResponseCollection is the type of the "chatapi" service "getchat"
// endpoint HTTP response body.
type GoaChatResponseCollection []*GoaChatResponse

// GoaChatResponse is used to define fields on response body types.
type GoaChatResponse struct {
	// room id
	ID int `form:"Id" json:"Id" xml:"Id"`
	// user id
	UserID string `form:"UserId" json:"UserId" xml:"UserId"`
	// room name
	RoomName string `form:"RoomName" json:"RoomName" xml:"RoomName"`
	// member
	Member string `form:"Member" json:"Member" xml:"Member"`
	// chat
	Chat   string `form:"Chat" json:"Chat" xml:"Chat"`
	PostDt string `form:"PostDt" json:"PostDt" xml:"PostDt"`
}

// NewGoaChatResponseCollection builds the HTTP response body from the result
// of the "getchat" endpoint of the "chatapi" service.
func NewGoaChatResponseCollection(res chatapiviews.GoaChatCollectionView) GoaChatResponseCollection {
	body := make([]*GoaChatResponse, len(res))
	for i, val := range res {
		body[i] = marshalChatapiviewsGoaChatViewToGoaChatResponse(val)
	}
	return body
}

// NewGetchatPayload builds a chatapi service getchat endpoint payload.
func NewGetchatPayload(id int, key string) *chatapi.GetchatPayload {
	v := &chatapi.GetchatPayload{}
	v.ID = id
	v.Key = key

	return v
}

// NewPostchatPayload builds a chatapi service postchat endpoint payload.
func NewPostchatPayload(body *PostchatRequestBody, key string) *chatapi.PostchatPayload {
	v := &chatapi.PostchatPayload{
		ID:       *body.ID,
		UserID:   *body.UserID,
		RoomName: *body.RoomName,
		Member:   *body.Member,
		Chat:     *body.Chat,
		PostDt:   *body.PostDt,
		Cookie:   *body.Cookie,
	}
	v.Key = key

	return v
}

// ValidatePostchatRequestBody runs the validations defined on
// PostchatRequestBody
func ValidatePostchatRequestBody(body *PostchatRequestBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("Id", "body"))
	}
	if body.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("UserId", "body"))
	}
	if body.RoomName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("RoomName", "body"))
	}
	if body.Member == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("Member", "body"))
	}
	if body.Chat == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("Chat", "body"))
	}
	if body.PostDt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("PostDt", "body"))
	}
	if body.Cookie == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("Cookie", "body"))
	}
	if body.PostDt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.PostDt", *body.PostDt, goa.FormatDateTime))
	}
	return
}
