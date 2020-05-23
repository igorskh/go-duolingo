// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StreakData streak data
// swagger:model streakData
type StreakData struct {

	// length
	Length int64 `json:"length,omitempty"`

	// updated time zone
	UpdatedTimeZone string `json:"updatedTimeZone,omitempty"`

	// updated timestamp
	UpdatedTimestamp int64 `json:"updatedTimestamp,omitempty"`

	// xp goal
	XpGoal int64 `json:"xpGoal,omitempty"`
}

// Validate validates this streak data
func (m *StreakData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StreakData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StreakData) UnmarshalBinary(b []byte) error {
	var res StreakData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
