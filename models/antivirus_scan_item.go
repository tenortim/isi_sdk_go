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

// AntivirusScanItem Parameters for a manual antivirus scan.
// swagger:model AntivirusScanItem
type AntivirusScanItem struct {

	// The full path of the file to scan.
	// Required: true
	// Min Length: 1
	File *string `json:"file"`

	// Forces the scan to run regardless of whether the files were recently scanned. The default value is true.
	ForceRun bool `json:"force_run,omitempty"`

	// The ID of the policy to use for the scan. By default, the scan will use the MANUAL policy.
	// Min Length: 1
	Policy string `json:"policy,omitempty"`

	// The ID for the report for this scan. A report ID will be generated if one is not provided.
	// Min Length: 1
	ReportID string `json:"report_id,omitempty"`
}

// Validate validates this antivirus scan item
func (m *AntivirusScanItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFile(formats); err != nil {
		res = append(res, err)
	}

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

func (m *AntivirusScanItem) validateFile(formats strfmt.Registry) error {

	if err := validate.Required("file", "body", m.File); err != nil {
		return err
	}

	if err := validate.MinLength("file", "body", string(*m.File), 1); err != nil {
		return err
	}

	return nil
}

func (m *AntivirusScanItem) validatePolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.Policy) { // not required
		return nil
	}

	if err := validate.MinLength("policy", "body", string(m.Policy), 1); err != nil {
		return err
	}

	return nil
}

func (m *AntivirusScanItem) validateReportID(formats strfmt.Registry) error {

	if swag.IsZero(m.ReportID) { // not required
		return nil
	}

	if err := validate.MinLength("report_id", "body", string(m.ReportID), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AntivirusScanItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AntivirusScanItem) UnmarshalBinary(b []byte) error {
	var res AntivirusScanItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
