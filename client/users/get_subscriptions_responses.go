// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/igorskh/go-duolingo/models"
)

// GetSubscriptionsReader is a Reader for the GetSubscriptions structure.
type GetSubscriptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubscriptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSubscriptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSubscriptionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSubscriptionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetSubscriptionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSubscriptionsOK creates a GetSubscriptionsOK with default headers values
func NewGetSubscriptionsOK() *GetSubscriptionsOK {
	return &GetSubscriptionsOK{}
}

/*GetSubscriptionsOK handles this case with default header values.

A Users object.
*/
type GetSubscriptionsOK struct {
	Payload *models.SubscriptionList
}

func (o *GetSubscriptionsOK) Error() string {
	return fmt.Sprintf("[GET /users/{userID}/subscriptions][%d] getSubscriptionsOK  %+v", 200, o.Payload)
}

func (o *GetSubscriptionsOK) GetPayload() *models.SubscriptionList {
	return o.Payload
}

func (o *GetSubscriptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriptionList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionsBadRequest creates a GetSubscriptionsBadRequest with default headers values
func NewGetSubscriptionsBadRequest() *GetSubscriptionsBadRequest {
	return &GetSubscriptionsBadRequest{}
}

/*GetSubscriptionsBadRequest handles this case with default header values.

The specified user ID is invalid (e.g. not a number).
*/
type GetSubscriptionsBadRequest struct {
}

func (o *GetSubscriptionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{userID}/subscriptions][%d] getSubscriptionsBadRequest ", 400)
}

func (o *GetSubscriptionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSubscriptionsNotFound creates a GetSubscriptionsNotFound with default headers values
func NewGetSubscriptionsNotFound() *GetSubscriptionsNotFound {
	return &GetSubscriptionsNotFound{}
}

/*GetSubscriptionsNotFound handles this case with default header values.

A user with the specified ID was not found.
*/
type GetSubscriptionsNotFound struct {
}

func (o *GetSubscriptionsNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{userID}/subscriptions][%d] getSubscriptionsNotFound ", 404)
}

func (o *GetSubscriptionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSubscriptionsDefault creates a GetSubscriptionsDefault with default headers values
func NewGetSubscriptionsDefault(code int) *GetSubscriptionsDefault {
	return &GetSubscriptionsDefault{
		_statusCode: code,
	}
}

/*GetSubscriptionsDefault handles this case with default header values.

Unexpected error
*/
type GetSubscriptionsDefault struct {
	_statusCode int
}

// Code gets the status code for the get subscriptions default response
func (o *GetSubscriptionsDefault) Code() int {
	return o._statusCode
}

func (o *GetSubscriptionsDefault) Error() string {
	return fmt.Sprintf("[GET /users/{userID}/subscriptions][%d] getSubscriptions default ", o._statusCode)
}

func (o *GetSubscriptionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
