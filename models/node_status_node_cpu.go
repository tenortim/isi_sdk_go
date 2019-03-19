// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// NodeStatusNodeCPU node status node Cpu
// swagger:model NodeStatusNodeCpu
type NodeStatusNodeCPU struct {

	// Manufacturer model description of this CPU.
	Model string `json:"model,omitempty"`

	// CPU overtemp state.
	Overtemp string `json:"overtemp,omitempty"`

	// Type of processor and core of this CPU.
	Proc string `json:"proc,omitempty"`

	// CPU throttling (expressed as a percentage).
	SpeedLimit string `json:"speed_limit,omitempty"`
}

// Validate validates this node status node Cpu
func (m *NodeStatusNodeCPU) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NodeStatusNodeCPU) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeStatusNodeCPU) UnmarshalBinary(b []byte) error {
	var res NodeStatusNodeCPU
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}