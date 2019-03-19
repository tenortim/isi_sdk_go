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

// NodeStateNodeServicelight node state node servicelight
// swagger:model NodeStateNodeServicelight
type NodeStateNodeServicelight struct {
	NodeStateServicelightExtended

	// The node service light state (True = on).
	// Required: true
	Enabled *bool `json:"enabled"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *NodeStateNodeServicelight) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 NodeStateServicelightExtended
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.NodeStateServicelightExtended = aO0

	// AO1
	var dataAO1 struct {
		Enabled *bool `json:"enabled"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Enabled = dataAO1.Enabled

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m NodeStateNodeServicelight) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.NodeStateServicelightExtended)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Enabled *bool `json:"enabled"`
	}

	dataAO1.Enabled = m.Enabled

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this node state node servicelight
func (m *NodeStateNodeServicelight) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with NodeStateServicelightExtended
	if err := m.NodeStateServicelightExtended.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeStateNodeServicelight) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodeStateNodeServicelight) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeStateNodeServicelight) UnmarshalBinary(b []byte) error {
	var res NodeStateNodeServicelight
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}