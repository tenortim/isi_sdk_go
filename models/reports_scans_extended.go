// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ReportsScansExtended reports scans extended
// swagger:model ReportsScansExtended
type ReportsScansExtended struct {
	ReportsScans

	// Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options).
	Resume string `json:"resume,omitempty"`

	// Total number of items available.
	Total int64 `json:"total,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ReportsScansExtended) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ReportsScans
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ReportsScans = aO0

	// AO1
	var dataAO1 struct {
		Resume string `json:"resume,omitempty"`

		Total int64 `json:"total,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Resume = dataAO1.Resume

	m.Total = dataAO1.Total

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ReportsScansExtended) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ReportsScans)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Resume string `json:"resume,omitempty"`

		Total int64 `json:"total,omitempty"`
	}

	dataAO1.Resume = m.Resume

	dataAO1.Total = m.Total

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this reports scans extended
func (m *ReportsScansExtended) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ReportsScans
	if err := m.ReportsScans.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ReportsScansExtended) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportsScansExtended) UnmarshalBinary(b []byte) error {
	var res ReportsScansExtended
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}