// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TrackingProperties tracking properties
// swagger:model trackingProperties
type TrackingProperties struct {

	// direction
	Direction string `json:"direction,omitempty"`

	// gold skill count
	GoldSkillCount int64 `json:"gold_skill_count,omitempty"`

	// gold tree percent
	GoldTreePercent float64 `json:"gold_tree_percent,omitempty"`

	// learning language
	LearningLanguage string `json:"learning_language,omitempty"`

	// max tree level
	MaxTreeLevel int64 `json:"max_tree_level,omitempty"`

	// num skills decayed
	NumSkillsDecayed int64 `json:"num_skills_decayed,omitempty"`

	// num skills newly decayed
	NumSkillsNewlyDecayed int64 `json:"num_skills_newly_decayed,omitempty"`

	// took placementtest
	TookPlacementtest bool `json:"took_placementtest,omitempty"`

	// total crowns
	TotalCrowns int64 `json:"total_crowns,omitempty"`

	// ui language
	UILanguage string `json:"ui_language,omitempty"`
}

// Validate validates this tracking properties
func (m *TrackingProperties) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TrackingProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrackingProperties) UnmarshalBinary(b []byte) error {
	var res TrackingProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
