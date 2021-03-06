// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NdmpSettingsVariables Get list of preferred environment variable.
// swagger:model NdmpSettingsVariables
type NdmpSettingsVariables struct {

	// Resume string returned by previous query.
	Resume string `json:"resume,omitempty"`

	// The number of backup paths.
	Total int64 `json:"total,omitempty"`

	// variables
	Variables []*NdmpSettingsVariablesVariable `json:"variables"`
}

// Validate validates this ndmp settings variables
func (m *NdmpSettingsVariables) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVariables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NdmpSettingsVariables) validateVariables(formats strfmt.Registry) error {

	if swag.IsZero(m.Variables) { // not required
		return nil
	}

	for i := 0; i < len(m.Variables); i++ {
		if swag.IsZero(m.Variables[i]) { // not required
			continue
		}

		if m.Variables[i] != nil {
			if err := m.Variables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NdmpSettingsVariables) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NdmpSettingsVariables) UnmarshalBinary(b []byte) error {
	var res NdmpSettingsVariables
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
