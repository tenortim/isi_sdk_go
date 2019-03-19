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

// SmbShareExtended smb share extended
// swagger:model SmbShareExtended
type SmbShareExtended struct {

	// Only enumerate files and folders the requesting user has access to.
	AccessBasedEnumeration bool `json:"access_based_enumeration,omitempty"`

	// Access-based enumeration on only the root directory of the share.
	AccessBasedEnumerationRootOnly bool `json:"access_based_enumeration_root_only,omitempty"`

	// Allow deletion of read-only files in the share.
	AllowDeleteReadonly bool `json:"allow_delete_readonly,omitempty"`

	// Allows users to execute files they have read rights for.
	AllowExecuteAlways bool `json:"allow_execute_always,omitempty"`

	// Allow automatic expansion of variables for home directories.
	AllowVariableExpansion bool `json:"allow_variable_expansion,omitempty"`

	// Automatically create home directories.
	AutoCreateDirectory bool `json:"auto_create_directory,omitempty"`

	// Share is visible in net view and the browse list.
	Browsable bool `json:"browsable,omitempty"`

	// Persistent open timeout for the share.
	// Minimum: 2
	CaTimeout int64 `json:"ca_timeout,omitempty"`

	// Specify the level of write-integrity on continuously available shares.
	// Enum: [none write-read-coherent full]
	CaWriteIntegrity string `json:"ca_write_integrity,omitempty"`

	// Level of change notification alerts on the share.
	// Enum: [all norecurse none]
	ChangeNotify string `json:"change_notify,omitempty"`

	// Specify if persistent opens are allowed on the share.
	ContinuouslyAvailable bool `json:"continuously_available,omitempty"`

	// Create permissions for new files and directories in share.
	// Enum: [default acl inherit mode bits use create mask and mode]
	CreatePermissions string `json:"create_permissions,omitempty"`

	// Client-side caching policy for the shares.
	// Enum: [manual documents programs none]
	CscPolicy string `json:"csc_policy,omitempty"`

	// Description for this SMB share.
	Description string `json:"description,omitempty"`

	// Directory create mask bits.
	DirectoryCreateMask int64 `json:"directory_create_mask,omitempty"`

	// Directory create mode bits.
	DirectoryCreateMode int64 `json:"directory_create_mode,omitempty"`

	// File create mask bits.
	FileCreateMask int64 `json:"file_create_mask,omitempty"`

	// File create mode bits.
	FileCreateMode int64 `json:"file_create_mode,omitempty"`

	// Specifies the list of file extensions.
	FileFilterExtensions []string `json:"file_filter_extensions"`

	// Specifies if filter list is for deny or allow. Default is deny.
	// Enum: [deny allow]
	FileFilterType string `json:"file_filter_type,omitempty"`

	// Enables file filtering on this zone.
	FileFilteringEnabled bool `json:"file_filtering_enabled,omitempty"`

	// Hide files and directories that begin with a period '.'.
	HideDotFiles bool `json:"hide_dot_files,omitempty"`

	// An ACL expressing which hosts are allowed access. A deny clause must be the final entry.
	HostACL []string `json:"host_acl"`

	// Share ID.
	ID string `json:"id,omitempty"`

	// Specify the condition in which user access is done as the guest account.
	// Enum: [always bad user never]
	ImpersonateGuest string `json:"impersonate_guest,omitempty"`

	// User account to be used as guest account.
	ImpersonateUser string `json:"impersonate_user,omitempty"`

	// Set the inheritable ACL on the share path.
	InheritablePathACL bool `json:"inheritable_path_acl,omitempty"`

	// Specifies the wchar_t starting point for automatic byte mangling.
	MangleByteStart int64 `json:"mangle_byte_start,omitempty"`

	// Character mangle map.
	MangleMap []string `json:"mangle_map"`

	// Share name.
	Name string `json:"name,omitempty"`

	// Support NTFS ACLs on files and directories.
	NtfsACLSupport bool `json:"ntfs_acl_support,omitempty"`

	// Support oplocks.
	Oplocks bool `json:"oplocks,omitempty"`

	// Path of share within /ifs.
	Path string `json:"path,omitempty"`

	// Specifies an ordered list of permission modifications.
	Permissions []*SmbSharePermission `json:"permissions"`

	// Allow account to run as root.
	RunAsRoot []*GroupMember `json:"run_as_root"`

	// Specifies if persistent opens would do strict lockout on the share.
	StrictCaLockout bool `json:"strict_ca_lockout,omitempty"`

	// Handle SMB flush operations.
	StrictFlush bool `json:"strict_flush,omitempty"`

	// Specifies whether byte range locks contend against SMB I/O.
	StrictLocking bool `json:"strict_locking,omitempty"`

	// Numeric ID of the access zone which contains this SMB share
	// Required: true
	Zid *int64 `json:"zid"`
}

