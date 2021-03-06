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

// SmbSettingsGlobalSettings smb settings global settings
// swagger:model SmbSettingsGlobalSettings
type SmbSettingsGlobalSettings struct {

	// Only enumerate files and folders the requesting user has access to.
	AccessBasedShareEnum bool `json:"access_based_share_enum,omitempty"`

	// Specify level of file share audit events to log.
	// Enum: [all success failure none]
	AuditFileshare string `json:"audit_fileshare,omitempty"`

	// Specifies a list of permissions to audit.
	AuditGlobalSacl []*SmbSettingsGlobalSettingsAuditGlobalSaclItem `json:"audit_global_sacl"`

	// Specify the level of logon audit events to log.
	// Enum: [all success failure none]
	AuditLogon string `json:"audit_logon,omitempty"`

	// Allow access to .snapshot directories in share subdirectories.
	DotSnapAccessibleChild bool `json:"dot_snap_accessible_child,omitempty"`

	// Allow access to the .snapshot directory in the root of the share.
	DotSnapAccessibleRoot bool `json:"dot_snap_accessible_root,omitempty"`

	// Show .snapshot directories in share subdirectories.
	DotSnapVisibleChild bool `json:"dot_snap_visible_child,omitempty"`

	// Show the .snapshot directory in the root of a share.
	DotSnapVisibleRoot bool `json:"dot_snap_visible_root,omitempty"`

	// Indicates whether the server supports signed SMB packets.
	EnableSecuritySignatures bool `json:"enable_security_signatures,omitempty"`

	// Specifies the fully-qualified user to use for guest access.
	GuestUser string `json:"guest_user,omitempty"`

	// Specify whether to ignore EAs on files.
	IgnoreEas bool `json:"ignore_eas,omitempty"`

	// Specify the number of OneFS driver worker threads per CPU.
	// Maximum: 4
	// Minimum: 1
	OnefsCPUMultiplier int64 `json:"onefs_cpu_multiplier,omitempty"`

	// Set the maximum number of OneFS driver worker threads.
	// Maximum: 1024
	// Minimum: 0
	OnefsNumWorkers *int64 `json:"onefs_num_workers,omitempty"`

	// Indicates whether the server requires signed SMB packets.
	RequireSecuritySignatures bool `json:"require_security_signatures,omitempty"`

	// Enable Server Side Copy.
	ServerSideCopy bool `json:"server_side_copy,omitempty"`

	// Provides a description of the server.
	ServerString string `json:"server_string,omitempty"`

	// Specify whether service is enabled.
	Service bool `json:"service,omitempty"`

	// Specify the number of SRV service worker threads per CPU.
	// Maximum: 8
	// Minimum: 1
	SrvCPUMultiplier int64 `json:"srv_cpu_multiplier,omitempty"`

	// Set the maximum number of SRV service worker threads.
	// Maximum: 1024
	// Minimum: 0
	SrvNumWorkers *int64 `json:"srv_num_workers,omitempty"`

	// Support multichannel.
	SupportMultichannel bool `json:"support_multichannel,omitempty"`

	// Support NetBIOS.
	SupportNetbios bool `json:"support_netbios,omitempty"`

	// Support the SMB2 protocol on the server.
	SupportSmb2 bool `json:"support_smb2,omitempty"`
}

