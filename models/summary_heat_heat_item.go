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

// SummaryHeatHeatItem summary heat heat item
// swagger:model SummaryHeatHeatItem
type SummaryHeatHeatItem struct {

	// The class of operation
	// Required: true
	ClassName *string `json:"class_name"`

	// The type of event
	// Required: true
	EventName *string `json:"event_name"`

	// The event type id
	EventType int64 `json:"event_type,omitempty"`

	// Logical inode (LIN)
	Lin string `json:"lin,omitempty"`

	// The node where this event occurred.
	Node int64 `json:"node,omitempty"`

	// Approximate operations per second for this lin.
	// Required: true
	OperationRate *float64 `json:"operation_rate"`

	// Canonical LIN path if known
	// Required: true
	Path *string `json:"path"`

	// Unix Epoch time in seconds of the request.
	// Required: true
	Time *int64 `json:"time"`
}

// Validate validates this summary heat heat item
func (m *SummaryHeatHeatItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClassName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
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

func (m *SummaryHeatHeatItem) validateClassName(formats strfmt.Registry) error {

	if err := validate.Required("class_name", "body", m.ClassName); err != nil {
		return err
	}

	return nil
}

func (m *SummaryHeatHeatItem) validateEventName(formats strfmt.Registry) error {

	if err := validate.Required("event_name", "body", m.EventName); err != nil {
		return err
	}

	return nil
}

func (m *SummaryHeatHeatItem) validateOperationRate(formats strfmt.Registry) error {

	if err := validate.Required("operation_rate", "body", m.OperationRate); err != nil {
		return err
	}

	return nil
}

func (m *SummaryHeatHeatItem) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *SummaryHeatHeatItem) validateTime(formats strfmt.Registry) error {

	if err := validate.Required("time", "body", m.Time); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SummaryHeatHeatItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SummaryHeatHeatItem) UnmarshalBinary(b []byte) error {
	var res SummaryHeatHeatItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
