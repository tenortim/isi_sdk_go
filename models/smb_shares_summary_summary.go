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

// SmbSharesSummarySummary smb shares summary summary
// swagger:model SmbSharesSummarySummary
type SmbSharesSummarySummary struct {

	// Total number of shares.
	// Required: true
	Count *int64 `json:"count"`
}

// Validate validates this smb shares summary summary
func (m *SmbSharesSummarySummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SmbSharesSummarySummary) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SmbSharesSummarySummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SmbSharesSummarySummary) UnmarshalBinary(b []byte) error {
	var res SmbSharesSummarySummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
