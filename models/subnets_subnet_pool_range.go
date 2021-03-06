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

// SubnetsSubnetPoolRange subnets subnet pool range
// swagger:model SubnetsSubnetPoolRange
type SubnetsSubnetPoolRange struct {

	// High IP
	// Required: true
	High *string `json:"high"`

	// Low IP
	// Required: true
	Low *string `json:"low"`
}

// Validate validates this subnets subnet pool range
func (m *SubnetsSubnetPoolRange) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHigh(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLow(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubnetsSubnetPoolRange) validateHigh(formats strfmt.Registry) error {

	if err := validate.Required("high", "body", m.High); err != nil {
		return err
	}

	return nil
}

func (m *SubnetsSubnetPoolRange) validateLow(formats strfmt.Registry) error {

	if err := validate.Required("low", "body", m.Low); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubnetsSubnetPoolRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubnetsSubnetPoolRange) UnmarshalBinary(b []byte) error {
	var res SubnetsSubnetPoolRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
