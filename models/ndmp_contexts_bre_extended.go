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

// NdmpContextsBreExtended ndmp contexts bre extended
// swagger:model NdmpContextsBreExtended
type NdmpContextsBreExtended struct {

	// contexts
	Contexts []*NdmpContextsBreContext `json:"contexts"`

	// Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options).
	Resume string `json:"resume,omitempty"`

	// Total number of items available.
	Total int64 `json:"total,omitempty"`
}

// Validate validates this ndmp contexts bre extended
func (m *NdmpContextsBreExtended) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContexts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NdmpContextsBreExtended) validateContexts(formats strfmt.Registry) error {

	if swag.IsZero(m.Contexts) { // not required
		return nil
	}

	for i := 0; i < len(m.Contexts); i++ {
		if swag.IsZero(m.Contexts[i]) { // not required
			continue
		}

		if m.Contexts[i] != nil {
			if err := m.Contexts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contexts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NdmpContextsBreExtended) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NdmpContextsBreExtended) UnmarshalBinary(b []byte) error {
	var res NdmpContextsBreExtended
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}