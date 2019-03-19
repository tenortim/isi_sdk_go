// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// FilepoolDefaultPolicy A default filepool policy object
// swagger:model FilepoolDefaultPolicy
type FilepoolDefaultPolicy struct {

	// A default filepool policy object
	DefaultPolicy *FilepoolDefaultPolicyDefaultPolicy `json:"default-policy,omitempty"`
}

// Validate validates this filepool default policy
func (m *FilepoolDefaultPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FilepoolDefaultPolicy) validateDefaultPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultPolicy) { // not required
		return nil
	}

	if m.DefaultPolicy != nil {
		if err := m.DefaultPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default-policy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FilepoolDefaultPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FilepoolDefaultPolicy) UnmarshalBinary(b []byte) error {
	var res FilepoolDefaultPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}