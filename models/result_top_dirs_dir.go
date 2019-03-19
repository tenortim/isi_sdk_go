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

// ResultTopDirsDir result top dirs dir
// swagger:model ResultTopDirsDir
type ResultTopDirsDir struct {

	// Directory access time
	// Required: true
	Atime *int64 `json:"atime"`

	// Directory creation begin time.
	// Required: true
	Btime *int64 `json:"btime"`

	// Unix inode change time.
	// Required: true
	Ctime *int64 `json:"ctime"`

	// Relative directory path under /ifs/.
	// Required: true
	Path *string `json:"path"`
}

// Validate validates this result top dirs dir
func (m *ResultTopDirsDir) Validate(formats strfmt.Registry) error {
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

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResultTopDirsDir) validateAtime(formats strfmt.Registry) error {

	if err := validate.Required("atime", "body", m.Atime); err != nil {
		return err
	}

	return nil
}

func (m *ResultTopDirsDir) validateBtime(formats strfmt.Registry) error {

	if err := validate.Required("btime", "body", m.Btime); err != nil {
		return err
	}

	return nil
}

func (m *ResultTopDirsDir) validateCtime(formats strfmt.Registry) error {

	if err := validate.Required("ctime", "body", m.Ctime); err != nil {
		return err
	}

	return nil
}

func (m *ResultTopDirsDir) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResultTopDirsDir) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResultTopDirsDir) UnmarshalBinary(b []byte) error {
	var res ResultTopDirsDir
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
