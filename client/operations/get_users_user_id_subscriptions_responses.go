// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/igorskh/go-duolingo/models"
)

// GetUsersUserIDSubscriptionsReader is a Reader for the GetUsersUserIDSubscriptions structure.
type GetUsersUserIDSubscriptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersUserIDSubscriptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersUserIDSubscriptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUsersUserIDSubscriptionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUsersUserIDSubscriptionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetUsersUserIDSubscriptionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUsersUserIDSubscriptionsOK creates a GetUsersUserIDSubscriptionsOK with default headers values
func NewGetUsersUserIDSubscriptionsOK() *GetUsersUserIDSubscriptionsOK {
	return &GetUsersUserIDSubscriptionsOK{}
}

/*GetUsersUserIDSubscriptionsOK handles this case with default header values.

A Users object.
*/
type GetUsersUserIDSubscriptionsOK struct {
	Payload *models.SubscriptionList
}

func (o *GetUsersUserIDSubscriptionsOK) Error() string {
	return fmt.Sprintf("[GET /users/{userID}/subscriptions][%d] getUsersUserIdSubscriptionsOK  %+v", 200, o.Payload)
}

func (o *GetUsersUserIDSubscriptionsOK) GetPayload() *models.SubscriptionList {
	return o.Payload
}

func (o *GetUsersUserIDSubscriptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriptionList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUserIDSubscriptionsBadRequest creates a GetUsersUserIDSubscriptionsBadRequest with default headers values
func NewGetUsersUserIDSubscriptionsBadRequest() *GetUsersUserIDSubscriptionsBadRequest {
	return &GetUsersUserIDSubscriptionsBadRequest{}
}

/*GetUsersUserIDSubscriptionsBadRequest handles this case with default header values.

The specified user ID is invalid (e.g. not a number).
*/
type GetUsersUserIDSubscriptionsBadRequest struct {
}

func (o *GetUsersUserIDSubscriptionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{userID}/subscriptions][%d] getUsersUserIdSubscriptionsBadRequest ", 400)
}

func (o *GetUsersUserIDSubscriptionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersUserIDSubscriptionsNotFound creates a GetUsersUserIDSubscriptionsNotFound with default headers values
func NewGetUsersUserIDSubscriptionsNotFound() *GetUsersUserIDSubscriptionsNotFound {
	return &GetUsersUserIDSubscriptionsNotFound{}
}

/*GetUsersUserIDSubscriptionsNotFound handles this case with default header values.

A user with the specified ID was not found.
*/
type GetUsersUserIDSubscriptionsNotFound struct {
}

func (o *GetUsersUserIDSubscriptionsNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{userID}/subscriptions][%d] getUsersUserIdSubscriptionsNotFound ", 404)
}

func (o *GetUsersUserIDSubscriptionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersUserIDSubscriptionsDefault creates a GetUsersUserIDSubscriptionsDefault with default headers values
func NewGetUsersUserIDSubscriptionsDefault(code int) *GetUsersUserIDSubscriptionsDefault {
	return &GetUsersUserIDSubscriptionsDefault{
		_statusCode: code,
	}
}

/*GetUsersUserIDSubscriptionsDefault handles this case with default header values.

Unexpected error
*/
type GetUsersUserIDSubscriptionsDefault struct {
	_statusCode int
}

// Code gets the status code for the get users user ID subscriptions default response
func (o *GetUsersUserIDSubscriptionsDefault) Code() int {
	return o._statusCode
}

func (o *GetUsersUserIDSubscriptionsDefault) Error() string {
	return fmt.Sprintf("[GET /users/{userID}/subscriptions][%d] GetUsersUserIDSubscriptions default ", o._statusCode)
}

func (o *GetUsersUserIDSubscriptionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
