// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ZonesSummary zones summary
// swagger:model ZonesSummary
type ZonesSummary struct {

	// The summary of a collection of objects.
	Summary *ZonesSummarySummary `json:"summary,omitempty"`
}

// Validate validates this zones summary
func (m *ZonesSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSummary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ZonesSummary) validateSummary(formats strfmt.Registry) error {

	if swag.IsZero(m.Summary) { // not required
		return nil
	}

	if m.Summary != nil {
		if err := m.Summary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("summary")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ZonesSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ZonesSummary) UnmarshalBinary(b []byte) error {
	var res ZonesSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
