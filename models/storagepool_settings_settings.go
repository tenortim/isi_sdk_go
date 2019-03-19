// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StoragepoolSettingsSettings storagepool settings settings
// swagger:model StoragepoolSettingsSettings
type StoragepoolSettingsSettings struct {

	// Automatically manage IO optimization settings on files.
	// Required: true
	// Enum: [all files_at_default none]
	AutomaticallyManageIoOptimization *string `json:"automatically_manage_io_optimization"`

	// Automatically manage protection settings on files.
	// Required: true
	// Enum: [all files_at_default none]
	AutomaticallyManageProtection *string `json:"automatically_manage_protection"`

	// Optimize namespace operations by storing metadata on SSDs.
	// Required: true
	GlobalNamespaceAccelerationEnabled *bool `json:"global_namespace_acceleration_enabled"`

	// Whether or not namespace operation optimizations are currently in effect.
	// Required: true
	// Enum: [honored inactive]
	GlobalNamespaceAccelerationState *string `json:"global_namespace_acceleration_state"`

	// Automatically add additional protection level to all directories.
	// Required: true
	ProtectDirectoriesOneLevelHigher *bool `json:"protect_directories_one_level_higher"`

	// Spill writes into other pools as needed.
	// Required: true
	SpilloverEnabled *bool `json:"spillover_enabled"`

	// Target pool for spilled writes.
	// Required: true
	SpilloverTarget *StoragepoolSettingsSettingsSpilloverTarget `json:"spillover_target"`

	// The L3 Cache default enabled state. This specifies whether L3 Cache should be enabled on new node pools.
	// Required: true
	SsdL3CacheDefaultEnabled *bool `json:"ssd_l3_cache_default_enabled"`

	// Deny writes into reserved virtual hot spare space.
	// Required: true
	VirtualHotSpareDenyWrites *bool `json:"virtual_hot_spare_deny_writes"`

	// Hide reserved virtual hot spare space from free space counts.
	// Required: true
	VirtualHotSpareHideSpare *bool `json:"virtual_hot_spare_hide_spare"`

	// The number of drives to reserve for the virtual hot spare, from 0-4.
	// Required: true
	// Maximum: 4
	// Minimum: 0
	VirtualHotSpareLimitDrives *int64 `json:"virtual_hot_spare_limit_drives"`

	// The percent space to reserve for the virtual hot spare, from 0-20.
	// Required: true
	// Maximum: 20
	// Minimum: 0
	VirtualHotSpareLimitPercent *int64 `json:"virtual_hot_spare_limit_percent"`
}

// Validate validates this storagepool settings settings
func (m *StoragepoolSettingsSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutomaticallyManageIoOptimization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutomaticallyManageProtection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGlobalNamespaceAccelerationEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGlobalNamespaceAccelerationState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtectDirectoriesOneLevelHigher(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpilloverEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpilloverTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSsdL3CacheDefaultEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualHotSpareDenyWrites(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualHotSpareHideSpare(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualHotSpareLimitDrives(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualHotSpareLimitPercent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var storagepoolSettingsSettingsTypeAutomaticallyManageIoOptimizationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["all","files_at_default","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storagepoolSettingsSettingsTypeAutomaticallyManageIoOptimizationPropEnum = append(storagepoolSettingsSettingsTypeAutomaticallyManageIoOptimizationPropEnum, v)
	}
}

const (

	// StoragepoolSettingsSettingsAutomaticallyManageIoOptimizationAll captures enum value "all"
	StoragepoolSettingsSettingsAutomaticallyManageIoOptimizationAll string = "all"

	// StoragepoolSettingsSettingsAutomaticallyManageIoOptimizationFilesAtDefault captures enum value "files_at_default"
	StoragepoolSettingsSettingsAutomaticallyManageIoOptimizationFilesAtDefault string = "files_at_default"

	// StoragepoolSettingsSettingsAutomaticallyManageIoOptimizationNone captures enum value "none"
	StoragepoolSettingsSettingsAutomaticallyManageIoOptimizationNone string = "none"
)

// prop value enum
func (m *StoragepoolSettingsSettings) validateAutomaticallyManageIoOptimizationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, storagepoolSettingsSettingsTypeAutomaticallyManageIoOptimizationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StoragepoolSettingsSettings) validateAutomaticallyManageIoOptimization(formats strfmt.Registry) error {

	if err := validate.Required("automatically_manage_io_optimization", "body", m.AutomaticallyManageIoOptimization); err != nil {
		return err
	}

	// value enum
	if err := m.validateAutomaticallyManageIoOptimizationEnum("automatically_manage_io_optimization", "body", *m.AutomaticallyManageIoOptimization); err != nil {
		return err
	}

	return nil
}

var storagepoolSettingsSettingsTypeAutomaticallyManageProtectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["all","files_at_default","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storagepoolSettingsSettingsTypeAutomaticallyManageProtectionPropEnum = append(storagepoolSettingsSettingsTypeAutomaticallyManageProtectionPropEnum, v)
	}
}

const (

	// StoragepoolSettingsSettingsAutomaticallyManageProtectionAll captures enum value "all"
	StoragepoolSettingsSettingsAutomaticallyManageProtectionAll string = "all"

	// StoragepoolSettingsSettingsAutomaticallyManageProtectionFilesAtDefault captures enum value "files_at_default"
	StoragepoolSettingsSettingsAutomaticallyManageProtectionFilesAtDefault string = "files_at_default"

	// StoragepoolSettingsSettingsAutomaticallyManageProtectionNone captures enum value "none"
	StoragepoolSettingsSettingsAutomaticallyManageProtectionNone string = "none"
)

// prop value enum
func (m *StoragepoolSettingsSettings) validateAutomaticallyManageProtectionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, storagepoolSettingsSettingsTypeAutomaticallyManageProtectionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StoragepoolSettingsSettings) validateAutomaticallyManageProtection(formats strfmt.Registry) error {

	if err := validate.Required("automatically_manage_protection", "body", m.AutomaticallyManageProtection); err != nil {
		return err
	}

	// value enum
	if err := m.validateAutomaticallyManageProtectionEnum("automatically_manage_protection", "body", *m.AutomaticallyManageProtection); err != nil {
		return err
	}

	return nil
}

func (m *StoragepoolSettingsSettings) validateGlobalNamespaceAccelerationEnabled(formats strfmt.Registry) error {

	if err := validate.Required("global_namespace_acceleration_enabled", "body", m.GlobalNamespaceAccelerationEnabled); err != nil {
		return err
	}

	return nil
}

var storagepoolSettingsSettingsTypeGlobalNamespaceAccelerationStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["honored","inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storagepoolSettingsSettingsTypeGlobalNamespaceAccelerationStatePropEnum = append(storagepoolSettingsSettingsTypeGlobalNamespaceAccelerationStatePropEnum, v)
	}
}

