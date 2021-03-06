// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AuthAccessAccessItemPermissions auth access access item permissions
// swagger:model AuthAccessAccessItemPermissions
type AuthAccessAccessItemPermissions struct {

	// Returns a status message if the Null ACL is set.
	Dacl string `json:"dacl,omitempty"`

	// Returns a status message if the parent directory has the delete_child property set for the user. If the delete_child property is set for a user, that user is able to delete the file.the delete_child for the user.
	DeleteChild string `json:"delete_child,omitempty"`

	// Specifies the Access Control Entity (ACE) for the user.
	Expected string `json:"expected,omitempty"`

	// Returns a status message if the user owns the file.
	Ownership string `json:"ownership,omitempty"`

	// Returns a status message if the user owns the file.
	Sticky string `json:"sticky,omitempty"`
}

// Validate validates this auth access access item permissions
func (m *AuthAccessAccessItemPermissions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuthAccessAccessItemPermissions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthAccessAccessItemPermissions) UnmarshalBinary(b []byte) error {
	var res AuthAccessAccessItemPermissions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
