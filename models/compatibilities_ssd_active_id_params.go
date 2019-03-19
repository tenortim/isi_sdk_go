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

// CompatibilitiesSsdActiveIDParams compatibilities ssd active Id params
// swagger:model CompatibilitiesSsdActiveIdParams
type CompatibilitiesSsdActiveIDParams struct {

	// Do not delete ssd compatibility, only assess if deletion is possible.
	Assess bool `json:"assess,omitempty"`

	// Are we enabling or disabling count
	// Required: true
	Count *bool `json:"count"`

	// The optional id of the second ssd compatibility.
	ID2 int64 `json:"id_2,omitempty"`
}

// Validate validates this compatibilities ssd active Id params
func (m *CompatibilitiesSsdActiveIDParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompatibilitiesSsdActiveIDParams) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CompatibilitiesSsdActiveIDParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompatibilitiesSsdActiveIDParams) UnmarshalBinary(b []byte) error {
	var res CompatibilitiesSsdActiveIDParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
