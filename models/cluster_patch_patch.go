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

// ClusterPatchPatch A software patch that can be applied to OneFS.
// swagger:model ClusterPatchPatch
type ClusterPatchPatch struct {

	// The path location of the patch file.
	Location string `json:"location,omitempty"`

	// The name or path of the patch to install.
	// Required: true
	Patch *string `json:"patch"`
}

// Validate validates this cluster patch patch
func (m *ClusterPatchPatch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePatch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterPatchPatch) validatePatch(formats strfmt.Registry) error {

	if err := validate.Required("patch", "body", m.Patch); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterPatchPatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterPatchPatch) UnmarshalBinary(b []byte) error {
	var res ClusterPatchPatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
