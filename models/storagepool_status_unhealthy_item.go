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

// StoragepoolStatusUnhealthyItem storagepool status unhealthy item
// swagger:model StoragepoolStatusUnhealthyItem
type StoragepoolStatusUnhealthyItem struct {

	// The affected nodes and/or drives.
	// Required: true
	Affected []*StoragepoolStatusUnhealthyItemAffectedItem `json:"affected"`

	// diskpool
	Diskpool *StoragepoolStatusUnhealthyItemDiskpool `json:"diskpool,omitempty"`

	// An array of containing any health issues with this pool.  If the pool is healthy, the list is empty.
	HealthFlags []string `json:"health_flags"`
}

// Validate validates this storagepool status unhealthy item
func (m *StoragepoolStatusUnhealthyItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAffected(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiskpool(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthFlags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoragepoolStatusUnhealthyItem) validateAffected(formats strfmt.Registry) error {

	if err := validate.Required("affected", "body", m.Affected); err != nil {
		return err
	}

	for i := 0; i < len(m.Affected); i++ {
		if swag.IsZero(m.Affected[i]) { // not required
			continue
		}

		if m.Affected[i] != nil {
			if err := m.Affected[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("affected" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StoragepoolStatusUnhealthyItem) validateDiskpool(formats strfmt.Registry) error {

	if swag.IsZero(m.Diskpool) { // not required
		return nil
	}

	if m.Diskpool != nil {
		if err := m.Diskpool.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("diskpool")
			}
			return err
		}
	}

	return nil
}

var storagepoolStatusUnhealthyItemHealthFlagsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["underprovisioned","missing_drives","devices_down","devices_smartfailed","waiting_repair"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storagepoolStatusUnhealthyItemHealthFlagsItemsEnum = append(storagepoolStatusUnhealthyItemHealthFlagsItemsEnum, v)
	}
}

func (m *StoragepoolStatusUnhealthyItem) validateHealthFlagsItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, storagepoolStatusUnhealthyItemHealthFlagsItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *StoragepoolStatusUnhealthyItem) validateHealthFlags(formats strfmt.Registry) error {

	if swag.IsZero(m.HealthFlags) { // not required
		return nil
	}

	for i := 0; i < len(m.HealthFlags); i++ {

		// value enum
		if err := m.validateHealthFlagsItemsEnum("health_flags"+"."+strconv.Itoa(i), "body", m.HealthFlags[i]); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StoragepoolStatusUnhealthyItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoragepoolStatusUnhealthyItem) UnmarshalBinary(b []byte) error {
	var res StoragepoolStatusUnhealthyItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}