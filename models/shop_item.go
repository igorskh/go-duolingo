// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ShopItem shop item
// swagger:model shopItem
type ShopItem struct {

	// id
	ID string `json:"id,omitempty"`

	// item name
	ItemName string `json:"itemName,omitempty"`

	// purchase date
	PurchaseDate int64 `json:"purchaseDate,omitempty"`

	// purchase Id
	PurchaseID string `json:"purchaseId,omitempty"`

	// purchase price
	PurchasePrice int64 `json:"purchasePrice,omitempty"`
}

// Validate validates this shop item
func (m *ShopItem) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ShopItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShopItem) UnmarshalBinary(b []byte) error {
	var res ShopItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
