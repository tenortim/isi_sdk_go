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

// JobJobPrepairParams job job prepair params
// swagger:model JobJobPrepairParams
type JobJobPrepairParams struct {

	// Type of permissions; not accepted with mode=clone or mode=inherit.
	// Enum: [global sid unix native]
	MappingType string `json:"mapping_type,omitempty"`

	// Type of PermissionRepair operation.
	// Required: true
	// Enum: [clone inherit convert]
	Mode *string `json:"mode"`

	// IFS file or directory to use as a template; required with mode=clone and mode=inherit, not accepted with mode=convert.
	Template string `json:"template,omitempty"`

	// Authentication zone; not accepted with mode=clone or mode=inherit.
	Zone string `json:"zone,omitempty"`
}

// Validate validates this job job prepair params
func (m *JobJobPrepairParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMappingType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var jobJobPrepairParamsTypeMappingTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["global","sid","unix","native"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		jobJobPrepairParamsTypeMappingTypePropEnum = append(jobJobPrepairParamsTypeMappingTypePropEnum, v)
	}
}

const (

	// JobJobPrepairParamsMappingTypeGlobal captures enum value "global"
	JobJobPrepairParamsMappingTypeGlobal string = "global"

	// JobJobPrepairParamsMappingTypeSid captures enum value "sid"
	JobJobPrepairParamsMappingTypeSid string = "sid"

	// JobJobPrepairParamsMappingTypeUnix captures enum value "unix"
	JobJobPrepairParamsMappingTypeUnix string = "unix"

	// JobJobPrepairParamsMappingTypeNative captures enum value "native"
	JobJobPrepairParamsMappingTypeNative string = "native"
)

// prop value enum
func (m *JobJobPrepairParams) validateMappingTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, jobJobPrepairParamsTypeMappingTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *JobJobPrepairParams) validateMappingType(formats strfmt.Registry) error {

	if swag.IsZero(m.MappingType) { // not required
		return nil
	}

	// value enum
	if err := m.validateMappingTypeEnum("mapping_type", "body", m.MappingType); err != nil {
		return err
	}

	return nil
}

var jobJobPrepairParamsTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["clone","inherit","convert"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		jobJobPrepairParamsTypeModePropEnum = append(jobJobPrepairParamsTypeModePropEnum, v)
	}
}

const (

	// JobJobPrepairParamsModeClone captures enum value "clone"
	JobJobPrepairParamsModeClone string = "clone"

	// JobJobPrepairParamsModeInherit captures enum value "inherit"
	JobJobPrepairParamsModeInherit string = "inherit"

	// JobJobPrepairParamsModeConvert captures enum value "convert"
	JobJobPrepairParamsModeConvert string = "convert"
)

// prop value enum
func (m *JobJobPrepairParams) validateModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, jobJobPrepairParamsTypeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *JobJobPrepairParams) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("mode", "body", m.Mode); err != nil {
		return err
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", *m.Mode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobJobPrepairParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobJobPrepairParams) UnmarshalBinary(b []byte) error {
	var res JobJobPrepairParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
