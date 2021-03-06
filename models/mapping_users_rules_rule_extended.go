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

// MappingUsersRulesRuleExtended mapping users rules rule extended
// swagger:model MappingUsersRulesRuleExtended
type MappingUsersRulesRuleExtended struct {

	// Specifies the operator to make rules on specified users or groups.
	// Enum: [append insert replace trim union]
	Operator string `json:"operator,omitempty"`

	// Specifies the properties for user mapping rules.
	Options *MappingUsersRulesRuleOptionsExtended `json:"options,omitempty"`

	// user1
	// Required: true
	User1 *MappingUsersRulesRuleUser1 `json:"user1"`

	// user2
	User2 *MappingUsersRulesRuleUser2Extended `json:"user2,omitempty"`
}

// Validate validates this mapping users rules rule extended
func (m *MappingUsersRulesRuleExtended) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser2(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var mappingUsersRulesRuleExtendedTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["append","insert","replace","trim","union"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mappingUsersRulesRuleExtendedTypeOperatorPropEnum = append(mappingUsersRulesRuleExtendedTypeOperatorPropEnum, v)
	}
}

const (

	// MappingUsersRulesRuleExtendedOperatorAppend captures enum value "append"
	MappingUsersRulesRuleExtendedOperatorAppend string = "append"

	// MappingUsersRulesRuleExtendedOperatorInsert captures enum value "insert"
	MappingUsersRulesRuleExtendedOperatorInsert string = "insert"

	// MappingUsersRulesRuleExtendedOperatorReplace captures enum value "replace"
	MappingUsersRulesRuleExtendedOperatorReplace string = "replace"

	// MappingUsersRulesRuleExtendedOperatorTrim captures enum value "trim"
	MappingUsersRulesRuleExtendedOperatorTrim string = "trim"

	// MappingUsersRulesRuleExtendedOperatorUnion captures enum value "union"
	MappingUsersRulesRuleExtendedOperatorUnion string = "union"
)

// prop value enum
func (m *MappingUsersRulesRuleExtended) validateOperatorEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, mappingUsersRulesRuleExtendedTypeOperatorPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MappingUsersRulesRuleExtended) validateOperator(formats strfmt.Registry) error {

	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

func (m *MappingUsersRulesRuleExtended) validateOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.Options) { // not required
		return nil
	}

	if m.Options != nil {
		if err := m.Options.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("options")
			}
			return err
		}
	}

	return nil
}

func (m *MappingUsersRulesRuleExtended) validateUser1(formats strfmt.Registry) error {

	if err := validate.Required("user1", "body", m.User1); err != nil {
		return err
	}

	if m.User1 != nil {
		if err := m.User1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user1")
			}
			return err
		}
	}

	return nil
}

func (m *MappingUsersRulesRuleExtended) validateUser2(formats strfmt.Registry) error {

	if swag.IsZero(m.User2) { // not required
		return nil
	}

	if m.User2 != nil {
		if err := m.User2.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user2")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MappingUsersRulesRuleExtended) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MappingUsersRulesRuleExtended) UnmarshalBinary(b []byte) error {
	var res MappingUsersRulesRuleExtended
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
