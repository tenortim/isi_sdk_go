// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NamespaceACL namespace Acl
// swagger:model NamespaceAcl
type NamespaceACL struct {

	// Provides the JSON array of access rights.
	// Min Items: 0
	ACL []*ACLObject `json:"acl"`

	// If the directory has access rights set, then this field is returned as acl. If the directory has POSIX permissions set, then this field is returned as mode.
	Authoritative string `json:"authoritative,omitempty"`

	// Provides the JSON object for the group persona of the owner.
	Group *MemberObject `json:"group,omitempty"`

	// Provides the POSIX mode.
	Mode string `json:"mode,omitempty"`

	// Provides the JSON object for the owner persona.
	Owner *MemberObject `json:"owner,omitempty"`
}

// Validate validates this namespace Acl
func (m *NamespaceACL) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateACL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NamespaceACL) validateACL(formats strfmt.Registry) error {

	if swag.IsZero(m.ACL) { // not required
		return nil
	}

	iACLSize := int64(len(m.ACL))

	if err := validate.MinItems("acl", "body", iACLSize, 0); err != nil {
		return err
	}

	for i := 0; i < len(m.ACL); i++ {
		if swag.IsZero(m.ACL[i]) { // not required
			continue
		}

		if m.ACL[i] != nil {
			if err := m.ACL[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("acl" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NamespaceACL) validateGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if m.Group != nil {
		if err := m.Group.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

func (m *NamespaceACL) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NamespaceACL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NamespaceACL) UnmarshalBinary(b []byte) error {
	var res NamespaceACL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}