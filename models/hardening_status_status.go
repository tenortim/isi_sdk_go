// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// HardeningStatusStatus hardening status status
// swagger:model HardeningStatusStatus
type HardeningStatusStatus struct {

	// Status text containing cluster-level and nodewise hardening status. Also includes hardening profile if hardening is enabled on at least one node.
	Message string `json:"message,omitempty"`
}

// Validate validates this hardening status status
func (m *HardeningStatusStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HardeningStatusStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HardeningStatusStatus) UnmarshalBinary(b []byte) error {
	var res HardeningStatusStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
