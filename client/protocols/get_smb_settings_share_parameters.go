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

// NewGetSmbSettingsShareParams creates a new GetSmbSettingsShareParams object
// with the default values initialized.
func NewGetSmbSettingsShareParams() *GetSmbSettingsShareParams {
	var ()
	return &GetSmbSettingsShareParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSmbSettingsShareParamsWithTimeout creates a new GetSmbSettingsShareParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSmbSettingsShareParamsWithTimeout(timeout time.Duration) *GetSmbSettingsShareParams {
	var ()
	return &GetSmbSettingsShareParams{

		timeout: timeout,
	}
}

// NewGetSmbSettingsShareParamsWithContext creates a new GetSmbSettingsShareParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSmbSettingsShareParamsWithContext(ctx context.Context) *GetSmbSettingsShareParams {
	var ()
	return &GetSmbSettingsShareParams{

		Context: ctx,
	}
}

// NewGetSmbSettingsShareParamsWithHTTPClient creates a new GetSmbSettingsShareParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSmbSettingsShareParamsWithHTTPClient(client *http.Client) *GetSmbSettingsShareParams {
	var ()
	return &GetSmbSettingsShareParams{
		HTTPClient: client,
	}
}

/*GetSmbSettingsShareParams contains all the parameters to send to the API endpoint
for the get smb settings share operation typically these are written to a http.Request
*/
type GetSmbSettingsShareParams struct {

	/*Scope
	  If specified as "effective" or not specified, all fields are returned.  If specified as "user", only fields with non-default values are shown.  If specified as "default", the original values are returned.

	*/
	Scope *string
	/*Zone
	  Zone which contains these share settings.

	*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get smb settings share params
func (o *GetSmbSettingsShareParams) WithTimeout(timeout time.Duration) *GetSmbSettingsShareParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get smb settings share params
func (o *GetSmbSettingsShareParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get smb settings share params
func (o *GetSmbSettingsShareParams) WithContext(ctx context.Context) *GetSmbSettingsShareParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get smb settings share params
func (o *GetSmbSettingsShareParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get smb settings share params
func (o *GetSmbSettingsShareParams) WithHTTPClient(client *http.Client) *GetSmbSettingsShareParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get smb settings share params
func (o *GetSmbSettingsShareParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithScope adds the scope to the get smb settings share params
func (o *GetSmbSettingsShareParams) WithScope(scope *string) *GetSmbSettingsShareParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the get smb settings share params
func (o *GetSmbSettingsShareParams) SetScope(scope *string) {
	o.Scope = scope
}

// WithZone adds the zone to the get smb settings share params
func (o *GetSmbSettingsShareParams) WithZone(zone *string) *GetSmbSettingsShareParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the get smb settings share params
func (o *GetSmbSettingsShareParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *GetSmbSettingsShareParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
