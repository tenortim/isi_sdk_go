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

// HardwareTapesDevices hardware tapes devices
// swagger:model HardwareTapesDevices
type HardwareTapesDevices struct {

	// media changers
	MediaChangers []*HardwareTapesDevicesMediaChanger `json:"media_changers"`

	// tapes
	Tapes []*HardwareTapesDevicesMediaChanger `json:"tapes"`
}

// Validate validates this hardware tapes devices
func (m *HardwareTapesDevices) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMediaChangers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTapes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HardwareTapesDevices) validateMediaChangers(formats strfmt.Registry) error {

	if swag.IsZero(m.MediaChangers) { // not required
		return nil
	}

	for i := 0; i < len(m.MediaChangers); i++ {
		if swag.IsZero(m.MediaChangers[i]) { // not required
			continue
		}

		if m.MediaChangers[i] != nil {
			if err := m.MediaChangers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("media_changers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HardwareTapesDevices) validateTapes(formats strfmt.Registry) error {

	if swag.IsZero(m.Tapes) { // not required
		return nil
	}

	for i := 0; i < len(m.Tapes); i++ {
		if swag.IsZero(m.Tapes[i]) { // not required
			continue
		}

		if m.Tapes[i] != nil {
			if err := m.Tapes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tapes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HardwareTapesDevices) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HardwareTapesDevices) UnmarshalBinary(b []byte) error {
	var res HardwareTapesDevices
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}