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

// FilepoolPolicyCreateParams filepool policy create params
// swagger:model FilepoolPolicyCreateParams
type FilepoolPolicyCreateParams struct {
	FilepoolPolicy

	// The file matching rules for this policy
	// Required: true
	FileMatchingPattern *FilepoolPolicyFileMatchingPattern `json:"file_matching_pattern"`

	// A unique name for this policy
	// Required: true
	Name *string `json:"name"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *FilepoolPolicyCreateParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 FilepoolPolicy
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.FilepoolPolicy = aO0

	// AO1
	var dataAO1 struct {
		FileMatchingPattern *FilepoolPolicyFileMatchingPattern `json:"file_matching_pattern"`

		Name *string `json:"name"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.FileMatchingPattern = dataAO1.FileMatchingPattern

	m.Name = dataAO1.Name

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m FilepoolPolicyCreateParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.FilepoolPolicy)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		FileMatchingPattern *FilepoolPolicyFileMatchingPattern `json:"file_matching_pattern"`

		Name *string `json:"name"`
	}

	dataAO1.FileMatchingPattern = m.FileMatchingPattern

	dataAO1.Name = m.Name

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this filepool policy create params
func (m *FilepoolPolicyCreateParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with FilepoolPolicy
	if err := m.FilepoolPolicy.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileMatchingPattern(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FilepoolPolicyCreateParams) validateFileMatchingPattern(formats strfmt.Registry) error {

	if err := validate.Required("file_matching_pattern", "body", m.FileMatchingPattern); err != nil {
		return err
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

func (m *FilepoolPolicyCreateParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FilepoolPolicyCreateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FilepoolPolicyCreateParams) UnmarshalBinary(b []byte) error {
	var res FilepoolPolicyCreateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
