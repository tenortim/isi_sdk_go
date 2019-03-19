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

// AuthPrivileges auth privileges
// swagger:model AuthPrivileges
type AuthPrivileges struct {

	// privileges
	Privileges []*AuthPrivilege `json:"privileges"`

	// Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options).
	Resume string `json:"resume,omitempty"`
}

// Validate validates this auth privileges
func (m *AuthPrivileges) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrivileges(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthPrivileges) validatePrivileges(formats strfmt.Registry) error {

	if swag.IsZero(m.Privileges) { // not required
		return nil
	}

	for i := 0; i < len(m.Privileges); i++ {
		if swag.IsZero(m.Privileges[i]) { // not required
			continue
		}

		if m.Privileges[i] != nil {
			if err := m.Privileges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("privileges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuthPrivileges) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthPrivileges) UnmarshalBinary(b []byte) error {
	var res AuthPrivileges
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}