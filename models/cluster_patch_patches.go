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

// ClusterPatchPatches cluster patch patches
// swagger:model ClusterPatchPatches
type ClusterPatchPatches struct {

	// patches
	Patches []*ClusterPatchPatchesPatch `json:"patches"`
}

// Validate validates this cluster patch patches
func (m *ClusterPatchPatches) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePatches(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterPatchPatches) validatePatches(formats strfmt.Registry) error {

	if swag.IsZero(m.Patches) { // not required
		return nil
	}

	for i := 0; i < len(m.Patches); i++ {
		if swag.IsZero(m.Patches[i]) { // not required
			continue
		}

		if m.Patches[i] != nil {
			if err := m.Patches[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("patches" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterPatchPatches) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterPatchPatches) UnmarshalBinary(b []byte) error {
	var res ClusterPatchPatches
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}