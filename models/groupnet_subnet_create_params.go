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

// GroupnetSubnetCreateParams groupnet subnet create params
// swagger:model GroupnetSubnetCreateParams
type GroupnetSubnetCreateParams struct {
	GroupnetSubnet

	// IP address format.
	// Required: true
	// Enum: [ipv4 ipv6]
	AddrFamily *string `json:"addr_family"`

	// The name of the subnet.
	// Required: true
	// Max Length: 32
	Name *string `json:"name"`

	// Subnet Prefix Length.
	// Required: true
	// Maximum: 128
	// Minimum: 1
	Prefixlen *int64 `json:"prefixlen"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GroupnetSubnetCreateParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 GroupnetSubnet
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.GroupnetSubnet = aO0

	// AO1
	var dataAO1 struct {
		AddrFamily *string `json:"addr_family"`

		Name *string `json:"name"`

		Prefixlen *int64 `json:"prefixlen"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AddrFamily = dataAO1.AddrFamily

	m.Name = dataAO1.Name

	m.Prefixlen = dataAO1.Prefixlen

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GroupnetSubnetCreateParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.GroupnetSubnet)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		AddrFamily *string `json:"addr_family"`

		Name *string `json:"name"`

		Prefixlen *int64 `json:"prefixlen"`
	}

	dataAO1.AddrFamily = m.AddrFamily

	dataAO1.Name = m.Name

	dataAO1.Prefixlen = m.Prefixlen

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this groupnet subnet create params
func (m *GroupnetSubnetCreateParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with GroupnetSubnet
	if err := m.GroupnetSubnet.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddrFamily(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrefixlen(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var groupnetSubnetCreateParamsTypeAddrFamilyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ipv4","ipv6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		groupnetSubnetCreateParamsTypeAddrFamilyPropEnum = append(groupnetSubnetCreateParamsTypeAddrFamilyPropEnum, v)
	}
}

// property enum
func (m *GroupnetSubnetCreateParams) validateAddrFamilyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, groupnetSubnetCreateParamsTypeAddrFamilyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GroupnetSubnetCreateParams) validateAddrFamily(formats strfmt.Registry) error {

	if err := validate.Required("addr_family", "body", m.AddrFamily); err != nil {
		return err
	}

	// value enum
	if err := m.validateAddrFamilyEnum("addr_family", "body", *m.AddrFamily); err != nil {
		return err
	}

	return nil
}

func (m *GroupnetSubnetCreateParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 32); err != nil {
		return err
	}

	return nil
}

func (m *GroupnetSubnetCreateParams) validatePrefixlen(formats strfmt.Registry) error {

	if err := validate.Required("prefixlen", "body", m.Prefixlen); err != nil {
		return err
	}

	if err := validate.MinimumInt("prefixlen", "body", int64(*m.Prefixlen), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("prefixlen", "body", int64(*m.Prefixlen), 128, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupnetSubnetCreateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupnetSubnetCreateParams) UnmarshalBinary(b []byte) error {
	var res GroupnetSubnetCreateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
