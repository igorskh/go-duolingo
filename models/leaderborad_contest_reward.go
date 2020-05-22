// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LeaderboradContestReward leaderborad contest reward
// swagger:model leaderboradContestReward
type LeaderboradContestReward struct {

	// item id
	ItemID string `json:"item_id,omitempty"`

	// item quantity
	ItemQuantity int64 `json:"item_quantity,omitempty"`

	// rank
	Rank int64 `json:"rank,omitempty"`

	// reward type
	RewardType string `json:"reward_type,omitempty"`

	// tier
	Tier int64 `json:"tier,omitempty"`
}

// Validate validates this leaderborad contest reward
func (m *LeaderboradContestReward) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LeaderboradContestReward) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LeaderboradContestReward) UnmarshalBinary(b []byte) error {
	var res LeaderboradContestReward
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}