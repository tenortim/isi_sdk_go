// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// NewUpdateSettingsKrb5DomainParams creates a new UpdateSettingsKrb5DomainParams object
// with the default values initialized.
func NewUpdateSettingsKrb5DomainParams() *UpdateSettingsKrb5DomainParams {
	var ()
	return &UpdateSettingsKrb5DomainParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSettingsKrb5DomainParamsWithTimeout creates a new UpdateSettingsKrb5DomainParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSettingsKrb5DomainParamsWithTimeout(timeout time.Duration) *UpdateSettingsKrb5DomainParams {
	var ()
	return &UpdateSettingsKrb5DomainParams{

		timeout: timeout,
	}
}

// NewUpdateSettingsKrb5DomainParamsWithContext creates a new UpdateSettingsKrb5DomainParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSettingsKrb5DomainParamsWithContext(ctx context.Context) *UpdateSettingsKrb5DomainParams {
	var ()
	return &UpdateSettingsKrb5DomainParams{

		Context: ctx,
	}
}

// NewUpdateSettingsKrb5DomainParamsWithHTTPClient creates a new UpdateSettingsKrb5DomainParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSettingsKrb5DomainParamsWithHTTPClient(client *http.Client) *UpdateSettingsKrb5DomainParams {
	var ()
	return &UpdateSettingsKrb5DomainParams{
		HTTPClient: client,
	}
}

/*UpdateSettingsKrb5DomainParams contains all the parameters to send to the API endpoint
for the update settings krb5 domain operation typically these are written to a http.Request
*/
type UpdateSettingsKrb5DomainParams struct {

	/*SettingsKrb5Domain*/
	SettingsKrb5Domain *models.SettingsKrb5Domain
	/*SettingsKrb5DomainID
	  Modify the krb5 domain settings.

	*/
	SettingsKrb5DomainID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update settings krb5 domain params
func (o *UpdateSettingsKrb5DomainParams) WithTimeout(timeout time.Duration) *UpdateSettingsKrb5DomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update settings krb5 domain params
func (o *UpdateSettingsKrb5DomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update settings krb5 domain params
func (o *UpdateSettingsKrb5DomainParams) WithContext(ctx context.Context) *UpdateSettingsKrb5DomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update settings krb5 domain params
func (o *UpdateSettingsKrb5DomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update settings krb5 domain params
func (o *UpdateSettingsKrb5DomainParams) WithHTTPClient(client *http.Client) *UpdateSettingsKrb5DomainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update settings krb5 domain params
func (o *UpdateSettingsKrb5DomainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSettingsKrb5Domain adds the settingsKrb5Domain to the update settings krb5 domain params
func (o *UpdateSettingsKrb5DomainParams) WithSettingsKrb5Domain(settingsKrb5Domain *models.SettingsKrb5Domain) *UpdateSettingsKrb5DomainParams {
	o.SetSettingsKrb5Domain(settingsKrb5Domain)
	return o
}

// SetSettingsKrb5Domain adds the settingsKrb5Domain to the update settings krb5 domain params
func (o *UpdateSettingsKrb5DomainParams) SetSettingsKrb5Domain(settingsKrb5Domain *models.SettingsKrb5Domain) {
	o.SettingsKrb5Domain = settingsKrb5Domain
}

// WithSettingsKrb5DomainID adds the settingsKrb5DomainID to the update settings krb5 domain params
func (o *UpdateSettingsKrb5DomainParams) WithSettingsKrb5DomainID(settingsKrb5DomainID string) *UpdateSettingsKrb5DomainParams {
	o.SetSettingsKrb5DomainID(settingsKrb5DomainID)
	return o
}

// SetSettingsKrb5DomainID adds the settingsKrb5DomainId to the update settings krb5 domain params
func (o *UpdateSettingsKrb5DomainParams) SetSettingsKrb5DomainID(settingsKrb5DomainID string) {
	o.SettingsKrb5DomainID = settingsKrb5DomainID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSettingsKrb5DomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SettingsKrb5Domain != nil {
		if err := r.SetBodyParam(o.SettingsKrb5Domain); err != nil {
			return err
		}
	}

	// path param SettingsKrb5DomainId
	if err := r.SetPathParam("SettingsKrb5DomainId", o.SettingsKrb5DomainID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
