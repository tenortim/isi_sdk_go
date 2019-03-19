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

// ClusterNodeStatus cluster node status
// swagger:model ClusterNodeStatus
type ClusterNodeStatus struct {

	// Battery status information.
	Batterystatus *NodeStatusNodeBatterystatus `json:"batterystatus,omitempty"`

	// Storage capacity of this node.
	Capacity []*NodeStatusNodeCapacityItem `json:"capacity"`

	// CPU status information for this node.
	CPU *NodeStatusNodeCPU `json:"cpu,omitempty"`

	// Node NVRAM information.
	Nvram *NodeStatusNodeNvram `json:"nvram,omitempty"`

	// Information about this node's power supplies.
	Powersupplies *NodeStatusNodePowersupplies `json:"powersupplies,omitempty"`

	// OneFS release.
	Release string `json:"release,omitempty"`

	// Seconds this node has been online.
	Uptime int64 `json:"uptime,omitempty"`

	// OneFS version.
	Version string `json:"version,omitempty"`
}

// Validate validates this cluster node status
func (m *ClusterNodeStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBatterystatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNvram(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePowersupplies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterNodeStatus) validateBatterystatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Batterystatus) { // not required
		return nil
	}

	if m.Batterystatus != nil {
		if err := m.Batterystatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("batterystatus")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterNodeStatus) validateCapacity(formats strfmt.Registry) error {

	if swag.IsZero(m.Capacity) { // not required
		return nil
	}

	for i := 0; i < len(m.Capacity); i++ {
		if swag.IsZero(m.Capacity[i]) { // not required
			continue
		}

		if m.Capacity[i] != nil {
			if err := m.Capacity[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("capacity" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterNodeStatus) validateCPU(formats strfmt.Registry) error {

	if swag.IsZero(m.CPU) { // not required
		return nil
	}

	if m.CPU != nil {
		if err := m.CPU.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cpu")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterNodeStatus) validateNvram(formats strfmt.Registry) error {

	if swag.IsZero(m.Nvram) { // not required
		return nil
	}

	if m.Nvram != nil {
		if err := m.Nvram.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nvram")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterNodeStatus) validatePowersupplies(formats strfmt.Registry) error {

	if swag.IsZero(m.Powersupplies) { // not required
		return nil
	}

	if m.Powersupplies != nil {
		if err := m.Powersupplies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("powersupplies")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterNodeStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterNodeStatus) UnmarshalBinary(b []byte) error {
	var res ClusterNodeStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
