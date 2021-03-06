// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AntivirusServerExtended antivirus server extended
// swagger:model AntivirusServerExtended
type AntivirusServerExtended struct {
	AntivirusServer

	// Virus definitions on the server.
	Definitions string `json:"definitions,omitempty"`

	// A unique identifier for the server.
	ID string `json:"id,omitempty"`

	// The status of the server.
	// Enum: [active inactive]
	Status string `json:"status,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AntivirusServerExtended) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 AntivirusServer
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.AntivirusServer = aO0

	// AO1
	var dataAO1 struct {
		Definitions string `json:"definitions,omitempty"`

		ID string `json:"id,omitempty"`

		Status string `json:"status,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Definitions = dataAO1.Definitions

	m.ID = dataAO1.ID

	m.Status = dataAO1.Status

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AntivirusServerExtended) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.AntivirusServer)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Definitions string `json:"definitions,omitempty"`

		ID string `json:"id,omitempty"`

		Status string `json:"status,omitempty"`
	}

	dataAO1.Definitions = m.Definitions

	dataAO1.ID = m.ID

	dataAO1.Status = m.Status

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this antivirus server extended
func (m *AntivirusServerExtended) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AntivirusServer
	if err := m.AntivirusServer.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var antivirusServerExtendedTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		antivirusServerExtendedTypeStatusPropEnum = append(antivirusServerExtendedTypeStatusPropEnum, v)
	}
}

// property enum
func (m *AntivirusServerExtended) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, antivirusServerExtendedTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AntivirusServerExtended) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AntivirusServerExtended) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AntivirusServerExtended) UnmarshalBinary(b []byte) error {
	var res AntivirusServerExtended
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
