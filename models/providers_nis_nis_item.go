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

// ProvidersNisNisItem providers nis nis item
// swagger:model ProvidersNisNisItem
type ProvidersNisNisItem struct {

	// If true, enables authentication and identity management through the authentication provider.
	Authentication bool `json:"authentication,omitempty"`

	// If true, connects the provider to a random server.
	BalanceServers bool `json:"balance_servers,omitempty"`

	// Specifies the time in seconds between provider online checks.
	CheckOnlineInterval int64 `json:"check_online_interval,omitempty"`

	// Automatically creates the home directory on the first login.
	CreateHomeDirectory bool `json:"create_home_directory,omitempty"`

	// If true, enables the NIS provider.
	Enabled bool `json:"enabled,omitempty"`

	// If true, allows the provider to enumerate groups.
	EnumerateGroups bool `json:"enumerate_groups,omitempty"`

	// If true, allows the provider to enumerate users.
	EnumerateUsers bool `json:"enumerate_users,omitempty"`

	// Specifies the list of groups that can be resolved.
	FindableGroups []string `json:"findable_groups"`

	// Specifies the list of users that can be resolved.
	FindableUsers []string `json:"findable_users"`

	// Specifies the domain for this provider through which groups are qualified.
	GroupDomain string `json:"group_domain,omitempty"`

	// Specifies the path to the home directory template.
	HomeDirectoryTemplate string `json:"home_directory_template,omitempty"`

	// If true, enables host name look ups.
	HostnameLookup bool `json:"hostname_lookup,omitempty"`

	// Specifies the NIS provider ID.
	ID string `json:"id,omitempty"`

	// Specifies the groups that can be viewed in the provider.
	ListableGroups []string `json:"listable_groups"`

	// Specifies the users that can be viewed in the provider.
	ListableUsers []string `json:"listable_users"`

	// Specifies the login shell path.
	LoginShell string `json:"login_shell,omitempty"`

	// Specifies the NIS provider name.
	Name string `json:"name,omitempty"`

	// Specifies the NIS domain name.
	NisDomain string `json:"nis_domain,omitempty"`

	// Normalizes group names to lowercase before look up.
	NormalizeGroups bool `json:"normalize_groups,omitempty"`

	// Normalizes user names to lowercase before look up.
	NormalizeUsers bool `json:"normalize_users,omitempty"`

	// Specifies which NTLM versions to support for users with NTLM-compatible credentials.
	// Enum: [all v2only none]
	NtlmSupport string `json:"ntlm_support,omitempty"`

	// Specifies the domain for the provider.
	ProviderDomain string `json:"provider_domain,omitempty"`

	// Specifies the request timeout interval in seconds.
	RequestTimeout int64 `json:"request_timeout,omitempty"`

	// If true, checks the provider for filtered lists of findable and unfindable users and groups.
	RestrictFindable bool `json:"restrict_findable,omitempty"`

	// If true, checks the provider for filtered lists of listable and unlistable users and groups.
	RestrictListable bool `json:"restrict_listable,omitempty"`

	// Specifies the timeout period in seconds after which a request will be retried.
	RetryTime int64 `json:"retry_time,omitempty"`

	// Adds an NIS server for this provider.
	Servers []string `json:"servers"`

	// Specifies the status of the provider.
	Status string `json:"status,omitempty"`

	// If true, indicates that this provider instance was created by OneFS and cannot be removed.
	System bool `json:"system,omitempty"`

	// Specifies groups that cannot be resolved by the provider.
	UnfindableGroups []string `json:"unfindable_groups"`

	// Specifies users that cannot be resolved by the provider.
	UnfindableUsers []string `json:"unfindable_users"`

	// Specifies a group that cannot be listed by the provider.
	UnlistableGroups []string `json:"unlistable_groups"`

	// Specifies a user that cannot be listed by the provider.
	UnlistableUsers []string `json:"unlistable_users"`

	// Specifies the domain for this provider through which users are qualified.
	UserDomain string `json:"user_domain,omitempty"`

	// If true, specifies TCP for YP Match operations.
	YpmatchUsingTCP bool `json:"ypmatch_using_tcp,omitempty"`
}

// Validate validates this providers nis nis item
func (m *ProvidersNisNisItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNtlmSupport(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var providersNisNisItemTypeNtlmSupportPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["all","v2only","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		providersNisNisItemTypeNtlmSupportPropEnum = append(providersNisNisItemTypeNtlmSupportPropEnum, v)
	}
}

const (

	// ProvidersNisNisItemNtlmSupportAll captures enum value "all"
	ProvidersNisNisItemNtlmSupportAll string = "all"

	// ProvidersNisNisItemNtlmSupportV2only captures enum value "v2only"
	ProvidersNisNisItemNtlmSupportV2only string = "v2only"

	// ProvidersNisNisItemNtlmSupportNone captures enum value "none"
	ProvidersNisNisItemNtlmSupportNone string = "none"
)

// prop value enum
func (m *ProvidersNisNisItem) validateNtlmSupportEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, providersNisNisItemTypeNtlmSupportPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ProvidersNisNisItem) validateNtlmSupport(formats strfmt.Registry) error {

	if swag.IsZero(m.NtlmSupport) { // not required
		return nil
	}

	// value enum
	if err := m.validateNtlmSupportEnum("ntlm_support", "body", m.NtlmSupport); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProvidersNisNisItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProvidersNisNisItem) UnmarshalBinary(b []byte) error {
	var res ProvidersNisNisItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
