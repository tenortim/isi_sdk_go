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

// ClusterEmailSettings cluster email settings
// swagger:model ClusterEmailSettings
type ClusterEmailSettings struct {

	// This setting determines how notifications will be batched together to be sent by email.  'none' means each notification will be sent separately.  'severity' means notifications of the same severity will be sent together.  'category' means notifications of the same category will be sent together.  'all' means all notifications will be batched together and sent in a single email.
	// Required: true
	// Enum: [all severity category none]
	BatchMode *string `json:"batch_mode"`

	// The address of the SMTP server to be used for relaying the notification messages.  An SMTP server is required in order to send notifications.  If this string is empty, no emails will be sent.
	// Required: true
	MailRelay *string `json:"mail_relay"`

	// The full email address that will appear as the sender of notification messages.
	// Required: true
	MailSender *string `json:"mail_sender"`

	// The subject line for notification messages from this cluster.
	// Required: true
	MailSubject *string `json:"mail_subject"`

	// Indicates if an SMTP authentication password is set.
	// Required: true
	SMTPAuthPasswdSet *bool `json:"smtp_auth_passwd_set"`

	// The type of secure communication protocol to use if SMTP is being used.  If 'none', plain text will be used, if 'starttls', the encrypted STARTTLS protocol will be used.
	// Required: true
	// Enum: [none starttls]
	SMTPAuthSecurity *string `json:"smtp_auth_security"`

	// Username to authenticate with if SMTP authentication is being used.
	// Required: true
	SMTPAuthUsername *string `json:"smtp_auth_username"`

	// The port on the SMTP server to be used for relaying the notification messages.
	// Required: true
	SMTPPort *int64 `json:"smtp_port"`

	// If true, this cluster will send SMTP authentication credentials to the SMTP relay server in order to send its notification emails.  If false, the cluster will attempt to send its notification emails without authentication.
	// Required: true
	UseSMTPAuth *bool `json:"use_smtp_auth"`

	// Location of a custom template file that can be used to specify the layout of the notification emails.
	UserTemplate string `json:"user_template,omitempty"`
}

// Validate validates this cluster email settings
func (m *ClusterEmailSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBatchMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMailRelay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMailSender(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMailSubject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSMTPAuthPasswdSet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSMTPAuthSecurity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSMTPAuthUsername(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSMTPPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseSMTPAuth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var clusterEmailSettingsTypeBatchModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["all","severity","category","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterEmailSettingsTypeBatchModePropEnum = append(clusterEmailSettingsTypeBatchModePropEnum, v)
	}
}

const (

	// ClusterEmailSettingsBatchModeAll captures enum value "all"
	ClusterEmailSettingsBatchModeAll string = "all"

	// ClusterEmailSettingsBatchModeSeverity captures enum value "severity"
	ClusterEmailSettingsBatchModeSeverity string = "severity"

	// ClusterEmailSettingsBatchModeCategory captures enum value "category"
	ClusterEmailSettingsBatchModeCategory string = "category"

	// ClusterEmailSettingsBatchModeNone captures enum value "none"
	ClusterEmailSettingsBatchModeNone string = "none"
)

// prop value enum
func (m *ClusterEmailSettings) validateBatchModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterEmailSettingsTypeBatchModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterEmailSettings) validateBatchMode(formats strfmt.Registry) error {

	if err := validate.Required("batch_mode", "body", m.BatchMode); err != nil {
		return err
	}

	// value enum
	if err := m.validateBatchModeEnum("batch_mode", "body", *m.BatchMode); err != nil {
		return err
	}

	return nil
}

func (m *ClusterEmailSettings) validateMailRelay(formats strfmt.Registry) error {

	if err := validate.Required("mail_relay", "body", m.MailRelay); err != nil {
		return err
	}

	return nil
}

func (m *ClusterEmailSettings) validateMailSender(formats strfmt.Registry) error {

	if err := validate.Required("mail_sender", "body", m.MailSender); err != nil {
		return err
	}

	return nil
}

func (m *ClusterEmailSettings) validateMailSubject(formats strfmt.Registry) error {

	if err := validate.Required("mail_subject", "body", m.MailSubject); err != nil {
		return err
	}

	return nil
}

func (m *ClusterEmailSettings) validateSMTPAuthPasswdSet(formats strfmt.Registry) error {

	if err := validate.Required("smtp_auth_passwd_set", "body", m.SMTPAuthPasswdSet); err != nil {
		return err
	}

	return nil
}

var clusterEmailSettingsTypeSMTPAuthSecurityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","starttls"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterEmailSettingsTypeSMTPAuthSecurityPropEnum = append(clusterEmailSettingsTypeSMTPAuthSecurityPropEnum, v)
	}
}

const (

	// ClusterEmailSettingsSMTPAuthSecurityNone captures enum value "none"
	ClusterEmailSettingsSMTPAuthSecurityNone string = "none"

	// ClusterEmailSettingsSMTPAuthSecurityStarttls captures enum value "starttls"
	ClusterEmailSettingsSMTPAuthSecurityStarttls string = "starttls"
)

// prop value enum
func (m *ClusterEmailSettings) validateSMTPAuthSecurityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterEmailSettingsTypeSMTPAuthSecurityPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterEmailSettings) validateSMTPAuthSecurity(formats strfmt.Registry) error {

	if err := validate.Required("smtp_auth_security", "body", m.SMTPAuthSecurity); err != nil {
		return err
	}

	// value enum
	if err := m.validateSMTPAuthSecurityEnum("smtp_auth_security", "body", *m.SMTPAuthSecurity); err != nil {
		return err
	}

	return nil
}

func (m *ClusterEmailSettings) validateSMTPAuthUsername(formats strfmt.Registry) error {

	if err := validate.Required("smtp_auth_username", "body", m.SMTPAuthUsername); err != nil {
		return err
	}

	return nil
}

func (m *ClusterEmailSettings) validateSMTPPort(formats strfmt.Registry) error {

	if err := validate.Required("smtp_port", "body", m.SMTPPort); err != nil {
		return err
	}

	return nil
}

func (m *ClusterEmailSettings) validateUseSMTPAuth(formats strfmt.Registry) error {

	if err := validate.Required("use_smtp_auth", "body", m.UseSMTPAuth); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterEmailSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterEmailSettings) UnmarshalBinary(b []byte) error {
	var res ClusterEmailSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
