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

// FilepoolPolicyFileMatchingPatternOrCriteriaItem filepool policy file matching pattern or criteria item
// swagger:model FilepoolPolicyFileMatchingPatternOrCriteriaItem
type FilepoolPolicyFileMatchingPatternOrCriteriaItem struct {

	// and criteria
	// Required: true
	AndCriteria []*FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItem `json:"and_criteria"`
}

// Validate validates this filepool policy file matching pattern or criteria item
func (m *FilepoolPolicyFileMatchingPatternOrCriteriaItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAndCriteria(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FilepoolPolicyFileMatchingPatternOrCriteriaItem) validateAndCriteria(formats strfmt.Registry) error {

	if err := validate.Required("and_criteria", "body", m.AndCriteria); err != nil {
		return err
	}

	for i := 0; i < len(m.AndCriteria); i++ {
		if swag.IsZero(m.AndCriteria[i]) { // not required
			continue
		}

		if m.AndCriteria[i] != nil {
			if err := m.AndCriteria[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("and_criteria" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FilepoolPolicyFileMatchingPatternOrCriteriaItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FilepoolPolicyFileMatchingPatternOrCriteriaItem) UnmarshalBinary(b []byte) error {
	var res FilepoolPolicyFileMatchingPatternOrCriteriaItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
