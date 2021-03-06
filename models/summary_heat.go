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

// SummaryHeat summary heat
// swagger:model SummaryHeat
type SummaryHeat struct {

	// heat
	Heat []*SummaryHeatHeatItem `json:"heat"`
}

// Validate validates this summary heat
func (m *SummaryHeat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeat(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SummaryHeat) validateHeat(formats strfmt.Registry) error {

	if swag.IsZero(m.Heat) { // not required
		return nil
	}

	for i := 0; i < len(m.Heat); i++ {
		if swag.IsZero(m.Heat[i]) { // not required
			continue
		}

		if m.Heat[i] != nil {
			if err := m.Heat[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("heat" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SummaryHeat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SummaryHeat) UnmarshalBinary(b []byte) error {
	var res SummaryHeat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
