// Code generated by go-swagger; DO NOT EDIT.

package achievements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/igorskh/go-duolingo/models"
)

// GetAchievementsReader is a Reader for the GetAchievements structure.
type GetAchievementsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAchievementsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAchievementsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAchievementsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAchievementsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAchievementsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAchievementsOK creates a GetAchievementsOK with default headers values
func NewGetAchievementsOK() *GetAchievementsOK {
	return &GetAchievementsOK{}
}

/*GetAchievementsOK handles this case with default header values.

Achievements object.
*/
type GetAchievementsOK struct {
	Payload *models.Achievements
}

func (o *GetAchievementsOK) Error() string {
	return fmt.Sprintf("[GET /users/{userID}/achievements][%d] getAchievementsOK  %+v", 200, o.Payload)
}

func (o *GetAchievementsOK) GetPayload() *models.Achievements {
	return o.Payload
}

func (o *GetAchievementsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Achievements)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAchievementsBadRequest creates a GetAchievementsBadRequest with default headers values
func NewGetAchievementsBadRequest() *GetAchievementsBadRequest {
	return &GetAchievementsBadRequest{}
}

/*GetAchievementsBadRequest handles this case with default header values.

Learning language or language from are not provided
*/
type GetAchievementsBadRequest struct {
}

func (o *GetAchievementsBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{userID}/achievements][%d] getAchievementsBadRequest ", 400)
}

func (o *GetAchievementsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAchievementsNotFound creates a GetAchievementsNotFound with default headers values
func NewGetAchievementsNotFound() *GetAchievementsNotFound {
	return &GetAchievementsNotFound{}
}

/*GetAchievementsNotFound handles this case with default header values.

User not found
*/
type GetAchievementsNotFound struct {
}

func (o *GetAchievementsNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{userID}/achievements][%d] getAchievementsNotFound ", 404)
}

func (o *GetAchievementsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAchievementsDefault creates a GetAchievementsDefault with default headers values
func NewGetAchievementsDefault(code int) *GetAchievementsDefault {
	return &GetAchievementsDefault{
		_statusCode: code,
	}
}

/*GetAchievementsDefault handles this case with default header values.

Unexpected error
*/
type GetAchievementsDefault struct {
	_statusCode int
}

// Code gets the status code for the get achievements default response
func (o *GetAchievementsDefault) Code() int {
	return o._statusCode
}

func (o *GetAchievementsDefault) Error() string {
	return fmt.Sprintf("[GET /users/{userID}/achievements][%d] getAchievements default ", o._statusCode)
}

func (o *GetAchievementsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}