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

// SyncPolicySourceNetwork sync policy source network
// swagger:model SyncPolicySourceNetwork
type SyncPolicySourceNetwork struct {

	// The pool to restrict replication policies to.
	// Required: true
	Pool *string `json:"pool"`

	// The subnet to restrict replication policies to.
	// Required: true
	Subnet *string `json:"subnet"`
}

// Validate validates this sync policy source network
func (m *SyncPolicySourceNetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePool(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SyncPolicySourceNetwork) validatePool(formats strfmt.Registry) error {

	if err := validate.Required("pool", "body", m.Pool); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicySourceNetwork) validateSubnet(formats strfmt.Registry) error {

	if err := validate.Required("subnet", "body", m.Subnet); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SyncPolicySourceNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SyncPolicySourceNetwork) UnmarshalBinary(b []byte) error {
	var res SyncPolicySourceNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
