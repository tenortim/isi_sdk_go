// Code generated by go-swagger; DO NOT EDIT.

package cloud

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

// NewUpdateCloudSettingsParams creates a new UpdateCloudSettingsParams object
// with the default values initialized.
func NewUpdateCloudSettingsParams() *UpdateCloudSettingsParams {
	var ()
	return &UpdateCloudSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCloudSettingsParamsWithTimeout creates a new UpdateCloudSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateCloudSettingsParamsWithTimeout(timeout time.Duration) *UpdateCloudSettingsParams {
	var ()
	return &UpdateCloudSettingsParams{

		timeout: timeout,
	}
}

// NewUpdateCloudSettingsParamsWithContext creates a new UpdateCloudSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateCloudSettingsParamsWithContext(ctx context.Context) *UpdateCloudSettingsParams {
	var ()
	return &UpdateCloudSettingsParams{

		Context: ctx,
	}
}

// NewUpdateCloudSettingsParamsWithHTTPClient creates a new UpdateCloudSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateCloudSettingsParamsWithHTTPClient(client *http.Client) *UpdateCloudSettingsParams {
	var ()
	return &UpdateCloudSettingsParams{
		HTTPClient: client,
	}
}

/*UpdateCloudSettingsParams contains all the parameters to send to the API endpoint
for the update cloud settings operation typically these are written to a http.Request
*/
type UpdateCloudSettingsParams struct {

	/*CloudSettings*/
	CloudSettings *models.CloudSettingsSettings

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update cloud settings params
func (o *UpdateCloudSettingsParams) WithTimeout(timeout time.Duration) *UpdateCloudSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update cloud settings params
func (o *UpdateCloudSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update cloud settings params
func (o *UpdateCloudSettingsParams) WithContext(ctx context.Context) *UpdateCloudSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update cloud settings params
func (o *UpdateCloudSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update cloud settings params
func (o *UpdateCloudSettingsParams) WithHTTPClient(client *http.Client) *UpdateCloudSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update cloud settings params
func (o *UpdateCloudSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudSettings adds the cloudSettings to the update cloud settings params
func (o *UpdateCloudSettingsParams) WithCloudSettings(cloudSettings *models.CloudSettingsSettings) *UpdateCloudSettingsParams {
	o.SetCloudSettings(cloudSettings)
	return o
}

// SetCloudSettings adds the cloudSettings to the update cloud settings params
func (o *UpdateCloudSettingsParams) SetCloudSettings(cloudSettings *models.CloudSettingsSettings) {
	o.CloudSettings = cloudSettings
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCloudSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CloudSettings != nil {
		if err := r.SetBodyParam(o.CloudSettings); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
