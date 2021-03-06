// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SettingsAclsACLPolicySettings settings acls Acl policy settings
// swagger:model SettingsAclsAclPolicySettings
type SettingsAclsACLPolicySettings struct {

	// Access checks (chmod, chown).
	Access string `json:"access,omitempty"`

	// Displayed mode bits.
	Calcmode string `json:"calcmode,omitempty"`

	// Approximate group mode bits when ACL exists.
	CalcmodeGroup string `json:"calcmode_group,omitempty"`

	// Approximate owner mode bits when ACL exists.
	CalcmodeOwner string `json:"calcmode_owner,omitempty"`

	// chmod on files with existing ACLs.
	Chmod string `json:"chmod,omitempty"`

	// chmod (007) on files with existing ACLs.
	Chmod007 string `json:"chmod_007,omitempty"`

	// ACLs created on directories by UNIX chmod.
	ChmodInheritable string `json:"chmod_inheritable,omitempty"`

	// chown/chgrp on files with existing ACLs.
	Chown string `json:"chown,omitempty"`

	// ACL creation over SMB.
	CreateOverSmb string `json:"create_over_smb,omitempty"`

	//  Read only DOS attribute.
	DosAttr string `json:"dos_attr,omitempty"`

	// Group owner inheritance.
	GroupOwnerInheritance string `json:"group_owner_inheritance,omitempty"`

	// Treatment of 'rwx' permissions.
	Rwx string `json:"rwx,omitempty"`

	// Synthetic 'deny' ACEs.
	SyntheticDenies string `json:"synthetic_denies,omitempty"`

	// Access check (utimes)
	Utimes string `json:"utimes,omitempty"`
}

// Validate validates this settings acls Acl policy settings
func (m *SettingsAclsACLPolicySettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SettingsAclsACLPolicySettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SettingsAclsACLPolicySettings) UnmarshalBinary(b []byte) error {
	var res SettingsAclsACLPolicySettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
