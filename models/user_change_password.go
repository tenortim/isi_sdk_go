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

// UserChangePassword Change Password Request
// swagger:model UserChangePassword
type UserChangePassword struct {

	// Specifies user's new password
	// Required: true
	NewPassword *string `json:"new_password"`

	// User's expired password
	// Required: true
	OldPassword *string `json:"old_password"`
}

// Validate validates this user change password
func (m *UserChangePassword) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNewPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOldPassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserChangePassword) validateNewPassword(formats strfmt.Registry) error {

	if err := validate.Required("new_password", "body", m.NewPassword); err != nil {
		return err
	}

	return nil
}

func (m *UserChangePassword) validateOldPassword(formats strfmt.Registry) error {

	if err := validate.Required("old_password", "body", m.OldPassword); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserChangePassword) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserChangePassword) UnmarshalBinary(b []byte) error {
	var res UserChangePassword
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
