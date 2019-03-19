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

// SnapshotScheduleCreateParams snapshot schedule create params
// swagger:model SnapshotScheduleCreateParams
type SnapshotScheduleCreateParams struct {
	SnapshotSchedule

	// The schedule name.
	// Required: true
	Name *string `json:"name"`

	// The /ifs path snapshotted.
	// Required: true
	Path *string `json:"path"`

	// Pattern expanded with strftime to create snapshot names.
	// Required: true
	Pattern *string `json:"pattern"`

	// The isidate compatible natural language description of the schedule.
	// Required: true
	Schedule *string `json:"schedule"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SnapshotScheduleCreateParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SnapshotSchedule
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SnapshotSchedule = aO0

	// AO1
	var dataAO1 struct {
		Name *string `json:"name"`

		Path *string `json:"path"`

		Pattern *string `json:"pattern"`

		Schedule *string `json:"schedule"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Name = dataAO1.Name

	m.Path = dataAO1.Path

	m.Pattern = dataAO1.Pattern

	m.Schedule = dataAO1.Schedule

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SnapshotScheduleCreateParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.SnapshotSchedule)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Name *string `json:"name"`

		Path *string `json:"path"`

		Pattern *string `json:"pattern"`

		Schedule *string `json:"schedule"`
	}

	dataAO1.Name = m.Name

	dataAO1.Path = m.Path

	dataAO1.Pattern = m.Pattern

	dataAO1.Schedule = m.Schedule

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this snapshot schedule create params
func (m *SnapshotScheduleCreateParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SnapshotSchedule
	if err := m.SnapshotSchedule.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePattern(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotScheduleCreateParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotScheduleCreateParams) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotScheduleCreateParams) validatePattern(formats strfmt.Registry) error {

	if err := validate.Required("pattern", "body", m.Pattern); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotScheduleCreateParams) validateSchedule(formats strfmt.Registry) error {

	if err := validate.Required("schedule", "body", m.Schedule); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotScheduleCreateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotScheduleCreateParams) UnmarshalBinary(b []byte) error {
	var res SnapshotScheduleCreateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
