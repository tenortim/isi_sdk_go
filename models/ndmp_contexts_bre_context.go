// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// NdmpContextsBreContext ndmp contexts bre context
// swagger:model NdmpContextsBreContext
type NdmpContextsBreContext struct {

	// Unique ID of NDMP BRE context
	BreContextID string `json:"bre_context_id,omitempty"`

	// Unique display id.
	ID string `json:"id,omitempty"`
}

// Validate validates this ndmp contexts bre context
func (m *NdmpContextsBreContext) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NdmpContextsBreContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NdmpContextsBreContext) UnmarshalBinary(b []byte) error {
	var res NdmpContextsBreContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
