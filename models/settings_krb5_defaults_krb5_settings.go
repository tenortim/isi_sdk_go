// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SettingsKrb5DefaultsKrb5Settings settings krb5 defaults krb5 settings
// swagger:model SettingsKrb5DefaultsKrb5Settings
type SettingsKrb5DefaultsKrb5Settings struct {

	// If true, always attempts to preauthenticate to the domain controller.
	AlwaysSendPreauth bool `json:"always_send_preauth,omitempty"`

	// Specifies the realm for unqualified names.
	DefaultRealm string `json:"default_realm,omitempty"`

	// If true, find KDCs through the DNS.
	DNSLookupKdc bool `json:"dns_lookup_kdc,omitempty"`

	// If true, find realm names through the DNS.
	DNSLookupRealm bool `json:"dns_lookup_realm,omitempty"`
}

// Validate validates this settings krb5 defaults krb5 settings
func (m *SettingsKrb5DefaultsKrb5Settings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SettingsKrb5DefaultsKrb5Settings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SettingsKrb5DefaultsKrb5Settings) UnmarshalBinary(b []byte) error {
	var res SettingsKrb5DefaultsKrb5Settings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}