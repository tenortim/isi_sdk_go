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

// NewGetProvidersLocalParams creates a new GetProvidersLocalParams object
// with the default values initialized.
func NewGetProvidersLocalParams() *GetProvidersLocalParams {
	var ()
	return &GetProvidersLocalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProvidersLocalParamsWithTimeout creates a new GetProvidersLocalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProvidersLocalParamsWithTimeout(timeout time.Duration) *GetProvidersLocalParams {
	var ()
	return &GetProvidersLocalParams{

		timeout: timeout,
	}
}

// NewGetProvidersLocalParamsWithContext creates a new GetProvidersLocalParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProvidersLocalParamsWithContext(ctx context.Context) *GetProvidersLocalParams {
	var ()
	return &GetProvidersLocalParams{

		Context: ctx,
	}
}

// NewGetProvidersLocalParamsWithHTTPClient creates a new GetProvidersLocalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProvidersLocalParamsWithHTTPClient(client *http.Client) *GetProvidersLocalParams {
	var ()
	return &GetProvidersLocalParams{
		HTTPClient: client,
	}
}

/*GetProvidersLocalParams contains all the parameters to send to the API endpoint
for the get providers local operation typically these are written to a http.Request
*/
type GetProvidersLocalParams struct {

	/*Scope
	  If specified as "effective" or not specified, all fields are returned.  If specified as "user", only fields with non-default values are shown.  If specified as "default", the original values are returned.

	*/
	Scope *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get providers local params
func (o *GetProvidersLocalParams) WithTimeout(timeout time.Duration) *GetProvidersLocalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get providers local params
func (o *GetProvidersLocalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get providers local params
func (o *GetProvidersLocalParams) WithContext(ctx context.Context) *GetProvidersLocalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get providers local params
func (o *GetProvidersLocalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get providers local params
func (o *GetProvidersLocalParams) WithHTTPClient(client *http.Client) *GetProvidersLocalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get providers local params
func (o *GetProvidersLocalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithScope adds the scope to the get providers local params
func (o *GetProvidersLocalParams) WithScope(scope *string) *GetProvidersLocalParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the get providers local params
func (o *GetProvidersLocalParams) SetScope(scope *string) {
	o.Scope = scope
}

// WriteToRequest writes these params to a swagger request
func (o *GetProvidersLocalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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