// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StoragepoolTierUsage storagepool tier usage
// swagger:model StoragepoolTierUsage
type StoragepoolTierUsage struct {

	// Available free bytes remaining in the pool when virtual hot spare is taken into account.
	// Required: true
	AvailBytes *string `json:"avail_bytes"`

	// Available free bytes remaining in the pool on SSD drives when virtual hot spare is taken into account.
	// Required: true
	AvailSsdBytes *string `json:"avail_ssd_bytes"`

	// Whether or not the pool usage is currently balanced.
	// Required: true
	Balanced *bool `json:"balanced"`

	// Free bytes remaining in the pool.
	// Required: true
	FreeBytes *string `json:"free_bytes"`

	// Free bytes remaining in the pool on SSD drives.
	// Required: true
	FreeSsdBytes *string `json:"free_ssd_bytes"`

	// Total bytes used in the pool.
	// Required: true
	TotalBytes *string `json:"total_bytes"`

	// Total bytes used in the pool on SSD drives.
	// Required: true
	TotalSsdBytes *string `json:"total_ssd_bytes"`

	// Bytes reserved for virtual hot spare in the pool.
	// Required: true
	VirtualHotSpareBytes *string `json:"virtual_hot_spare_bytes"`
}

// Validate validates this storagepool tier usage
func (m *StoragepoolTierUsage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailBytes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvailSsdBytes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBalanced(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFreeBytes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFreeSsdBytes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalBytes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalSsdBytes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualHotSpareBytes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoragepoolTierUsage) validateAvailBytes(formats strfmt.Registry) error {

	if err := validate.Required("avail_bytes", "body", m.AvailBytes); err != nil {
		return err
	}

	return nil
}

func (m *StoragepoolTierUsage) validateAvailSsdBytes(formats strfmt.Registry) error {

	if err := validate.Required("avail_ssd_bytes", "body", m.AvailSsdBytes); err != nil {
		return err
	}

	return nil
}

func (m *StoragepoolTierUsage) validateBalanced(formats strfmt.Registry) error {

	if err := validate.Required("balanced", "body", m.Balanced); err != nil {
		return err
	}

	return nil
}

func (m *StoragepoolTierUsage) validateFreeBytes(formats strfmt.Registry) error {

	if err := validate.Required("free_bytes", "body", m.FreeBytes); err != nil {
		return err
	}

	return nil
}

func (m *StoragepoolTierUsage) validateFreeSsdBytes(formats strfmt.Registry) error {

	if err := validate.Required("free_ssd_bytes", "body", m.FreeSsdBytes); err != nil {
		return err
	}

	return nil
}

func (m *StoragepoolTierUsage) validateTotalBytes(formats strfmt.Registry) error {

	if err := validate.Required("total_bytes", "body", m.TotalBytes); err != nil {
		return err
	}

	return nil
}

func (m *StoragepoolTierUsage) validateTotalSsdBytes(formats strfmt.Registry) error {

	if err := validate.Required("total_ssd_bytes", "body", m.TotalSsdBytes); err != nil {
		return err
	}

	return nil
}

func (m *StoragepoolTierUsage) validateVirtualHotSpareBytes(formats strfmt.Registry) error {

	if err := validate.Required("virtual_hot_spare_bytes", "body", m.VirtualHotSpareBytes); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StoragepoolTierUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoragepoolTierUsage) UnmarshalBinary(b []byte) error {
	var res StoragepoolTierUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
