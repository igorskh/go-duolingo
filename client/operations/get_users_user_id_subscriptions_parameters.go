// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetUsersUserIDSubscriptionsParams creates a new GetUsersUserIDSubscriptionsParams object
// with the default values initialized.
func NewGetUsersUserIDSubscriptionsParams() *GetUsersUserIDSubscriptionsParams {
	var ()
	return &GetUsersUserIDSubscriptionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersUserIDSubscriptionsParamsWithTimeout creates a new GetUsersUserIDSubscriptionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUsersUserIDSubscriptionsParamsWithTimeout(timeout time.Duration) *GetUsersUserIDSubscriptionsParams {
	var ()
	return &GetUsersUserIDSubscriptionsParams{

		timeout: timeout,
	}
}

// NewGetUsersUserIDSubscriptionsParamsWithContext creates a new GetUsersUserIDSubscriptionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUsersUserIDSubscriptionsParamsWithContext(ctx context.Context) *GetUsersUserIDSubscriptionsParams {
	var ()
	return &GetUsersUserIDSubscriptionsParams{

		Context: ctx,
	}
}

// NewGetUsersUserIDSubscriptionsParamsWithHTTPClient creates a new GetUsersUserIDSubscriptionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUsersUserIDSubscriptionsParamsWithHTTPClient(client *http.Client) *GetUsersUserIDSubscriptionsParams {
	var ()
	return &GetUsersUserIDSubscriptionsParams{
		HTTPClient: client,
	}
}

/*GetUsersUserIDSubscriptionsParams contains all the parameters to send to the API endpoint
for the get users user ID subscriptions operation typically these are written to a http.Request
*/
type GetUsersUserIDSubscriptionsParams struct {

	/*UserID
	  UserID

	*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get users user ID subscriptions params
func (o *GetUsersUserIDSubscriptionsParams) WithTimeout(timeout time.Duration) *GetUsersUserIDSubscriptionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users user ID subscriptions params
func (o *GetUsersUserIDSubscriptionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users user ID subscriptions params
func (o *GetUsersUserIDSubscriptionsParams) WithContext(ctx context.Context) *GetUsersUserIDSubscriptionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users user ID subscriptions params
func (o *GetUsersUserIDSubscriptionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users user ID subscriptions params
func (o *GetUsersUserIDSubscriptionsParams) WithHTTPClient(client *http.Client) *GetUsersUserIDSubscriptionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users user ID subscriptions params
func (o *GetUsersUserIDSubscriptionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the get users user ID subscriptions params
func (o *GetUsersUserIDSubscriptionsParams) WithUserID(userID int64) *GetUsersUserIDSubscriptionsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get users user ID subscriptions params
func (o *GetUsersUserIDSubscriptionsParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersUserIDSubscriptionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param userID
	if err := r.SetPathParam("userID", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}