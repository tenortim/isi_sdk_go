// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SmbLogLevelFiltersFilter smb log level filters filter
// swagger:model SmbLogLevelFiltersFilter
type SmbLogLevelFiltersFilter struct {

	// Unique ID of the log filter.
	ID int64 `json:"id,omitempty"`

	// Array of client IP addresses to filter against.
	IPAddrs []string `json:"ip_addrs"`

	// Logging level of the filter.
	// Required: true
	Level *string `json:"level"`

	// Array of SMB operations to filter against.
	Ops []string `json:"ops"`
}

// Validate validates this smb log level filters filter
func (m *SmbLogLevelFiltersFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SmbLogLevelFiltersFilter) validateLevel(formats strfmt.Registry) error {

	if err := validate.Required("level", "body", m.Level); err != nil {
		return err
	}

	return nil
}

var smbLogLevelFiltersFilterOpsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["read","write","session-setup","logoff","flush","notify","tree-connect","tree-disconnect","create","delete","oplock","locking","set-info","query","close","create-directory","delete-directory"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		smbLogLevelFiltersFilterOpsItemsEnum = append(smbLogLevelFiltersFilterOpsItemsEnum, v)
	}
}

func (m *SmbLogLevelFiltersFilter) validateOpsItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, smbLogLevelFiltersFilterOpsItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *SmbLogLevelFiltersFilter) validateOps(formats strfmt.Registry) error {

	if swag.IsZero(m.Ops) { // not required
		return nil
	}

	for i := 0; i < len(m.Ops); i++ {

		// value enum
		if err := m.validateOpsItemsEnum("ops"+"."+strconv.Itoa(i), "body", m.Ops[i]); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SmbLogLevelFiltersFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SmbLogLevelFiltersFilter) UnmarshalBinary(b []byte) error {
	var res SmbLogLevelFiltersFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
