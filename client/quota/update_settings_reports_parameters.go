// Code generated by go-swagger; DO NOT EDIT.

package quota

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

// NewUpdateSettingsReportsParams creates a new UpdateSettingsReportsParams object
// with the default values initialized.
func NewUpdateSettingsReportsParams() *UpdateSettingsReportsParams {
	var ()
	return &UpdateSettingsReportsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSettingsReportsParamsWithTimeout creates a new UpdateSettingsReportsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSettingsReportsParamsWithTimeout(timeout time.Duration) *UpdateSettingsReportsParams {
	var ()
	return &UpdateSettingsReportsParams{

		timeout: timeout,
	}
}

// NewUpdateSettingsReportsParamsWithContext creates a new UpdateSettingsReportsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSettingsReportsParamsWithContext(ctx context.Context) *UpdateSettingsReportsParams {
	var ()
	return &UpdateSettingsReportsParams{

		Context: ctx,
	}
}

// NewUpdateSettingsReportsParamsWithHTTPClient creates a new UpdateSettingsReportsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSettingsReportsParamsWithHTTPClient(client *http.Client) *UpdateSettingsReportsParams {
	var ()
	return &UpdateSettingsReportsParams{
		HTTPClient: client,
	}
}

/*UpdateSettingsReportsParams contains all the parameters to send to the API endpoint
for the update settings reports operation typically these are written to a http.Request
*/
type UpdateSettingsReportsParams struct {

	/*SettingsReports*/
	SettingsReports *models.SettingsReportsExtended

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update settings reports params
func (o *UpdateSettingsReportsParams) WithTimeout(timeout time.Duration) *UpdateSettingsReportsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update settings reports params
func (o *UpdateSettingsReportsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update settings reports params
func (o *UpdateSettingsReportsParams) WithContext(ctx context.Context) *UpdateSettingsReportsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update settings reports params
func (o *UpdateSettingsReportsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update settings reports params
func (o *UpdateSettingsReportsParams) WithHTTPClient(client *http.Client) *UpdateSettingsReportsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update settings reports params
func (o *UpdateSettingsReportsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSettingsReports adds the settingsReports to the update settings reports params
func (o *UpdateSettingsReportsParams) WithSettingsReports(settingsReports *models.SettingsReportsExtended) *UpdateSettingsReportsParams {
	o.SetSettingsReports(settingsReports)
	return o
}

// SetSettingsReports adds the settingsReports to the update settings reports params
func (o *UpdateSettingsReportsParams) SetSettingsReports(settingsReports *models.SettingsReportsExtended) {
	o.SettingsReports = settingsReports
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSettingsReportsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SettingsReports != nil {
		if err := r.SetBodyParam(o.SettingsReports); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
