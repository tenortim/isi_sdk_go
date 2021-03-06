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

// StoragepoolNodepoolExtended storagepool nodepool extended
// swagger:model StoragepoolNodepoolExtended
type StoragepoolNodepoolExtended struct {

	// Indicates if enabling L3 is possible. L3 cannot be enabled if there are unprovisioned drives.
	// Required: true
	CanEnableL3 *bool `json:"can_enable_l3"`

	// An array of containing any health issues with this pool.  If the pool is healthy, the list is empty.
	HealthFlags []string `json:"health_flags"`

	// The system ID given to the node pool.
	// Required: true
	ID *int64 `json:"id"`

	// Use SSDs in this node pool for L3 cache.
	// Required: true
	L3 *bool `json:"l3"`

	// 'storage' if the 'l3' option is disabled. If the l3 option is enabled, 'migrating' if any SSDs in this node pool have not yet been migrated to L3. If all SSDs have been migrated, 'l3'.
	// Required: true
	// Enum: [l3 storage migrating]
	L3Status *string `json:"l3_status"`

	// The nodes that are part of this node pool.
	// Required: true
	Lnns []int64 `json:"lnns"`

	// Whether or not the node pool is manually managed.
	// Required: true
	Manual *bool `json:"manual"`

	// The node pool name.
	// Required: true
	Name *string `json:"name"`

	// The underlying protection policy.
	ProtectionPolicy string `json:"protection_policy,omitempty"`

	// The name (if named) or system ID of the node pool's tier, if it is in a tier. Otherwise null.
	Tier string `json:"tier,omitempty"`

	// Total pool usage.
	// Required: true
	Usage *StoragepoolTierUsage `json:"usage"`
}

// Validate validates this storagepool nodepool extended
func (m *StoragepoolNodepoolExtended) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCanEnableL3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthFlags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateL3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateL3Status(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLnns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManual(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoragepoolNodepoolExtended) validateCanEnableL3(formats strfmt.Registry) error {

	if err := validate.Required("can_enable_l3", "body", m.CanEnableL3); err != nil {
		return err
	}

	return nil
}

var storagepoolNodepoolExtendedHealthFlagsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["underprovisioned","missing_drives","devices_down","devices_smartfailed","waiting_repair"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storagepoolNodepoolExtendedHealthFlagsItemsEnum = append(storagepoolNodepoolExtendedHealthFlagsItemsEnum, v)
	}
}

func (m *StoragepoolNodepoolExtended) validateHealthFlagsItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, storagepoolNodepoolExtendedHealthFlagsItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *StoragepoolNodepoolExtended) validateHealthFlags(formats strfmt.Registry) error {

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

func (m *StoragepoolNodepoolExtended) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *StoragepoolNodepoolExtended) validateL3(formats strfmt.Registry) error {

	if err := validate.Required("l3", "body", m.L3); err != nil {
		return err
	}

	return nil
}

var storagepoolNodepoolExtendedTypeL3StatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["l3","storage","migrating"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storagepoolNodepoolExtendedTypeL3StatusPropEnum = append(storagepoolNodepoolExtendedTypeL3StatusPropEnum, v)
	}
}

const (

	// StoragepoolNodepoolExtendedL3StatusL3 captures enum value "l3"
	StoragepoolNodepoolExtendedL3StatusL3 string = "l3"

	// StoragepoolNodepoolExtendedL3StatusStorage captures enum value "storage"
	StoragepoolNodepoolExtendedL3StatusStorage string = "storage"

	// StoragepoolNodepoolExtendedL3StatusMigrating captures enum value "migrating"
	StoragepoolNodepoolExtendedL3StatusMigrating string = "migrating"
)

// prop value enum
func (m *StoragepoolNodepoolExtended) validateL3StatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, storagepoolNodepoolExtendedTypeL3StatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StoragepoolNodepoolExtended) validateL3Status(formats strfmt.Registry) error {

	if err := validate.Required("l3_status", "body", m.L3Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateL3StatusEnum("l3_status", "body", *m.L3Status); err != nil {
		return err
	}

	return nil
}

func (m *StoragepoolNodepoolExtended) validateLnns(formats strfmt.Registry) error {

	if err := validate.Required("lnns", "body", m.Lnns); err != nil {
		return err
	}

	return nil
}

func (m *StoragepoolNodepoolExtended) validateManual(formats strfmt.Registry) error {

	if err := validate.Required("manual", "body", m.Manual); err != nil {
		return err
	}

	return nil
}

func (m *StoragepoolNodepoolExtended) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *StoragepoolNodepoolExtended) validateUsage(formats strfmt.Registry) error {

	if err := validate.Required("usage", "body", m.Usage); err != nil {
		return err
	}

	if m.Usage != nil {
		if err := m.Usage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StoragepoolNodepoolExtended) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoragepoolNodepoolExtended) UnmarshalBinary(b []byte) error {
	var res StoragepoolNodepoolExtended
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
