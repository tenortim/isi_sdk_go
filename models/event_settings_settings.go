// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// EventSettingsSettings event settings settings
// swagger:model EventSettingsSettings
type EventSettingsSettings struct {

	// Interval between heartbeat events. "daily", "weekly", or "monthly".
	HeartbeatInterval string `json:"heartbeat_interval,omitempty"`

	// Specifies start and duration of maintenance period during which no alerts will be sent for new eventgroups.
	Maintenance *EventSettingsSettingsMaintenance `json:"maintenance,omitempty"`

	// Retention period in days
	RetentionDays int64 `json:"retention_days,omitempty"`

	// Storage limit in megabytes per terabyte of available storage
	StorageLimit int64 `json:"storage_limit,omitempty"`
}

// Validate validates this event settings settings
func (m *EventSettingsSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaintenance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventSettingsSettings) validateMaintenance(formats strfmt.Registry) error {

	if swag.IsZero(m.Maintenance) { // not required
		return nil
	}

	if m.Maintenance != nil {
		if err := m.Maintenance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maintenance")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EventSettingsSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventSettingsSettings) UnmarshalBinary(b []byte) error {
	var res EventSettingsSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}