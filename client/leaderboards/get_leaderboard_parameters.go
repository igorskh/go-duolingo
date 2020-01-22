// Code generated by go-swagger; DO NOT EDIT.

package leaderboards

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

// NewGetLeaderboardParams creates a new GetLeaderboardParams object
// with the default values initialized.
func NewGetLeaderboardParams() *GetLeaderboardParams {
	var ()
	return &GetLeaderboardParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLeaderboardParamsWithTimeout creates a new GetLeaderboardParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLeaderboardParamsWithTimeout(timeout time.Duration) *GetLeaderboardParams {
	var ()
	return &GetLeaderboardParams{

		timeout: timeout,
	}
}

// NewGetLeaderboardParamsWithContext creates a new GetLeaderboardParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLeaderboardParamsWithContext(ctx context.Context) *GetLeaderboardParams {
	var ()
	return &GetLeaderboardParams{

		Context: ctx,
	}
}

// NewGetLeaderboardParamsWithHTTPClient creates a new GetLeaderboardParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLeaderboardParamsWithHTTPClient(client *http.Client) *GetLeaderboardParams {
	var ()
	return &GetLeaderboardParams{
		HTTPClient: client,
	}
}

/*GetLeaderboardParams contains all the parameters to send to the API endpoint
for the get leaderboard operation typically these are written to a http.Request
*/
type GetLeaderboardParams struct {

	/*LeaderboardID
	  LeaderboardID

	*/
	LeaderboardID string
	/*UserID
	  UserID

	*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get leaderboard params
func (o *GetLeaderboardParams) WithTimeout(timeout time.Duration) *GetLeaderboardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get leaderboard params
func (o *GetLeaderboardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get leaderboard params
func (o *GetLeaderboardParams) WithContext(ctx context.Context) *GetLeaderboardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get leaderboard params
func (o *GetLeaderboardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get leaderboard params
func (o *GetLeaderboardParams) WithHTTPClient(client *http.Client) *GetLeaderboardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get leaderboard params
func (o *GetLeaderboardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLeaderboardID adds the leaderboardID to the get leaderboard params
func (o *GetLeaderboardParams) WithLeaderboardID(leaderboardID string) *GetLeaderboardParams {
	o.SetLeaderboardID(leaderboardID)
	return o
}

// SetLeaderboardID adds the leaderboardId to the get leaderboard params
func (o *GetLeaderboardParams) SetLeaderboardID(leaderboardID string) {
	o.LeaderboardID = leaderboardID
}

// WithUserID adds the userID to the get leaderboard params
func (o *GetLeaderboardParams) WithUserID(userID int64) *GetLeaderboardParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get leaderboard params
func (o *GetLeaderboardParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLeaderboardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param leaderboardID
	if err := r.SetPathParam("leaderboardID", o.LeaderboardID); err != nil {
		return err
	}

	// path param userID
	if err := r.SetPathParam("userID", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
