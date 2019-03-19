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

// ClusterTimeExtendedExtended cluster time extended extended
// swagger:model ClusterTimeExtendedExtended
type ClusterTimeExtendedExtended struct {

	// The current time on the cluster as a UNIX epoch (seconds since 1/1/1970), as reported by this node.
	// Maximum: 9.223372036854776e+18
	// Minimum: -9.223372036854776e+18
	Time *int64 `json:"time,omitempty"`
}

// Validate validates this cluster time extended extended
func (m *ClusterTimeExtendedExtended) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterTimeExtendedExtended) validateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.MinimumInt("time", "body", int64(*m.Time), -9.223372036854776e+18, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("time", "body", int64(*m.Time), 9.223372036854776e+18, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterTimeExtendedExtended) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterTimeExtendedExtended) UnmarshalBinary(b []byte) error {
	var res ClusterTimeExtendedExtended
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}