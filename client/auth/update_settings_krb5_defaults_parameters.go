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

// NewUpdateSettingsKrb5DefaultsParams creates a new UpdateSettingsKrb5DefaultsParams object
// with the default values initialized.
func NewUpdateSettingsKrb5DefaultsParams() *UpdateSettingsKrb5DefaultsParams {
	var ()
	return &UpdateSettingsKrb5DefaultsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSettingsKrb5DefaultsParamsWithTimeout creates a new UpdateSettingsKrb5DefaultsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSettingsKrb5DefaultsParamsWithTimeout(timeout time.Duration) *UpdateSettingsKrb5DefaultsParams {
	var ()
	return &UpdateSettingsKrb5DefaultsParams{

		timeout: timeout,
	}
}

// NewUpdateSettingsKrb5DefaultsParamsWithContext creates a new UpdateSettingsKrb5DefaultsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSettingsKrb5DefaultsParamsWithContext(ctx context.Context) *UpdateSettingsKrb5DefaultsParams {
	var ()
	return &UpdateSettingsKrb5DefaultsParams{

		Context: ctx,
	}
}

// NewUpdateSettingsKrb5DefaultsParamsWithHTTPClient creates a new UpdateSettingsKrb5DefaultsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSettingsKrb5DefaultsParamsWithHTTPClient(client *http.Client) *UpdateSettingsKrb5DefaultsParams {
	var ()
	return &UpdateSettingsKrb5DefaultsParams{
		HTTPClient: client,
	}
}

/*UpdateSettingsKrb5DefaultsParams contains all the parameters to send to the API endpoint
for the update settings krb5 defaults operation typically these are written to a http.Request
*/
type UpdateSettingsKrb5DefaultsParams struct {

	/*SettingsKrb5Defaults*/
	SettingsKrb5Defaults *models.SettingsKrb5DefaultsKrb5Settings

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update settings krb5 defaults params
func (o *UpdateSettingsKrb5DefaultsParams) WithTimeout(timeout time.Duration) *UpdateSettingsKrb5DefaultsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update settings krb5 defaults params
func (o *UpdateSettingsKrb5DefaultsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update settings krb5 defaults params
func (o *UpdateSettingsKrb5DefaultsParams) WithContext(ctx context.Context) *UpdateSettingsKrb5DefaultsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update settings krb5 defaults params
func (o *UpdateSettingsKrb5DefaultsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update settings krb5 defaults params
func (o *UpdateSettingsKrb5DefaultsParams) WithHTTPClient(client *http.Client) *UpdateSettingsKrb5DefaultsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update settings krb5 defaults params
func (o *UpdateSettingsKrb5DefaultsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSettingsKrb5Defaults adds the settingsKrb5Defaults to the update settings krb5 defaults params
func (o *UpdateSettingsKrb5DefaultsParams) WithSettingsKrb5Defaults(settingsKrb5Defaults *models.SettingsKrb5DefaultsKrb5Settings) *UpdateSettingsKrb5DefaultsParams {
	o.SetSettingsKrb5Defaults(settingsKrb5Defaults)
	return o
}

// SetSettingsKrb5Defaults adds the settingsKrb5Defaults to the update settings krb5 defaults params
func (o *UpdateSettingsKrb5DefaultsParams) SetSettingsKrb5Defaults(settingsKrb5Defaults *models.SettingsKrb5DefaultsKrb5Settings) {
	o.SettingsKrb5Defaults = settingsKrb5Defaults
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSettingsKrb5DefaultsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SettingsKrb5Defaults != nil {
		if err := r.SetBodyParam(o.SettingsKrb5Defaults); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
