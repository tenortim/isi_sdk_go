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

// JobJobSnaprevertParams job job snaprevert params
// swagger:model JobJobSnaprevertParams
type JobJobSnaprevertParams struct {

	// Snapshot to revert.
	// Required: true
	// Minimum: 1
	Snapid *int64 `json:"snapid"`
}

// Validate validates this job job snaprevert params
func (m *JobJobSnaprevertParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSnapid(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobJobSnaprevertParams) validateSnapid(formats strfmt.Registry) error {

	if err := validate.Required("snapid", "body", m.Snapid); err != nil {
		return err
	}

	if err := validate.MinimumInt("snapid", "body", int64(*m.Snapid), 1, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobJobSnaprevertParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobJobSnaprevertParams) UnmarshalBinary(b []byte) error {
	var res JobJobSnaprevertParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}