// Validate validates this smb settings global settings
func (m *SmbSettingsGlobalSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuditFileshare(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuditGlobalSacl(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuditLogon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnefsCPUMultiplier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnefsNumWorkers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrvCPUMultiplier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrvNumWorkers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var smbSettingsGlobalSettingsTypeAuditFilesharePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["all","success","failure","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		smbSettingsGlobalSettingsTypeAuditFilesharePropEnum = append(smbSettingsGlobalSettingsTypeAuditFilesharePropEnum, v)
	}
}

const (

	// SmbSettingsGlobalSettingsAuditFileshareAll captures enum value "all"
	SmbSettingsGlobalSettingsAuditFileshareAll string = "all"

	// SmbSettingsGlobalSettingsAuditFileshareSuccess captures enum value "success"
	SmbSettingsGlobalSettingsAuditFileshareSuccess string = "success"

	// SmbSettingsGlobalSettingsAuditFileshareFailure captures enum value "failure"
	SmbSettingsGlobalSettingsAuditFileshareFailure string = "failure"

	// SmbSettingsGlobalSettingsAuditFileshareNone captures enum value "none"
	SmbSettingsGlobalSettingsAuditFileshareNone string = "none"
)

// prop value enum
func (m *SmbSettingsGlobalSettings) validateAuditFileshareEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, smbSettingsGlobalSettingsTypeAuditFilesharePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SmbSettingsGlobalSettings) validateAuditFileshare(formats strfmt.Registry) error {

	if swag.IsZero(m.AuditFileshare) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuditFileshareEnum("audit_fileshare", "body", m.AuditFileshare); err != nil {
		return err
	}

	return nil
}

func (m *SmbSettingsGlobalSettings) validateAuditGlobalSacl(formats strfmt.Registry) error {

	if swag.IsZero(m.AuditGlobalSacl) { // not required
		return nil
	}

	for i := 0; i < len(m.AuditGlobalSacl); i++ {
		if swag.IsZero(m.AuditGlobalSacl[i]) { // not required
			continue
		}

		if m.AuditGlobalSacl[i] != nil {
			if err := m.AuditGlobalSacl[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("audit_global_sacl" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var smbSettingsGlobalSettingsTypeAuditLogonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["all","success","failure","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		smbSettingsGlobalSettingsTypeAuditLogonPropEnum = append(smbSettingsGlobalSettingsTypeAuditLogonPropEnum, v)
	}
}

const (

	// SmbSettingsGlobalSettingsAuditLogonAll captures enum value "all"
	SmbSettingsGlobalSettingsAuditLogonAll string = "all"

	// SmbSettingsGlobalSettingsAuditLogonSuccess captures enum value "success"
	SmbSettingsGlobalSettingsAuditLogonSuccess string = "success"

	// SmbSettingsGlobalSettingsAuditLogonFailure captures enum value "failure"
	SmbSettingsGlobalSettingsAuditLogonFailure string = "failure"

	// SmbSettingsGlobalSettingsAuditLogonNone captures enum value "none"
	SmbSettingsGlobalSettingsAuditLogonNone string = "none"
)

// prop value enum
func (m *SmbSettingsGlobalSettings) validateAuditLogonEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, smbSettingsGlobalSettingsTypeAuditLogonPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SmbSettingsGlobalSettings) validateAuditLogon(formats strfmt.Registry) error {

	if swag.IsZero(m.AuditLogon) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuditLogonEnum("audit_logon", "body", m.AuditLogon); err != nil {
		return err
	}

	return nil
}

func (m *SmbSettingsGlobalSettings) validateOnefsCPUMultiplier(formats strfmt.Registry) error {

	if swag.IsZero(m.OnefsCPUMultiplier) { // not required
		return nil
	}

	if err := validate.MinimumInt("onefs_cpu_multiplier", "body", int64(m.OnefsCPUMultiplier), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("onefs_cpu_multiplier", "body", int64(m.OnefsCPUMultiplier), 4, false); err != nil {
		return err
	}

	return nil
}

func (m *SmbSettingsGlobalSettings) validateOnefsNumWorkers(formats strfmt.Registry) error {

	if swag.IsZero(m.OnefsNumWorkers) { // not required
		return nil
	}

	if err := validate.MinimumInt("onefs_num_workers", "body", int64(*m.OnefsNumWorkers), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("onefs_num_workers", "body", int64(*m.OnefsNumWorkers), 1024, false); err != nil {
		return err
	}

	return nil
}

func (m *SmbSettingsGlobalSettings) validateSrvCPUMultiplier(formats strfmt.Registry) error {

	if swag.IsZero(m.SrvCPUMultiplier) { // not required
		return nil
	}

	if err := validate.MinimumInt("srv_cpu_multiplier", "body", int64(m.SrvCPUMultiplier), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("srv_cpu_multiplier", "body", int64(m.SrvCPUMultiplier), 8, false); err != nil {
		return err
	}

	return nil
}

func (m *SmbSettingsGlobalSettings) validateSrvNumWorkers(formats strfmt.Registry) error {

	if swag.IsZero(m.SrvNumWorkers) { // not required
		return nil
	}

	if err := validate.MinimumInt("srv_num_workers", "body", int64(*m.SrvNumWorkers), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("srv_num_workers", "body", int64(*m.SrvNumWorkers), 1024, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SmbSettingsGlobalSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SmbSettingsGlobalSettings) UnmarshalBinary(b []byte) error {
	var res SmbSettingsGlobalSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
