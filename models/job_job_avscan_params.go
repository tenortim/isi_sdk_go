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

// JobJobAvscanParams job job avscan params
// swagger:model JobJobAvscanParams
type JobJobAvscanParams struct {

	// Force files to be scanned, even if excluded by the policy.
	ForceRun bool `json:"force_run,omitempty"`

	// The antivirus scan policy to run.
	// Required: true
	// Min Length: 1
	Policy *string `json:"policy"`

	// An optional report id for the scan.
	// Max Length: 15
	// Min Length: 1
	ReportID string `json:"report_id,omitempty"`

	// Update the last run time for the policy.
	Update bool `json:"update,omitempty"`
}

// Validate validates this job job avscan params
func (m *JobJobAvscanParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobJobAvscanParams) validatePolicy(formats strfmt.Registry) error {

	if err := validate.Required("policy", "body", m.Policy); err != nil {
		return err
	}

	if err := validate.MinLength("policy", "body", string(*m.Policy), 1); err != nil {
		return err
	}

	return nil
}

func (m *JobJobAvscanParams) validateReportID(formats strfmt.Registry) error {

	if swag.IsZero(m.ReportID) { // not required
		return nil
	}

	if err := validate.MinLength("report_id", "body", string(m.ReportID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("report_id", "body", string(m.ReportID), 15); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobJobAvscanParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobJobAvscanParams) UnmarshalBinary(b []byte) error {
	var res JobJobAvscanParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
