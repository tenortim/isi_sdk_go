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

// CloudPools cloud pools
// swagger:model CloudPools
type CloudPools struct {

	// pools
	Pools []*CloudPoolExtended `json:"pools"`

	// Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options).
	Resume string `json:"resume,omitempty"`
}

// Validate validates this cloud pools
func (m *CloudPools) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePools(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudPools) validatePools(formats strfmt.Registry) error {

	if swag.IsZero(m.Pools) { // not required
		return nil
	}

	for i := 0; i < len(m.Pools); i++ {
		if swag.IsZero(m.Pools[i]) { // not required
			continue
		}

		if m.Pools[i] != nil {
			if err := m.Pools[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pools" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudPools) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudPools) UnmarshalBinary(b []byte) error {
	var res CloudPools
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}