// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FilepoolPolicyFileMatchingPattern filepool policy file matching pattern
// swagger:model FilepoolPolicyFileMatchingPattern
type FilepoolPolicyFileMatchingPattern struct {

	// or criteria
	// Required: true
	OrCriteria []*FilepoolPolicyFileMatchingPatternOrCriteriaItem `json:"or_criteria"`
}

// Validate validates this filepool policy file matching pattern
func (m *FilepoolPolicyFileMatchingPattern) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrCriteria(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FilepoolPolicyFileMatchingPattern) validateOrCriteria(formats strfmt.Registry) error {

	if err := validate.Required("or_criteria", "body", m.OrCriteria); err != nil {
		return err
	}

	for i := 0; i < len(m.OrCriteria); i++ {
		if swag.IsZero(m.OrCriteria[i]) { // not required
			continue
		}

		if m.OrCriteria[i] != nil {
			if err := m.OrCriteria[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("or_criteria" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FilepoolPolicyFileMatchingPattern) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FilepoolPolicyFileMatchingPattern) UnmarshalBinary(b []byte) error {
	var res FilepoolPolicyFileMatchingPattern
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
