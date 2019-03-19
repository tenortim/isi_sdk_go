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

// AuthAccessAccessItem auth access access item
// swagger:model AuthAccessAccessItem
type AuthAccessAccessItem struct {

	// Specifies properties for access rights.
	File *AuthAccessAccessItemFile `json:"file,omitempty"`

	// Specifies the ID of the user.
	ID string `json:"id,omitempty"`

	// Specifies the permissions that the user has on the file.
	Permissions *AuthAccessAccessItemPermissions `json:"permissions,omitempty"`

	// Specifies a list of the relevant Access Control Entries for the user.
	RelevantAces []*AuthAccessAccessItemRelevantAce `json:"relevant_aces"`

	// Specifies the persona for the user.
	User *AuthAccessAccessItemUser `json:"user,omitempty"`
}

// Validate validates this auth access access item
func (m *AuthAccessAccessItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelevantAces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthAccessAccessItem) validateFile(formats strfmt.Registry) error {

	if swag.IsZero(m.File) { // not required
		return nil
	}

	if m.File != nil {
		if err := m.File.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("file")
			}
			return err
		}
	}

	return nil
}

func (m *AuthAccessAccessItem) validatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	if m.Permissions != nil {
		if err := m.Permissions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("permissions")
			}
			return err
		}
	}

	return nil
}

func (m *AuthAccessAccessItem) validateRelevantAces(formats strfmt.Registry) error {

	if swag.IsZero(m.RelevantAces) { // not required
		return nil
	}

	for i := 0; i < len(m.RelevantAces); i++ {
		if swag.IsZero(m.RelevantAces[i]) { // not required
			continue
		}

		if m.RelevantAces[i] != nil {
			if err := m.RelevantAces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relevant_aces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AuthAccessAccessItem) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuthAccessAccessItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthAccessAccessItem) UnmarshalBinary(b []byte) error {
	var res AuthAccessAccessItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}