// Code generated by go-swagger; DO NOT EDIT.

package shop

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new shop API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for shop API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetShopItemList(params *GetShopItemListParams, authInfo runtime.ClientAuthInfoWriter) (*GetShopItemListOK, error)

	PostShopItemList(params *PostShopItemListParams, authInfo runtime.ClientAuthInfoWriter) (*PostShopItemListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetShopItemList returns shop item list

  List Shop items.
*/
func (a *Client) GetShopItemList(params *GetShopItemListParams, authInfo runtime.ClientAuthInfoWriter) (*GetShopItemListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetShopItemListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getShopItemList",
		Method:             "GET",
		PathPattern:        "/shop-items",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetShopItemListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetShopItemListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetShopItemListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostShopItemList buys an item from a shop

  Buys an item from a shop
*/
func (a *Client) PostShopItemList(params *PostShopItemListParams, authInfo runtime.ClientAuthInfoWriter) (*PostShopItemListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostShopItemListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postShopItemList",
		Method:             "POST",
		PathPattern:        "/shop-items",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostShopItemListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostShopItemListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostShopItemListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
