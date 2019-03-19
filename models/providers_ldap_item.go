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

// ProvidersLdapItem Specifies the properties for the LDAP authentication provider.
// swagger:model ProvidersLdapItem
type ProvidersLdapItem struct {

	// Specifies the attribute name used when searching for alternate security identities.
	AlternateSecurityIdentitiesAttribute string `json:"alternate_security_identities_attribute,omitempty"`

	// If true, enables authentication and identity management through the authentication provider.
	Authentication bool `json:"authentication,omitempty"`

	// If true, connects the provider to a random server.
	BalanceServers bool `json:"balance_servers,omitempty"`

	// Specifies the root of the tree in which to search identities.
	// Required: true
	BaseDn *string `json:"base_dn"`

	// Specifies the distinguished name for binding to the LDAP server.
	BindDn string `json:"bind_dn,omitempty"`

	// Specifies which bind mechanism to use when connecting to an LDAP server. The only supported option is the 'simple' value.
	// Enum: [simple gssapi digest-md5]
	BindMechanism string `json:"bind_mechanism,omitempty"`

	// Specifies the password for the distinguished name for binding to the LDAP server.
	BindPassword string `json:"bind_password,omitempty"`

	// Specifies the timeout in seconds when binding to an LDAP server.
	BindTimeout int64 `json:"bind_timeout,omitempty"`

	// Specifies the path to the root certificates file.
	CertificateAuthorityFile string `json:"certificate_authority_file,omitempty"`

	// Specifies the time in seconds between provider online checks.
	CheckOnlineInterval int64 `json:"check_online_interval,omitempty"`

	// Specifies the canonical name.
	CnAttribute string `json:"cn_attribute,omitempty"`

	// Automatically create the home directory on the first login.
	CreateHomeDirectory bool `json:"create_home_directory,omitempty"`

	// Specifies the hashed password value.
	CryptPasswordAttribute string `json:"crypt_password_attribute,omitempty"`

	// Specifies the LDAP Email attribute.
	EmailAttribute string `json:"email_attribute,omitempty"`

	// If true, enables the LDAP provider.
	Enabled bool `json:"enabled,omitempty"`

	// If true, allows the provider to enumerate groups.
	EnumerateGroups bool `json:"enumerate_groups,omitempty"`

	// If true, allows the provider to enumerate users.
	EnumerateUsers bool `json:"enumerate_users,omitempty"`

	// Specifies the list of groups that can be resolved.
	FindableGroups []string `json:"findable_groups"`

	// Specifies the list of users that can be resolved.
	FindableUsers []string `json:"findable_users"`

	// Specifies the LDAP GECOS attribute.
	GecosAttribute string `json:"gecos_attribute,omitempty"`

	// Specifies the LDAP GID attribute.
	GidAttribute string `json:"gid_attribute,omitempty"`

	// Specifies the distinguished name of the entry where LDAP searches for groups are started.
	GroupBaseDn string `json:"group_base_dn,omitempty"`

	// Specifies the domain for this provider through which groups are qualified.
	GroupDomain string `json:"group_domain,omitempty"`

	// Specifies the LDAP filter for group objects.
	GroupFilter string `json:"group_filter,omitempty"`

	// Specifies the LDAP Group Members attribute.
	GroupMembersAttribute string `json:"group_members_attribute,omitempty"`

	// Specifies the depth from the base DN to perform LDAP searches.
	// Enum: [default base onelevel subtree children]
	GroupSearchScope string `json:"group_search_scope,omitempty"`

	// Groupnet identifier.
	Groupnet string `json:"groupnet,omitempty"`

	// Specifies the path to the home directory template.
	HomeDirectoryTemplate string `json:"home_directory_template,omitempty"`

	// Specifies the LDAP Homedir attribute.
	HomedirAttribute string `json:"homedir_attribute,omitempty"`

	// If true, continues over secure connections even if identity checks fail.
	IgnoreTLSErrors bool `json:"ignore_tls_errors,omitempty"`

	// Specifies the groups that can be viewed in the provider.
	ListableGroups []string `json:"listable_groups"`

	// Specifies the users that can be viewed in the provider.
	ListableUsers []string `json:"listable_users"`

	// Specifies the login shell path.
	LoginShell string `json:"login_shell,omitempty"`

	// Specifies the LDAP Query Member Of attribute, which performs reverse membership queries.
	MemberOfAttribute string `json:"member_of_attribute,omitempty"`

	// Specifies the name of the LDAP provider.
	// Required: true
	Name *string `json:"name"`

	// Specifies the LDAP UID attribute, which is used as the login name.
	NameAttribute string `json:"name_attribute,omitempty"`

	// Specifies the distinguished name of the entry where LDAP searches for netgroups are started.
	NetgroupBaseDn string `json:"netgroup_base_dn,omitempty"`

	// Specifies the LDAP filter for netgroup objects.
	NetgroupFilter string `json:"netgroup_filter,omitempty"`

	// Specifies the LDAP Netgroup Members attribute.
	NetgroupMembersAttribute string `json:"netgroup_members_attribute,omitempty"`

	// Specifies the depth from the base DN to perform LDAP searches.
	// Enum: [default base onelevel subtree children]
	NetgroupSearchScope string `json:"netgroup_search_scope,omitempty"`

	// Specifies the LDAP Netgroup Triple attribute.
	NetgroupTripleAttribute string `json:"netgroup_triple_attribute,omitempty"`

	// Normalizes group names to lowercase before look up.
	NormalizeGroups bool `json:"normalize_groups,omitempty"`

	// Normalizes user names to lowercase before look up.
	NormalizeUsers bool `json:"normalize_users,omitempty"`

	// Specifies the LDAP NT Password attribute.
	NtPasswordAttribute string `json:"nt_password_attribute,omitempty"`

	// Specifies which NTLM versions to support for users with NTLM-compatible credentials.
	// Enum: [all v2only none]
	NtlmSupport string `json:"ntlm_support,omitempty"`

	// Specifies the provider domain.
	ProviderDomain string `json:"provider_domain,omitempty"`

	// Determines whether to continue over a non-TLS connection.
	RequireSecureConnection bool `json:"require_secure_connection,omitempty"`

	// If true, checks the provider for filtered lists of findable and unfindable users and groups.
	RestrictFindable bool `json:"restrict_findable,omitempty"`

	// If true, checks the provider for filtered lists of listable and unlistable users and groups.
	RestrictListable bool `json:"restrict_listable,omitempty"`

	// Specifies the default depth from the base DN to perform LDAP searches.
	// Enum: [base onelevel subtree children]
	SearchScope string `json:"search_scope,omitempty"`

	// Specifies the search timeout period in seconds.
	SearchTimeout int64 `json:"search_timeout,omitempty"`

	// Specifies the server URIs.
	// Required: true
	ServerUris []string `json:"server_uris"`

	// Specifies the the LDAP Shell attribute.
	ShellAttribute string `json:"shell_attribute,omitempty"`

	// Specifies the the LDAP UID Number attribute.
	UIDAttribute string `json:"uid_attribute,omitempty"`

	// Specifies the groups that cannot be resolved by the provider.
	UnfindableGroups []string `json:"unfindable_groups"`

	// Specifies users that cannot be resolved by the provider.
	UnfindableUsers []string `json:"unfindable_users"`

	// Sets the LDAP Unique Group Members attribute.
	UniqueGroupMembersAttribute string `json:"unique_group_members_attribute,omitempty"`

	// Specifies a group that cannot be listed by the provider.
	UnlistableGroups []string `json:"unlistable_groups"`

	// Specifies a user that cannot be listed by the provider.
	UnlistableUsers []string `json:"unlistable_users"`

	// Specifies the distinguished name of the entry at which to start LDAP searches for users.
	UserBaseDn string `json:"user_base_dn,omitempty"`

	// Specifies the domain for this provider through which users are qualified.
	UserDomain string `json:"user_domain,omitempty"`

	// Specifies the LDAP filter for user objects.
	UserFilter string `json:"user_filter,omitempty"`

	// Specifies the depth from the base DN to perform LDAP searches.
	// Enum: [default base onelevel subtree children]
	UserSearchScope string `json:"user_search_scope,omitempty"`
}

