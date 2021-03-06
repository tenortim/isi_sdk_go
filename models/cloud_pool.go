// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CloudPool A group of cloud accounts which will be the destination for file stubbing
// swagger:model CloudPool
type CloudPool struct {

	// A list of valid names for the accounts in this pool.  There is currently only one account allowed per pool.
	Accounts []string `json:"accounts"`

	// The guid of the cluster where this pool was created
	BirthClusterID string `json:"birth_cluster_id,omitempty"`

	// A brief description of this pool
	Description string `json:"description,omitempty"`

	// A unique name for this pool
	Name string `json:"name,omitempty"`

	// A string identifier of the cloud services vendor
	Vendor string `json:"vendor,omitempty"`
}

// Validate validates this cloud pool
func (m *CloudPool) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CloudPool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudPool) UnmarshalBinary(b []byte) error {
	var res CloudPool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
