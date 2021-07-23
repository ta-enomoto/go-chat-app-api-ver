// Code generated by goa v3.4.3, DO NOT EDIT.
//
// chatapi service
//
// Command:
// $ goa gen chat-api/design

package chatapi

import (
	chatapiviews "chat-api/gen/chatapi/views"
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// The service performs get chat.
type Service interface {
	// Getchat implements getchat.
	Getchat(context.Context, *GetchatPayload) (res GoaChatCollection, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// APIKeyAuth implements the authorization logic for the APIKey security scheme.
	APIKeyAuth(ctx context.Context, key string, schema *security.APIKeyScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "chatapi"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"getchat"}

// GetchatPayload is the payload type of the chatapi service getchat method.
type GetchatPayload struct {
	// API key used to perform authorization
	Key string
	// room id
	ID int
}

// GoaChatCollection is the result type of the chatapi service getchat method.
type GoaChatCollection []*GoaChat

// All chat
type GoaChat struct {
	// room id
	ID int
	// user id
	UserID string
	// room name
	RoomName string
	// member
	Member string
	// chat
	Chat   string
	PostDt string
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "NotFound",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeBadRequest builds a goa.ServiceError from an error.
func MakeBadRequest(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "BadRequest",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// NewGoaChatCollection initializes result type GoaChatCollection from viewed
// result type GoaChatCollection.
func NewGoaChatCollection(vres chatapiviews.GoaChatCollection) GoaChatCollection {
	return newGoaChatCollection(vres.Projected)
}

// NewViewedGoaChatCollection initializes viewed result type GoaChatCollection
// from result type GoaChatCollection using the given view.
func NewViewedGoaChatCollection(res GoaChatCollection, view string) chatapiviews.GoaChatCollection {
	p := newGoaChatCollectionView(res)
	return chatapiviews.GoaChatCollection{Projected: p, View: "default"}
}

// newGoaChatCollection converts projected type GoaChatCollection to service
// type GoaChatCollection.
func newGoaChatCollection(vres chatapiviews.GoaChatCollectionView) GoaChatCollection {
	res := make(GoaChatCollection, len(vres))
	for i, n := range vres {
		res[i] = newGoaChat(n)
	}
	return res
}

// newGoaChatCollectionView projects result type GoaChatCollection to projected
// type GoaChatCollectionView using the "default" view.
func newGoaChatCollectionView(res GoaChatCollection) chatapiviews.GoaChatCollectionView {
	vres := make(chatapiviews.GoaChatCollectionView, len(res))
	for i, n := range res {
		vres[i] = newGoaChatView(n)
	}
	return vres
}

// newGoaChat converts projected type GoaChat to service type GoaChat.
func newGoaChat(vres *chatapiviews.GoaChatView) *GoaChat {
	res := &GoaChat{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.UserID != nil {
		res.UserID = *vres.UserID
	}
	if vres.RoomName != nil {
		res.RoomName = *vres.RoomName
	}
	if vres.Member != nil {
		res.Member = *vres.Member
	}
	if vres.Chat != nil {
		res.Chat = *vres.Chat
	}
	if vres.PostDt != nil {
		res.PostDt = *vres.PostDt
	}
	return res
}

// newGoaChatView projects result type GoaChat to projected type GoaChatView
// using the "default" view.
func newGoaChatView(res *GoaChat) *chatapiviews.GoaChatView {
	vres := &chatapiviews.GoaChatView{
		ID:       &res.ID,
		UserID:   &res.UserID,
		RoomName: &res.RoomName,
		Member:   &res.Member,
		Chat:     &res.Chat,
		PostDt:   &res.PostDt,
	}
	return vres
}
