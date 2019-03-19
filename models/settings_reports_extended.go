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

// SettingsReportsExtended settings reports extended
// swagger:model SettingsReportsExtended
type SettingsReportsExtended struct {

	// The directory on /ifs where manual or live reports will be placed.
	LiveDir string `json:"live_dir,omitempty"`

	// The number of manual reports to keep.
	// Minimum: 1
	LiveRetain int64 `json:"live_retain,omitempty"`

	// The isidate schedule used to generate reports.
	Schedule string `json:"schedule,omitempty"`

	// The directory on /ifs where schedule reports will be placed.
	ScheduledDir string `json:"scheduled_dir,omitempty"`

	// The number of scheduled reports to keep.
	// Minimum: 1
	ScheduledRetain int64 `json:"scheduled_retain,omitempty"`
}

// Validate validates this settings reports extended
func (m *SettingsReportsExtended) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLiveRetain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduledRetain(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SettingsReportsExtended) validateLiveRetain(formats strfmt.Registry) error {

	if swag.IsZero(m.LiveRetain) { // not required
		return nil
	}

	if err := validate.MinimumInt("live_retain", "body", int64(m.LiveRetain), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *SettingsReportsExtended) validateScheduledRetain(formats strfmt.Registry) error {

	if swag.IsZero(m.ScheduledRetain) { // not required
		return nil
	}

	if err := validate.MinimumInt("scheduled_retain", "body", int64(m.ScheduledRetain), 1, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SettingsReportsExtended) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SettingsReportsExtended) UnmarshalBinary(b []byte) error {
	var res SettingsReportsExtended
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
