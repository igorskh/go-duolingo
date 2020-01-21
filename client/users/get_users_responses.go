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

// GetUsersReader is a Reader for the GetUsers structure.
type GetUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetUsersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUsersOK creates a GetUsersOK with default headers values
func NewGetUsersOK() *GetUsersOK {
	return &GetUsersOK{}
}

/*GetUsersOK handles this case with default header values.

A Users object.
*/
type GetUsersOK struct {
	Payload *models.UserList
}

func (o *GetUsersOK) Error() string {
	return fmt.Sprintf("[GET /users][%d] getUsersOK  %+v", 200, o.Payload)
}

func (o *GetUsersOK) GetPayload() *models.UserList {
	return o.Payload
}

func (o *GetUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersBadRequest creates a GetUsersBadRequest with default headers values
func NewGetUsersBadRequest() *GetUsersBadRequest {
	return &GetUsersBadRequest{}
}

/*GetUsersBadRequest handles this case with default header values.

The specified user ID is invalid (e.g. not a number).
*/
type GetUsersBadRequest struct {
}

func (o *GetUsersBadRequest) Error() string {
	return fmt.Sprintf("[GET /users][%d] getUsersBadRequest ", 400)
}

func (o *GetUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersNotFound creates a GetUsersNotFound with default headers values
func NewGetUsersNotFound() *GetUsersNotFound {
	return &GetUsersNotFound{}
}

/*GetUsersNotFound handles this case with default header values.

A user with the specified ID was not found.
*/
type GetUsersNotFound struct {
}

func (o *GetUsersNotFound) Error() string {
	return fmt.Sprintf("[GET /users][%d] getUsersNotFound ", 404)
}

func (o *GetUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersDefault creates a GetUsersDefault with default headers values
func NewGetUsersDefault(code int) *GetUsersDefault {
	return &GetUsersDefault{
		_statusCode: code,
	}
}

/*GetUsersDefault handles this case with default header values.

Unexpected error
*/
type GetUsersDefault struct {
	_statusCode int
}

// Code gets the status code for the get users default response
func (o *GetUsersDefault) Code() int {
	return o._statusCode
}

func (o *GetUsersDefault) Error() string {
	return fmt.Sprintf("[GET /users][%d] getUsers default ", o._statusCode)
}

func (o *GetUsersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
