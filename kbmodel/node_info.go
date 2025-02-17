// Code generated by go-swagger; DO NOT EDIT.

package kbmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NodeInfo node info
// swagger:model NodeInfo
type NodeInfo struct {

	// api version
	APIVersion string `json:"apiVersion,omitempty"`

	// boot time
	// Format: date-time
	BootTime strfmt.DateTime `json:"bootTime,omitempty"`

	// common version
	CommonVersion string `json:"commonVersion,omitempty"`

	// kb version
	KbVersion string `json:"kbVersion,omitempty"`

	// last updated date
	// Format: date-time
	LastUpdatedDate strfmt.DateTime `json:"lastUpdatedDate,omitempty"`

	// node name
	NodeName string `json:"nodeName,omitempty"`

	// platform version
	PlatformVersion string `json:"platformVersion,omitempty"`

	// plugin Api version
	PluginAPIVersion string `json:"pluginApiVersion,omitempty"`

	// plugins info
	PluginsInfo []*PluginInfo `json:"pluginsInfo"`
}

// Validate validates this node info
func (m *NodeInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBootTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePluginsInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeInfo) validateBootTime(formats strfmt.Registry) error {

	if swag.IsZero(m.BootTime) { // not required
		return nil
	}

	if err := validate.FormatOf("bootTime", "body", "date-time", m.BootTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NodeInfo) validateLastUpdatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdatedDate", "body", "date-time", m.LastUpdatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NodeInfo) validatePluginsInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.PluginsInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.PluginsInfo); i++ {
		if swag.IsZero(m.PluginsInfo[i]) { // not required
			continue
		}

		if m.PluginsInfo[i] != nil {
			if err := m.PluginsInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pluginsInfo" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodeInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeInfo) UnmarshalBinary(b []byte) error {
	var res NodeInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
