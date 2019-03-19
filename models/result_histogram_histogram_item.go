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

// ResultHistogramHistogramItem result histogram histogram item
// swagger:model ResultHistogramHistogramItem
type ResultHistogramHistogramItem struct {

	// Bucket for holding file counts within a range.
	// Required: true
	Bucket *int64 `json:"bucket"`

	// Number of files within the bucket.
	// Required: true
	Value *int64 `json:"value"`
}

// Validate validates this result histogram histogram item
func (m *ResultHistogramHistogramItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBucket(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResultHistogramHistogramItem) validateBucket(formats strfmt.Registry) error {

	if err := validate.Required("bucket", "body", m.Bucket); err != nil {
		return err
	}

	return nil
}

func (m *ResultHistogramHistogramItem) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResultHistogramHistogramItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResultHistogramHistogramItem) UnmarshalBinary(b []byte) error {
	var res ResultHistogramHistogramItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
