// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ClusterNodeStateSmartfail cluster node state smartfail
// swagger:model ClusterNodeStateSmartfail
type ClusterNodeStateSmartfail struct {

	// This node is smartfailed (soft_devs).
	Smartfailed bool `json:"smartfailed,omitempty"`
}

// Validate validates this cluster node state smartfail
func (m *ClusterNodeStateSmartfail) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterNodeStateSmartfail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterNodeStateSmartfail) UnmarshalBinary(b []byte) error {
	var res ClusterNodeStateSmartfail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
