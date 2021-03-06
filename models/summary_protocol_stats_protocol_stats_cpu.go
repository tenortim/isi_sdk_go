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

// SummaryProtocolStatsProtocolStatsCPU summary protocol stats protocol stats Cpu
// swagger:model SummaryProtocolStatsProtocol-StatsCpu
type SummaryProtocolStatsProtocolStatsCPU struct {

	// Percentage of CPU idle time
	// Required: true
	IDLE *float64 `json:"idle"`

	// Percentage of CPU consumed by the system
	// Required: true
	System *float64 `json:"system"`

	// Percentage of CPU consumed by user activities
	// Required: true
	User *float64 `json:"user"`
}

// Validate validates this summary protocol stats protocol stats Cpu
func (m *SummaryProtocolStatsProtocolStatsCPU) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIDLE(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SummaryProtocolStatsProtocolStatsCPU) validateIDLE(formats strfmt.Registry) error {

	if err := validate.Required("idle", "body", m.IDLE); err != nil {
		return err
	}

	return nil
}

func (m *SummaryProtocolStatsProtocolStatsCPU) validateSystem(formats strfmt.Registry) error {

	if err := validate.Required("system", "body", m.System); err != nil {
		return err
	}

	return nil
}

func (m *SummaryProtocolStatsProtocolStatsCPU) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("user", "body", m.User); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SummaryProtocolStatsProtocolStatsCPU) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SummaryProtocolStatsProtocolStatsCPU) UnmarshalBinary(b []byte) error {
	var res SummaryProtocolStatsProtocolStatsCPU
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
