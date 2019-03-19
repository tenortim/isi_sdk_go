// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// EventSettingsSettingsMaintenance event settings settings maintenance
// swagger:model EventSettingsSettingsMaintenance
type EventSettingsSettingsMaintenance struct {

	// Duration of maintenance period in seconds.
	Duration int64 `json:"duration,omitempty"`

	// Start of maintenance period.
	Start int64 `json:"start,omitempty"`
}

// Validate validates this event settings settings maintenance
func (m *EventSettingsSettingsMaintenance) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventSettingsSettingsMaintenance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventSettingsSettingsMaintenance) UnmarshalBinary(b []byte) error {
	var res EventSettingsSettingsMaintenance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
