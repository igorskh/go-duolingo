// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Achievement achievement
// swagger:model achievement
type Achievement struct {

	// count
	Count int64 `json:"count,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// should show unlock
	ShouldShowUnlock bool `json:"shouldShowUnlock,omitempty"`

	// tier
	Tier int64 `json:"tier,omitempty"`

	// tier counts
	TierCounts []int64 `json:"tierCounts"`
}

// Validate validates this achievement
func (m *Achievement) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Achievement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Achievement) UnmarshalBinary(b []byte) error {
	var res Achievement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
