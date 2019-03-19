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

// NewGetNfsCheckParams creates a new GetNfsCheckParams object
// with the default values initialized.
func NewGetNfsCheckParams() *GetNfsCheckParams {
	var ()
	return &GetNfsCheckParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNfsCheckParamsWithTimeout creates a new GetNfsCheckParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNfsCheckParamsWithTimeout(timeout time.Duration) *GetNfsCheckParams {
	var ()
	return &GetNfsCheckParams{

		timeout: timeout,
	}
}

// NewGetNfsCheckParamsWithContext creates a new GetNfsCheckParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNfsCheckParamsWithContext(ctx context.Context) *GetNfsCheckParams {
	var ()
	return &GetNfsCheckParams{

		Context: ctx,
	}
}

// NewGetNfsCheckParamsWithHTTPClient creates a new GetNfsCheckParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNfsCheckParamsWithHTTPClient(client *http.Client) *GetNfsCheckParams {
	var ()
	return &GetNfsCheckParams{
		HTTPClient: client,
	}
}

/*GetNfsCheckParams contains all the parameters to send to the API endpoint
for the get nfs check operation typically these are written to a http.Request
*/
type GetNfsCheckParams struct {

	/*IgnoreBadAuth
	  Ignore invalid users.

	*/
	IgnoreBadAuth *bool
	/*IgnoreBadPaths
	  Ignore nonexistent or otherwise bad paths.

	*/
	IgnoreBadPaths *bool
	/*IgnoreUnresolvableHosts
	  Ignore unresolvable hosts.

	*/
	IgnoreUnresolvableHosts *bool
	/*Zone
	  Access zone

	*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get nfs check params
func (o *GetNfsCheckParams) WithTimeout(timeout time.Duration) *GetNfsCheckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nfs check params
func (o *GetNfsCheckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nfs check params
func (o *GetNfsCheckParams) WithContext(ctx context.Context) *GetNfsCheckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nfs check params
func (o *GetNfsCheckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nfs check params
func (o *GetNfsCheckParams) WithHTTPClient(client *http.Client) *GetNfsCheckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nfs check params
func (o *GetNfsCheckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIgnoreBadAuth adds the ignoreBadAuth to the get nfs check params
func (o *GetNfsCheckParams) WithIgnoreBadAuth(ignoreBadAuth *bool) *GetNfsCheckParams {
	o.SetIgnoreBadAuth(ignoreBadAuth)
	return o
}

// SetIgnoreBadAuth adds the ignoreBadAuth to the get nfs check params
func (o *GetNfsCheckParams) SetIgnoreBadAuth(ignoreBadAuth *bool) {
	o.IgnoreBadAuth = ignoreBadAuth
}

// WithIgnoreBadPaths adds the ignoreBadPaths to the get nfs check params
func (o *GetNfsCheckParams) WithIgnoreBadPaths(ignoreBadPaths *bool) *GetNfsCheckParams {
	o.SetIgnoreBadPaths(ignoreBadPaths)
	return o
}

// SetIgnoreBadPaths adds the ignoreBadPaths to the get nfs check params
func (o *GetNfsCheckParams) SetIgnoreBadPaths(ignoreBadPaths *bool) {
	o.IgnoreBadPaths = ignoreBadPaths
}

// WithIgnoreUnresolvableHosts adds the ignoreUnresolvableHosts to the get nfs check params
func (o *GetNfsCheckParams) WithIgnoreUnresolvableHosts(ignoreUnresolvableHosts *bool) *GetNfsCheckParams {
	o.SetIgnoreUnresolvableHosts(ignoreUnresolvableHosts)
	return o
}

// SetIgnoreUnresolvableHosts adds the ignoreUnresolvableHosts to the get nfs check params
func (o *GetNfsCheckParams) SetIgnoreUnresolvableHosts(ignoreUnresolvableHosts *bool) {
	o.IgnoreUnresolvableHosts = ignoreUnresolvableHosts
}

// WithZone adds the zone to the get nfs check params
func (o *GetNfsCheckParams) WithZone(zone *string) *GetNfsCheckParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the get nfs check params
func (o *GetNfsCheckParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *GetNfsCheckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IgnoreBadAuth != nil {

		// query param ignore_bad_auth
		var qrIgnoreBadAuth bool
		if o.IgnoreBadAuth != nil {
			qrIgnoreBadAuth = *o.IgnoreBadAuth
		}
		qIgnoreBadAuth := swag.FormatBool(qrIgnoreBadAuth)
		if qIgnoreBadAuth != "" {
			if err := r.SetQueryParam("ignore_bad_auth", qIgnoreBadAuth); err != nil {
				return err
			}
		}

	}

	if o.IgnoreBadPaths != nil {

		// query param ignore_bad_paths
		var qrIgnoreBadPaths bool
		if o.IgnoreBadPaths != nil {
			qrIgnoreBadPaths = *o.IgnoreBadPaths
		}
		qIgnoreBadPaths := swag.FormatBool(qrIgnoreBadPaths)
		if qIgnoreBadPaths != "" {
			if err := r.SetQueryParam("ignore_bad_paths", qIgnoreBadPaths); err != nil {
				return err
			}
		}

	}

	if o.IgnoreUnresolvableHosts != nil {

		// query param ignore_unresolvable_hosts
		var qrIgnoreUnresolvableHosts bool
		if o.IgnoreUnresolvableHosts != nil {
			qrIgnoreUnresolvableHosts = *o.IgnoreUnresolvableHosts
		}
		qIgnoreUnresolvableHosts := swag.FormatBool(qrIgnoreUnresolvableHosts)
		if qIgnoreUnresolvableHosts != "" {
			if err := r.SetQueryParam("ignore_unresolvable_hosts", qIgnoreUnresolvableHosts); err != nil {
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
