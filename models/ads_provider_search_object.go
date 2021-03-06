// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AdsProviderSearchObject ads provider search object
// swagger:model AdsProviderSearchObject
type AdsProviderSearchObject struct {

	// description
	Description string `json:"description,omitempty"`

	// display name
	DisplayName string `json:"display_name,omitempty"`

	// Specifies properties for a persona, which consists of either a 'type' and a 'name' or an 'ID'.
	ID *GroupMember `json:"id,omitempty"`
}

// Validate validates this ads provider search object
func (m *AdsProviderSearchObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdsProviderSearchObject) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if m.ID != nil {
		if err := m.ID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("id")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdsProviderSearchObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdsProviderSearchObject) UnmarshalBinary(b []byte) error {
	var res AdsProviderSearchObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
