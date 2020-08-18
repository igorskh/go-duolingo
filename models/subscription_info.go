// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SubscriptionInfo subscription info
// swagger:model subscriptionInfo
type SubscriptionInfo struct {

	// currency
	Currency string `json:"currency,omitempty"`

	// expected expiration
	ExpectedExpiration int64 `json:"expectedExpiration,omitempty"`

	// is free trial period
	IsFreeTrialPeriod bool `json:"isFreeTrialPeriod,omitempty"`

	// is in billing retry period
	IsInBillingRetryPeriod bool `json:"isInBillingRetryPeriod,omitempty"`

	// is intro offer period
	IsIntroOfferPeriod bool `json:"isIntroOfferPeriod,omitempty"`

	// period length
	PeriodLength int64 `json:"periodLength,omitempty"`

	// price
	Price int64 `json:"price,omitempty"`

	// product Id
	ProductID string `json:"productId,omitempty"`

	// renewer
	Renewer string `json:"renewer,omitempty"`

	// renewing
	Renewing bool `json:"renewing,omitempty"`

	// tier
	Tier interface{} `json:"tier,omitempty"`
}

// Validate validates this subscription info
func (m *SubscriptionInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SubscriptionInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionInfo) UnmarshalBinary(b []byte) error {
	var res SubscriptionInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
