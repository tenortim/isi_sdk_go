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

// HistoryFileStatistic history file statistic
// swagger:model HistoryFileStatistic
type HistoryFileStatistic struct {

	// Nodes allocated for the sync action.
	// Required: true
	Allocated *int64 `json:"allocated"`

	// An ID for a single performance report.
	// Required: true
	ID *int64 `json:"id"`

	// Sync action limit.
	// Required: true
	Limit *int64 `json:"limit"`

	// Timestamp for the performance report.
	// Required: true
	Timestamp *int64 `json:"timestamp"`

	// Total usage for the performance report.
	// Required: true
	Total *int64 `json:"total"`
}

// Validate validates this history file statistic
func (m *HistoryFileStatistic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HistoryFileStatistic) validateAllocated(formats strfmt.Registry) error {

	if err := validate.Required("allocated", "body", m.Allocated); err != nil {
		return err
	}

	return nil
}

func (m *HistoryFileStatistic) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *HistoryFileStatistic) validateLimit(formats strfmt.Registry) error {

	if err := validate.Required("limit", "body", m.Limit); err != nil {
		return err
	}

	return nil
}

func (m *HistoryFileStatistic) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

func (m *HistoryFileStatistic) validateTotal(formats strfmt.Registry) error {

	if err := validate.Required("total", "body", m.Total); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HistoryFileStatistic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HistoryFileStatistic) UnmarshalBinary(b []byte) error {
	var res HistoryFileStatistic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
