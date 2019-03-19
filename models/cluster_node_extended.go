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

// ClusterNodeExtended cluster node extended
// swagger:model ClusterNodeExtended
type ClusterNodeExtended struct {

	// List of the drives in this node.
	Drives []*NodeDrivesNodeDrive `json:"drives"`

	// Node hardware identifying information (static).
	Hardware *ClusterNodeHardware `json:"hardware,omitempty"`

	// Node ID (Device Number) of this node.
	ID int64 `json:"id,omitempty"`

	// Logical Node Number (LNN) of this node.
	Lnn int64 `json:"lnn,omitempty"`

	// Node partition information.
	Partitions *ClusterNodePartitions `json:"partitions,omitempty"`

	// Node sensor information (hardware reported).
	Sensors *ClusterNodeSensors `json:"sensors,omitempty"`

	// Node state information (reported and modifiable).
	State *ClusterNodeStateExtended `json:"state,omitempty"`

	// Node status information (hardware reported).
	Status *ClusterNodeStatus `json:"status,omitempty"`
}

// Validate validates this cluster node extended
func (m *ClusterNodeExtended) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDrives(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHardware(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartitions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSensors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterNodeExtended) validateDrives(formats strfmt.Registry) error {

	if swag.IsZero(m.Drives) { // not required
		return nil
	}

	for i := 0; i < len(m.Drives); i++ {
		if swag.IsZero(m.Drives[i]) { // not required
			continue
		}

		if m.Drives[i] != nil {
			if err := m.Drives[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("drives" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterNodeExtended) validateHardware(formats strfmt.Registry) error {

	if swag.IsZero(m.Hardware) { // not required
		return nil
	}

	if m.Hardware != nil {
		if err := m.Hardware.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hardware")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterNodeExtended) validatePartitions(formats strfmt.Registry) error {

	if swag.IsZero(m.Partitions) { // not required
		return nil
	}

	if m.Partitions != nil {
		if err := m.Partitions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("partitions")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterNodeExtended) validateSensors(formats strfmt.Registry) error {

	if swag.IsZero(m.Sensors) { // not required
		return nil
	}

	if m.Sensors != nil {
		if err := m.Sensors.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sensors")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterNodeExtended) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterNodeExtended) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterNodeExtended) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterNodeExtended) UnmarshalBinary(b []byte) error {
	var res ClusterNodeExtended
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}