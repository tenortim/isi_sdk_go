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

// FtpSettingsExtended FTP settings.
// swagger:model FtpSettingsExtended
type FtpSettingsExtended struct {

	// The timeout, in seconds, for a remote client to establish a PASV style data connection.
	// Maximum: 600
	// Minimum: 30
	AcceptTimeout int64 `json:"accept_timeout,omitempty"`

	// Controls whether anonymous logins are permitted or not.
	AllowAnonAccess bool `json:"allow_anon_access,omitempty"`

	// Controls whether anonymous users will be permitted to upload files.
	AllowAnonUpload bool `json:"allow_anon_upload,omitempty"`

	// If set to false, all directory list commands will return a permission denied error.
	AllowDirlists bool `json:"allow_dirlists,omitempty"`

	// If set to false, all downloads requests will return a permission denied error.
	AllowDownloads bool `json:"allow_downloads,omitempty"`

	// Controls whether local logins are permitted or not.
	AllowLocalAccess bool `json:"allow_local_access,omitempty"`

	// This controls whether any FTP commands which change the filesystem are allowed or not.
	AllowWrites bool `json:"allow_writes,omitempty"`

	// This controls whether FTP will always initially change directories to the home directory of the user, regardless of whether it is chroot-ing.
	AlwaysChdirHomedir bool `json:"always_chdir_homedir,omitempty"`

	// This is the name of the user who is given ownership of anonymously uploaded files.
	AnonChownUsername string `json:"anon_chown_username,omitempty"`

	// A list of passwords for anonymous users.
	AnonPasswordList []string `json:"anon_password_list"`

	// This option represents a directory in /ifs which vsftpd will try to change into after an anonymous login.
	AnonRootPath string `json:"anon_root_path,omitempty"`

	// The value that the umask for file creation is set to for anonymous users.
	// Maximum: 511
	// Minimum: 0
	AnonUmask *int64 `json:"anon_umask,omitempty"`

	// Controls whether ascii mode data transfers are honored for various types of requests.
	ASCIIMode string `json:"ascii_mode,omitempty"`

	// A list of users that are not chrooted when logging in.
	ChrootExceptionList []string `json:"chroot_exception_list"`

	// If set to 'all', all local users will be (by default) placed in a chroot() jail in their home directory after login. If set to 'all-with-exceptions', all local users except those listed in the chroot exception list (isi ftp chroot-exception-list) will be placed in a chroot() jail in their home directory after login. If set to 'none', no local users will be chrooted by default. If set to 'none-with-exceptions', only the local users listed in the chroot exception list (isi ftp chroot-exception-list) will be place in a chroot() jail in their home directory after login.
	ChrootLocalMode string `json:"chroot_local_mode,omitempty"`

	// The timeout, in seconds, for a remote client to respond to our PORT style data connection.
	// Maximum: 600
	// Minimum: 30
	ConnectTimeout int64 `json:"connect_timeout,omitempty"`

	// The timeout, in seconds, which is roughly the maximum time we permit data transfers to stall for with no progress. If the timeout triggers, the remote client is kicked off.
	// Maximum: 600
	// Minimum: 30
	DataTimeout int64 `json:"data_timeout,omitempty"`

	// A list of uses that will be denied access.
	DeniedUserList []string `json:"denied_user_list"`

	// If enabled, display directory listings with the time in your local time zone. The default is to display GMT. The times returned by the MDTM FTP command are also affected by this option.
	DirlistLocaltime bool `json:"dirlist_localtime,omitempty"`

	// When set to 'hide',  all user and group information in directory listings will be displayed as 'ftp'. When set to 'textual', textual names are shown in the user and group fields of directory listings. When set to 'numeric', numeric IDs are show in the user and group fields of directory listings.
	DirlistNames string `json:"dirlist_names,omitempty"`

	// The permissions with which uploaded files are created. Umasks are applied on top of this value.
	// Maximum: 511
	// Minimum: 0
	FileCreatePerm *int64 `json:"file_create_perm,omitempty"`

	// This field determines whether the anon_password_list is used.
	LimitAnonPasswords bool `json:"limit_anon_passwords,omitempty"`

	// This option represents a directory in /ifs which vsftpd will try to change into after a local login.
	LocalRootPath string `json:"local_root_path,omitempty"`

	// The value that the umask for file creation is set to for local users.
	// Maximum: 511
	// Minimum: 0
	LocalUmask *int64 `json:"local_umask,omitempty"`

	// If enabled, allow server-to-server (FXP) transfers.
	ServerToServer bool `json:"server_to_server,omitempty"`

	// This field controls whether the FTP daemon is running.
	Service bool `json:"service,omitempty"`

	// If enabled, maintain login sessions for each user through Pluggable Authentication Modules (PAM). Disabling this option prevents the ability to do automatic home directory creation if that functionality were otherwise available.
	SessionSupport bool `json:"session_support,omitempty"`

	// The timeout, in seconds, which is roughly the maximum time we permit data transfers to stall for with no progress. If the timeout triggers, the remote client is kicked off.
	// Maximum: 600
	// Minimum: 30
	SessionTimeout int64 `json:"session_timeout,omitempty"`

	// Specifies the directory where per-user config overrides can be found.
	UserConfigDir string `json:"user_config_dir,omitempty"`
}

