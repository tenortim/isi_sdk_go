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

// CloudJobs cloud jobs
// swagger:model CloudJobs
type CloudJobs struct {

	// jobs
	Jobs []*CloudJobExtended `json:"jobs"`

	// Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options).
	Resume string `json:"resume,omitempty"`
}

// Validate validates this cloud jobs
func (m *CloudJobs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJobs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudJobs) validateJobs(formats strfmt.Registry) error {

	if swag.IsZero(m.Jobs) { // not required
		return nil
	}

	for i := 0; i < len(m.Jobs); i++ {
		if swag.IsZero(m.Jobs[i]) { // not required
			continue
		}

		if m.Jobs[i] != nil {
			if err := m.Jobs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("jobs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudJobs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudJobs) UnmarshalBinary(b []byte) error {
	var res CloudJobs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
