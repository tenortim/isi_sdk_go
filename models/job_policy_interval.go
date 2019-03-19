// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// JobPolicyInterval job policy interval
// swagger:model JobPolicyInterval
type JobPolicyInterval struct {

	// Beginning time for the corresponding impact, in the format 'WWWW HH:MM', where 'WWWW' is the full English name of the day of the week, 'HH' is the hour (00-23), and 'MM' is the minute (00-59).
	// Required: true
	Begin *string `json:"begin"`

	// Ending time for the corresponding impact, in the format 'WWWW HH:MM', where 'WWWW' is the full English name of the day of the week, 'HH' is the hour (00-23), and 'MM' is the minute (00-59).
	// Required: true
	End *string `json:"end"`

	// Impact for the corresponding time span.
	// Required: true
	// Enum: [Low Medium High Paused]
	Impact *string `json:"impact"`
}

// Validate validates this job policy interval
func (m *JobPolicyInterval) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBegin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImpact(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobPolicyInterval) validateBegin(formats strfmt.Registry) error {

	if err := validate.Required("begin", "body", m.Begin); err != nil {
		return err
	}

	return nil
}

func (m *JobPolicyInterval) validateEnd(formats strfmt.Registry) error {

	if err := validate.Required("end", "body", m.End); err != nil {
		return err
	}

	return nil
}

var jobPolicyIntervalTypeImpactPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Low","Medium","High","Paused"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		jobPolicyIntervalTypeImpactPropEnum = append(jobPolicyIntervalTypeImpactPropEnum, v)
	}
}

const (

	// JobPolicyIntervalImpactLow captures enum value "Low"
	JobPolicyIntervalImpactLow string = "Low"

	// JobPolicyIntervalImpactMedium captures enum value "Medium"
	JobPolicyIntervalImpactMedium string = "Medium"

	// JobPolicyIntervalImpactHigh captures enum value "High"
	JobPolicyIntervalImpactHigh string = "High"

	// JobPolicyIntervalImpactPaused captures enum value "Paused"
	JobPolicyIntervalImpactPaused string = "Paused"
)

// prop value enum
func (m *JobPolicyInterval) validateImpactEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, jobPolicyIntervalTypeImpactPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *JobPolicyInterval) validateImpact(formats strfmt.Registry) error {

	if err := validate.Required("impact", "body", m.Impact); err != nil {
		return err
	}

	// value enum
	if err := m.validateImpactEnum("impact", "body", *m.Impact); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobPolicyInterval) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobPolicyInterval) UnmarshalBinary(b []byte) error {
	var res JobPolicyInterval
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}