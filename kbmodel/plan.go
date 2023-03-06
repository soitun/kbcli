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

// Plan plan
//
// swagger:model Plan
type Plan struct {

	// billing period
	// Enum: [DAILY WEEKLY BIWEEKLY THIRTY_DAYS THIRTY_ONE_DAYS SIXTY_DAYS NINETY_DAYS MONTHLY BIMESTRIAL QUARTERLY TRIANNUAL BIANNUAL ANNUAL SESQUIENNIAL BIENNIAL TRIENNIAL NO_BILLING_PERIOD]
	BillingPeriod PlanBillingPeriodEnum `json:"billingPeriod,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// phases
	Phases []*Phase `json:"phases"`

	// pretty name
	PrettyName string `json:"prettyName,omitempty"`

	// recurring billing mode
	// Enum: [IN_ADVANCE IN_ARREAR]
	RecurringBillingMode PlanRecurringBillingModeEnum `json:"recurringBillingMode,omitempty"`
}

// Validate validates this plan
func (m *Plan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecurringBillingMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var planTypeBillingPeriodPropEnum []interface{}

func init() {
	var res []PlanBillingPeriodEnum
	if err := json.Unmarshal([]byte(`["DAILY","WEEKLY","BIWEEKLY","THIRTY_DAYS","THIRTY_ONE_DAYS","SIXTY_DAYS","NINETY_DAYS","MONTHLY","BIMESTRIAL","QUARTERLY","TRIANNUAL","BIANNUAL","ANNUAL","SESQUIENNIAL","BIENNIAL","TRIENNIAL","NO_BILLING_PERIOD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		planTypeBillingPeriodPropEnum = append(planTypeBillingPeriodPropEnum, v)
	}
}

type PlanBillingPeriodEnum string

const (

	// PlanBillingPeriodDAILY captures enum value "DAILY"
	PlanBillingPeriodDAILY PlanBillingPeriodEnum = "DAILY"

	// PlanBillingPeriodWEEKLY captures enum value "WEEKLY"
	PlanBillingPeriodWEEKLY PlanBillingPeriodEnum = "WEEKLY"

	// PlanBillingPeriodBIWEEKLY captures enum value "BIWEEKLY"
	PlanBillingPeriodBIWEEKLY PlanBillingPeriodEnum = "BIWEEKLY"

	// PlanBillingPeriodTHIRTYDAYS captures enum value "THIRTY_DAYS"
	PlanBillingPeriodTHIRTYDAYS PlanBillingPeriodEnum = "THIRTY_DAYS"

	// PlanBillingPeriodTHIRTYONEDAYS captures enum value "THIRTY_ONE_DAYS"
	PlanBillingPeriodTHIRTYONEDAYS PlanBillingPeriodEnum = "THIRTY_ONE_DAYS"

	// PlanBillingPeriodSIXTYDAYS captures enum value "SIXTY_DAYS"
	PlanBillingPeriodSIXTYDAYS PlanBillingPeriodEnum = "SIXTY_DAYS"

	// PlanBillingPeriodNINETYDAYS captures enum value "NINETY_DAYS"
	PlanBillingPeriodNINETYDAYS PlanBillingPeriodEnum = "NINETY_DAYS"

	// PlanBillingPeriodMONTHLY captures enum value "MONTHLY"
	PlanBillingPeriodMONTHLY PlanBillingPeriodEnum = "MONTHLY"

	// PlanBillingPeriodBIMESTRIAL captures enum value "BIMESTRIAL"
	PlanBillingPeriodBIMESTRIAL PlanBillingPeriodEnum = "BIMESTRIAL"

	// PlanBillingPeriodQUARTERLY captures enum value "QUARTERLY"
	PlanBillingPeriodQUARTERLY PlanBillingPeriodEnum = "QUARTERLY"

	// PlanBillingPeriodTRIANNUAL captures enum value "TRIANNUAL"
	PlanBillingPeriodTRIANNUAL PlanBillingPeriodEnum = "TRIANNUAL"

	// PlanBillingPeriodBIANNUAL captures enum value "BIANNUAL"
	PlanBillingPeriodBIANNUAL PlanBillingPeriodEnum = "BIANNUAL"

	// PlanBillingPeriodANNUAL captures enum value "ANNUAL"
	PlanBillingPeriodANNUAL PlanBillingPeriodEnum = "ANNUAL"

	// PlanBillingPeriodSESQUIENNIAL captures enum value "SESQUIENNIAL"
	PlanBillingPeriodSESQUIENNIAL PlanBillingPeriodEnum = "SESQUIENNIAL"

	// PlanBillingPeriodBIENNIAL captures enum value "BIENNIAL"
	PlanBillingPeriodBIENNIAL PlanBillingPeriodEnum = "BIENNIAL"

	// PlanBillingPeriodTRIENNIAL captures enum value "TRIENNIAL"
	PlanBillingPeriodTRIENNIAL PlanBillingPeriodEnum = "TRIENNIAL"

	// PlanBillingPeriodNOBILLINGPERIOD captures enum value "NO_BILLING_PERIOD"
	PlanBillingPeriodNOBILLINGPERIOD PlanBillingPeriodEnum = "NO_BILLING_PERIOD"
)

