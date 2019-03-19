// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// QuotaQuota quota quota
// swagger:model QuotaQuota
type QuotaQuota struct {

	// If true, SMB shares using the quota directory see the quota thresholds as share size.
	Container bool `json:"container,omitempty"`

	// True if the quota provides enforcement, otherwise a accounting quota.
	Enforced bool `json:"enforced,omitempty"`

	// If false and the quota is linked, attempt to unlink.
	Linked bool `json:"linked,omitempty"`

	// thresholds
	Thresholds *QuotaQuotaThresholds `json:"thresholds,omitempty"`

	// If true, thresholds apply to data plus filesystem overhead required to store the data (i.e. 'physical' usage).
	ThresholdsIncludeOverhead bool `json:"thresholds_include_overhead,omitempty"`
}

// Validate validates this quota quota
func (m *QuotaQuota) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateThresholds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuotaQuota) validateThresholds(formats strfmt.Registry) error {

	if swag.IsZero(m.Thresholds) { // not required
		return nil
	}

	if m.Thresholds != nil {
		if err := m.Thresholds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("thresholds")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QuotaQuota) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuotaQuota) UnmarshalBinary(b []byte) error {
	var res QuotaQuota
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
