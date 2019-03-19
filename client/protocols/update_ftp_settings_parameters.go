// Code generated by go-swagger; DO NOT EDIT.

package protocols

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

// NewUpdateFtpSettingsParams creates a new UpdateFtpSettingsParams object
// with the default values initialized.
func NewUpdateFtpSettingsParams() *UpdateFtpSettingsParams {
	var ()
	return &UpdateFtpSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateFtpSettingsParamsWithTimeout creates a new UpdateFtpSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateFtpSettingsParamsWithTimeout(timeout time.Duration) *UpdateFtpSettingsParams {
	var ()
	return &UpdateFtpSettingsParams{

		timeout: timeout,
	}
}

// NewUpdateFtpSettingsParamsWithContext creates a new UpdateFtpSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateFtpSettingsParamsWithContext(ctx context.Context) *UpdateFtpSettingsParams {
	var ()
	return &UpdateFtpSettingsParams{

		Context: ctx,
	}
}

// NewUpdateFtpSettingsParamsWithHTTPClient creates a new UpdateFtpSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateFtpSettingsParamsWithHTTPClient(client *http.Client) *UpdateFtpSettingsParams {
	var ()
	return &UpdateFtpSettingsParams{
		HTTPClient: client,
	}
}

/*UpdateFtpSettingsParams contains all the parameters to send to the API endpoint
for the update ftp settings operation typically these are written to a http.Request
*/
type UpdateFtpSettingsParams struct {

	/*FtpSettings*/
	FtpSettings *models.FtpSettingsExtended

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update ftp settings params
func (o *UpdateFtpSettingsParams) WithTimeout(timeout time.Duration) *UpdateFtpSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update ftp settings params
func (o *UpdateFtpSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update ftp settings params
func (o *UpdateFtpSettingsParams) WithContext(ctx context.Context) *UpdateFtpSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update ftp settings params
func (o *UpdateFtpSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update ftp settings params
func (o *UpdateFtpSettingsParams) WithHTTPClient(client *http.Client) *UpdateFtpSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update ftp settings params
func (o *UpdateFtpSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFtpSettings adds the ftpSettings to the update ftp settings params
func (o *UpdateFtpSettingsParams) WithFtpSettings(ftpSettings *models.FtpSettingsExtended) *UpdateFtpSettingsParams {
	o.SetFtpSettings(ftpSettings)
	return o
}

// SetFtpSettings adds the ftpSettings to the update ftp settings params
func (o *UpdateFtpSettingsParams) SetFtpSettings(ftpSettings *models.FtpSettingsExtended) {
	o.FtpSettings = ftpSettings
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateFtpSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FtpSettings != nil {
		if err := r.SetBodyParam(o.FtpSettings); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
