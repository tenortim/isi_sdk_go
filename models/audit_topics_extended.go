// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AuditTopicsExtended audit topics extended
// swagger:model AuditTopicsExtended
type AuditTopicsExtended struct {
	AuditTopics

	// Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options).
	Resume string `json:"resume,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AuditTopicsExtended) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 AuditTopics
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.AuditTopics = aO0

	// AO1
	var dataAO1 struct {
		Resume string `json:"resume,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Resume = dataAO1.Resume

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AuditTopicsExtended) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.AuditTopics)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Resume string `json:"resume,omitempty"`
	}

	dataAO1.Resume = m.Resume

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this audit topics extended
func (m *AuditTopicsExtended) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AuditTopics
	if err := m.AuditTopics.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AuditTopicsExtended) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditTopicsExtended) UnmarshalBinary(b []byte) error {
	var res AuditTopicsExtended
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
