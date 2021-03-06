// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CreateAntivirusScanItemResponse The result of a manual antivirus scan.
// swagger:model CreateAntivirusScanItemResponse
type CreateAntivirusScanItemResponse struct {

	// The ID for the report for this scan. A report ID will be generated if one is not provided.
	ReportID string `json:"report_id,omitempty"`

	// The result of the scan.
	Result string `json:"result,omitempty"`
}

// Validate validates this create antivirus scan item response
func (m *CreateAntivirusScanItemResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateAntivirusScanItemResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAntivirusScanItemResponse) UnmarshalBinary(b []byte) error {
	var res CreateAntivirusScanItemResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
