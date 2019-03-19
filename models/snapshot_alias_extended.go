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

// SnapshotAliasExtended snapshot alias extended
// swagger:model SnapshotAliasExtended
type SnapshotAliasExtended struct {

	// The Unix Epoch time the snapshot alias was created.
	// Required: true
	Created *int64 `json:"created"`

	// The system ID given to the snapshot alias.
	// Required: true
	ID *int64 `json:"id"`

	// The user or system supplied snapshot alias name.
	// Required: true
	Name *string `json:"name"`

	// The ID of the snapshot pointed to.
	// Required: true
	TargetID *int64 `json:"target_id"`

	// The name of the snapshot pointed to.
	// Required: true
	TargetName *string `json:"target_name"`
}

// Validate validates this snapshot alias extended
func (m *SnapshotAliasExtended) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotAliasExtended) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotAliasExtended) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotAliasExtended) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotAliasExtended) validateTargetID(formats strfmt.Registry) error {

	if err := validate.Required("target_id", "body", m.TargetID); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotAliasExtended) validateTargetName(formats strfmt.Registry) error {

	if err := validate.Required("target_name", "body", m.TargetName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotAliasExtended) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotAliasExtended) UnmarshalBinary(b []byte) error {
	var res SnapshotAliasExtended
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}