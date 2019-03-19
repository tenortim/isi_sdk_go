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

// PoolsPoolRule pools pool rule
// swagger:model PoolsPoolRule
type PoolsPoolRule struct {

	// Description for the provisioning rule.
	// Max Length: 128
	Description string `json:"description,omitempty"`

	// Interface name the provisioning rule applies to.
	Iface string `json:"iface,omitempty"`

	// Name of the provisioning rule.
	// Max Length: 32
	Name string `json:"name,omitempty"`

	// Node type the provisioning rule applies to.
	NodeType string `json:"node_type,omitempty"`
}

// Validate validates this pools pool rule
func (m *PoolsPoolRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoolsPoolRule) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 128); err != nil {
		return err
	}

	return nil
}

func (m *PoolsPoolRule) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 32); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PoolsPoolRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PoolsPoolRule) UnmarshalBinary(b []byte) error {
	var res PoolsPoolRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}