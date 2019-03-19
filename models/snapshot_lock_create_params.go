// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SnapshotLockCreateParams snapshot lock create params
// swagger:model SnapshotLockCreateParams
type SnapshotLockCreateParams struct {
	SnapshotLock

	// Free form comment.
	Comment string `json:"comment,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SnapshotLockCreateParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SnapshotLock
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SnapshotLock = aO0

	// AO1
	var dataAO1 struct {
		Comment string `json:"comment,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Comment = dataAO1.Comment

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SnapshotLockCreateParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.SnapshotLock)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Comment string `json:"comment,omitempty"`
	}

	dataAO1.Comment = m.Comment

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this snapshot lock create params
func (m *SnapshotLockCreateParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SnapshotLock
	if err := m.SnapshotLock.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotLockCreateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotLockCreateParams) UnmarshalBinary(b []byte) error {
	var res SnapshotLockCreateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}