// Validate validates this ftp settings extended
func (m *FtpSettingsExtended) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptTimeout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAnonUmask(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectTimeout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataTimeout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileCreatePerm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalUmask(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionTimeout(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FtpSettingsExtended) validateAcceptTimeout(formats strfmt.Registry) error {

	if swag.IsZero(m.AcceptTimeout) { // not required
		return nil
	}

	if err := validate.MinimumInt("accept_timeout", "body", int64(m.AcceptTimeout), 30, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("accept_timeout", "body", int64(m.AcceptTimeout), 600, false); err != nil {
		return err
	}

	return nil
}

func (m *FtpSettingsExtended) validateAnonUmask(formats strfmt.Registry) error {

	if swag.IsZero(m.AnonUmask) { // not required
		return nil
	}

	if err := validate.MinimumInt("anon_umask", "body", int64(*m.AnonUmask), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("anon_umask", "body", int64(*m.AnonUmask), 511, false); err != nil {
		return err
	}

	return nil
}

func (m *FtpSettingsExtended) validateConnectTimeout(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectTimeout) { // not required
		return nil
	}

	if err := validate.MinimumInt("connect_timeout", "body", int64(m.ConnectTimeout), 30, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("connect_timeout", "body", int64(m.ConnectTimeout), 600, false); err != nil {
		return err
	}

	return nil
}

func (m *FtpSettingsExtended) validateDataTimeout(formats strfmt.Registry) error {

	if swag.IsZero(m.DataTimeout) { // not required
		return nil
	}

	if err := validate.MinimumInt("data_timeout", "body", int64(m.DataTimeout), 30, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("data_timeout", "body", int64(m.DataTimeout), 600, false); err != nil {
		return err
	}

	return nil
}

func (m *FtpSettingsExtended) validateFileCreatePerm(formats strfmt.Registry) error {

	if swag.IsZero(m.FileCreatePerm) { // not required
		return nil
	}

	if err := validate.MinimumInt("file_create_perm", "body", int64(*m.FileCreatePerm), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("file_create_perm", "body", int64(*m.FileCreatePerm), 511, false); err != nil {
		return err
	}

	return nil
}

func (m *FtpSettingsExtended) validateLocalUmask(formats strfmt.Registry) error {

	if swag.IsZero(m.LocalUmask) { // not required
		return nil
	}

	if err := validate.MinimumInt("local_umask", "body", int64(*m.LocalUmask), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("local_umask", "body", int64(*m.LocalUmask), 511, false); err != nil {
		return err
	}

	return nil
}

func (m *FtpSettingsExtended) validateSessionTimeout(formats strfmt.Registry) error {

	if swag.IsZero(m.SessionTimeout) { // not required
		return nil
	}

	if err := validate.MinimumInt("session_timeout", "body", int64(m.SessionTimeout), 30, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("session_timeout", "body", int64(m.SessionTimeout), 600, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FtpSettingsExtended) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FtpSettingsExtended) UnmarshalBinary(b []byte) error {
	var res FtpSettingsExtended
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}