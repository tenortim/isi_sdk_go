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

// NdmpUser ndmp user
// swagger:model NdmpUser
type NdmpUser struct {

	// The password for the NDMP administrator.
	// Required: true
	Password *string `json:"password"`
}

// Validate validates this ndmp user
func (m *NdmpUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NdmpUser) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NdmpUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NdmpUser) UnmarshalBinary(b []byte) error {
	var res NdmpUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
