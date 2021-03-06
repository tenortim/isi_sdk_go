// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// JobStatisticsJobNode job statistics job node
// swagger:model JobStatisticsJobNode
type JobStatisticsJobNode struct {

	// cpu
	// Required: true
	CPU *JobStatisticsJobNodeCPU `json:"cpu"`

	// io
	// Required: true
	Io *JobStatisticsJobNodeIo `json:"io"`

	// memory
	// Required: true
	Memory *JobStatisticsJobNodeMemory `json:"memory"`

	// The devid of the node.
	// Required: true
	Node *int64 `json:"node"`

	// The process ID of the job on this node.
	// Required: true
	Pid *int64 `json:"pid"`

	// The number of workers for this job on this node.
	// Required: true
	TotalWorkers *int64 `json:"total_workers"`

	// workers
	// Required: true
	Workers []*JobStatisticsJobNodeWorker `json:"workers"`
}

// Validate validates this job statistics job node
func (m *JobStatisticsJobNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalWorkers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobStatisticsJobNode) validateCPU(formats strfmt.Registry) error {

	if err := validate.Required("cpu", "body", m.CPU); err != nil {
		return err
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

func (m *JobStatisticsJobNode) validateIo(formats strfmt.Registry) error {

	if err := validate.Required("io", "body", m.Io); err != nil {
		return err
	}

	if m.Io != nil {
		if err := m.Io.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("io")
			}
			return err
		}
	}

	return nil
}

func (m *JobStatisticsJobNode) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("memory", "body", m.Memory); err != nil {
		return err
	}

	if m.Memory != nil {
		if err := m.Memory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("memory")
			}
			return err
		}
	}

	return nil
}

func (m *JobStatisticsJobNode) validateNode(formats strfmt.Registry) error {

	if err := validate.Required("node", "body", m.Node); err != nil {
		return err
	}

	return nil
}

func (m *JobStatisticsJobNode) validatePid(formats strfmt.Registry) error {

	if err := validate.Required("pid", "body", m.Pid); err != nil {
		return err
	}

	return nil
}

func (m *JobStatisticsJobNode) validateTotalWorkers(formats strfmt.Registry) error {

	if err := validate.Required("total_workers", "body", m.TotalWorkers); err != nil {
		return err
	}

	return nil
}

func (m *JobStatisticsJobNode) validateWorkers(formats strfmt.Registry) error {

	if err := validate.Required("workers", "body", m.Workers); err != nil {
		return err
	}

	for i := 0; i < len(m.Workers); i++ {
		if swag.IsZero(m.Workers[i]) { // not required
			continue
		}

		if m.Workers[i] != nil {
			if err := m.Workers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobStatisticsJobNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobStatisticsJobNode) UnmarshalBinary(b []byte) error {
	var res JobStatisticsJobNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
