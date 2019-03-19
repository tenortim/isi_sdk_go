// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DrivesDriveFirmwareUpdateItem Drive firmware update information.
// swagger:model DrivesDriveFirmwareUpdateItem
type DrivesDriveFirmwareUpdateItem struct {

	// Indicates whether this is a cluster wide drive firwmare update or not
	// Required: true
	ClusterWide *bool `json:"cluster_wide"`
}

// Validate validates this drives drive firmware update item
func (m *DrivesDriveFirmwareUpdateItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterWide(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DrivesDriveFirmwareUpdateItem) validateClusterWide(formats strfmt.Registry) error {

	if err := validate.Required("cluster_wide", "body", m.ClusterWide); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DrivesDriveFirmwareUpdateItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DrivesDriveFirmwareUpdateItem) UnmarshalBinary(b []byte) error {
	var res DrivesDriveFirmwareUpdateItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
