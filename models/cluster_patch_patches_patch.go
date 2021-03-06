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

// ClusterPatchPatchesPatch cluster patch patches patch
// swagger:model ClusterPatchPatchesPatch
type ClusterPatchPatchesPatch struct {

	// A long comment about the patch.
	Comment string `json:"comment,omitempty"`

	// Other patches that this patch conflicts with.
	Conflicts []string `json:"conflicts"`

	// Other patches that this patch depends on.
	Dependencies []string `json:"dependencies"`

	// A short description of the patch.
	Description string `json:"description,omitempty"`

	// The files contained in this patch.
	Files []*ClusterPatchPatchesPatchFile `json:"files"`

	// A unique identifier for the patch.
	ID string `json:"id,omitempty"`

	// The name of the patch.
	Name string `json:"name,omitempty"`

	// The nodes that this patch is installed on.
	Nodes []int64 `json:"nodes"`

	// The intallation status of this patch on the cluster.
	Status string `json:"status,omitempty"`
}

// Validate validates this cluster patch patches patch
func (m *ClusterPatchPatchesPatch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterPatchPatchesPatch) validateFiles(formats strfmt.Registry) error {

	if swag.IsZero(m.Files) { // not required
		return nil
	}

	for i := 0; i < len(m.Files); i++ {
		if swag.IsZero(m.Files[i]) { // not required
			continue
		}

		if m.Files[i] != nil {
			if err := m.Files[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterPatchPatchesPatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterPatchPatchesPatch) UnmarshalBinary(b []byte) error {
	var res ClusterPatchPatchesPatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
