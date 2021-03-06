// Code generated by go-swagger; DO NOT EDIT.

package audit

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

// NewUpdateSettingsGlobalParams creates a new UpdateSettingsGlobalParams object
// with the default values initialized.
func NewUpdateSettingsGlobalParams() *UpdateSettingsGlobalParams {
	var ()
	return &UpdateSettingsGlobalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSettingsGlobalParamsWithTimeout creates a new UpdateSettingsGlobalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSettingsGlobalParamsWithTimeout(timeout time.Duration) *UpdateSettingsGlobalParams {
	var ()
	return &UpdateSettingsGlobalParams{

		timeout: timeout,
	}
}

// NewUpdateSettingsGlobalParamsWithContext creates a new UpdateSettingsGlobalParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSettingsGlobalParamsWithContext(ctx context.Context) *UpdateSettingsGlobalParams {
	var ()
	return &UpdateSettingsGlobalParams{

		Context: ctx,
	}
}

// NewUpdateSettingsGlobalParamsWithHTTPClient creates a new UpdateSettingsGlobalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSettingsGlobalParamsWithHTTPClient(client *http.Client) *UpdateSettingsGlobalParams {
	var ()
	return &UpdateSettingsGlobalParams{
		HTTPClient: client,
	}
}

/*UpdateSettingsGlobalParams contains all the parameters to send to the API endpoint
for the update settings global operation typically these are written to a http.Request
*/
type UpdateSettingsGlobalParams struct {

	/*SettingsGlobal*/
	SettingsGlobal *models.SettingsGlobalSettings

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update settings global params
func (o *UpdateSettingsGlobalParams) WithTimeout(timeout time.Duration) *UpdateSettingsGlobalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update settings global params
func (o *UpdateSettingsGlobalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update settings global params
func (o *UpdateSettingsGlobalParams) WithContext(ctx context.Context) *UpdateSettingsGlobalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update settings global params
func (o *UpdateSettingsGlobalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update settings global params
func (o *UpdateSettingsGlobalParams) WithHTTPClient(client *http.Client) *UpdateSettingsGlobalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update settings global params
func (o *UpdateSettingsGlobalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSettingsGlobal adds the settingsGlobal to the update settings global params
func (o *UpdateSettingsGlobalParams) WithSettingsGlobal(settingsGlobal *models.SettingsGlobalSettings) *UpdateSettingsGlobalParams {
	o.SetSettingsGlobal(settingsGlobal)
	return o
}

// SetSettingsGlobal adds the settingsGlobal to the update settings global params
func (o *UpdateSettingsGlobalParams) SetSettingsGlobal(settingsGlobal *models.SettingsGlobalSettings) {
	o.SettingsGlobal = settingsGlobal
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSettingsGlobalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SettingsGlobal != nil {
		if err := r.SetBodyParam(o.SettingsGlobal); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
