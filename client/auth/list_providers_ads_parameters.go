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

	strfmt "github.com/go-openapi/strfmt"
)

// NewListProvidersAdsParams creates a new ListProvidersAdsParams object
// with the default values initialized.
func NewListProvidersAdsParams() *ListProvidersAdsParams {
	var ()
	return &ListProvidersAdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListProvidersAdsParamsWithTimeout creates a new ListProvidersAdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListProvidersAdsParamsWithTimeout(timeout time.Duration) *ListProvidersAdsParams {
	var ()
	return &ListProvidersAdsParams{

		timeout: timeout,
	}
}

// NewListProvidersAdsParamsWithContext creates a new ListProvidersAdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListProvidersAdsParamsWithContext(ctx context.Context) *ListProvidersAdsParams {
	var ()
	return &ListProvidersAdsParams{

		Context: ctx,
	}
}

// NewListProvidersAdsParamsWithHTTPClient creates a new ListProvidersAdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListProvidersAdsParamsWithHTTPClient(client *http.Client) *ListProvidersAdsParams {
	var ()
	return &ListProvidersAdsParams{
		HTTPClient: client,
	}
}

/*ListProvidersAdsParams contains all the parameters to send to the API endpoint
for the list providers ads operation typically these are written to a http.Request
*/
type ListProvidersAdsParams struct {

	/*Scope
	  If specified as "effective" or not specified, all fields are returned.  If specified as "user", only fields with non-default values are shown.  If specified as "default", the original values are returned.

	*/
	Scope *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list providers ads params
func (o *ListProvidersAdsParams) WithTimeout(timeout time.Duration) *ListProvidersAdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list providers ads params
func (o *ListProvidersAdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list providers ads params
func (o *ListProvidersAdsParams) WithContext(ctx context.Context) *ListProvidersAdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list providers ads params
func (o *ListProvidersAdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list providers ads params
func (o *ListProvidersAdsParams) WithHTTPClient(client *http.Client) *ListProvidersAdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list providers ads params
func (o *ListProvidersAdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithScope adds the scope to the list providers ads params
func (o *ListProvidersAdsParams) WithScope(scope *string) *ListProvidersAdsParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the list providers ads params
func (o *ListProvidersAdsParams) SetScope(scope *string) {
	o.Scope = scope
}

// WriteToRequest writes these params to a swagger request
func (o *ListProvidersAdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
