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

// NewUpdateNfsSettingsZoneParams creates a new UpdateNfsSettingsZoneParams object
// with the default values initialized.
func NewUpdateNfsSettingsZoneParams() *UpdateNfsSettingsZoneParams {
	var ()
	return &UpdateNfsSettingsZoneParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNfsSettingsZoneParamsWithTimeout creates a new UpdateNfsSettingsZoneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateNfsSettingsZoneParamsWithTimeout(timeout time.Duration) *UpdateNfsSettingsZoneParams {
	var ()
	return &UpdateNfsSettingsZoneParams{

		timeout: timeout,
	}
}

// NewUpdateNfsSettingsZoneParamsWithContext creates a new UpdateNfsSettingsZoneParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateNfsSettingsZoneParamsWithContext(ctx context.Context) *UpdateNfsSettingsZoneParams {
	var ()
	return &UpdateNfsSettingsZoneParams{

		Context: ctx,
	}
}

// NewUpdateNfsSettingsZoneParamsWithHTTPClient creates a new UpdateNfsSettingsZoneParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateNfsSettingsZoneParamsWithHTTPClient(client *http.Client) *UpdateNfsSettingsZoneParams {
	var ()
	return &UpdateNfsSettingsZoneParams{
		HTTPClient: client,
	}
}

/*UpdateNfsSettingsZoneParams contains all the parameters to send to the API endpoint
for the update nfs settings zone operation typically these are written to a http.Request
*/
type UpdateNfsSettingsZoneParams struct {

	/*NfsSettingsZone*/
	NfsSettingsZone *models.NfsSettingsZoneSettings
	/*Zone
	  Access zone

	*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update nfs settings zone params
func (o *UpdateNfsSettingsZoneParams) WithTimeout(timeout time.Duration) *UpdateNfsSettingsZoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update nfs settings zone params
func (o *UpdateNfsSettingsZoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update nfs settings zone params
func (o *UpdateNfsSettingsZoneParams) WithContext(ctx context.Context) *UpdateNfsSettingsZoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update nfs settings zone params
func (o *UpdateNfsSettingsZoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update nfs settings zone params
func (o *UpdateNfsSettingsZoneParams) WithHTTPClient(client *http.Client) *UpdateNfsSettingsZoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update nfs settings zone params
func (o *UpdateNfsSettingsZoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNfsSettingsZone adds the nfsSettingsZone to the update nfs settings zone params
func (o *UpdateNfsSettingsZoneParams) WithNfsSettingsZone(nfsSettingsZone *models.NfsSettingsZoneSettings) *UpdateNfsSettingsZoneParams {
	o.SetNfsSettingsZone(nfsSettingsZone)
	return o
}

// SetNfsSettingsZone adds the nfsSettingsZone to the update nfs settings zone params
func (o *UpdateNfsSettingsZoneParams) SetNfsSettingsZone(nfsSettingsZone *models.NfsSettingsZoneSettings) {
	o.NfsSettingsZone = nfsSettingsZone
}

// WithZone adds the zone to the update nfs settings zone params
func (o *UpdateNfsSettingsZoneParams) WithZone(zone *string) *UpdateNfsSettingsZoneParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the update nfs settings zone params
func (o *UpdateNfsSettingsZoneParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNfsSettingsZoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.NfsSettingsZone != nil {
		if err := r.SetBodyParam(o.NfsSettingsZone); err != nil {
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
