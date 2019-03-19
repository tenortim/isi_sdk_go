// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PoolsPoolInterfacesInterface pools pool interfaces interface
// swagger:model PoolsPoolInterfacesInterface
type PoolsPoolInterfacesInterface struct {

	// Unique interface ID.
	// Required: true
	ID *string `json:"id"`

	// List of IP addresses
	// Required: true
	IPAddrs []string `json:"ip_addrs"`

	// Logical Node Number
	// Required: true
	// Minimum: 1
	Lnn *int64 `json:"lnn"`

	// The name of the interface.
	// Required: true
	Name *string `json:"name"`

	// NIC name
	// Required: true
	NicName *string `json:"nic_name"`

	// List of owners (membership)
	// Required: true
	Owners []*PoolsPoolInterfacesInterfaceOwner `json:"owners"`

	// Status of the interface
	// Required: true
	Status *string `json:"status"`

	// Interface type.  The '*gige' types stand for 'gigabit ethernet'.  'gige' itself is occasionally also referred to in other places as 'ext' for 'external'.  'ib' and 'ib_qdr' are internal Infiniband interface types.  'vlan' and 'vmxnet3' are virtual interface types that appear on virtual nodes.  'loopback' is an interface for failover addresses and should only appear if failover is configured.
	// Required: true
	// Enum: [any gige fastgige 10gige 40gige ib ib_qdr aggregated vlan vmxnet3 loopback]
	Type *string `json:"type"`
}

// Validate validates this pools pool interfaces interface
func (m *PoolsPoolInterfacesInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPAddrs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLnn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNicName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwners(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoolsPoolInterfacesInterface) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *PoolsPoolInterfacesInterface) validateIPAddrs(formats strfmt.Registry) error {

	if err := validate.Required("ip_addrs", "body", m.IPAddrs); err != nil {
		return err
	}

	return nil
}

func (m *PoolsPoolInterfacesInterface) validateLnn(formats strfmt.Registry) error {

	if err := validate.Required("lnn", "body", m.Lnn); err != nil {
		return err
	}

	if err := validate.MinimumInt("lnn", "body", int64(*m.Lnn), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *PoolsPoolInterfacesInterface) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PoolsPoolInterfacesInterface) validateNicName(formats strfmt.Registry) error {

	if err := validate.Required("nic_name", "body", m.NicName); err != nil {
		return err
	}

	return nil
}

func (m *PoolsPoolInterfacesInterface) validateOwners(formats strfmt.Registry) error {

	if err := validate.Required("owners", "body", m.Owners); err != nil {
		return err
	}

	for i := 0; i < len(m.Owners); i++ {
		if swag.IsZero(m.Owners[i]) { // not required
			continue
		}

		if m.Owners[i] != nil {
			if err := m.Owners[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("owners" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PoolsPoolInterfacesInterface) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var poolsPoolInterfacesInterfaceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["any","gige","fastgige","10gige","40gige","ib","ib_qdr","aggregated","vlan","vmxnet3","loopback"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		poolsPoolInterfacesInterfaceTypeTypePropEnum = append(poolsPoolInterfacesInterfaceTypeTypePropEnum, v)
	}
}

const (

	// PoolsPoolInterfacesInterfaceTypeAny captures enum value "any"
	PoolsPoolInterfacesInterfaceTypeAny string = "any"

	// PoolsPoolInterfacesInterfaceTypeGige captures enum value "gige"
	PoolsPoolInterfacesInterfaceTypeGige string = "gige"

	// PoolsPoolInterfacesInterfaceTypeFastgige captures enum value "fastgige"
	PoolsPoolInterfacesInterfaceTypeFastgige string = "fastgige"

	// PoolsPoolInterfacesInterfaceTypeNr10gige captures enum value "10gige"
	PoolsPoolInterfacesInterfaceTypeNr10gige string = "10gige"

	// PoolsPoolInterfacesInterfaceTypeNr40gige captures enum value "40gige"
	PoolsPoolInterfacesInterfaceTypeNr40gige string = "40gige"

	// PoolsPoolInterfacesInterfaceTypeIb captures enum value "ib"
	PoolsPoolInterfacesInterfaceTypeIb string = "ib"

	// PoolsPoolInterfacesInterfaceTypeIbQdr captures enum value "ib_qdr"
	PoolsPoolInterfacesInterfaceTypeIbQdr string = "ib_qdr"

	// PoolsPoolInterfacesInterfaceTypeAggregated captures enum value "aggregated"
	PoolsPoolInterfacesInterfaceTypeAggregated string = "aggregated"

	// PoolsPoolInterfacesInterfaceTypeVlan captures enum value "vlan"
	PoolsPoolInterfacesInterfaceTypeVlan string = "vlan"

	// PoolsPoolInterfacesInterfaceTypeVmxnet3 captures enum value "vmxnet3"
	PoolsPoolInterfacesInterfaceTypeVmxnet3 string = "vmxnet3"

	// PoolsPoolInterfacesInterfaceTypeLoopback captures enum value "loopback"
	PoolsPoolInterfacesInterfaceTypeLoopback string = "loopback"
)

// prop value enum
func (m *PoolsPoolInterfacesInterface) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, poolsPoolInterfacesInterfaceTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PoolsPoolInterfacesInterface) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PoolsPoolInterfacesInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PoolsPoolInterfacesInterface) UnmarshalBinary(b []byte) error {
	var res PoolsPoolInterfacesInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}