// Validate validates this smb share extended
func (m *SmbShareExtended) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCaTimeout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCaWriteIntegrity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChangeNotify(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCscPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileFilterType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImpersonateGuest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunAsRoot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZid(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SmbShareExtended) validateCaTimeout(formats strfmt.Registry) error {

	if swag.IsZero(m.CaTimeout) { // not required
		return nil
	}

	if err := validate.MinimumInt("ca_timeout", "body", int64(m.CaTimeout), 2, false); err != nil {
		return err
	}

	return nil
}

var smbShareExtendedTypeCaWriteIntegrityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","write-read-coherent","full"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		smbShareExtendedTypeCaWriteIntegrityPropEnum = append(smbShareExtendedTypeCaWriteIntegrityPropEnum, v)
	}
}

const (

	// SmbShareExtendedCaWriteIntegrityNone captures enum value "none"
	SmbShareExtendedCaWriteIntegrityNone string = "none"

	// SmbShareExtendedCaWriteIntegrityWriteReadCoherent captures enum value "write-read-coherent"
	SmbShareExtendedCaWriteIntegrityWriteReadCoherent string = "write-read-coherent"

	// SmbShareExtendedCaWriteIntegrityFull captures enum value "full"
	SmbShareExtendedCaWriteIntegrityFull string = "full"
)

// prop value enum
func (m *SmbShareExtended) validateCaWriteIntegrityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, smbShareExtendedTypeCaWriteIntegrityPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SmbShareExtended) validateCaWriteIntegrity(formats strfmt.Registry) error {

	if swag.IsZero(m.CaWriteIntegrity) { // not required
		return nil
	}

	// value enum
	if err := m.validateCaWriteIntegrityEnum("ca_write_integrity", "body", m.CaWriteIntegrity); err != nil {
		return err
	}

	return nil
}

var smbShareExtendedTypeChangeNotifyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["all","norecurse","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		smbShareExtendedTypeChangeNotifyPropEnum = append(smbShareExtendedTypeChangeNotifyPropEnum, v)
	}
}

const (

	// SmbShareExtendedChangeNotifyAll captures enum value "all"
	SmbShareExtendedChangeNotifyAll string = "all"

	// SmbShareExtendedChangeNotifyNorecurse captures enum value "norecurse"
	SmbShareExtendedChangeNotifyNorecurse string = "norecurse"

	// SmbShareExtendedChangeNotifyNone captures enum value "none"
	SmbShareExtendedChangeNotifyNone string = "none"
)

// prop value enum
func (m *SmbShareExtended) validateChangeNotifyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, smbShareExtendedTypeChangeNotifyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SmbShareExtended) validateChangeNotify(formats strfmt.Registry) error {

	if swag.IsZero(m.ChangeNotify) { // not required
		return nil
	}

	// value enum
	if err := m.validateChangeNotifyEnum("change_notify", "body", m.ChangeNotify); err != nil {
		return err
	}

	return nil
}

var smbShareExtendedTypeCreatePermissionsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["default acl","inherit mode bits","use create mask and mode"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		smbShareExtendedTypeCreatePermissionsPropEnum = append(smbShareExtendedTypeCreatePermissionsPropEnum, v)
	}
}

const (

	// SmbShareExtendedCreatePermissionsDefaultACL captures enum value "default acl"
	SmbShareExtendedCreatePermissionsDefaultACL string = "default acl"

	// SmbShareExtendedCreatePermissionsInheritModeBits captures enum value "inherit mode bits"
	SmbShareExtendedCreatePermissionsInheritModeBits string = "inherit mode bits"

	// SmbShareExtendedCreatePermissionsUseCreateMaskAndMode captures enum value "use create mask and mode"
	SmbShareExtendedCreatePermissionsUseCreateMaskAndMode string = "use create mask and mode"
)

