// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SyncRuleExtended sync rule extended
// swagger:model SyncRuleExtended
type SyncRuleExtended struct {
	SyncRule

	// User-entered description of this performance rule.
	// Required: true
	Description *string `json:"description"`

	// Whether this performance rule is currently in effect during its specified intervals.
	// Required: true
	Enabled *bool `json:"enabled"`

	// The system ID given to this performance rule.
	// Required: true
	ID *string `json:"id"`

	// Amount the specified system resource type is limited by this rule.  Units are kb/s for bandwidth, files/s for file-count, processing percentage used for cpu, or percentage of maximum available workers.
	// Required: true
	Limit *int64 `json:"limit"`

	// The type of system resource this rule limits.
	// Required: true
	// Enum: [bandwidth file_count cpu worker]
	Type *string `json:"type"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SyncRuleExtended) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SyncRule
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SyncRule = aO0

	// AO1
	var dataAO1 struct {
		Description *string `json:"description"`

		Enabled *bool `json:"enabled"`

		ID *string `json:"id"`

		Limit *int64 `json:"limit"`

		Type *string `json:"type"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Description = dataAO1.Description

	m.Enabled = dataAO1.Enabled

	m.ID = dataAO1.ID

	m.Limit = dataAO1.Limit

	m.Type = dataAO1.Type

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SyncRuleExtended) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.SyncRule)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Description *string `json:"description"`

		Enabled *bool `json:"enabled"`

		ID *string `json:"id"`

		Limit *int64 `json:"limit"`

		Type *string `json:"type"`
	}

	dataAO1.Description = m.Description

	dataAO1.Enabled = m.Enabled

	dataAO1.ID = m.ID

	dataAO1.Limit = m.Limit

	dataAO1.Type = m.Type

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this sync rule extended
func (m *SyncRuleExtended) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SyncRule
	if err := m.SyncRule.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SyncRuleExtended) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *SyncRuleExtended) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *SyncRuleExtended) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *SyncRuleExtended) validateLimit(formats strfmt.Registry) error {

	if err := validate.Required("limit", "body", m.Limit); err != nil {
		return err
	}

	return nil
}

var syncRuleExtendedTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bandwidth","file_count","cpu","worker"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		syncRuleExtendedTypeTypePropEnum = append(syncRuleExtendedTypeTypePropEnum, v)
	}
}

// property enum
func (m *SyncRuleExtended) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, syncRuleExtendedTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SyncRuleExtended) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SyncRuleExtended) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SyncRuleExtended) UnmarshalBinary(b []byte) error {
	var res SyncRuleExtended
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
