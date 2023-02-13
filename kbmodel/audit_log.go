// Code generated by go-swagger; DO NOT EDIT.

package kbmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AuditLog audit log
//
// swagger:model AuditLog
type AuditLog struct {

	// change date
	// Format: date-time
	ChangeDate strfmt.DateTime `json:"changeDate,omitempty"`

	// change type
	ChangeType string `json:"changeType,omitempty"`

	// changed by
	ChangedBy string `json:"changedBy,omitempty"`

	// comments
	Comments string `json:"comments,omitempty"`

	// history
	History *Entity `json:"history,omitempty"`

	// object Id
	// Format: uuid
	ObjectID strfmt.UUID `json:"objectId,omitempty"`

	// object type
	// Enum: [ACCOUNT ACCOUNT_EMAIL BLOCKING_STATES BUNDLE CUSTOM_FIELD INVOICE PAYMENT TRANSACTION INVOICE_ITEM INVOICE_PAYMENT SUBSCRIPTION SUBSCRIPTION_EVENT SERVICE_BROADCAST PAYMENT_ATTEMPT PAYMENT_METHOD TAG TAG_DEFINITION TENANT TENANT_KVS]
	ObjectType string `json:"objectType,omitempty"`

	// reason code
	ReasonCode string `json:"reasonCode,omitempty"`

	// user token
	UserToken string `json:"userToken,omitempty"`
}

// Validate validates this audit log
func (m *AuditLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChangeDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHistory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditLog) validateChangeDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ChangeDate) { // not required
		return nil
	}

	if err := validate.FormatOf("changeDate", "body", "date-time", m.ChangeDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AuditLog) validateHistory(formats strfmt.Registry) error {
	if swag.IsZero(m.History) { // not required
		return nil
	}

	if m.History != nil {
		if err := m.History.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("history")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("history")
			}
			return err
		}
	}

	return nil
}

func (m *AuditLog) validateObjectID(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectID) { // not required
		return nil
	}

	if err := validate.FormatOf("objectId", "body", "uuid", m.ObjectID.String(), formats); err != nil {
		return err
	}

	return nil
}

var auditLogTypeObjectTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACCOUNT","ACCOUNT_EMAIL","BLOCKING_STATES","BUNDLE","CUSTOM_FIELD","INVOICE","PAYMENT","TRANSACTION","INVOICE_ITEM","INVOICE_PAYMENT","SUBSCRIPTION","SUBSCRIPTION_EVENT","SERVICE_BROADCAST","PAYMENT_ATTEMPT","PAYMENT_METHOD","TAG","TAG_DEFINITION","TENANT","TENANT_KVS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		auditLogTypeObjectTypePropEnum = append(auditLogTypeObjectTypePropEnum, v)
	}
}

const (

	// AuditLogObjectTypeACCOUNT captures enum value "ACCOUNT"
	AuditLogObjectTypeACCOUNT string = "ACCOUNT"

	// AuditLogObjectTypeACCOUNTEMAIL captures enum value "ACCOUNT_EMAIL"
	AuditLogObjectTypeACCOUNTEMAIL string = "ACCOUNT_EMAIL"

	// AuditLogObjectTypeBLOCKINGSTATES captures enum value "BLOCKING_STATES"
	AuditLogObjectTypeBLOCKINGSTATES string = "BLOCKING_STATES"

	// AuditLogObjectTypeBUNDLE captures enum value "BUNDLE"
	AuditLogObjectTypeBUNDLE string = "BUNDLE"

	// AuditLogObjectTypeCUSTOMFIELD captures enum value "CUSTOM_FIELD"
	AuditLogObjectTypeCUSTOMFIELD string = "CUSTOM_FIELD"

	// AuditLogObjectTypeINVOICE captures enum value "INVOICE"
	AuditLogObjectTypeINVOICE string = "INVOICE"

	// AuditLogObjectTypePAYMENT captures enum value "PAYMENT"
	AuditLogObjectTypePAYMENT string = "PAYMENT"

	// AuditLogObjectTypeTRANSACTION captures enum value "TRANSACTION"
	AuditLogObjectTypeTRANSACTION string = "TRANSACTION"

	// AuditLogObjectTypeINVOICEITEM captures enum value "INVOICE_ITEM"
	AuditLogObjectTypeINVOICEITEM string = "INVOICE_ITEM"

	// AuditLogObjectTypeINVOICEPAYMENT captures enum value "INVOICE_PAYMENT"
	AuditLogObjectTypeINVOICEPAYMENT string = "INVOICE_PAYMENT"

	// AuditLogObjectTypeSUBSCRIPTION captures enum value "SUBSCRIPTION"
	AuditLogObjectTypeSUBSCRIPTION string = "SUBSCRIPTION"

	// AuditLogObjectTypeSUBSCRIPTIONEVENT captures enum value "SUBSCRIPTION_EVENT"
	AuditLogObjectTypeSUBSCRIPTIONEVENT string = "SUBSCRIPTION_EVENT"

	// AuditLogObjectTypeSERVICEBROADCAST captures enum value "SERVICE_BROADCAST"
	AuditLogObjectTypeSERVICEBROADCAST string = "SERVICE_BROADCAST"

	// AuditLogObjectTypePAYMENTATTEMPT captures enum value "PAYMENT_ATTEMPT"
	AuditLogObjectTypePAYMENTATTEMPT string = "PAYMENT_ATTEMPT"

	// AuditLogObjectTypePAYMENTMETHOD captures enum value "PAYMENT_METHOD"
	AuditLogObjectTypePAYMENTMETHOD string = "PAYMENT_METHOD"

	// AuditLogObjectTypeTAG captures enum value "TAG"
	AuditLogObjectTypeTAG string = "TAG"

	// AuditLogObjectTypeTAGDEFINITION captures enum value "TAG_DEFINITION"
	AuditLogObjectTypeTAGDEFINITION string = "TAG_DEFINITION"

	// AuditLogObjectTypeTENANT captures enum value "TENANT"
	AuditLogObjectTypeTENANT string = "TENANT"

	// AuditLogObjectTypeTENANTKVS captures enum value "TENANT_KVS"
	AuditLogObjectTypeTENANTKVS string = "TENANT_KVS"
)

// prop value enum
func (m *AuditLog) validateObjectTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, auditLogTypeObjectTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AuditLog) validateObjectType(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectType) { // not required
		return nil
	}

	// value enum
	if err := m.validateObjectTypeEnum("objectType", "body", m.ObjectType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this audit log based on the context it is used
func (m *AuditLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHistory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditLog) contextValidateHistory(ctx context.Context, formats strfmt.Registry) error {

	if m.History != nil {
		if err := m.History.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("history")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("history")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuditLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditLog) UnmarshalBinary(b []byte) error {
	var res AuditLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
