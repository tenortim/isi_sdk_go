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

// SwiftAccount This is an account for the Swift protocol.
// swagger:model SwiftAccount
type SwiftAccount struct {

	// Unique id of swift account
	ID string `json:"id,omitempty"`

	// Name of Swift account
	// Required: true
	Name *string `json:"name"`

	// Group with filesystem ownership of this account
	Swiftgroup string `json:"swiftgroup,omitempty"`

	// User with filesystem ownership of this account
	Swiftuser string `json:"swiftuser,omitempty"`

	// Users who are allowed to access Swift account
	Users []string `json:"users"`

	// Name of access zone for account
	Zone string `json:"zone,omitempty"`
}

// Validate validates this swift account
func (m *SwiftAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SwiftAccount) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SwiftAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SwiftAccount) UnmarshalBinary(b []byte) error {
	var res SwiftAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}