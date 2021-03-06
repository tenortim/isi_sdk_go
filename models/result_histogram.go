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

// ResultHistogram result histogram
// swagger:model ResultHistogram
type ResultHistogram struct {

	// Access time enabled.
	// Required: true
	AtimeEnabled *bool `json:"atime_enabled"`

	// User attribute count.
	// Required: true
	AttributeCount *int64 `json:"attribute_count"`

	// Unix Epoch time of start of results collection job.
	// Required: true
	BeginTime *int64 `json:"begin_time"`

	// Histogram data of specified file count parameter.
	// Required: true
	Histogram []*ResultHistogramHistogramItem `json:"histogram"`
}

// Validate validates this result histogram
func (m *ResultHistogram) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAtimeEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttributeCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBeginTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHistogram(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResultHistogram) validateAtimeEnabled(formats strfmt.Registry) error {

	if err := validate.Required("atime_enabled", "body", m.AtimeEnabled); err != nil {
		return err
	}

	return nil
}

func (m *ResultHistogram) validateAttributeCount(formats strfmt.Registry) error {

	if err := validate.Required("attribute_count", "body", m.AttributeCount); err != nil {
		return err
	}

	return nil
}

func (m *ResultHistogram) validateBeginTime(formats strfmt.Registry) error {

	if err := validate.Required("begin_time", "body", m.BeginTime); err != nil {
		return err
	}

	return nil
}

func (m *ResultHistogram) validateHistogram(formats strfmt.Registry) error {

	if err := validate.Required("histogram", "body", m.Histogram); err != nil {
		return err
	}

	for i := 0; i < len(m.Histogram); i++ {
		if swag.IsZero(m.Histogram[i]) { // not required
			continue
		}

		if m.Histogram[i] != nil {
			if err := m.Histogram[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("histogram" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResultHistogram) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResultHistogram) UnmarshalBinary(b []byte) error {
	var res ResultHistogram
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
