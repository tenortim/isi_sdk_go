// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SummaryProtocolStats summary protocol stats
// swagger:model SummaryProtocolStats
type SummaryProtocolStats struct {

	// protocol stats
	ProtocolStats *SummaryProtocolStatsProtocolStats `json:"protocol-stats,omitempty"`
}

// Validate validates this summary protocol stats
func (m *SummaryProtocolStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProtocolStats(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SummaryProtocolStats) validateProtocolStats(formats strfmt.Registry) error {

	if swag.IsZero(m.ProtocolStats) { // not required
		return nil
	}

	if m.ProtocolStats != nil {
		if err := m.ProtocolStats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protocol-stats")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SummaryProtocolStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SummaryProtocolStats) UnmarshalBinary(b []byte) error {
	var res SummaryProtocolStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
