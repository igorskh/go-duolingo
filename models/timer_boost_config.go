// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TimerBoostConfig timer boost config
// swagger:model timerBoostConfig
type TimerBoostConfig struct {

	// has free timer boost
	HasFreeTimerBoost bool `json:"hasFreeTimerBoost,omitempty"`

	// time per boost
	TimePerBoost int64 `json:"timePerBoost,omitempty"`

	// timer boosts
	TimerBoosts int64 `json:"timerBoosts,omitempty"`
}

// Validate validates this timer boost config
func (m *TimerBoostConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TimerBoostConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimerBoostConfig) UnmarshalBinary(b []byte) error {
	var res TimerBoostConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
