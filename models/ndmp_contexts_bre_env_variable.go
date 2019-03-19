// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// NdmpContextsBreEnvVariable ndmp contexts bre env variable
// swagger:model NdmpContextsBreEnvVariable
type NdmpContextsBreEnvVariable struct {

	// Environment variable name
	Name string `json:"name,omitempty"`

	// Environment variable value
	Value string `json:"value,omitempty"`
}

// Validate validates this ndmp contexts bre env variable
func (m *NdmpContextsBreEnvVariable) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NdmpContextsBreEnvVariable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NdmpContextsBreEnvVariable) UnmarshalBinary(b []byte) error {
	var res NdmpContextsBreEnvVariable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}