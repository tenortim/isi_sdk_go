// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CloudJobCreateParams A cloud job for archiving or recalling files or restoring COI
// swagger:model CloudJobCreateParams
type CloudJobCreateParams struct {

	// The names of accounts for COI restore
	Accounts []string `json:"accounts"`

	// Directories addressed by this job
	Directories []string `json:"directories"`

	// The new expiration date in seconds
	ExpirationDate int64 `json:"expiration_date,omitempty"`

	// The file filtering logic to find files for this job. (Only applicable for 'recall' jobs)
	FileMatchingPattern Empty `json:"file_matching_pattern,omitempty"`

	// Filenames addressed by this job
	Files []string `json:"files"`

	// The name of an existing cloudpool policy to apply to this job. (Only applicable for 'archive' jobs)
	Policy string `json:"policy,omitempty"`

	// The type of cloud action to be performed by this job.
	// Required: true
	// Enum: [archive recall local-garbage-collection cloud-garbage-collection cache-writeback cache-on-access cache-invalidation restore-coi]
	Type *string `json:"type"`
}

// Validate validates this cloud job create params
func (m *CloudJobCreateParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var cloudJobCreateParamsTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["archive","recall","local-garbage-collection","cloud-garbage-collection","cache-writeback","cache-on-access","cache-invalidation","restore-coi"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloudJobCreateParamsTypeTypePropEnum = append(cloudJobCreateParamsTypeTypePropEnum, v)
	}
}

const (

	// CloudJobCreateParamsTypeArchive captures enum value "archive"
	CloudJobCreateParamsTypeArchive string = "archive"

	// CloudJobCreateParamsTypeRecall captures enum value "recall"
	CloudJobCreateParamsTypeRecall string = "recall"

	// CloudJobCreateParamsTypeLocalGarbageCollection captures enum value "local-garbage-collection"
	CloudJobCreateParamsTypeLocalGarbageCollection string = "local-garbage-collection"

	// CloudJobCreateParamsTypeCloudGarbageCollection captures enum value "cloud-garbage-collection"
	CloudJobCreateParamsTypeCloudGarbageCollection string = "cloud-garbage-collection"

	// CloudJobCreateParamsTypeCacheWriteback captures enum value "cache-writeback"
	CloudJobCreateParamsTypeCacheWriteback string = "cache-writeback"

	// CloudJobCreateParamsTypeCacheOnAccess captures enum value "cache-on-access"
	CloudJobCreateParamsTypeCacheOnAccess string = "cache-on-access"

	// CloudJobCreateParamsTypeCacheInvalidation captures enum value "cache-invalidation"
	CloudJobCreateParamsTypeCacheInvalidation string = "cache-invalidation"

	// CloudJobCreateParamsTypeRestoreCoi captures enum value "restore-coi"
	CloudJobCreateParamsTypeRestoreCoi string = "restore-coi"
)

// prop value enum
func (m *CloudJobCreateParams) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, cloudJobCreateParamsTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CloudJobCreateParams) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudJobCreateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudJobCreateParams) UnmarshalBinary(b []byte) error {
	var res CloudJobCreateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
