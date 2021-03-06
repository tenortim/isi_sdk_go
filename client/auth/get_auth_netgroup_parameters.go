// Code generated by go-swagger; DO NOT EDIT.

package auth

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

// NewGetAuthNetgroupParams creates a new GetAuthNetgroupParams object
// with the default values initialized.
func NewGetAuthNetgroupParams() *GetAuthNetgroupParams {
	var ()
	return &GetAuthNetgroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthNetgroupParamsWithTimeout creates a new GetAuthNetgroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAuthNetgroupParamsWithTimeout(timeout time.Duration) *GetAuthNetgroupParams {
	var ()
	return &GetAuthNetgroupParams{

		timeout: timeout,
	}
}

// NewGetAuthNetgroupParamsWithContext creates a new GetAuthNetgroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAuthNetgroupParamsWithContext(ctx context.Context) *GetAuthNetgroupParams {
	var ()
	return &GetAuthNetgroupParams{

		Context: ctx,
	}
}

// NewGetAuthNetgroupParamsWithHTTPClient creates a new GetAuthNetgroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAuthNetgroupParamsWithHTTPClient(client *http.Client) *GetAuthNetgroupParams {
	var ()
	return &GetAuthNetgroupParams{
		HTTPClient: client,
	}
}

/*GetAuthNetgroupParams contains all the parameters to send to the API endpoint
for the get auth netgroup operation typically these are written to a http.Request
*/
type GetAuthNetgroupParams struct {

	/*AuthNetgroupID
	  Retrieve the user information.

	*/
	AuthNetgroupID string
	/*IgnoreErrors
	  Ignore netgroup errors.

	*/
	IgnoreErrors *bool
	/*Provider
	  Filter users by provider.

	*/
	Provider *string
	/*Recursive
	  Perform recursive search.

	*/
	Recursive *bool
	/*Zone
	  Filter users by zone.

	*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get auth netgroup params
func (o *GetAuthNetgroupParams) WithTimeout(timeout time.Duration) *GetAuthNetgroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get auth netgroup params
func (o *GetAuthNetgroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get auth netgroup params
func (o *GetAuthNetgroupParams) WithContext(ctx context.Context) *GetAuthNetgroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get auth netgroup params
func (o *GetAuthNetgroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get auth netgroup params
func (o *GetAuthNetgroupParams) WithHTTPClient(client *http.Client) *GetAuthNetgroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get auth netgroup params
func (o *GetAuthNetgroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthNetgroupID adds the authNetgroupID to the get auth netgroup params
func (o *GetAuthNetgroupParams) WithAuthNetgroupID(authNetgroupID string) *GetAuthNetgroupParams {
	o.SetAuthNetgroupID(authNetgroupID)
	return o
}

// SetAuthNetgroupID adds the authNetgroupId to the get auth netgroup params
func (o *GetAuthNetgroupParams) SetAuthNetgroupID(authNetgroupID string) {
	o.AuthNetgroupID = authNetgroupID
}

// WithIgnoreErrors adds the ignoreErrors to the get auth netgroup params
func (o *GetAuthNetgroupParams) WithIgnoreErrors(ignoreErrors *bool) *GetAuthNetgroupParams {
	o.SetIgnoreErrors(ignoreErrors)
	return o
}

// SetIgnoreErrors adds the ignoreErrors to the get auth netgroup params
func (o *GetAuthNetgroupParams) SetIgnoreErrors(ignoreErrors *bool) {
	o.IgnoreErrors = ignoreErrors
}

// WithProvider adds the provider to the get auth netgroup params
func (o *GetAuthNetgroupParams) WithProvider(provider *string) *GetAuthNetgroupParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the get auth netgroup params
func (o *GetAuthNetgroupParams) SetProvider(provider *string) {
	o.Provider = provider
}

// WithRecursive adds the recursive to the get auth netgroup params
func (o *GetAuthNetgroupParams) WithRecursive(recursive *bool) *GetAuthNetgroupParams {
	o.SetRecursive(recursive)
	return o
}

// SetRecursive adds the recursive to the get auth netgroup params
func (o *GetAuthNetgroupParams) SetRecursive(recursive *bool) {
	o.Recursive = recursive
}

// WithZone adds the zone to the get auth netgroup params
func (o *GetAuthNetgroupParams) WithZone(zone *string) *GetAuthNetgroupParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the get auth netgroup params
func (o *GetAuthNetgroupParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthNetgroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param AuthNetgroupId
	if err := r.SetPathParam("AuthNetgroupId", o.AuthNetgroupID); err != nil {
		return err
	}

	if o.IgnoreErrors != nil {

		// query param ignore_errors
		var qrIgnoreErrors bool
		if o.IgnoreErrors != nil {
			qrIgnoreErrors = *o.IgnoreErrors
		}
		qIgnoreErrors := swag.FormatBool(qrIgnoreErrors)
		if qIgnoreErrors != "" {
			if err := r.SetQueryParam("ignore_errors", qIgnoreErrors); err != nil {
				return err
			}
		}

	}

	if o.Provider != nil {

		// query param provider
		var qrProvider string
		if o.Provider != nil {
			qrProvider = *o.Provider
		}
		qProvider := qrProvider
		if qProvider != "" {
			if err := r.SetQueryParam("provider", qProvider); err != nil {
				return err
			}
		}

	}

	if o.Recursive != nil {

		// query param recursive
		var qrRecursive bool
		if o.Recursive != nil {
			qrRecursive = *o.Recursive
		}
		qRecursive := swag.FormatBool(qrRecursive)
		if qRecursive != "" {
			if err := r.SetQueryParam("recursive", qRecursive); err != nil {
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
