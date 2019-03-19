// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Zone Specifies properties for access zones.
// swagger:model Zone
type Zone struct {

	// Specifies an alternate system provider.
	AlternateSystemProvider string `json:"alternate_system_provider,omitempty"`

	// Specifies the list of authentication providers available on this access zone.
	AuthProviders []string `json:"auth_providers"`

	// Specifies amount of time in seconds to cache a user/group.
	CacheEntryExpiry int64 `json:"cache_entry_expiry,omitempty"`

	// Determines if a path is created when a path does not exist.
	CreatePath bool `json:"create_path,omitempty"`

	// Allow for overlapping base path.
	ForceOverlap bool `json:"force_overlap,omitempty"`

	// Specifies the permissions set on automatically created user home directories.
	HomeDirectoryUmask int64 `json:"home_directory_umask,omitempty"`

	// Specifies a list of users and groups that have read and write access to /ifs.
	IfsRestricted []*GroupMember `json:"ifs_restricted"`

	// Maps untrusted domains to this NetBIOS domain during authentication.
	MapUntrusted string `json:"map_untrusted,omitempty"`

	// Specifies the access zone name.
	Name string `json:"name,omitempty"`

	// Specifies the NetBIOS name.
	NetbiosName string `json:"netbios_name,omitempty"`

	// Specifies the access zone base directory path.
	Path string `json:"path,omitempty"`

	// Specifies the skeleton directory that is used for user home directories.
	SkeletonDirectory string `json:"skeleton_directory,omitempty"`

	// Specifies the system provider for the access zone.
	SystemProvider string `json:"system_provider,omitempty"`

	// Specifies the current ID mapping rules.
	UserMappingRules []string `json:"user_mapping_rules"`
}

// Validate validates this zone
func (m *Zone) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIfsRestricted(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Zone) validateIfsRestricted(formats strfmt.Registry) error {

	if swag.IsZero(m.IfsRestricted) { // not required
		return nil
	}

	for i := 0; i < len(m.IfsRestricted); i++ {
		if swag.IsZero(m.IfsRestricted[i]) { // not required
			continue
		}

		if m.IfsRestricted[i] != nil {
			if err := m.IfsRestricted[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ifs_restricted" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Zone) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Zone) UnmarshalBinary(b []byte) error {
	var res Zone
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
