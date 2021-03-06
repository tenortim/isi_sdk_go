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

// AuthGroupExtended auth group extended
// swagger:model AuthGroupExtended
type AuthGroupExtended struct {

	// Specifies the distinguished name for the user.
	Dn string `json:"dn,omitempty"`

	// Specifies the DNS domain.
	DNSDomain string `json:"dns_domain,omitempty"`

	// Specifies the domain that the object is part of.
	Domain string `json:"domain,omitempty"`

	// If true, the GID was generated.
	GeneratedGid bool `json:"generated_gid,omitempty"`

	// Specifies properties for a persona, which consists of either a 'type' and a 'name' or an 'ID'.
	Gid *GroupMember `json:"gid,omitempty"`

	// Specifies the user or group ID.
	// Required: true
	ID *string `json:"id"`

	// member of
	MemberOf []*GroupMember `json:"member_of"`

	// Specifies a user or group name.
	// Required: true
	Name *string `json:"name"`

	// Specifies the authentication provider that the object belongs to.
	Provider string `json:"provider,omitempty"`

	// Specifies a user or group name.
	SamAccountName string `json:"sam_account_name,omitempty"`

	// Specifies properties for a persona, which consists of either a 'type' and a 'name' or an 'ID'.
	Sid *GroupMember `json:"sid,omitempty"`

	// Specifies the object type.
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this auth group extended
func (m *AuthGroupExtended) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemberOf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSid(formats); err != nil {
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

func (m *AuthGroupExtended) validateGid(formats strfmt.Registry) error {

	if swag.IsZero(m.Gid) { // not required
		return nil
	}

	if m.Gid != nil {
		if err := m.Gid.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gid")
			}
			return err
		}
	}

	return nil
}

func (m *AuthGroupExtended) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *AuthGroupExtended) validateMemberOf(formats strfmt.Registry) error {

	if swag.IsZero(m.MemberOf) { // not required
		return nil
	}

	for i := 0; i < len(m.MemberOf); i++ {
		if swag.IsZero(m.MemberOf[i]) { // not required
			continue
		}

		if m.MemberOf[i] != nil {
			if err := m.MemberOf[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("member_of" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AuthGroupExtended) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *AuthGroupExtended) validateSid(formats strfmt.Registry) error {

	if swag.IsZero(m.Sid) { // not required
		return nil
	}

	if m.Sid != nil {
		if err := m.Sid.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sid")
			}
			return err
		}
	}

	return nil
}

func (m *AuthGroupExtended) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuthGroupExtended) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthGroupExtended) UnmarshalBinary(b []byte) error {
	var res AuthGroupExtended
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
