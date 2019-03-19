// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// JobJobSmartpoolstreeParams job job smartpoolstree params
// swagger:model JobJobSmartpoolstreeParams
type JobJobSmartpoolstreeParams struct {

	// Skip processing of regular files.
	DirectoryOnly bool `json:"directory_only,omitempty"`

	// Calculate what would be done (dry run).
	Nop bool `json:"nop,omitempty"`

	// Apply policies but skip restriping.
	PolicyOnly bool `json:"policy_only,omitempty"`

	// Process children, recursively.
	Recurse bool `json:"recurse,omitempty"`
}

// Validate validates this job job smartpoolstree params
func (m *JobJobSmartpoolstreeParams) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JobJobSmartpoolstreeParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobJobSmartpoolstreeParams) UnmarshalBinary(b []byte) error {
	var res JobJobSmartpoolstreeParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}