// Validate validates this providers ldap item
func (m *ProvidersLdapItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseDn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBindMechanism(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupSearchScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetgroupSearchScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNtlmSupport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearchScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerUris(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserSearchScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProvidersLdapItem) validateBaseDn(formats strfmt.Registry) error {

	if err := validate.Required("base_dn", "body", m.BaseDn); err != nil {
		return err
	}

	return nil
}

var providersLdapItemTypeBindMechanismPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["simple","gssapi","digest-md5"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		providersLdapItemTypeBindMechanismPropEnum = append(providersLdapItemTypeBindMechanismPropEnum, v)
	}
}

const (

	// ProvidersLdapItemBindMechanismSimple captures enum value "simple"
	ProvidersLdapItemBindMechanismSimple string = "simple"

	// ProvidersLdapItemBindMechanismGssapi captures enum value "gssapi"
	ProvidersLdapItemBindMechanismGssapi string = "gssapi"

	// ProvidersLdapItemBindMechanismDigestMd5 captures enum value "digest-md5"
	ProvidersLdapItemBindMechanismDigestMd5 string = "digest-md5"
)

// prop value enum
func (m *ProvidersLdapItem) validateBindMechanismEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, providersLdapItemTypeBindMechanismPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ProvidersLdapItem) validateBindMechanism(formats strfmt.Registry) error {

	if swag.IsZero(m.BindMechanism) { // not required
		return nil
	}

	// value enum
	if err := m.validateBindMechanismEnum("bind_mechanism", "body", m.BindMechanism); err != nil {
		return err
	}

	return nil
}

var providersLdapItemTypeGroupSearchScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["default","base","onelevel","subtree","children"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		providersLdapItemTypeGroupSearchScopePropEnum = append(providersLdapItemTypeGroupSearchScopePropEnum, v)
	}
}

const (

	// ProvidersLdapItemGroupSearchScopeDefault captures enum value "default"
	ProvidersLdapItemGroupSearchScopeDefault string = "default"

	// ProvidersLdapItemGroupSearchScopeBase captures enum value "base"
	ProvidersLdapItemGroupSearchScopeBase string = "base"

	// ProvidersLdapItemGroupSearchScopeOnelevel captures enum value "onelevel"
	ProvidersLdapItemGroupSearchScopeOnelevel string = "onelevel"

	// ProvidersLdapItemGroupSearchScopeSubtree captures enum value "subtree"
	ProvidersLdapItemGroupSearchScopeSubtree string = "subtree"

	// ProvidersLdapItemGroupSearchScopeChildren captures enum value "children"
	ProvidersLdapItemGroupSearchScopeChildren string = "children"
)

// prop value enum
func (m *ProvidersLdapItem) validateGroupSearchScopeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, providersLdapItemTypeGroupSearchScopePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ProvidersLdapItem) validateGroupSearchScope(formats strfmt.Registry) error {

	if swag.IsZero(m.GroupSearchScope) { // not required
		return nil
	}

	// value enum
	if err := m.validateGroupSearchScopeEnum("group_search_scope", "body", m.GroupSearchScope); err != nil {
		return err
	}

	return nil
}

func (m *ProvidersLdapItem) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var providersLdapItemTypeNetgroupSearchScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["default","base","onelevel","subtree","children"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		providersLdapItemTypeNetgroupSearchScopePropEnum = append(providersLdapItemTypeNetgroupSearchScopePropEnum, v)
	}
}

const (

	// ProvidersLdapItemNetgroupSearchScopeDefault captures enum value "default"
	ProvidersLdapItemNetgroupSearchScopeDefault string = "default"

	// ProvidersLdapItemNetgroupSearchScopeBase captures enum value "base"
	ProvidersLdapItemNetgroupSearchScopeBase string = "base"

	// ProvidersLdapItemNetgroupSearchScopeOnelevel captures enum value "onelevel"
	ProvidersLdapItemNetgroupSearchScopeOnelevel string = "onelevel"

	// ProvidersLdapItemNetgroupSearchScopeSubtree captures enum value "subtree"
	ProvidersLdapItemNetgroupSearchScopeSubtree string = "subtree"

	// ProvidersLdapItemNetgroupSearchScopeChildren captures enum value "children"
	ProvidersLdapItemNetgroupSearchScopeChildren string = "children"
)

// prop value enum
func (m *ProvidersLdapItem) validateNetgroupSearchScopeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, providersLdapItemTypeNetgroupSearchScopePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ProvidersLdapItem) validateNetgroupSearchScope(formats strfmt.Registry) error {

	if swag.IsZero(m.NetgroupSearchScope) { // not required
		return nil
	}

	// value enum
	if err := m.validateNetgroupSearchScopeEnum("netgroup_search_scope", "body", m.NetgroupSearchScope); err != nil {
		return err
	}

	return nil
}

var providersLdapItemTypeNtlmSupportPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["all","v2only","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		providersLdapItemTypeNtlmSupportPropEnum = append(providersLdapItemTypeNtlmSupportPropEnum, v)
	}
}

const (

	// ProvidersLdapItemNtlmSupportAll captures enum value "all"
	ProvidersLdapItemNtlmSupportAll string = "all"

	// ProvidersLdapItemNtlmSupportV2only captures enum value "v2only"
	ProvidersLdapItemNtlmSupportV2only string = "v2only"

	// ProvidersLdapItemNtlmSupportNone captures enum value "none"
	ProvidersLdapItemNtlmSupportNone string = "none"
)

// prop value enum
func (m *ProvidersLdapItem) validateNtlmSupportEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, providersLdapItemTypeNtlmSupportPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ProvidersLdapItem) validateNtlmSupport(formats strfmt.Registry) error {

	if swag.IsZero(m.NtlmSupport) { // not required
		return nil
	}

	// value enum
	if err := m.validateNtlmSupportEnum("ntlm_support", "body", m.NtlmSupport); err != nil {
		return err
	}

	return nil
}

var providersLdapItemTypeSearchScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["base","onelevel","subtree","children"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		providersLdapItemTypeSearchScopePropEnum = append(providersLdapItemTypeSearchScopePropEnum, v)
	}
}

const (

	// ProvidersLdapItemSearchScopeBase captures enum value "base"
	ProvidersLdapItemSearchScopeBase string = "base"

	// ProvidersLdapItemSearchScopeOnelevel captures enum value "onelevel"
	ProvidersLdapItemSearchScopeOnelevel string = "onelevel"

	// ProvidersLdapItemSearchScopeSubtree captures enum value "subtree"
	ProvidersLdapItemSearchScopeSubtree string = "subtree"

	// ProvidersLdapItemSearchScopeChildren captures enum value "children"
	ProvidersLdapItemSearchScopeChildren string = "children"
)

// prop value enum
func (m *ProvidersLdapItem) validateSearchScopeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, providersLdapItemTypeSearchScopePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ProvidersLdapItem) validateSearchScope(formats strfmt.Registry) error {

	if swag.IsZero(m.SearchScope) { // not required
		return nil
	}

	// value enum
	if err := m.validateSearchScopeEnum("search_scope", "body", m.SearchScope); err != nil {
		return err
	}

	return nil
}

func (m *ProvidersLdapItem) validateServerUris(formats strfmt.Registry) error {

	if err := validate.Required("server_uris", "body", m.ServerUris); err != nil {
		return err
	}

	return nil
}

var providersLdapItemTypeUserSearchScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["default","base","onelevel","subtree","children"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		providersLdapItemTypeUserSearchScopePropEnum = append(providersLdapItemTypeUserSearchScopePropEnum, v)
	}
}

const (

	// ProvidersLdapItemUserSearchScopeDefault captures enum value "default"
	ProvidersLdapItemUserSearchScopeDefault string = "default"

	// ProvidersLdapItemUserSearchScopeBase captures enum value "base"
	ProvidersLdapItemUserSearchScopeBase string = "base"

	// ProvidersLdapItemUserSearchScopeOnelevel captures enum value "onelevel"
	ProvidersLdapItemUserSearchScopeOnelevel string = "onelevel"

	// ProvidersLdapItemUserSearchScopeSubtree captures enum value "subtree"
	ProvidersLdapItemUserSearchScopeSubtree string = "subtree"

	// ProvidersLdapItemUserSearchScopeChildren captures enum value "children"
	ProvidersLdapItemUserSearchScopeChildren string = "children"
)

// prop value enum
func (m *ProvidersLdapItem) validateUserSearchScopeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, providersLdapItemTypeUserSearchScopePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ProvidersLdapItem) validateUserSearchScope(formats strfmt.Registry) error {

	if swag.IsZero(m.UserSearchScope) { // not required
		return nil
	}

	// value enum
	if err := m.validateUserSearchScopeEnum("user_search_scope", "body", m.UserSearchScope); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProvidersLdapItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProvidersLdapItem) UnmarshalBinary(b []byte) error {
	var res ProvidersLdapItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
