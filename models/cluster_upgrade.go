// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ClusterUpgrade Add nodes to a running upgrade.
// swagger:model ClusterUpgrade
type ClusterUpgrade struct {

	// The nodes (to be) scheduled for an existing upgrade ordered by queue position number. [<lnn-1>, <lnn-2>, ... ], 'All', null
	NodesToRollingUpgrade []int64 `json:"nodes_to_rolling_upgrade"`
}

// Validate validates this cluster upgrade
func (m *ClusterUpgrade) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterUpgrade) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterUpgrade) UnmarshalBinary(b []byte) error {
	var res ClusterUpgrade
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
