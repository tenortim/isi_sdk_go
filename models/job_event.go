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

// JobEvent job event
// swagger:model JobEvent
type JobEvent struct {

	// Event flags.
	// Required: true
	Flags *string `json:"flags"`

	// fmt type
	FmtType string `json:"fmt_type,omitempty"`

	// Job event ID.
	// Required: true
	ID *string `json:"id"`

	// Job ID.
	// Required: true
	// Minimum: 1
	JobID *int64 `json:"job_id"`

	// Job Type.
	// Required: true
	JobType *string `json:"job_type"`

	// Event key name.
	// Required: true
	Key *string `json:"key"`

	// Job phase number at time of event.
	// Required: true
	Phase *int64 `json:"phase"`

	// raw type
	RawType string `json:"raw_type,omitempty"`

	// Time of event in Unix epoch seconds.
	// Required: true
	Time *int64 `json:"time"`

	// Event value.
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this job event
func (m *JobEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobEvent) validateFlags(formats strfmt.Registry) error {

	if err := validate.Required("flags", "body", m.Flags); err != nil {
		return err
	}

	return nil
}

func (m *JobEvent) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *JobEvent) validateJobID(formats strfmt.Registry) error {

	if err := validate.Required("job_id", "body", m.JobID); err != nil {
		return err
	}

	if err := validate.MinimumInt("job_id", "body", int64(*m.JobID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *JobEvent) validateJobType(formats strfmt.Registry) error {

	if err := validate.Required("job_type", "body", m.JobType); err != nil {
		return err
	}

	return nil
}

func (m *JobEvent) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

func (m *JobEvent) validatePhase(formats strfmt.Registry) error {

	if err := validate.Required("phase", "body", m.Phase); err != nil {
		return err
	}

	return nil
}

func (m *JobEvent) validateTime(formats strfmt.Registry) error {

	if err := validate.Required("time", "body", m.Time); err != nil {
		return err
	}

	return nil
}

func (m *JobEvent) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobEvent) UnmarshalBinary(b []byte) error {
	var res JobEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
