// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ProgressQuizItem progress quiz item
// swagger:model progressQuizItem
type ProgressQuizItem struct {

	// end time
	EndTime int64 `json:"endTime,omitempty"`

	// score
	Score float64 `json:"score,omitempty"`

	// start time
	StartTime int64 `json:"startTime,omitempty"`
}

// Validate validates this progress quiz item
func (m *ProgressQuizItem) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProgressQuizItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProgressQuizItem) UnmarshalBinary(b []byte) error {
	var res ProgressQuizItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
