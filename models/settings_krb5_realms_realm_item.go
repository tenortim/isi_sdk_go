// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SettingsKrb5RealmsRealmItem settings krb5 realms realm item
// swagger:model SettingsKrb5RealmsRealmItem
type SettingsKrb5RealmsRealmItem struct {

	// Specifies the administrative server hostname.
	AdminServer string `json:"admin_server,omitempty"`

	// Specifies the default domain mapped to the realm.
	DefaultDomain string `json:"default_domain,omitempty"`

	// ID of realm
	ID string `json:"id,omitempty"`

	// If true, indicates that the realm is the default.
	IsDefaultRealm bool `json:"is_default_realm,omitempty"`

	// If true, indicates that the realm is joined.
	IsJoined bool `json:"is_joined,omitempty"`

	// Specifies the list of KDC hostnames.
	Kdc []string `json:"kdc"`

	// Specifies the name of the realm.
	Realm string `json:"realm,omitempty"`
}

// Validate validates this settings krb5 realms realm item
func (m *SettingsKrb5RealmsRealmItem) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SettingsKrb5RealmsRealmItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SettingsKrb5RealmsRealmItem) UnmarshalBinary(b []byte) error {
	var res SettingsKrb5RealmsRealmItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
