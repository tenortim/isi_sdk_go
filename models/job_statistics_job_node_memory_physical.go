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

// JobStatisticsJobNodeMemoryPhysical job statistics job node memory physical
// swagger:model JobStatisticsJobNodeMemoryPhysical
type JobStatisticsJobNodeMemoryPhysical struct {

	// The average physical memory utilization of the job on this node, in KB.
	Average float64 `json:"average,omitempty"`

	// The current physical memory utilization of the job on this node, in KB.
	// Required: true
	Current *float64 `json:"current"`

	// The maximum physical memory utilization of the job on this node, in KB.
	// Required: true
	Maximum *float64 `json:"maximum"`

	// The minimum physical memory utilization of the job on this node, in KB.
	// Required: true
	Minimum *float64 `json:"minimum"`
}

// Validate validates this job statistics job node memory physical
func (m *JobStatisticsJobNodeMemoryPhysical) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaximum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinimum(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobStatisticsJobNodeMemoryPhysical) validateCurrent(formats strfmt.Registry) error {

	if err := validate.Required("current", "body", m.Current); err != nil {
		return err
	}

	return nil
}

func (m *JobStatisticsJobNodeMemoryPhysical) validateMaximum(formats strfmt.Registry) error {

	if err := validate.Required("maximum", "body", m.Maximum); err != nil {
		return err
	}

	return nil
}

func (m *JobStatisticsJobNodeMemoryPhysical) validateMinimum(formats strfmt.Registry) error {

	if err := validate.Required("minimum", "body", m.Minimum); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobStatisticsJobNodeMemoryPhysical) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobStatisticsJobNodeMemoryPhysical) UnmarshalBinary(b []byte) error {
	var res JobStatisticsJobNodeMemoryPhysical
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
