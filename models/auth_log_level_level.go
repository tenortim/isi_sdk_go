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

// AuthLogLevelLevel auth log level level
// swagger:model AuthLogLevelLevel
type AuthLogLevelLevel struct {

	// Valid auth logging levels
	// Enum: [always error warning info verbose debug trace]
	LsassLevel string `json:"lsass-level,omitempty"`

	// Valid auth logging levels
	// Enum: [always error warning info verbose debug trace]
	NetlogonLevel string `json:"netlogon-level,omitempty"`
}

// Validate validates this auth log level level
func (m *AuthLogLevelLevel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLsassLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetlogonLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var authLogLevelLevelTypeLsassLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["always","error","warning","info","verbose","debug","trace"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		authLogLevelLevelTypeLsassLevelPropEnum = append(authLogLevelLevelTypeLsassLevelPropEnum, v)
	}
}

const (

	// AuthLogLevelLevelLsassLevelAlways captures enum value "always"
	AuthLogLevelLevelLsassLevelAlways string = "always"

	// AuthLogLevelLevelLsassLevelError captures enum value "error"
	AuthLogLevelLevelLsassLevelError string = "error"

	// AuthLogLevelLevelLsassLevelWarning captures enum value "warning"
	AuthLogLevelLevelLsassLevelWarning string = "warning"

	// AuthLogLevelLevelLsassLevelInfo captures enum value "info"
	AuthLogLevelLevelLsassLevelInfo string = "info"

	// AuthLogLevelLevelLsassLevelVerbose captures enum value "verbose"
	AuthLogLevelLevelLsassLevelVerbose string = "verbose"

	// AuthLogLevelLevelLsassLevelDebug captures enum value "debug"
	AuthLogLevelLevelLsassLevelDebug string = "debug"

	// AuthLogLevelLevelLsassLevelTrace captures enum value "trace"
	AuthLogLevelLevelLsassLevelTrace string = "trace"
)

// prop value enum
func (m *AuthLogLevelLevel) validateLsassLevelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, authLogLevelLevelTypeLsassLevelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AuthLogLevelLevel) validateLsassLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.LsassLevel) { // not required
		return nil
	}

	// value enum
	if err := m.validateLsassLevelEnum("lsass-level", "body", m.LsassLevel); err != nil {
		return err
	}

	return nil
}

var authLogLevelLevelTypeNetlogonLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["always","error","warning","info","verbose","debug","trace"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		authLogLevelLevelTypeNetlogonLevelPropEnum = append(authLogLevelLevelTypeNetlogonLevelPropEnum, v)
	}
}

const (

	// AuthLogLevelLevelNetlogonLevelAlways captures enum value "always"
	AuthLogLevelLevelNetlogonLevelAlways string = "always"

	// AuthLogLevelLevelNetlogonLevelError captures enum value "error"
	AuthLogLevelLevelNetlogonLevelError string = "error"

	// AuthLogLevelLevelNetlogonLevelWarning captures enum value "warning"
	AuthLogLevelLevelNetlogonLevelWarning string = "warning"

	// AuthLogLevelLevelNetlogonLevelInfo captures enum value "info"
	AuthLogLevelLevelNetlogonLevelInfo string = "info"

	// AuthLogLevelLevelNetlogonLevelVerbose captures enum value "verbose"
	AuthLogLevelLevelNetlogonLevelVerbose string = "verbose"

	// AuthLogLevelLevelNetlogonLevelDebug captures enum value "debug"
	AuthLogLevelLevelNetlogonLevelDebug string = "debug"

	// AuthLogLevelLevelNetlogonLevelTrace captures enum value "trace"
	AuthLogLevelLevelNetlogonLevelTrace string = "trace"
)

// prop value enum
func (m *AuthLogLevelLevel) validateNetlogonLevelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, authLogLevelLevelTypeNetlogonLevelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AuthLogLevelLevel) validateNetlogonLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.NetlogonLevel) { // not required
		return nil
	}

	// value enum
	if err := m.validateNetlogonLevelEnum("netlogon-level", "body", m.NetlogonLevel); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuthLogLevelLevel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthLogLevelLevel) UnmarshalBinary(b []byte) error {
	var res AuthLogLevelLevel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