// prop value enum
func (m *SmbShareExtended) validateCreatePermissionsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, smbShareExtendedTypeCreatePermissionsPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SmbShareExtended) validateCreatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatePermissions) { // not required
		return nil
	}

	// value enum
	if err := m.validateCreatePermissionsEnum("create_permissions", "body", m.CreatePermissions); err != nil {
		return err
	}

	return nil
}

var smbShareExtendedTypeCscPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["manual","documents","programs","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		smbShareExtendedTypeCscPolicyPropEnum = append(smbShareExtendedTypeCscPolicyPropEnum, v)
	}
}

const (

	// SmbShareExtendedCscPolicyManual captures enum value "manual"
	SmbShareExtendedCscPolicyManual string = "manual"

	// SmbShareExtendedCscPolicyDocuments captures enum value "documents"
	SmbShareExtendedCscPolicyDocuments string = "documents"

	// SmbShareExtendedCscPolicyPrograms captures enum value "programs"
	SmbShareExtendedCscPolicyPrograms string = "programs"

	// SmbShareExtendedCscPolicyNone captures enum value "none"
	SmbShareExtendedCscPolicyNone string = "none"
)

// prop value enum
func (m *SmbShareExtended) validateCscPolicyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, smbShareExtendedTypeCscPolicyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SmbShareExtended) validateCscPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.CscPolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateCscPolicyEnum("csc_policy", "body", m.CscPolicy); err != nil {
		return err
	}

	return nil
}

var smbShareExtendedTypeFileFilterTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["deny","allow"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		smbShareExtendedTypeFileFilterTypePropEnum = append(smbShareExtendedTypeFileFilterTypePropEnum, v)
	}
}

const (

	// SmbShareExtendedFileFilterTypeDeny captures enum value "deny"
	SmbShareExtendedFileFilterTypeDeny string = "deny"

	// SmbShareExtendedFileFilterTypeAllow captures enum value "allow"
	SmbShareExtendedFileFilterTypeAllow string = "allow"
)

// prop value enum
func (m *SmbShareExtended) validateFileFilterTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, smbShareExtendedTypeFileFilterTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SmbShareExtended) validateFileFilterType(formats strfmt.Registry) error {

	if swag.IsZero(m.FileFilterType) { // not required
		return nil
	}

	// value enum
	if err := m.validateFileFilterTypeEnum("file_filter_type", "body", m.FileFilterType); err != nil {
		return err
	}

	return nil
}

var smbShareExtendedTypeImpersonateGuestPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["always","bad user","never"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		smbShareExtendedTypeImpersonateGuestPropEnum = append(smbShareExtendedTypeImpersonateGuestPropEnum, v)
	}
}

const (

	// SmbShareExtendedImpersonateGuestAlways captures enum value "always"
	SmbShareExtendedImpersonateGuestAlways string = "always"

	// SmbShareExtendedImpersonateGuestBadUser captures enum value "bad user"
	SmbShareExtendedImpersonateGuestBadUser string = "bad user"

	// SmbShareExtendedImpersonateGuestNever captures enum value "never"
	SmbShareExtendedImpersonateGuestNever string = "never"
)

// prop value enum
func (m *SmbShareExtended) validateImpersonateGuestEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, smbShareExtendedTypeImpersonateGuestPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SmbShareExtended) validateImpersonateGuest(formats strfmt.Registry) error {

	if swag.IsZero(m.ImpersonateGuest) { // not required
		return nil
	}

	// value enum
	if err := m.validateImpersonateGuestEnum("impersonate_guest", "body", m.ImpersonateGuest); err != nil {
		return err
	}

	return nil
}

func (m *SmbShareExtended) validatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	for i := 0; i < len(m.Permissions); i++ {
		if swag.IsZero(m.Permissions[i]) { // not required
			continue
		}

		if m.Permissions[i] != nil {
			if err := m.Permissions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("permissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SmbShareExtended) validateRunAsRoot(formats strfmt.Registry) error {

	if swag.IsZero(m.RunAsRoot) { // not required
		return nil
	}

	for i := 0; i < len(m.RunAsRoot); i++ {
		if swag.IsZero(m.RunAsRoot[i]) { // not required
			continue
		}

		if m.RunAsRoot[i] != nil {
			if err := m.RunAsRoot[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("run_as_root" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SmbShareExtended) validateZid(formats strfmt.Registry) error {

	if err := validate.Required("zid", "body", m.Zid); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SmbShareExtended) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SmbShareExtended) UnmarshalBinary(b []byte) error {
	var res SmbShareExtended
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}