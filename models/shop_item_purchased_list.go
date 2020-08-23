// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ShopItemPurchasedList shop item purchased list
// swagger:model shopItemPurchasedList
type ShopItemPurchasedList struct {

	// shop items
	ShopItems []*ShopItemPurchased `json:"shopItems"`
}

// Validate validates this shop item purchased list
func (m *ShopItemPurchasedList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateShopItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShopItemPurchasedList) validateShopItems(formats strfmt.Registry) error {

	if swag.IsZero(m.ShopItems) { // not required
		return nil
	}

	for i := 0; i < len(m.ShopItems); i++ {
		if swag.IsZero(m.ShopItems[i]) { // not required
			continue
		}

		if m.ShopItems[i] != nil {
			if err := m.ShopItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shopItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ShopItemPurchasedList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShopItemPurchasedList) UnmarshalBinary(b []byte) error {
	var res ShopItemPurchasedList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
