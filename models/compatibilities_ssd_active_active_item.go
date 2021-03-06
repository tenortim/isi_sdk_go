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

// CompatibilitiesSsdActiveActiveItem compatibilities ssd active active item
// swagger:model CompatibilitiesSsdActiveActiveItem
type CompatibilitiesSsdActiveActiveItem struct {

	// The node class of an ssd compatibility
	// Required: true
	Class1 *string `json:"class_1"`

	// Is this SSD Compatibility Count Compatible.
	// Required: true
	Count *bool `json:"count"`

	// The id of this ssd compatibility
	// Required: true
	ID *int64 `json:"id"`
}

// Validate validates this compatibilities ssd active active item
func (m *CompatibilitiesSsdActiveActiveItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClass1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompatibilitiesSsdActiveActiveItem) validateClass1(formats strfmt.Registry) error {

	if err := validate.Required("class_1", "body", m.Class1); err != nil {
		return err
	}

	return nil
}

func (m *CompatibilitiesSsdActiveActiveItem) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

func (m *CompatibilitiesSsdActiveActiveItem) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CompatibilitiesSsdActiveActiveItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompatibilitiesSsdActiveActiveItem) UnmarshalBinary(b []byte) error {
	var res CompatibilitiesSsdActiveActiveItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
