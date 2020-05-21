// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReferralInfo referral info
// swagger:model referralInfo
type ReferralInfo struct {

	// has reached cap
	HasReachedCap bool `json:"hasReachedCap,omitempty"`

	// inviter name
	InviterName string `json:"inviterName,omitempty"`

	// is eligible for bonus
	IsEligibleForBonus bool `json:"isEligibleForBonus,omitempty"`

	// num bonuses ready
	NumBonusesReady int64 `json:"numBonusesReady,omitempty"`

	// unconsumed invitee name
	UnconsumedInviteeName string `json:"unconsumedInviteeName,omitempty"`
}

// Validate validates this referral info
func (m *ReferralInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReferralInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReferralInfo) UnmarshalBinary(b []byte) error {
	var res ReferralInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
