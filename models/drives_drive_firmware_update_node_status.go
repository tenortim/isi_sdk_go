// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DrivesDriveFirmwareUpdateNodeStatus drives drive firmware update node status
// swagger:model DrivesDriveFirmwareUpdateNodeStatus
type DrivesDriveFirmwareUpdateNodeStatus struct {

	// The number of drives that did not successfully complete firmware updates update on the node.
	Failed int64 `json:"failed,omitempty"`

	// Time when drive firmware update finished on node.
	FinishTime string `json:"finish_time,omitempty"`

	// Number of drives remaining to be updated on node.
	Remaining int64 `json:"remaining,omitempty"`

	// Time when drive firmware update started on node.
	StartTime string `json:"start_time,omitempty"`

	// Overall status of this nodes firmware updates.
	Status string `json:"status,omitempty"`

	// The number of drives that completed firmware updates on the node.
	Updated int64 `json:"updated,omitempty"`
}

// Validate validates this drives drive firmware update node status
func (m *DrivesDriveFirmwareUpdateNodeStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DrivesDriveFirmwareUpdateNodeStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DrivesDriveFirmwareUpdateNodeStatus) UnmarshalBinary(b []byte) error {
	var res DrivesDriveFirmwareUpdateNodeStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}