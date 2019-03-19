// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SyncReportPolicy sync report policy
// swagger:model SyncReportPolicy
type SyncReportPolicy struct {

	// If 'copy', source files will be copied to the target cluster.  If 'sync', the target directory will be made an image of the source directory:  Files and directories that have been deleted on the source, have been moved within the target directory, or no longer match the selection criteria will be deleted from the target directory.
	Action string `json:"action,omitempty"`

	// A file matching pattern, organized as an OR'ed set of AND'ed file criteria, for example ((a AND b) OR (x AND y)) used to define a set of files with specific properties.  Policies of type 'sync' cannot use 'path' or time criteria in their matching patterns, but policies of type 'copy' can use all listed criteria.
	FileMatchingPattern *ReportSubreportPolicyFileMatchingPattern `json:"file_matching_pattern,omitempty"`

	// User-assigned name of this sync policy.
	Name string `json:"name,omitempty"`

	// Directories that will be excluded from the sync.  Modifying this field will result in a full synchronization of all data.
	SourceExcludeDirectories []string `json:"source_exclude_directories"`

	// Directories that will be included in the sync.  Modifying this field will result in a full synchronization of all data.
	SourceIncludeDirectories []string `json:"source_include_directories"`

	// The root directory on the source cluster the files will be synced from.  Modifying this field will result in a full synchronization of all data.
	SourceRootPath string `json:"source_root_path,omitempty"`

	// Hostname or IP address of sync target cluster.  Modifying the target cluster host can result in the policy being unrunnable if the new target does not match the current target association.
	TargetHost string `json:"target_host,omitempty"`

	// Absolute filesystem path on the target cluster for the sync destination.
	TargetPath string `json:"target_path,omitempty"`
}

// Validate validates this sync report policy
func (m *SyncReportPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFileMatchingPattern(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SyncReportPolicy) validateFileMatchingPattern(formats strfmt.Registry) error {

	if swag.IsZero(m.FileMatchingPattern) { // not required
		return nil
	}

	if m.FileMatchingPattern != nil {
		if err := m.FileMatchingPattern.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("file_matching_pattern")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SyncReportPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SyncReportPolicy) UnmarshalBinary(b []byte) error {
	var res SyncReportPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}