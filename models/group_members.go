// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GroupMembers group members
// swagger:model GroupMembers
type GroupMembers struct {

	// members
	Members []*GroupMember `json:"members"`

	// Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options).
	Resume string `json:"resume,omitempty"`
}

// Validate validates this group members
func (m *GroupMembers) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMembers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupMembers) validateMembers(formats strfmt.Registry) error {

	if swag.IsZero(m.Members) { // not required
		return nil
	}

	for i := 0; i < len(m.Members); i++ {
		if swag.IsZero(m.Members[i]) { // not required
			continue
		}

		if m.Members[i] != nil {
			if err := m.Members[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupMembers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupMembers) UnmarshalBinary(b []byte) error {
	var res GroupMembers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}