const (

	// StoragepoolSettingsSettingsGlobalNamespaceAccelerationStateHonored captures enum value "honored"
	StoragepoolSettingsSettingsGlobalNamespaceAccelerationStateHonored string = "honored"

	// StoragepoolSettingsSettingsGlobalNamespaceAccelerationStateInactive captures enum value "inactive"
	StoragepoolSettingsSettingsGlobalNamespaceAccelerationStateInactive string = "inactive"
)

// prop value enum
func (m *StoragepoolSettingsSettings) validateGlobalNamespaceAccelerationStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, storagepoolSettingsSettingsTypeGlobalNamespaceAccelerationStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StoragepoolSettingsSettings) validateGlobalNamespaceAccelerationState(formats strfmt.Registry) error {

	if err := validate.Required("global_namespace_acceleration_state", "body", m.GlobalNamespaceAccelerationState); err != nil {
		return err
	}

	// value enum
	if err := m.validateGlobalNamespaceAccelerationStateEnum("global_namespace_acceleration_state", "body", *m.GlobalNamespaceAccelerationState); err != nil {
		return err
	}

	return nil
}

func (m *StoragepoolSettingsSettings) validateProtectDirectoriesOneLevelHigher(formats strfmt.Registry) error {

	if err := validate.Required("protect_directories_one_level_higher", "body", m.ProtectDirectoriesOneLevelHigher); err != nil {
		return err
	}

	return nil
}

func (m *StoragepoolSettingsSettings) validateSpilloverEnabled(formats strfmt.Registry) error {

	if err := validate.Required("spillover_enabled", "body", m.SpilloverEnabled); err != nil {
		return err
	}

	return nil
}

func (m *StoragepoolSettingsSettings) validateSpilloverTarget(formats strfmt.Registry) error {

	if err := validate.Required("spillover_target", "body", m.SpilloverTarget); err != nil {
		return err
	}

	if m.SpilloverTarget != nil {
		if err := m.SpilloverTarget.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spillover_target")
			}
			return err
		}
	}

	return nil
}

func (m *StoragepoolSettingsSettings) validateSsdL3CacheDefaultEnabled(formats strfmt.Registry) error {

	if err := validate.Required("ssd_l3_cache_default_enabled", "body", m.SsdL3CacheDefaultEnabled); err != nil {
		return err
	}

	return nil
}

func (m *StoragepoolSettingsSettings) validateVirtualHotSpareDenyWrites(formats strfmt.Registry) error {

	if err := validate.Required("virtual_hot_spare_deny_writes", "body", m.VirtualHotSpareDenyWrites); err != nil {
		return err
	}

	return nil
}

func (m *StoragepoolSettingsSettings) validateVirtualHotSpareHideSpare(formats strfmt.Registry) error {

	if err := validate.Required("virtual_hot_spare_hide_spare", "body", m.VirtualHotSpareHideSpare); err != nil {
		return err
	}

	return nil
}

func (m *StoragepoolSettingsSettings) validateVirtualHotSpareLimitDrives(formats strfmt.Registry) error {

	if err := validate.Required("virtual_hot_spare_limit_drives", "body", m.VirtualHotSpareLimitDrives); err != nil {
		return err
	}

	if err := validate.MinimumInt("virtual_hot_spare_limit_drives", "body", int64(*m.VirtualHotSpareLimitDrives), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("virtual_hot_spare_limit_drives", "body", int64(*m.VirtualHotSpareLimitDrives), 4, false); err != nil {
		return err
	}

	return nil
}

func (m *StoragepoolSettingsSettings) validateVirtualHotSpareLimitPercent(formats strfmt.Registry) error {

	if err := validate.Required("virtual_hot_spare_limit_percent", "body", m.VirtualHotSpareLimitPercent); err != nil {
		return err
	}

	if err := validate.MinimumInt("virtual_hot_spare_limit_percent", "body", int64(*m.VirtualHotSpareLimitPercent), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("virtual_hot_spare_limit_percent", "body", int64(*m.VirtualHotSpareLimitPercent), 20, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StoragepoolSettingsSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoragepoolSettingsSettings) UnmarshalBinary(b []byte) error {
	var res StoragepoolSettingsSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}