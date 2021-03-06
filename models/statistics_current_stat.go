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

// StatisticsCurrentStat statistics current stat
// swagger:model StatisticsCurrentStat
type StatisticsCurrentStat struct {

	// Devid of node of statistic or 0 for cluster scoped statistics.
	// Required: true
	Devid *int64 `json:"devid"`

	// Key specific error string, if applicable.
	Error string `json:"error,omitempty"`

	// Key specific error number, if applicable.
	ErrorCode int64 `json:"error_code,omitempty"`

	// Key name of statistic.
	// Required: true
	Key *string `json:"key"`

	// Unix Epoch time in seconds that statistic was collected.
	// Required: true
	Time *int64 `json:"time"`

	// Key dependent value.
	Value string `json:"value,omitempty"`
}

// Validate validates this statistics current stat
func (m *StatisticsCurrentStat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatisticsCurrentStat) validateDevid(formats strfmt.Registry) error {

	if err := validate.Required("devid", "body", m.Devid); err != nil {
		return err
	}

	return nil
}

func (m *StatisticsCurrentStat) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

func (m *StatisticsCurrentStat) validateTime(formats strfmt.Registry) error {

	if err := validate.Required("time", "body", m.Time); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StatisticsCurrentStat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatisticsCurrentStat) UnmarshalBinary(b []byte) error {
	var res StatisticsCurrentStat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
