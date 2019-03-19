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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSmbShareParams creates a new GetSmbShareParams object
// with the default values initialized.
func NewGetSmbShareParams() *GetSmbShareParams {
	var ()
	return &GetSmbShareParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSmbShareParamsWithTimeout creates a new GetSmbShareParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSmbShareParamsWithTimeout(timeout time.Duration) *GetSmbShareParams {
	var ()
	return &GetSmbShareParams{

		timeout: timeout,
	}
}

// NewGetSmbShareParamsWithContext creates a new GetSmbShareParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSmbShareParamsWithContext(ctx context.Context) *GetSmbShareParams {
	var ()
	return &GetSmbShareParams{

		Context: ctx,
	}
}

// NewGetSmbShareParamsWithHTTPClient creates a new GetSmbShareParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSmbShareParamsWithHTTPClient(client *http.Client) *GetSmbShareParams {
	var ()
	return &GetSmbShareParams{
		HTTPClient: client,
	}
}

/*GetSmbShareParams contains all the parameters to send to the API endpoint
for the get smb share operation typically these are written to a http.Request
*/
type GetSmbShareParams struct {

	/*SmbShareID
	  Retrieve share.

	*/
	SmbShareID string
	/*ResolveNames
	  If true, resolve group and user names in personas.

	*/
	ResolveNames *bool
	/*Scope
	  If specified as "effective" or not specified, all fields are returned.  If specified as "user", only fields with non-default values are shown.  If specified as "default", the original values are returned.

	*/
	Scope *string
	/*Zone
	  Zone which contains this share.

	*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get smb share params
func (o *GetSmbShareParams) WithTimeout(timeout time.Duration) *GetSmbShareParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get smb share params
func (o *GetSmbShareParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get smb share params
func (o *GetSmbShareParams) WithContext(ctx context.Context) *GetSmbShareParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get smb share params
func (o *GetSmbShareParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get smb share params
func (o *GetSmbShareParams) WithHTTPClient(client *http.Client) *GetSmbShareParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get smb share params
func (o *GetSmbShareParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSmbShareID adds the smbShareID to the get smb share params
func (o *GetSmbShareParams) WithSmbShareID(smbShareID string) *GetSmbShareParams {
	o.SetSmbShareID(smbShareID)
	return o
}

// SetSmbShareID adds the smbShareId to the get smb share params
func (o *GetSmbShareParams) SetSmbShareID(smbShareID string) {
	o.SmbShareID = smbShareID
}

// WithResolveNames adds the resolveNames to the get smb share params
func (o *GetSmbShareParams) WithResolveNames(resolveNames *bool) *GetSmbShareParams {
	o.SetResolveNames(resolveNames)
	return o
}

// SetResolveNames adds the resolveNames to the get smb share params
func (o *GetSmbShareParams) SetResolveNames(resolveNames *bool) {
	o.ResolveNames = resolveNames
}

// WithScope adds the scope to the get smb share params
func (o *GetSmbShareParams) WithScope(scope *string) *GetSmbShareParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the get smb share params
func (o *GetSmbShareParams) SetScope(scope *string) {
	o.Scope = scope
}

// WithZone adds the zone to the get smb share params
func (o *GetSmbShareParams) WithZone(zone *string) *GetSmbShareParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the get smb share params
func (o *GetSmbShareParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *GetSmbShareParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param SmbShareId
	if err := r.SetPathParam("SmbShareId", o.SmbShareID); err != nil {
		return err
	}

	if o.ResolveNames != nil {

		// query param resolve_names
		var qrResolveNames bool
		if o.ResolveNames != nil {
			qrResolveNames = *o.ResolveNames
		}
		qResolveNames := swag.FormatBool(qrResolveNames)
		if qResolveNames != "" {
			if err := r.SetQueryParam("resolve_names", qResolveNames); err != nil {
				return err
			}
		}

	}

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