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

// ClusterConfigDevice cluster config device
// swagger:model ClusterConfigDevice
type ClusterConfigDevice struct {

	// Device ID.
	// Required: true
	Devid *int64 `json:"devid"`

	// Device GUID.
	// Required: true
	GUID *string `json:"guid"`

	// If true, this node is online and communicating with the local node and every other node with the is_up property normally.  If false, this node is not currently communicating with the local node or other nodes with the is_up property.  It may be shut down, rebooting, disconnected from the backend network, or connected only to other nodes.
	// Required: true
	IsUp *bool `json:"is_up"`

	// Device logical node number.
	// Required: true
	Lnn *int64 `json:"lnn"`
}

// Validate validates this cluster config device
func (m *ClusterConfigDevice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsUp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLnn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterConfigDevice) validateDevid(formats strfmt.Registry) error {

	if err := validate.Required("devid", "body", m.Devid); err != nil {
		return err
	}

	return nil
}

func (m *ClusterConfigDevice) validateGUID(formats strfmt.Registry) error {

	if err := validate.Required("guid", "body", m.GUID); err != nil {
		return err
	}

	return nil
}

func (m *ClusterConfigDevice) validateIsUp(formats strfmt.Registry) error {

	if err := validate.Required("is_up", "body", m.IsUp); err != nil {
		return err
	}

	return nil
}

func (m *ClusterConfigDevice) validateLnn(formats strfmt.Registry) error {

	if err := validate.Required("lnn", "body", m.Lnn); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterConfigDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterConfigDevice) UnmarshalBinary(b []byte) error {
	var res ClusterConfigDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}