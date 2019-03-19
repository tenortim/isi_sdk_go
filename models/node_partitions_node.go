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

// NodePartitionsNode node partitions node
// swagger:model NodePartitionsNode
type NodePartitionsNode struct {

	// Count of how many partitions are included.
	Count int64 `json:"count,omitempty"`

	// Node ID (Device Number) of this node.
	ID int64 `json:"id,omitempty"`

	// Logical Node Number (LNN) of this node.
	Lnn int64 `json:"lnn,omitempty"`

	// Partition information.
	Partitions []*NodePartitionsNodePartition `json:"partitions"`
}

// Validate validates this node partitions node
func (m *NodePartitionsNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePartitions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodePartitionsNode) validatePartitions(formats strfmt.Registry) error {

	if swag.IsZero(m.Partitions) { // not required
		return nil
	}

	for i := 0; i < len(m.Partitions); i++ {
		if swag.IsZero(m.Partitions[i]) { // not required
			continue
		}

		if m.Partitions[i] != nil {
			if err := m.Partitions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("partitions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodePartitionsNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodePartitionsNode) UnmarshalBinary(b []byte) error {
	var res NodePartitionsNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
