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

// SnapshotAliasCreateParams snapshot alias create params
// swagger:model SnapshotAliasCreateParams
type SnapshotAliasCreateParams struct {

	// The user or system supplied snapshot name.
	// Required: true
	Name *string `json:"name"`

	// Snapshot name target for the alias.
	// Required: true
	Target *string `json:"target"`
}

// Validate validates this snapshot alias create params
func (m *SnapshotAliasCreateParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotAliasCreateParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotAliasCreateParams) validateTarget(formats strfmt.Registry) error {

	if err := validate.Required("target", "body", m.Target); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotAliasCreateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotAliasCreateParams) UnmarshalBinary(b []byte) error {
	var res SnapshotAliasCreateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
