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

// WormDomainExtended worm domain extended
// swagger:model WormDomainExtended
type WormDomainExtended struct {
	WormDomain

	// Specifies the system-assigned ID for the protection domain.
	// Required: true
	ID *int64 `json:"id"`

	// True if OneFS is still in the process of creating this domain and is unable to prevent files from being modified or deleted. If false, indicates that the domain is fully created and is able to prevent files from being modified or deleted.
	// Required: true
	Incomplete *bool `json:"incomplete"`

	// Specifies the logical inode number (LIN) for the root of this domain.
	// Required: true
	Lin *int64 `json:"lin"`

	// Specifies the maximum amount of time, in seconds, that a file in this domain will be protected. This setting will override the retention period of any file committed with a longer retention period. If this parameter is set to null, an infinite length retention period is set.
	// Required: true
	MaxModifies *int64 `json:"max_modifies"`

	// Specifies the root path of this domain. Files in this directory and all sub-directories will be protected.
	// Required: true
	Path *string `json:"path"`

	// When this value is set to 'on', files in this domain can be deleted through the privileged delete feature. If this value is set to 'disabled', privileged file deletes are permanently disabled and cannot be turned on again.
	// Required: true
	// Enum: [on off disabled]
	PrivilegedDelete *string `json:"privileged_delete"`

	// Specifies the number of times this domain has been modified and the number of times the attributes for the domain have changed. A SmartLock domain can be modified a fixed number of times as defined by the 'max_modifies' parameter.
	// Required: true
	TotalModifies *int64 `json:"total_modifies"`

	// Specifies whether the domain is an enterprise domain or a compliance domain. Compliance domains can not be created on enterprise clusters. Enterprise and compliance domains can be created on compliance clusters.
	// Required: true
	// Enum: [enterprise compliance]
	Type *string `json:"type"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WormDomainExtended) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WormDomain
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WormDomain = aO0

	// AO1
	var dataAO1 struct {
		ID *int64 `json:"id"`

		Incomplete *bool `json:"incomplete"`

		Lin *int64 `json:"lin"`

		MaxModifies *int64 `json:"max_modifies"`

		Path *string `json:"path"`

		PrivilegedDelete *string `json:"privileged_delete"`

		TotalModifies *int64 `json:"total_modifies"`

		Type *string `json:"type"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ID = dataAO1.ID

	m.Incomplete = dataAO1.Incomplete

	m.Lin = dataAO1.Lin

	m.MaxModifies = dataAO1.MaxModifies

	m.Path = dataAO1.Path

	m.PrivilegedDelete = dataAO1.PrivilegedDelete

	m.TotalModifies = dataAO1.TotalModifies

	m.Type = dataAO1.Type

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WormDomainExtended) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.WormDomain)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		ID *int64 `json:"id"`

		Incomplete *bool `json:"incomplete"`

		Lin *int64 `json:"lin"`

		MaxModifies *int64 `json:"max_modifies"`

		Path *string `json:"path"`

		PrivilegedDelete *string `json:"privileged_delete"`

		TotalModifies *int64 `json:"total_modifies"`

		Type *string `json:"type"`
	}

	dataAO1.ID = m.ID

	dataAO1.Incomplete = m.Incomplete

	dataAO1.Lin = m.Lin

	dataAO1.MaxModifies = m.MaxModifies

	dataAO1.Path = m.Path

	dataAO1.PrivilegedDelete = m.PrivilegedDelete

	dataAO1.TotalModifies = m.TotalModifies

	dataAO1.Type = m.Type

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this worm domain extended
func (m *WormDomainExtended) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WormDomain
	if err := m.WormDomain.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncomplete(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxModifies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivilegedDelete(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalModifies(formats); err != nil {
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

func (m *WormDomainExtended) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *WormDomainExtended) validateIncomplete(formats strfmt.Registry) error {

	if err := validate.Required("incomplete", "body", m.Incomplete); err != nil {
		return err
	}

	return nil
}

func (m *WormDomainExtended) validateLin(formats strfmt.Registry) error {

	if err := validate.Required("lin", "body", m.Lin); err != nil {
		return err
	}

	return nil
}

func (m *WormDomainExtended) validateMaxModifies(formats strfmt.Registry) error {

	if err := validate.Required("max_modifies", "body", m.MaxModifies); err != nil {
		return err
	}

	return nil
}

func (m *WormDomainExtended) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

var wormDomainExtendedTypePrivilegedDeletePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["on","off","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wormDomainExtendedTypePrivilegedDeletePropEnum = append(wormDomainExtendedTypePrivilegedDeletePropEnum, v)
	}
}

// property enum
func (m *WormDomainExtended) validatePrivilegedDeleteEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, wormDomainExtendedTypePrivilegedDeletePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WormDomainExtended) validatePrivilegedDelete(formats strfmt.Registry) error {

	if err := validate.Required("privileged_delete", "body", m.PrivilegedDelete); err != nil {
		return err
	}

	// value enum
	if err := m.validatePrivilegedDeleteEnum("privileged_delete", "body", *m.PrivilegedDelete); err != nil {
		return err
	}

	return nil
}

func (m *WormDomainExtended) validateTotalModifies(formats strfmt.Registry) error {

	if err := validate.Required("total_modifies", "body", m.TotalModifies); err != nil {
		return err
	}

	return nil
}

var wormDomainExtendedTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enterprise","compliance"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wormDomainExtendedTypeTypePropEnum = append(wormDomainExtendedTypeTypePropEnum, v)
	}
}

// property enum
func (m *WormDomainExtended) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, wormDomainExtendedTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WormDomainExtended) validateType(formats strfmt.Registry) error {

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
func (m *WormDomainExtended) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WormDomainExtended) UnmarshalBinary(b []byte) error {
	var res WormDomainExtended
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}