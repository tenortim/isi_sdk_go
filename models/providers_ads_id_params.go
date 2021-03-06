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

// ProvidersAdsIDParams Specifies the properties for an ADS authentication provider.
// swagger:model ProvidersAdsIdParams
type ProvidersAdsIDParams struct {

	// Allocates an ID for an unmapped Active Directory (ADS) group. ADS groups without GIDs can be proactively assigned a GID by the ID mapper. If the ID mapper option is disabled, GIDs are not proactively assigned, and when a primary group for a user does not include a GID, the system may allocate one.
	AllocateGids bool `json:"allocate_gids,omitempty"`

	// Allocates a user ID for an unmapped Active Directory (ADS) user. ADS users without UIDs can be proactively assigned a UID by the ID mapper. IF the ID mapper option is disabled, UIDs are not proactively assigned, and when an identify for a user does not include a UID, the system may allocate one.
	AllocateUIDS bool `json:"allocate_uids,omitempty"`

	// Enables lookup of unqualified user names in the primary domain.
	AssumeDefaultDomain bool `json:"assume_default_domain,omitempty"`

	// Enables authentication and identity management through the authentication provider.
	Authentication bool `json:"authentication,omitempty"`

	// Specifies the time in seconds between provider online checks.
	CheckOnlineInterval int64 `json:"check_online_interval,omitempty"`

	// Specifies the current time for the domain controllers.
	ControllerTime int64 `json:"controller_time,omitempty"`

	// Automatically creates a home directory on the first login.
	CreateHomeDirectory bool `json:"create_home_directory,omitempty"`

	// Specifies the domain controller to which the authentication service should send requests
	DomainController string `json:"domain_controller,omitempty"`

	// Sends an alert if the domain goes offline.
	DomainOfflineAlerts bool `json:"domain_offline_alerts,omitempty"`

	// Sets list of groups that can be resolved.
	FindableGroups []string `json:"findable_groups"`

	// Sets list of users that can be resolved.
	FindableUsers []string `json:"findable_users"`

	// Specifies the path to the home directory template.
	HomeDirectoryTemplate string `json:"home_directory_template,omitempty"`

	// If set to true, ignores all trusted domains.
	IgnoreAllTrusts bool `json:"ignore_all_trusts,omitempty"`

	// Includes trusted domains when 'ignore_all_trusts' is set to false.
	IgnoredTrustedDomains []string `json:"ignored_trusted_domains"`

	// Includes trusted domains when 'ignore_all_trusts' is set to true.
	IncludeTrustedDomains []string `json:"include_trusted_domains"`

	// Specifies Active Directory provider instnace.
	Instance string `json:"instance,omitempty"`

	// Enables encryption and signing on LDAP requests.
	LdapSignAndSeal bool `json:"ldap_sign_and_seal,omitempty"`

	// Specifies the login shell path.
	LoginShell string `json:"login_shell,omitempty"`

	// Limits user and group lookups to the specified domains.
	LookupDomains []string `json:"lookup_domains"`

	// Looks up AD groups in other providers before allocating a group ID.
	LookupGroups bool `json:"lookup_groups,omitempty"`

	// Normalizes AD group names to lowercase before look up.
	LookupNormalizeGroups bool `json:"lookup_normalize_groups,omitempty"`

	// Normalize AD user names to lowercase before look up.
	LookupNormalizeUsers bool `json:"lookup_normalize_users,omitempty"`

	// Looks up AD users in other providers before allocating a user ID.
	LookupUsers bool `json:"lookup_users,omitempty"`

	// Specifies name to join AD as.
	MachineName string `json:"machine_name,omitempty"`

	// Enables periodic changes of the machine password for security.
	MachinePasswordChanges bool `json:"machine_password_changes,omitempty"`

	// Sets maximum age of a password in seconds.
	// Maximum: 3.1536e+07
	// Minimum: 3600
	MachinePasswordLifespan int64 `json:"machine_password_lifespan,omitempty"`

	// Specifies the domain controller for which the node has affinity.
	NodeDcAffinity string `json:"node_dc_affinity,omitempty"`

	// Specifies the timeout for the domain controller for which the local node has affinity.
	NodeDcAffinityTimeout int64 `json:"node_dc_affinity_timeout,omitempty"`

	// Enables the Active Directory provider to respond to 'getpwent' and 'getgrent' requests.
	NssEnumeration bool `json:"nss_enumeration,omitempty"`

	// Specifies the password used during domain join.
	Password string `json:"password,omitempty"`

	// Resets the secure channel to the primary domain.
	ResetSchannel bool `json:"reset_schannel,omitempty"`

	// Check the provider for filtered lists of findable and unfindable users and groups.
	RestrictFindable bool `json:"restrict_findable,omitempty"`

	// Specifies whether to support RFC 2307 attributes on ADS domain controllers.
	SfuSupport string `json:"sfu_support,omitempty"`

	// Currently configured SPNs.
	Spns []string `json:"spns"`

	// Stores SFU mappings permanently in the ID mapper.
	StoreSfuMappings bool `json:"store_sfu_mappings,omitempty"`

	// Specifies groups that cannot be resolved by the provider.
	UnfindableGroups []string `json:"unfindable_groups"`

	// Specifies users that cannot be resolved by the provider.
	UnfindableUsers []string `json:"unfindable_users"`

	// Specifies the user name that has permission to join a machine to the given domain.
	User string `json:"user,omitempty"`
}

// Validate validates this providers ads Id params
func (m *ProvidersAdsIDParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMachinePasswordLifespan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProvidersAdsIDParams) validateMachinePasswordLifespan(formats strfmt.Registry) error {

	if swag.IsZero(m.MachinePasswordLifespan) { // not required
		return nil
	}

	if err := validate.MinimumInt("machine_password_lifespan", "body", int64(m.MachinePasswordLifespan), 3600, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("machine_password_lifespan", "body", int64(m.MachinePasswordLifespan), 3.1536e+07, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProvidersAdsIDParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProvidersAdsIDParams) UnmarshalBinary(b []byte) error {
	var res ProvidersAdsIDParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