var PlanBillingPeriodEnumValues = []string{
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

func (e PlanBillingPeriodEnum) IsValid() bool {
	for _, v := range PlanBillingPeriodEnumValues {
		if v == string(e) {
			return true
		}
	}
	return false
}

// prop value enum
func (m *Plan) validateBillingPeriodEnum(path, location string, value PlanBillingPeriodEnum) error {
	if err := validate.EnumCase(path, location, value, planTypeBillingPeriodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Plan) validateBillingPeriod(formats strfmt.Registry) error {
	if swag.IsZero(m.BillingPeriod) { // not required
		return nil
	}

	// value enum
	if err := m.validateBillingPeriodEnum("billingPeriod", "body", m.BillingPeriod); err != nil {
		return err
	}

	return nil
}

func (m *Plan) validatePhases(formats strfmt.Registry) error {
	if swag.IsZero(m.Phases) { // not required
		return nil
	}

	for i := 0; i < len(m.Phases); i++ {
		if swag.IsZero(m.Phases[i]) { // not required
			continue
		}

		if m.Phases[i] != nil {
			if err := m.Phases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("phases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("phases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var planTypeRecurringBillingModePropEnum []interface{}

func init() {
	var res []PlanRecurringBillingModeEnum
	if err := json.Unmarshal([]byte(`["IN_ADVANCE","IN_ARREAR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		planTypeRecurringBillingModePropEnum = append(planTypeRecurringBillingModePropEnum, v)
	}
}

type PlanRecurringBillingModeEnum string

const (

	// PlanRecurringBillingModeINADVANCE captures enum value "IN_ADVANCE"
	PlanRecurringBillingModeINADVANCE PlanRecurringBillingModeEnum = "IN_ADVANCE"

	// PlanRecurringBillingModeINARREAR captures enum value "IN_ARREAR"
	PlanRecurringBillingModeINARREAR PlanRecurringBillingModeEnum = "IN_ARREAR"
)

var PlanRecurringBillingModeEnumValues = []string{
	"IN_ADVANCE",
	"IN_ARREAR",
}

func (e PlanRecurringBillingModeEnum) IsValid() bool {
	for _, v := range PlanRecurringBillingModeEnumValues {
		if v == string(e) {
			return true
		}
	}
	return false
}

// prop value enum
func (m *Plan) validateRecurringBillingModeEnum(path, location string, value PlanRecurringBillingModeEnum) error {
	if err := validate.EnumCase(path, location, value, planTypeRecurringBillingModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Plan) validateRecurringBillingMode(formats strfmt.Registry) error {
	if swag.IsZero(m.RecurringBillingMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateRecurringBillingModeEnum("recurringBillingMode", "body", m.RecurringBillingMode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this plan based on the context it is used
func (m *Plan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePhases(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Plan) contextValidatePhases(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Phases); i++ {

		if m.Phases[i] != nil {
			if err := m.Phases[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("phases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("phases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Plan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Plan) UnmarshalBinary(b []byte) error {
	var res Plan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
