// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CourseSection course section
// swagger:model courseSection
type CourseSection struct {

	// author Id
	AuthorID string `json:"authorId,omitempty"`

	// checkpoint accessible
	CheckpointAccessible bool `json:"checkpointAccessible,omitempty"`

	// checkpoint finished
	CheckpointFinished bool `json:"checkpointFinished,omitempty"`

	// extra crowns
	ExtraCrowns int64 `json:"extraCrowns,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// num rows
	NumRows int64 `json:"numRows,omitempty"`

	// placement test available
	PlacementTestAvailable bool `json:"placementTestAvailable,omitempty"`

	// preload
	Preload bool `json:"preload,omitempty"`

	// progress version
	ProgressVersion int64 `json:"progressVersion,omitempty"`

	// summary
	Summary string `json:"summary,omitempty"`
}

// Validate validates this course section
func (m *CourseSection) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CourseSection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CourseSection) UnmarshalBinary(b []byte) error {
	var res CourseSection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
