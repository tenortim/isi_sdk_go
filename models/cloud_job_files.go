// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CloudJobFiles cloud job files
// swagger:model CloudJobFiles
type CloudJobFiles struct {

	// The file filtering logic to find files for this job
	FileMatchingPattern Empty `json:"file_matching_pattern,omitempty"`

	// A list of files to be addressed by this job.  (Note* these are only reported when audit_level is 'high'
	Names []*CloudJobFilesName `json:"names"`

	// The total number of files addressed by this job
	Total int64 `json:"total,omitempty"`

	// The number of canceled files
	TotalCanceled int64 `json:"total_canceled,omitempty"`

	// The number of files which failed
	TotalFailed int64 `json:"total_failed,omitempty"`

	// The number of files pending action
	TotalPending int64 `json:"total_pending,omitempty"`

	// The number of files currently being processed
	TotalProcessing int64 `json:"total_processing,omitempty"`

	// The total number of files successfully completed
	TotalSucceeded int64 `json:"total_succeeded,omitempty"`
}

// Validate validates this cloud job files
func (m *CloudJobFiles) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNames(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudJobFiles) validateNames(formats strfmt.Registry) error {

	if swag.IsZero(m.Names) { // not required
		return nil
	}

	for i := 0; i < len(m.Names); i++ {
		if swag.IsZero(m.Names[i]) { // not required
			continue
		}

		if m.Names[i] != nil {
			if err := m.Names[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("names" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudJobFiles) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudJobFiles) UnmarshalBinary(b []byte) error {
	var res CloudJobFiles
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
