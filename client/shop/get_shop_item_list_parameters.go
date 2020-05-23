// Code generated by go-swagger; DO NOT EDIT.

package shop

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetShopItemListParams creates a new GetShopItemListParams object
// with the default values initialized.
func NewGetShopItemListParams() *GetShopItemListParams {

	return &GetShopItemListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetShopItemListParamsWithTimeout creates a new GetShopItemListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetShopItemListParamsWithTimeout(timeout time.Duration) *GetShopItemListParams {

	return &GetShopItemListParams{

		timeout: timeout,
	}
}

// NewGetShopItemListParamsWithContext creates a new GetShopItemListParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetShopItemListParamsWithContext(ctx context.Context) *GetShopItemListParams {

	return &GetShopItemListParams{

		Context: ctx,
	}
}

// NewGetShopItemListParamsWithHTTPClient creates a new GetShopItemListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetShopItemListParamsWithHTTPClient(client *http.Client) *GetShopItemListParams {

	return &GetShopItemListParams{
		HTTPClient: client,
	}
}

/*GetShopItemListParams contains all the parameters to send to the API endpoint
for the get shop item list operation typically these are written to a http.Request
*/
type GetShopItemListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get shop item list params
func (o *GetShopItemListParams) WithTimeout(timeout time.Duration) *GetShopItemListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get shop item list params
func (o *GetShopItemListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get shop item list params
func (o *GetShopItemListParams) WithContext(ctx context.Context) *GetShopItemListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get shop item list params
func (o *GetShopItemListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get shop item list params
func (o *GetShopItemListParams) WithHTTPClient(client *http.Client) *GetShopItemListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get shop item list params
func (o *GetShopItemListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetShopItemListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
