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

// Tag tag
//
// swagger:model Tag
type Tag struct {

	// audit logs
	AuditLogs []*AuditLog `json:"auditLogs"`

	// object Id
	// Format: uuid
	ObjectID strfmt.UUID `json:"objectId,omitempty"`

	// object type
	// Enum: [ACCOUNT ACCOUNT_EMAIL BLOCKING_STATES BUNDLE CUSTOM_FIELD INVOICE PAYMENT TRANSACTION INVOICE_ITEM INVOICE_PAYMENT SUBSCRIPTION SUBSCRIPTION_EVENT SERVICE_BROADCAST PAYMENT_ATTEMPT PAYMENT_METHOD TAG TAG_DEFINITION TENANT TENANT_KVS]
	ObjectType string `json:"objectType,omitempty"`

	// tag definition Id
	// Format: uuid
	TagDefinitionID strfmt.UUID `json:"tagDefinitionId,omitempty"`

	// tag definition name
	TagDefinitionName string `json:"tagDefinitionName,omitempty"`

	// tag Id
	// Format: uuid
	TagID strfmt.UUID `json:"tagId,omitempty"`
}

// Validate validates this tag
func (m *Tag) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuditLogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTagDefinitionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTagID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Tag) validateAuditLogs(formats strfmt.Registry) error {
	if swag.IsZero(m.AuditLogs) { // not required
		return nil
	}

	for i := 0; i < len(m.AuditLogs); i++ {
		if swag.IsZero(m.AuditLogs[i]) { // not required
			continue
		}

		if m.AuditLogs[i] != nil {
			if err := m.AuditLogs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Tag) validateObjectID(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectID) { // not required
		return nil
	}

	if err := validate.FormatOf("objectId", "body", "uuid", m.ObjectID.String(), formats); err != nil {
		return err
	}

	return nil
}

var tagTypeObjectTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACCOUNT","ACCOUNT_EMAIL","BLOCKING_STATES","BUNDLE","CUSTOM_FIELD","INVOICE","PAYMENT","TRANSACTION","INVOICE_ITEM","INVOICE_PAYMENT","SUBSCRIPTION","SUBSCRIPTION_EVENT","SERVICE_BROADCAST","PAYMENT_ATTEMPT","PAYMENT_METHOD","TAG","TAG_DEFINITION","TENANT","TENANT_KVS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tagTypeObjectTypePropEnum = append(tagTypeObjectTypePropEnum, v)
	}
}

const (

	// TagObjectTypeACCOUNT captures enum value "ACCOUNT"
	TagObjectTypeACCOUNT string = "ACCOUNT"

	// TagObjectTypeACCOUNTEMAIL captures enum value "ACCOUNT_EMAIL"
	TagObjectTypeACCOUNTEMAIL string = "ACCOUNT_EMAIL"

	// TagObjectTypeBLOCKINGSTATES captures enum value "BLOCKING_STATES"
	TagObjectTypeBLOCKINGSTATES string = "BLOCKING_STATES"

	// TagObjectTypeBUNDLE captures enum value "BUNDLE"
	TagObjectTypeBUNDLE string = "BUNDLE"

	// TagObjectTypeCUSTOMFIELD captures enum value "CUSTOM_FIELD"
	TagObjectTypeCUSTOMFIELD string = "CUSTOM_FIELD"

	// TagObjectTypeINVOICE captures enum value "INVOICE"
	TagObjectTypeINVOICE string = "INVOICE"

	// TagObjectTypePAYMENT captures enum value "PAYMENT"
	TagObjectTypePAYMENT string = "PAYMENT"

	// TagObjectTypeTRANSACTION captures enum value "TRANSACTION"
	TagObjectTypeTRANSACTION string = "TRANSACTION"

	// TagObjectTypeINVOICEITEM captures enum value "INVOICE_ITEM"
	TagObjectTypeINVOICEITEM string = "INVOICE_ITEM"

	// TagObjectTypeINVOICEPAYMENT captures enum value "INVOICE_PAYMENT"
	TagObjectTypeINVOICEPAYMENT string = "INVOICE_PAYMENT"

	// TagObjectTypeSUBSCRIPTION captures enum value "SUBSCRIPTION"
	TagObjectTypeSUBSCRIPTION string = "SUBSCRIPTION"

	// TagObjectTypeSUBSCRIPTIONEVENT captures enum value "SUBSCRIPTION_EVENT"
	TagObjectTypeSUBSCRIPTIONEVENT string = "SUBSCRIPTION_EVENT"

	// TagObjectTypeSERVICEBROADCAST captures enum value "SERVICE_BROADCAST"
	TagObjectTypeSERVICEBROADCAST string = "SERVICE_BROADCAST"

	// TagObjectTypePAYMENTATTEMPT captures enum value "PAYMENT_ATTEMPT"
	TagObjectTypePAYMENTATTEMPT string = "PAYMENT_ATTEMPT"

	// TagObjectTypePAYMENTMETHOD captures enum value "PAYMENT_METHOD"
	TagObjectTypePAYMENTMETHOD string = "PAYMENT_METHOD"

	// TagObjectTypeTAG captures enum value "TAG"
	TagObjectTypeTAG string = "TAG"

	// TagObjectTypeTAGDEFINITION captures enum value "TAG_DEFINITION"
	TagObjectTypeTAGDEFINITION string = "TAG_DEFINITION"

	// TagObjectTypeTENANT captures enum value "TENANT"
	TagObjectTypeTENANT string = "TENANT"

	// TagObjectTypeTENANTKVS captures enum value "TENANT_KVS"
	TagObjectTypeTENANTKVS string = "TENANT_KVS"
)

// prop value enum
func (m *Tag) validateObjectTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tagTypeObjectTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Tag) validateObjectType(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectType) { // not required
		return nil
	}

	// value enum
	if err := m.validateObjectTypeEnum("objectType", "body", m.ObjectType); err != nil {
		return err
	}

	return nil
}

func (m *Tag) validateTagDefinitionID(formats strfmt.Registry) error {
	if swag.IsZero(m.TagDefinitionID) { // not required
		return nil
	}

	if err := validate.FormatOf("tagDefinitionId", "body", "uuid", m.TagDefinitionID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Tag) validateTagID(formats strfmt.Registry) error {
	if swag.IsZero(m.TagID) { // not required
		return nil
	}

	if err := validate.FormatOf("tagId", "body", "uuid", m.TagID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this tag based on the context it is used
func (m *Tag) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuditLogs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Tag) contextValidateAuditLogs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AuditLogs); i++ {

		if m.AuditLogs[i] != nil {
			if err := m.AuditLogs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Tag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Tag) UnmarshalBinary(b []byte) error {
	var res Tag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
