// Code generated by go-swagger; DO NOT EDIT.

package kbmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UsagePrice usage price
//
// swagger:model UsagePrice
type UsagePrice struct {

	// billing mode
	// Enum: [IN_ADVANCE IN_ARREAR]
	BillingMode string `json:"billingMode,omitempty"`

	// tier block policy
	// Enum: [ALL_TIERS TOP_TIER]
	TierBlockPolicy string `json:"tierBlockPolicy,omitempty"`

	// tier prices
	TierPrices []*TierPrice `json:"tierPrices"`

	// usage name
	UsageName string `json:"usageName,omitempty"`

	// usage type
	// Enum: [CAPACITY CONSUMABLE]
	UsageType string `json:"usageType,omitempty"`
}

// Validate validates this usage price
func (m *UsagePrice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTierBlockPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTierPrices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsageType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var usagePriceTypeBillingModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IN_ADVANCE","IN_ARREAR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		usagePriceTypeBillingModePropEnum = append(usagePriceTypeBillingModePropEnum, v)
	}
}

const (

	// UsagePriceBillingModeINADVANCE captures enum value "IN_ADVANCE"
	UsagePriceBillingModeINADVANCE string = "IN_ADVANCE"

	// UsagePriceBillingModeINARREAR captures enum value "IN_ARREAR"
	UsagePriceBillingModeINARREAR string = "IN_ARREAR"
)

// prop value enum
func (m *UsagePrice) validateBillingModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, usagePriceTypeBillingModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UsagePrice) validateBillingMode(formats strfmt.Registry) error {
	if swag.IsZero(m.BillingMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateBillingModeEnum("billingMode", "body", m.BillingMode); err != nil {
		return err
	}

	return nil
}

var usagePriceTypeTierBlockPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ALL_TIERS","TOP_TIER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		usagePriceTypeTierBlockPolicyPropEnum = append(usagePriceTypeTierBlockPolicyPropEnum, v)
	}
}

const (

	// UsagePriceTierBlockPolicyALLTIERS captures enum value "ALL_TIERS"
	UsagePriceTierBlockPolicyALLTIERS string = "ALL_TIERS"

	// UsagePriceTierBlockPolicyTOPTIER captures enum value "TOP_TIER"
	UsagePriceTierBlockPolicyTOPTIER string = "TOP_TIER"
)

// prop value enum
func (m *UsagePrice) validateTierBlockPolicyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, usagePriceTypeTierBlockPolicyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UsagePrice) validateTierBlockPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.TierBlockPolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateTierBlockPolicyEnum("tierBlockPolicy", "body", m.TierBlockPolicy); err != nil {
		return err
	}

	return nil
}

func (m *UsagePrice) validateTierPrices(formats strfmt.Registry) error {
	if swag.IsZero(m.TierPrices) { // not required
		return nil
	}

	for i := 0; i < len(m.TierPrices); i++ {
		if swag.IsZero(m.TierPrices[i]) { // not required
			continue
		}

		if m.TierPrices[i] != nil {
			if err := m.TierPrices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tierPrices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tierPrices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var usagePriceTypeUsageTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CAPACITY","CONSUMABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		usagePriceTypeUsageTypePropEnum = append(usagePriceTypeUsageTypePropEnum, v)
	}
}

const (

	// UsagePriceUsageTypeCAPACITY captures enum value "CAPACITY"
	UsagePriceUsageTypeCAPACITY string = "CAPACITY"

	// UsagePriceUsageTypeCONSUMABLE captures enum value "CONSUMABLE"
	UsagePriceUsageTypeCONSUMABLE string = "CONSUMABLE"
)

// prop value enum
func (m *UsagePrice) validateUsageTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, usagePriceTypeUsageTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UsagePrice) validateUsageType(formats strfmt.Registry) error {
	if swag.IsZero(m.UsageType) { // not required
		return nil
	}

	// value enum
	if err := m.validateUsageTypeEnum("usageType", "body", m.UsageType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this usage price based on the context it is used
func (m *UsagePrice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTierPrices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UsagePrice) contextValidateTierPrices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TierPrices); i++ {

		if m.TierPrices[i] != nil {
			if err := m.TierPrices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tierPrices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tierPrices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UsagePrice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UsagePrice) UnmarshalBinary(b []byte) error {
	var res UsagePrice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
