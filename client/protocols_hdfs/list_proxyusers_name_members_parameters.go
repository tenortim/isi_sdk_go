// Code generated by go-swagger; DO NOT EDIT.

package protocols_hdfs

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

// NewListProxyusersNameMembersParams creates a new ListProxyusersNameMembersParams object
// with the default values initialized.
func NewListProxyusersNameMembersParams() *ListProxyusersNameMembersParams {
	var ()
	return &ListProxyusersNameMembersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListProxyusersNameMembersParamsWithTimeout creates a new ListProxyusersNameMembersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListProxyusersNameMembersParamsWithTimeout(timeout time.Duration) *ListProxyusersNameMembersParams {
	var ()
	return &ListProxyusersNameMembersParams{

		timeout: timeout,
	}
}

// NewListProxyusersNameMembersParamsWithContext creates a new ListProxyusersNameMembersParams object
// with the default values initialized, and the ability to set a context for a request
func NewListProxyusersNameMembersParamsWithContext(ctx context.Context) *ListProxyusersNameMembersParams {
	var ()
	return &ListProxyusersNameMembersParams{

		Context: ctx,
	}
}

// NewListProxyusersNameMembersParamsWithHTTPClient creates a new ListProxyusersNameMembersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListProxyusersNameMembersParamsWithHTTPClient(client *http.Client) *ListProxyusersNameMembersParams {
	var ()
	return &ListProxyusersNameMembersParams{
		HTTPClient: client,
	}
}

/*ListProxyusersNameMembersParams contains all the parameters to send to the API endpoint
for the list proxyusers name members operation typically these are written to a http.Request
*/
type ListProxyusersNameMembersParams struct {

	/*Name*/
	Name string
	/*Zone
	  Specifies which access zone or zones to use.

	*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list proxyusers name members params
func (o *ListProxyusersNameMembersParams) WithTimeout(timeout time.Duration) *ListProxyusersNameMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list proxyusers name members params
func (o *ListProxyusersNameMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list proxyusers name members params
func (o *ListProxyusersNameMembersParams) WithContext(ctx context.Context) *ListProxyusersNameMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list proxyusers name members params
func (o *ListProxyusersNameMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list proxyusers name members params
func (o *ListProxyusersNameMembersParams) WithHTTPClient(client *http.Client) *ListProxyusersNameMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list proxyusers name members params
func (o *ListProxyusersNameMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the list proxyusers name members params
func (o *ListProxyusersNameMembersParams) WithName(name string) *ListProxyusersNameMembersParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the list proxyusers name members params
func (o *ListProxyusersNameMembersParams) SetName(name string) {
	o.Name = name
}

// WithZone adds the zone to the list proxyusers name members params
func (o *ListProxyusersNameMembersParams) WithZone(zone *string) *ListProxyusersNameMembersParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the list proxyusers name members params
func (o *ListProxyusersNameMembersParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *ListProxyusersNameMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Name
	if err := r.SetPathParam("Name", o.Name); err != nil {
		return err
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
