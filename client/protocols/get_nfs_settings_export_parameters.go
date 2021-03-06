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
)

// NewGetNfsSettingsExportParams creates a new GetNfsSettingsExportParams object
// with the default values initialized.
func NewGetNfsSettingsExportParams() *GetNfsSettingsExportParams {
	var ()
	return &GetNfsSettingsExportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNfsSettingsExportParamsWithTimeout creates a new GetNfsSettingsExportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNfsSettingsExportParamsWithTimeout(timeout time.Duration) *GetNfsSettingsExportParams {
	var ()
	return &GetNfsSettingsExportParams{

		timeout: timeout,
	}
}

// NewGetNfsSettingsExportParamsWithContext creates a new GetNfsSettingsExportParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNfsSettingsExportParamsWithContext(ctx context.Context) *GetNfsSettingsExportParams {
	var ()
	return &GetNfsSettingsExportParams{

		Context: ctx,
	}
}

// NewGetNfsSettingsExportParamsWithHTTPClient creates a new GetNfsSettingsExportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNfsSettingsExportParamsWithHTTPClient(client *http.Client) *GetNfsSettingsExportParams {
	var ()
	return &GetNfsSettingsExportParams{
		HTTPClient: client,
	}
}

/*GetNfsSettingsExportParams contains all the parameters to send to the API endpoint
for the get nfs settings export operation typically these are written to a http.Request
*/
type GetNfsSettingsExportParams struct {

	/*Scope
	  If specified as "effective" or not specified, all fields are returned.  If specified as "user", only fields with non-default values are shown.  If specified as "default", the original values are returned.

	*/
	Scope *string
	/*Zone
	  Access zone

	*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get nfs settings export params
func (o *GetNfsSettingsExportParams) WithTimeout(timeout time.Duration) *GetNfsSettingsExportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nfs settings export params
func (o *GetNfsSettingsExportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nfs settings export params
func (o *GetNfsSettingsExportParams) WithContext(ctx context.Context) *GetNfsSettingsExportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nfs settings export params
func (o *GetNfsSettingsExportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nfs settings export params
func (o *GetNfsSettingsExportParams) WithHTTPClient(client *http.Client) *GetNfsSettingsExportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nfs settings export params
func (o *GetNfsSettingsExportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithScope adds the scope to the get nfs settings export params
func (o *GetNfsSettingsExportParams) WithScope(scope *string) *GetNfsSettingsExportParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the get nfs settings export params
func (o *GetNfsSettingsExportParams) SetScope(scope *string) {
	o.Scope = scope
}

// WithZone adds the zone to the get nfs settings export params
func (o *GetNfsSettingsExportParams) WithZone(zone *string) *GetNfsSettingsExportParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the get nfs settings export params
func (o *GetNfsSettingsExportParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *GetNfsSettingsExportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Scope != nil {

		// query param scope
		var qrScope string
		if o.Scope != nil {
			qrScope = *o.Scope
		}
		qScope := qrScope
		if qScope != "" {
			if err := r.SetQueryParam("scope", qScope); err != nil {
				return err
			}
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
