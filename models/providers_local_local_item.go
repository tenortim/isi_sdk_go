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

// ProvidersLocalLocalItem providers local local item
// swagger:model ProvidersLocalLocalItem
type ProvidersLocalLocalItem struct {

	// If true, enables authentication and identity management through the authentication provider.
	Authentication bool `json:"authentication,omitempty"`

	// Automatically creates the home directory on the first login.
	CreateHomeDirectory bool `json:"create_home_directory,omitempty"`

	// Specifies the path to the home directory template.
	HomeDirectoryTemplate string `json:"home_directory_template,omitempty"`

	// Specifies the local provider ID.
	ID string `json:"id,omitempty"`

	// Specifies the length of time in seconds that an account will be inaccessible after multiple failed login attempts.
	LockoutDuration int64 `json:"lockout_duration,omitempty"`

	// Specifies the number of failed login attempts necessary before an account is locked.
	LockoutThreshold int64 `json:"lockout_threshold,omitempty"`

	// Specifies the duration of time in seconds in which the number of failed attempts set in the 'lockout_threshold' parameter must be made before an account is locked.
	LockoutWindow int64 `json:"lockout_window,omitempty"`

	// Specifies the login shell path.
	LoginShell string `json:"login_shell,omitempty"`

	// Specifies the domain for this provider through which users and groups are qualified.
	MachineName string `json:"machine_name,omitempty"`

	// Specifies the maximum password age in seconds.
	MaxPasswordAge int64 `json:"max_password_age,omitempty"`

	// Specifies the minimum password age in seconds.
	MinPasswordAge int64 `json:"min_password_age,omitempty"`

	// Specifies the minimum password length.
	MinPasswordLength int64 `json:"min_password_length,omitempty"`

	// Specifies the local provider name.
	Name string `json:"name,omitempty"`

	// Specifies the conditions required for a password.
	PasswordComplexity []string `json:"password_complexity"`

	// Specifies the number of previous passwords to store.
	PasswordHistoryLength int64 `json:"password_history_length,omitempty"`

	// Specifies the time in seconds remaining before a user will be prompted for a password change.
	PasswordPromptTime int64 `json:"password_prompt_time,omitempty"`

	// Specifies the status of the provider.
	Status string `json:"status,omitempty"`

	// If true, indicates that this provider instance was created by OneFS and cannot be removed.
	System bool `json:"system,omitempty"`
}

// Validate validates this providers local local item
func (m *ProvidersLocalLocalItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePasswordComplexity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var providersLocalLocalItemPasswordComplexityItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["lowercase","uppercase","numeric","symbol"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		providersLocalLocalItemPasswordComplexityItemsEnum = append(providersLocalLocalItemPasswordComplexityItemsEnum, v)
	}
}

func (m *ProvidersLocalLocalItem) validatePasswordComplexityItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, providersLocalLocalItemPasswordComplexityItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *ProvidersLocalLocalItem) validatePasswordComplexity(formats strfmt.Registry) error {

	if swag.IsZero(m.PasswordComplexity) { // not required
		return nil
	}

	for i := 0; i < len(m.PasswordComplexity); i++ {

		// value enum
		if err := m.validatePasswordComplexityItemsEnum("password_complexity"+"."+strconv.Itoa(i), "body", m.PasswordComplexity[i]); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProvidersLocalLocalItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProvidersLocalLocalItem) UnmarshalBinary(b []byte) error {
	var res ProvidersLocalLocalItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
