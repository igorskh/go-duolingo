// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetLeaderboardsLeaderboardIDUsersUserID returns user information when not logged in

User information when not logged in.
*/
func (a *Client) GetLeaderboardsLeaderboardIDUsersUserID(params *GetLeaderboardsLeaderboardIDUsersUserIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetLeaderboardsLeaderboardIDUsersUserIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLeaderboardsLeaderboardIDUsersUserIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLeaderboardsLeaderboardIDUsersUserID",
		Method:             "GET",
		PathPattern:        "/leaderboards/{leaderboardID}/users/{userID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLeaderboardsLeaderboardIDUsersUserIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLeaderboardsLeaderboardIDUsersUserIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLeaderboardsLeaderboardIDUsersUserIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetUsers returns user information when not logged in

User information when not logged in.
*/
func (a *Client) GetUsers(params *GetUsersParams, authInfo runtime.ClientAuthInfoWriter) (*GetUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetUsers",
		Method:             "GET",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetUsersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetUsersUserIDSubscriptions returns user subscription list

User subscription.
*/
func (a *Client) GetUsersUserIDSubscriptions(params *GetUsersUserIDSubscriptionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetUsersUserIDSubscriptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersUserIDSubscriptionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetUsersUserIDSubscriptions",
		Method:             "GET",
		PathPattern:        "/users/{userID}/subscriptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsersUserIDSubscriptionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUsersUserIDSubscriptionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetUsersUserIDSubscriptionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostLogin returns token in header

Login user
*/
func (a *Client) PostLogin(params *PostLoginParams, authInfo runtime.ClientAuthInfoWriter) (*PostLoginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLoginParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostLogin",
		Method:             "POST",
		PathPattern:        "/login",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLoginReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostLoginOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostLoginDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
