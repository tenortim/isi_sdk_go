// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkExternalSettings network external settings
// swagger:model NetworkExternalSettings
type NetworkExternalSettings struct {

	// Default client-side DNS settings for non-multitenancy aware programs
	// Required: true
	DefaultGroupnet *string `json:"default_groupnet"`

	// Enable or disable Source Based Routing (Defaults to false)
	// Required: true
	Sbr *bool `json:"sbr"`

	// Delay in seconds for IP rebalance.
	// Required: true
	// Maximum: 10
	// Minimum: 0
	ScRebalanceDelay *int64 `json:"sc_rebalance_delay"`

	// List of client TCP ports.
	// Required: true
	TCPPorts []*int64 `json:"tcp_ports"`
}

// Validate validates this network external settings
func (m *NetworkExternalSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultGroupnet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScRebalanceDelay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTCPPorts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkExternalSettings) validateDefaultGroupnet(formats strfmt.Registry) error {

	if err := validate.Required("default_groupnet", "body", m.DefaultGroupnet); err != nil {
		return err
	}

	return nil
}

func (m *NetworkExternalSettings) validateSbr(formats strfmt.Registry) error {

	if err := validate.Required("sbr", "body", m.Sbr); err != nil {
		return err
	}

	return nil
}

func (m *NetworkExternalSettings) validateScRebalanceDelay(formats strfmt.Registry) error {

	if err := validate.Required("sc_rebalance_delay", "body", m.ScRebalanceDelay); err != nil {
		return err
	}

	if err := validate.MinimumInt("sc_rebalance_delay", "body", int64(*m.ScRebalanceDelay), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("sc_rebalance_delay", "body", int64(*m.ScRebalanceDelay), 10, false); err != nil {
		return err
	}

	return nil
}

func (m *NetworkExternalSettings) validateTCPPorts(formats strfmt.Registry) error {

	if err := validate.Required("tcp_ports", "body", m.TCPPorts); err != nil {
		return err
	}

	for i := 0; i < len(m.TCPPorts); i++ {
		if swag.IsZero(m.TCPPorts[i]) { // not required
			continue
		}

		if err := validate.MinimumInt("tcp_ports"+"."+strconv.Itoa(i), "body", int64(*m.TCPPorts[i]), 0, false); err != nil {
			return err
		}

		if err := validate.MaximumInt("tcp_ports"+"."+strconv.Itoa(i), "body", int64(*m.TCPPorts[i]), 65535, false); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkExternalSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkExternalSettings) UnmarshalBinary(b []byte) error {
	var res NetworkExternalSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
