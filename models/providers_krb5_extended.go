// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ProvidersKrb5Extended providers krb5 extended
// swagger:model ProvidersKrb5Extended
type ProvidersKrb5Extended struct {

	// krb5
	Krb5 []*ProvidersKrb5Krb5ItemExtended `json:"krb5"`
}

// Validate validates this providers krb5 extended
func (m *ProvidersKrb5Extended) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKrb5(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProvidersKrb5Extended) validateKrb5(formats strfmt.Registry) error {

	if swag.IsZero(m.Krb5) { // not required
		return nil
	}

	for i := 0; i < len(m.Krb5); i++ {
		if swag.IsZero(m.Krb5[i]) { // not required
			continue
		}

		if m.Krb5[i] != nil {
			if err := m.Krb5[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("krb5" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProvidersKrb5Extended) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProvidersKrb5Extended) UnmarshalBinary(b []byte) error {
	var res ProvidersKrb5Extended
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}