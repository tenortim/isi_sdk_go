// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HistogramStatBy histogram stat by
// swagger:model HistogramStatBy
type HistogramStatBy struct {

	// Unix Epoch time of start of results collection job.
	// Required: true
	BeginTime *int64 `json:"begin_time"`

	// Histogram breakout data according to breakout parameter.
	// Required: true
	Breakouts []*HistogramStatByBreakout `json:"breakouts"`

	// Total length of the result list.
	// Required: true
	ResultLength *int64 `json:"result_length"`
}

// Validate validates this histogram stat by
func (m *HistogramStatBy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBeginTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBreakouts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResultLength(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HistogramStatBy) validateBeginTime(formats strfmt.Registry) error {

	if err := validate.Required("begin_time", "body", m.BeginTime); err != nil {
		return err
	}

	return nil
}

func (m *HistogramStatBy) validateBreakouts(formats strfmt.Registry) error {

	if err := validate.Required("breakouts", "body", m.Breakouts); err != nil {
		return err
	}

	for i := 0; i < len(m.Breakouts); i++ {
		if swag.IsZero(m.Breakouts[i]) { // not required
			continue
		}

		if m.Breakouts[i] != nil {
			if err := m.Breakouts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("breakouts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HistogramStatBy) validateResultLength(formats strfmt.Registry) error {

	if err := validate.Required("result_length", "body", m.ResultLength); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HistogramStatBy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HistogramStatBy) UnmarshalBinary(b []byte) error {
	var res HistogramStatBy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
