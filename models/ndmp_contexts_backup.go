// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NdmpContextsBackup View a NDMP Context
// swagger:model NdmpContextsBackup
type NdmpContextsBackup struct {

	// Context expiration time
	ContextExpirationTime int64 `json:"context_expiration_time,omitempty"`

	// Context ID
	ContextID string `json:"context_id,omitempty"`

	// Unique display id.
	ID string `json:"id,omitempty"`

	// The directory of the backup. This is not applicable to restore contexts.
	Path string `json:"path,omitempty"`

	// sessions
	Sessions []*NdmpContextsBackupSession `json:"sessions"`

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

// Validate validates this ndmp contexts backup
func (m *NdmpContextsBackup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSessions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NdmpContextsBackup) validateSessions(formats strfmt.Registry) error {

	if swag.IsZero(m.Sessions) { // not required
		return nil
	}

	for i := 0; i < len(m.Sessions); i++ {
		if swag.IsZero(m.Sessions[i]) { // not required
			continue
		}

		if m.Sessions[i] != nil {
			if err := m.Sessions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sessions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NdmpContextsBackup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NdmpContextsBackup) UnmarshalBinary(b []byte) error {
	var res NdmpContextsBackup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}