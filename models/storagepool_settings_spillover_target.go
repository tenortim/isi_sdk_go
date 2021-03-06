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

// StoragepoolSettingsSpilloverTarget storagepool settings spillover target
// swagger:model StoragepoolSettingsSpilloverTarget
type StoragepoolSettingsSpilloverTarget struct {

	// Target pool ID if target specified, otherwise null.
	NameOrID string `json:"name_or_id,omitempty"`

	// Type of target pool.
	// Required: true
	// Enum: [storagepool anywhere]
	Type *string `json:"type"`
}

// Validate validates this storagepool settings spillover target
func (m *StoragepoolSettingsSpilloverTarget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var storagepoolSettingsSpilloverTargetTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["storagepool","anywhere"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storagepoolSettingsSpilloverTargetTypeTypePropEnum = append(storagepoolSettingsSpilloverTargetTypeTypePropEnum, v)
	}
}

const (

	// StoragepoolSettingsSpilloverTargetTypeStoragepool captures enum value "storagepool"
	StoragepoolSettingsSpilloverTargetTypeStoragepool string = "storagepool"

	// StoragepoolSettingsSpilloverTargetTypeAnywhere captures enum value "anywhere"
	StoragepoolSettingsSpilloverTargetTypeAnywhere string = "anywhere"
)

// prop value enum
func (m *StoragepoolSettingsSpilloverTarget) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, storagepoolSettingsSpilloverTargetTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StoragepoolSettingsSpilloverTarget) validateType(formats strfmt.Registry) error {

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
func (m *StoragepoolSettingsSpilloverTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoragepoolSettingsSpilloverTarget) UnmarshalBinary(b []byte) error {
	var res StoragepoolSettingsSpilloverTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
