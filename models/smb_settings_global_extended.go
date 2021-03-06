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

// SmbSettingsGlobalExtended smb settings global extended
// swagger:model SmbSettingsGlobalExtended
type SmbSettingsGlobalExtended struct {

	// Only enumerate files and folders the requesting user has access to.
	AccessBasedShareEnum bool `json:"access_based_share_enum,omitempty"`

	// Specify level of file share audit events to log.
	AuditFileshare string `json:"audit_fileshare,omitempty"`

	// Specifies a list of permissions to audit.
	AuditGlobalSacl []*SmbSettingsGlobalSettingsAuditGlobalSaclItem `json:"audit_global_sacl"`

	// Specify the level of logon audit events to log.
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

// Validate validates this smb settings global extended
func (m *SmbSettingsGlobalExtended) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuditGlobalSacl(formats); err != nil {
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

func (m *SmbSettingsGlobalExtended) validateAuditGlobalSacl(formats strfmt.Registry) error {

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

func (m *SmbSettingsGlobalExtended) validateOnefsCPUMultiplier(formats strfmt.Registry) error {

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

func (m *SmbSettingsGlobalExtended) validateOnefsNumWorkers(formats strfmt.Registry) error {

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

func (m *SmbSettingsGlobalExtended) validateSrvCPUMultiplier(formats strfmt.Registry) error {

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

func (m *SmbSettingsGlobalExtended) validateSrvNumWorkers(formats strfmt.Registry) error {

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
func (m *SmbSettingsGlobalExtended) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SmbSettingsGlobalExtended) UnmarshalBinary(b []byte) error {
	var res SmbSettingsGlobalExtended
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
