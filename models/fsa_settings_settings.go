// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// FsaSettingsSettings fsa settings settings
// swagger:model FsaSettingsSettings
type FsaSettingsSettings struct {

	// Name of question template to use for new FSA jobs.
	DefaultTemplate string `json:"default_template,omitempty"`

	// Maximum directory depth used for disk_usage question if not specified in the question.
	DiskUsageDepth int64 `json:"disk_usage_depth,omitempty"`

	// Maximum age of non-pinned results in seconds.
	MaxAge int64 `json:"max_age,omitempty"`

	// Maximum number of non-pinned result sets to keep.
	MaxCount int64 `json:"max_count,omitempty"`

	// Squash depth to use for squash binning questions if not specified in the question.
	SquashDepth int64 `json:"squash_depth,omitempty"`

	// Maximum number of items in a Top-N question result if not specified in the question.
	TopNMax int64 `json:"top_n_max,omitempty"`

	// If true, use a snapshot for consistency, otherwise analyze head.
	UseSnapshot bool `json:"use_snapshot,omitempty"`
}

// Validate validates this fsa settings settings
func (m *FsaSettingsSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FsaSettingsSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FsaSettingsSettings) UnmarshalBinary(b []byte) error {
	var res FsaSettingsSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
