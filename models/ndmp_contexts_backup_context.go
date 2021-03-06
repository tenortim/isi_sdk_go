// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// NdmpContextsBackupContext ndmp contexts backup context
// swagger:model NdmpContextsBackupContext
type NdmpContextsBackupContext struct {

	// Context ID
	ContextID string `json:"context_id,omitempty"`

	// Unique display id.
	ID string `json:"id,omitempty"`

	// The directory of the backup. This is not applicable to restore contexts.
	Path string `json:"path,omitempty"`

	// Snapshot ID reserved for the context. This is not applicable to restore contexts.
	Snapid int64 `json:"snapid,omitempty"`

	// Context creation time
	StartTime int64 `json:"start_time,omitempty"`

	// Whether the context is active.
	Status string `json:"status,omitempty"`

	// The number of sessions in the context
	TotalSessions int64 `json:"total_sessions,omitempty"`

	// Type of context; It can be bre, backup, and restore
	Type string `json:"type,omitempty"`
}

// Validate validates this ndmp contexts backup context
func (m *NdmpContextsBackupContext) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NdmpContextsBackupContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NdmpContextsBackupContext) UnmarshalBinary(b []byte) error {
	var res NdmpContextsBackupContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
