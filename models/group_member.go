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

// GroupMember Specifies properties for a persona, which consists of either a 'type' and a 'name' or an 'ID'.
// swagger:model GroupMember
type GroupMember struct {

	// Specifies the serialized form of a persona, which can be 'UID:0', 'USER:name', 'GID:0', 'GROUP:wheel', or 'SID:S-1-1'.
	ID string `json:"id,omitempty"`

	// Specifies the persona name, which must be combined with a type.
	Name string `json:"name,omitempty"`

	// Specifies the type of persona, which must be combined with a name.
	// Enum: [user group wellknown]
	Type string `json:"type,omitempty"`
}

// Validate validates this group member
func (m *GroupMember) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var groupMemberTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["user","group","wellknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		groupMemberTypeTypePropEnum = append(groupMemberTypeTypePropEnum, v)
	}
}

const (

	// GroupMemberTypeUser captures enum value "user"
	GroupMemberTypeUser string = "user"

	// GroupMemberTypeGroup captures enum value "group"
	GroupMemberTypeGroup string = "group"

	// GroupMemberTypeWellknown captures enum value "wellknown"
	GroupMemberTypeWellknown string = "wellknown"
)

// prop value enum
func (m *GroupMember) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, groupMemberTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GroupMember) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupMember) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupMember) UnmarshalBinary(b []byte) error {
	var res GroupMember
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
