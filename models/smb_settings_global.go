// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SmbSettingsGlobal smb settings global
// swagger:model SmbSettingsGlobal
type SmbSettingsGlobal struct {

	// settings
	// Required: true
	Settings *SmbSettingsGlobalSettings `json:"settings"`
}

// Validate validates this smb settings global
func (m *SmbSettingsGlobal) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SmbSettingsGlobal) validateSettings(formats strfmt.Registry) error {

	if err := validate.Required("settings", "body", m.Settings); err != nil {
		return err
	}

	if m.Settings != nil {
		if err := m.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("settings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SmbSettingsGlobal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SmbSettingsGlobal) UnmarshalBinary(b []byte) error {
	var res SmbSettingsGlobal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
