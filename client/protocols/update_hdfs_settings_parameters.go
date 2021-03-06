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

// NewUpdateHdfsSettingsParams creates a new UpdateHdfsSettingsParams object
// with the default values initialized.
func NewUpdateHdfsSettingsParams() *UpdateHdfsSettingsParams {
	var ()
	return &UpdateHdfsSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateHdfsSettingsParamsWithTimeout creates a new UpdateHdfsSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateHdfsSettingsParamsWithTimeout(timeout time.Duration) *UpdateHdfsSettingsParams {
	var ()
	return &UpdateHdfsSettingsParams{

		timeout: timeout,
	}
}

// NewUpdateHdfsSettingsParamsWithContext creates a new UpdateHdfsSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateHdfsSettingsParamsWithContext(ctx context.Context) *UpdateHdfsSettingsParams {
	var ()
	return &UpdateHdfsSettingsParams{

		Context: ctx,
	}
}

// NewUpdateHdfsSettingsParamsWithHTTPClient creates a new UpdateHdfsSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateHdfsSettingsParamsWithHTTPClient(client *http.Client) *UpdateHdfsSettingsParams {
	var ()
	return &UpdateHdfsSettingsParams{
		HTTPClient: client,
	}
}

/*UpdateHdfsSettingsParams contains all the parameters to send to the API endpoint
for the update hdfs settings operation typically these are written to a http.Request
*/
type UpdateHdfsSettingsParams struct {

	/*HdfsSettings*/
	HdfsSettings *models.HdfsSettingsSettings
	/*Zone
	  Access zone which contains HDFS settings.

	*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update hdfs settings params
func (o *UpdateHdfsSettingsParams) WithTimeout(timeout time.Duration) *UpdateHdfsSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update hdfs settings params
func (o *UpdateHdfsSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update hdfs settings params
func (o *UpdateHdfsSettingsParams) WithContext(ctx context.Context) *UpdateHdfsSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update hdfs settings params
func (o *UpdateHdfsSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update hdfs settings params
func (o *UpdateHdfsSettingsParams) WithHTTPClient(client *http.Client) *UpdateHdfsSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update hdfs settings params
func (o *UpdateHdfsSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHdfsSettings adds the hdfsSettings to the update hdfs settings params
func (o *UpdateHdfsSettingsParams) WithHdfsSettings(hdfsSettings *models.HdfsSettingsSettings) *UpdateHdfsSettingsParams {
	o.SetHdfsSettings(hdfsSettings)
	return o
}

// SetHdfsSettings adds the hdfsSettings to the update hdfs settings params
func (o *UpdateHdfsSettingsParams) SetHdfsSettings(hdfsSettings *models.HdfsSettingsSettings) {
	o.HdfsSettings = hdfsSettings
}

// WithZone adds the zone to the update hdfs settings params
func (o *UpdateHdfsSettingsParams) WithZone(zone *string) *UpdateHdfsSettingsParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the update hdfs settings params
func (o *UpdateHdfsSettingsParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateHdfsSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.HdfsSettings != nil {
		if err := r.SetBodyParam(o.HdfsSettings); err != nil {
			return err
		}
	}

	if o.Zone != nil {

		// query param zone
		var qrZone string
		if o.Zone != nil {
			qrZone = *o.Zone
		}
		qZone := qrZone
		if qZone != "" {
			if err := r.SetQueryParam("zone", qZone); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
