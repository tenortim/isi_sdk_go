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

// SummaryProtocolStatsProtocolStatsNetworkIn summary protocol stats protocol stats network in
// swagger:model SummaryProtocolStatsProtocol-StatsNetworkIn
type SummaryProtocolStatsProtocolStatsNetworkIn struct {

	// Network input errors per-second
	// Required: true
	ErrorsPerSec *float64 `json:"errors_per_sec"`

	// Network input megabytes per-second
	// Required: true
	MegabytesPerSec *float64 `json:"megabytes_per_sec"`

	// Network input packets per-second
	// Required: true
	PacketsPerSec *float64 `json:"packets_per_sec"`
}

// Validate validates this summary protocol stats protocol stats network in
func (m *SummaryProtocolStatsProtocolStatsNetworkIn) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrorsPerSec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMegabytesPerSec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePacketsPerSec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SummaryProtocolStatsProtocolStatsNetworkIn) validateErrorsPerSec(formats strfmt.Registry) error {

	if err := validate.Required("errors_per_sec", "body", m.ErrorsPerSec); err != nil {
		return err
	}

	return nil
}

func (m *SummaryProtocolStatsProtocolStatsNetworkIn) validateMegabytesPerSec(formats strfmt.Registry) error {

	if err := validate.Required("megabytes_per_sec", "body", m.MegabytesPerSec); err != nil {
		return err
	}

	return nil
}

func (m *SummaryProtocolStatsProtocolStatsNetworkIn) validatePacketsPerSec(formats strfmt.Registry) error {

	if err := validate.Required("packets_per_sec", "body", m.PacketsPerSec); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SummaryProtocolStatsProtocolStatsNetworkIn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SummaryProtocolStatsProtocolStatsNetworkIn) UnmarshalBinary(b []byte) error {
	var res SummaryProtocolStatsProtocolStatsNetworkIn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
