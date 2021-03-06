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

// SmbOpenfile smb openfile
// swagger:model SmbOpenfile
type SmbOpenfile struct {

	// Path of file within /ifs.
	// Required: true
	File *string `json:"file"`

	// The file ID.
	// Required: true
	ID *int64 `json:"id"`

	// Number of locks user holds on file.
	// Required: true
	Locks *int64 `json:"locks"`

	// The user's permissions on file.
	// Required: true
	Permissions []string `json:"permissions"`

	// User holding file open.
	// Required: true
	User *string `json:"user"`
}

// Validate validates this smb openfile
func (m *SmbOpenfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SmbOpenfile) validateFile(formats strfmt.Registry) error {

	if err := validate.Required("file", "body", m.File); err != nil {
		return err
	}

	return nil
}

func (m *SmbOpenfile) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *SmbOpenfile) validateLocks(formats strfmt.Registry) error {

	if err := validate.Required("locks", "body", m.Locks); err != nil {
		return err
	}

	return nil
}

var smbOpenfilePermissionsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["read","write","create"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		smbOpenfilePermissionsItemsEnum = append(smbOpenfilePermissionsItemsEnum, v)
	}
}

func (m *SmbOpenfile) validatePermissionsItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, smbOpenfilePermissionsItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *SmbOpenfile) validatePermissions(formats strfmt.Registry) error {

	if err := validate.Required("permissions", "body", m.Permissions); err != nil {
		return err
	}

	for i := 0; i < len(m.Permissions); i++ {

		// value enum
		if err := m.validatePermissionsItemsEnum("permissions"+"."+strconv.Itoa(i), "body", m.Permissions[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *SmbOpenfile) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("user", "body", m.User); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SmbOpenfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SmbOpenfile) UnmarshalBinary(b []byte) error {
	var res SmbOpenfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
