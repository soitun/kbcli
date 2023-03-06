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

// PlanDetail plan detail
//
// swagger:model PlanDetail
type PlanDetail struct {

	// final phase billing period
	// Enum: [DAILY WEEKLY BIWEEKLY THIRTY_DAYS THIRTY_ONE_DAYS SIXTY_DAYS NINETY_DAYS MONTHLY BIMESTRIAL QUARTERLY TRIANNUAL BIANNUAL ANNUAL SESQUIENNIAL BIENNIAL TRIENNIAL NO_BILLING_PERIOD]
	FinalPhaseBillingPeriod PlanDetailFinalPhaseBillingPeriodEnum `json:"finalPhaseBillingPeriod,omitempty"`

	// final phase recurring price
	FinalPhaseRecurringPrice []*Price `json:"finalPhaseRecurringPrice"`

	// plan
	Plan string `json:"plan,omitempty"`

	// price list
	PriceList string `json:"priceList,omitempty"`

	// product
	Product string `json:"product,omitempty"`
}

// Validate validates this plan detail
func (m *PlanDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFinalPhaseBillingPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFinalPhaseRecurringPrice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var planDetailTypeFinalPhaseBillingPeriodPropEnum []interface{}

func init() {
	var res []PlanDetailFinalPhaseBillingPeriodEnum
	if err := json.Unmarshal([]byte(`["DAILY","WEEKLY","BIWEEKLY","THIRTY_DAYS","THIRTY_ONE_DAYS","SIXTY_DAYS","NINETY_DAYS","MONTHLY","BIMESTRIAL","QUARTERLY","TRIANNUAL","BIANNUAL","ANNUAL","SESQUIENNIAL","BIENNIAL","TRIENNIAL","NO_BILLING_PERIOD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		planDetailTypeFinalPhaseBillingPeriodPropEnum = append(planDetailTypeFinalPhaseBillingPeriodPropEnum, v)
	}
}

type PlanDetailFinalPhaseBillingPeriodEnum string

const (

	// PlanDetailFinalPhaseBillingPeriodDAILY captures enum value "DAILY"
	PlanDetailFinalPhaseBillingPeriodDAILY PlanDetailFinalPhaseBillingPeriodEnum = "DAILY"

	// PlanDetailFinalPhaseBillingPeriodWEEKLY captures enum value "WEEKLY"
	PlanDetailFinalPhaseBillingPeriodWEEKLY PlanDetailFinalPhaseBillingPeriodEnum = "WEEKLY"

	// PlanDetailFinalPhaseBillingPeriodBIWEEKLY captures enum value "BIWEEKLY"
	PlanDetailFinalPhaseBillingPeriodBIWEEKLY PlanDetailFinalPhaseBillingPeriodEnum = "BIWEEKLY"

	// PlanDetailFinalPhaseBillingPeriodTHIRTYDAYS captures enum value "THIRTY_DAYS"
	PlanDetailFinalPhaseBillingPeriodTHIRTYDAYS PlanDetailFinalPhaseBillingPeriodEnum = "THIRTY_DAYS"

	// PlanDetailFinalPhaseBillingPeriodTHIRTYONEDAYS captures enum value "THIRTY_ONE_DAYS"
	PlanDetailFinalPhaseBillingPeriodTHIRTYONEDAYS PlanDetailFinalPhaseBillingPeriodEnum = "THIRTY_ONE_DAYS"

	// PlanDetailFinalPhaseBillingPeriodSIXTYDAYS captures enum value "SIXTY_DAYS"
	PlanDetailFinalPhaseBillingPeriodSIXTYDAYS PlanDetailFinalPhaseBillingPeriodEnum = "SIXTY_DAYS"

	// PlanDetailFinalPhaseBillingPeriodNINETYDAYS captures enum value "NINETY_DAYS"
	PlanDetailFinalPhaseBillingPeriodNINETYDAYS PlanDetailFinalPhaseBillingPeriodEnum = "NINETY_DAYS"

	// PlanDetailFinalPhaseBillingPeriodMONTHLY captures enum value "MONTHLY"
	PlanDetailFinalPhaseBillingPeriodMONTHLY PlanDetailFinalPhaseBillingPeriodEnum = "MONTHLY"

	// PlanDetailFinalPhaseBillingPeriodBIMESTRIAL captures enum value "BIMESTRIAL"
	PlanDetailFinalPhaseBillingPeriodBIMESTRIAL PlanDetailFinalPhaseBillingPeriodEnum = "BIMESTRIAL"

	// PlanDetailFinalPhaseBillingPeriodQUARTERLY captures enum value "QUARTERLY"
	PlanDetailFinalPhaseBillingPeriodQUARTERLY PlanDetailFinalPhaseBillingPeriodEnum = "QUARTERLY"

	// PlanDetailFinalPhaseBillingPeriodTRIANNUAL captures enum value "TRIANNUAL"
	PlanDetailFinalPhaseBillingPeriodTRIANNUAL PlanDetailFinalPhaseBillingPeriodEnum = "TRIANNUAL"

	// PlanDetailFinalPhaseBillingPeriodBIANNUAL captures enum value "BIANNUAL"
	PlanDetailFinalPhaseBillingPeriodBIANNUAL PlanDetailFinalPhaseBillingPeriodEnum = "BIANNUAL"

	// PlanDetailFinalPhaseBillingPeriodANNUAL captures enum value "ANNUAL"
	PlanDetailFinalPhaseBillingPeriodANNUAL PlanDetailFinalPhaseBillingPeriodEnum = "ANNUAL"

	// PlanDetailFinalPhaseBillingPeriodSESQUIENNIAL captures enum value "SESQUIENNIAL"
	PlanDetailFinalPhaseBillingPeriodSESQUIENNIAL PlanDetailFinalPhaseBillingPeriodEnum = "SESQUIENNIAL"

	// PlanDetailFinalPhaseBillingPeriodBIENNIAL captures enum value "BIENNIAL"
	PlanDetailFinalPhaseBillingPeriodBIENNIAL PlanDetailFinalPhaseBillingPeriodEnum = "BIENNIAL"

	// PlanDetailFinalPhaseBillingPeriodTRIENNIAL captures enum value "TRIENNIAL"
	PlanDetailFinalPhaseBillingPeriodTRIENNIAL PlanDetailFinalPhaseBillingPeriodEnum = "TRIENNIAL"

	// PlanDetailFinalPhaseBillingPeriodNOBILLINGPERIOD captures enum value "NO_BILLING_PERIOD"
	PlanDetailFinalPhaseBillingPeriodNOBILLINGPERIOD PlanDetailFinalPhaseBillingPeriodEnum = "NO_BILLING_PERIOD"
)

var PlanDetailFinalPhaseBillingPeriodEnumValues = []string{
	"DAILY",
	"WEEKLY",
	"BIWEEKLY",
	"THIRTY_DAYS",
	"THIRTY_ONE_DAYS",
	"SIXTY_DAYS",
	"NINETY_DAYS",
	"MONTHLY",
	"BIMESTRIAL",
	"QUARTERLY",
	"TRIANNUAL",
	"BIANNUAL",
	"ANNUAL",
	"SESQUIENNIAL",
	"BIENNIAL",
	"TRIENNIAL",
	"NO_BILLING_PERIOD",
}

func (e PlanDetailFinalPhaseBillingPeriodEnum) IsValid() bool {
	for _, v := range PlanDetailFinalPhaseBillingPeriodEnumValues {
		if v == string(e) {
			return true
		}
	}
	return false
}

// prop value enum
func (m *PlanDetail) validateFinalPhaseBillingPeriodEnum(path, location string, value PlanDetailFinalPhaseBillingPeriodEnum) error {
	if err := validate.EnumCase(path, location, value, planDetailTypeFinalPhaseBillingPeriodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PlanDetail) validateFinalPhaseBillingPeriod(formats strfmt.Registry) error {
	if swag.IsZero(m.FinalPhaseBillingPeriod) { // not required
		return nil
	}

	// value enum
	if err := m.validateFinalPhaseBillingPeriodEnum("finalPhaseBillingPeriod", "body", m.FinalPhaseBillingPeriod); err != nil {
		return err
	}

	return nil
}

func (m *PlanDetail) validateFinalPhaseRecurringPrice(formats strfmt.Registry) error {
	if swag.IsZero(m.FinalPhaseRecurringPrice) { // not required
		return nil
	}

	for i := 0; i < len(m.FinalPhaseRecurringPrice); i++ {
		if swag.IsZero(m.FinalPhaseRecurringPrice[i]) { // not required
			continue
		}

		if m.FinalPhaseRecurringPrice[i] != nil {
			if err := m.FinalPhaseRecurringPrice[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("finalPhaseRecurringPrice" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("finalPhaseRecurringPrice" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this plan detail based on the context it is used
func (m *PlanDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFinalPhaseRecurringPrice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlanDetail) contextValidateFinalPhaseRecurringPrice(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FinalPhaseRecurringPrice); i++ {

		if m.FinalPhaseRecurringPrice[i] != nil {
			if err := m.FinalPhaseRecurringPrice[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("finalPhaseRecurringPrice" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("finalPhaseRecurringPrice" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlanDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlanDetail) UnmarshalBinary(b []byte) error {
	var res PlanDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
