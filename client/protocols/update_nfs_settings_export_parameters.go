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

// NewUpdateNfsSettingsExportParams creates a new UpdateNfsSettingsExportParams object
// with the default values initialized.
func NewUpdateNfsSettingsExportParams() *UpdateNfsSettingsExportParams {
	var ()
	return &UpdateNfsSettingsExportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNfsSettingsExportParamsWithTimeout creates a new UpdateNfsSettingsExportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateNfsSettingsExportParamsWithTimeout(timeout time.Duration) *UpdateNfsSettingsExportParams {
	var ()
	return &UpdateNfsSettingsExportParams{

		timeout: timeout,
	}
}

// NewUpdateNfsSettingsExportParamsWithContext creates a new UpdateNfsSettingsExportParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateNfsSettingsExportParamsWithContext(ctx context.Context) *UpdateNfsSettingsExportParams {
	var ()
	return &UpdateNfsSettingsExportParams{

		Context: ctx,
	}
}

// NewUpdateNfsSettingsExportParamsWithHTTPClient creates a new UpdateNfsSettingsExportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateNfsSettingsExportParamsWithHTTPClient(client *http.Client) *UpdateNfsSettingsExportParams {
	var ()
	return &UpdateNfsSettingsExportParams{
		HTTPClient: client,
	}
}

/*UpdateNfsSettingsExportParams contains all the parameters to send to the API endpoint
for the update nfs settings export operation typically these are written to a http.Request
*/
type UpdateNfsSettingsExportParams struct {

	/*NfsSettingsExport*/
	NfsSettingsExport *models.NfsSettingsExportSettings
	/*Zone
	  Access zone

	*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update nfs settings export params
func (o *UpdateNfsSettingsExportParams) WithTimeout(timeout time.Duration) *UpdateNfsSettingsExportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update nfs settings export params
func (o *UpdateNfsSettingsExportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update nfs settings export params
func (o *UpdateNfsSettingsExportParams) WithContext(ctx context.Context) *UpdateNfsSettingsExportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update nfs settings export params
func (o *UpdateNfsSettingsExportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update nfs settings export params
func (o *UpdateNfsSettingsExportParams) WithHTTPClient(client *http.Client) *UpdateNfsSettingsExportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update nfs settings export params
func (o *UpdateNfsSettingsExportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNfsSettingsExport adds the nfsSettingsExport to the update nfs settings export params
func (o *UpdateNfsSettingsExportParams) WithNfsSettingsExport(nfsSettingsExport *models.NfsSettingsExportSettings) *UpdateNfsSettingsExportParams {
	o.SetNfsSettingsExport(nfsSettingsExport)
	return o
}

// SetNfsSettingsExport adds the nfsSettingsExport to the update nfs settings export params
func (o *UpdateNfsSettingsExportParams) SetNfsSettingsExport(nfsSettingsExport *models.NfsSettingsExportSettings) {
	o.NfsSettingsExport = nfsSettingsExport
}

// WithZone adds the zone to the update nfs settings export params
func (o *UpdateNfsSettingsExportParams) WithZone(zone *string) *UpdateNfsSettingsExportParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the update nfs settings export params
func (o *UpdateNfsSettingsExportParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNfsSettingsExportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.NfsSettingsExport != nil {
		if err := r.SetBodyParam(o.NfsSettingsExport); err != nil {
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
