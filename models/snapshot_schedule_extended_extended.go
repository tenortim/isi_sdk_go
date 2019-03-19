// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SnapshotScheduleExtendedExtended snapshot schedule extended extended
// swagger:model SnapshotScheduleExtendedExtended
type SnapshotScheduleExtendedExtended struct {

	// Alias name to create for each snapshot.
	Alias string `json:"alias,omitempty"`

	// Time in seconds added to creation time to construction expiration time.
	Duration int64 `json:"duration,omitempty"`

	// The system ID given to the schedule.
	ID int64 `json:"id,omitempty"`

	// The schedule name.
	Name string `json:"name,omitempty"`

	// Unix Epoch time of next snapshot to be created.
	NextRun int64 `json:"next_run,omitempty"`

	// Formatted name (see pattern) of next snapshot to be created.
	NextSnapshot string `json:"next_snapshot,omitempty"`

	// The /ifs path snapshotted.
	Path string `json:"path,omitempty"`

	// Pattern expanded with strftime to create snapshot names.
	Pattern string `json:"pattern,omitempty"`

	// The isidate compatible natural language description of the schedule.
	Schedule string `json:"schedule,omitempty"`
}

// Validate validates this snapshot schedule extended extended
func (m *SnapshotScheduleExtendedExtended) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotScheduleExtendedExtended) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotScheduleExtendedExtended) UnmarshalBinary(b []byte) error {
	var res SnapshotScheduleExtendedExtended
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}