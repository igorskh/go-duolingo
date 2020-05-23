// Code generated by go-swagger; DO NOT EDIT.

package achievements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new achievements API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for achievements API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetAchievements(params *GetAchievementsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAchievementsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAchievements returns user achievements

  User achievements.
*/
func (a *Client) GetAchievements(params *GetAchievementsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAchievementsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAchievementsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAchievements",
		Method:             "GET",
		PathPattern:        "/users/{userID}/achievements",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAchievementsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAchievementsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAchievementsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
