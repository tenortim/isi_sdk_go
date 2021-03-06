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

// ResultTopFilesFile result top files file
// swagger:model ResultTopFilesFile
type ResultTopFilesFile struct {

	// File access time.
	// Required: true
	Atime *int64 `json:"atime"`

	// File creation begin time.
	// Required: true
	Btime *int64 `json:"btime"`

	// Unix inode change time.
	// Required: true
	Ctime *int64 `json:"ctime"`

	// Logical file size in bytes.
	// Required: true
	LogSize *int64 `json:"log_size"`

	// Relative file path under /ifs/.
	// Required: true
	Path *string `json:"path"`

	// Physical file size in bytes.
	// Required: true
	PhysSize *int64 `json:"phys_size"`
}

// Validate validates this result top files file
func (m *ResultTopFilesFile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAtime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBtime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCtime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhysSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResultTopFilesFile) validateAtime(formats strfmt.Registry) error {

	if err := validate.Required("atime", "body", m.Atime); err != nil {
		return err
	}

	return nil
}

func (m *ResultTopFilesFile) validateBtime(formats strfmt.Registry) error {

	if err := validate.Required("btime", "body", m.Btime); err != nil {
		return err
	}

	return nil
}

func (m *ResultTopFilesFile) validateCtime(formats strfmt.Registry) error {

	if err := validate.Required("ctime", "body", m.Ctime); err != nil {
		return err
	}

	return nil
}

func (m *ResultTopFilesFile) validateLogSize(formats strfmt.Registry) error {

	if err := validate.Required("log_size", "body", m.LogSize); err != nil {
		return err
	}

	return nil
}

func (m *ResultTopFilesFile) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *ResultTopFilesFile) validatePhysSize(formats strfmt.Registry) error {

	if err := validate.Required("phys_size", "body", m.PhysSize); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResultTopFilesFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResultTopFilesFile) UnmarshalBinary(b []byte) error {
	var res ResultTopFilesFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
