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

// EventEventlists event eventlists
// swagger:model EventEventlists
type EventEventlists struct {

	// eventlists
	Eventlists []*EventEventlist `json:"eventlists"`
}

// Validate validates this event eventlists
func (m *EventEventlists) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventlists(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventEventlists) validateEventlists(formats strfmt.Registry) error {

	if swag.IsZero(m.Eventlists) { // not required
		return nil
	}

	for i := 0; i < len(m.Eventlists); i++ {
		if swag.IsZero(m.Eventlists[i]) { // not required
			continue
		}

		if m.Eventlists[i] != nil {
			if err := m.Eventlists[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("eventlists" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EventEventlists) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventEventlists) UnmarshalBinary(b []byte) error {
	var res EventEventlists
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
