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

// JobStatisticsJobNodeMemory job statistics job node memory
// swagger:model JobStatisticsJobNodeMemory
type JobStatisticsJobNodeMemory struct {

	// physical
	// Required: true
	Physical *JobStatisticsJobNodeMemoryPhysical `json:"physical"`

	// virtual
	// Required: true
	Virtual *JobStatisticsJobNodeMemoryVirtual `json:"virtual"`
}

// Validate validates this job statistics job node memory
func (m *JobStatisticsJobNodeMemory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePhysical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtual(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobStatisticsJobNodeMemory) validatePhysical(formats strfmt.Registry) error {

	if err := validate.Required("physical", "body", m.Physical); err != nil {
		return err
	}

	if m.Physical != nil {
		if err := m.Physical.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("physical")
			}
			return err
		}
	}

	return nil
}

func (m *JobStatisticsJobNodeMemory) validateVirtual(formats strfmt.Registry) error {

	if err := validate.Required("virtual", "body", m.Virtual); err != nil {
		return err
	}

	if m.Virtual != nil {
		if err := m.Virtual.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtual")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobStatisticsJobNodeMemory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobStatisticsJobNodeMemory) UnmarshalBinary(b []byte) error {
	var res JobStatisticsJobNodeMemory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
