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

// JobEvents job events
// swagger:model JobEvents
type JobEvents struct {

	// events
	Events []*JobEvent `json:"events"`

	// Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options).
	Resume string `json:"resume,omitempty"`

	// Total number of items available.
	Total int64 `json:"total,omitempty"`
}

// Validate validates this job events
func (m *JobEvents) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobEvents) validateEvents(formats strfmt.Registry) error {

	if swag.IsZero(m.Events) { // not required
		return nil
	}

	for i := 0; i < len(m.Events); i++ {
		if swag.IsZero(m.Events[i]) { // not required
			continue
		}

		if m.Events[i] != nil {
			if err := m.Events[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobEvents) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobEvents) UnmarshalBinary(b []byte) error {
	var res JobEvents
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
