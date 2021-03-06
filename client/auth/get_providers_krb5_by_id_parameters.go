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

// NewGetProvidersKrb5ByIDParams creates a new GetProvidersKrb5ByIDParams object
// with the default values initialized.
func NewGetProvidersKrb5ByIDParams() *GetProvidersKrb5ByIDParams {
	var ()
	return &GetProvidersKrb5ByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProvidersKrb5ByIDParamsWithTimeout creates a new GetProvidersKrb5ByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProvidersKrb5ByIDParamsWithTimeout(timeout time.Duration) *GetProvidersKrb5ByIDParams {
	var ()
	return &GetProvidersKrb5ByIDParams{

		timeout: timeout,
	}
}

// NewGetProvidersKrb5ByIDParamsWithContext creates a new GetProvidersKrb5ByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProvidersKrb5ByIDParamsWithContext(ctx context.Context) *GetProvidersKrb5ByIDParams {
	var ()
	return &GetProvidersKrb5ByIDParams{

		Context: ctx,
	}
}

// NewGetProvidersKrb5ByIDParamsWithHTTPClient creates a new GetProvidersKrb5ByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProvidersKrb5ByIDParamsWithHTTPClient(client *http.Client) *GetProvidersKrb5ByIDParams {
	var ()
	return &GetProvidersKrb5ByIDParams{
		HTTPClient: client,
	}
}

/*GetProvidersKrb5ByIDParams contains all the parameters to send to the API endpoint
for the get providers krb5 by Id operation typically these are written to a http.Request
*/
type GetProvidersKrb5ByIDParams struct {

	/*ProvidersKrb5ID
	  Retrieve the KRB5 provider.

	*/
	ProvidersKrb5ID string
	/*Scope
	  If specified as "effective" or not specified, all fields are returned.  If specified as "user", only fields with non-default values are shown.  If specified as "default", the original values are returned.

	*/
	Scope *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get providers krb5 by Id params
func (o *GetProvidersKrb5ByIDParams) WithTimeout(timeout time.Duration) *GetProvidersKrb5ByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get providers krb5 by Id params
func (o *GetProvidersKrb5ByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get providers krb5 by Id params
func (o *GetProvidersKrb5ByIDParams) WithContext(ctx context.Context) *GetProvidersKrb5ByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get providers krb5 by Id params
func (o *GetProvidersKrb5ByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get providers krb5 by Id params
func (o *GetProvidersKrb5ByIDParams) WithHTTPClient(client *http.Client) *GetProvidersKrb5ByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get providers krb5 by Id params
func (o *GetProvidersKrb5ByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProvidersKrb5ID adds the providersKrb5ID to the get providers krb5 by Id params
func (o *GetProvidersKrb5ByIDParams) WithProvidersKrb5ID(providersKrb5ID string) *GetProvidersKrb5ByIDParams {
	o.SetProvidersKrb5ID(providersKrb5ID)
	return o
}

// SetProvidersKrb5ID adds the providersKrb5Id to the get providers krb5 by Id params
func (o *GetProvidersKrb5ByIDParams) SetProvidersKrb5ID(providersKrb5ID string) {
	o.ProvidersKrb5ID = providersKrb5ID
}

// WithScope adds the scope to the get providers krb5 by Id params
func (o *GetProvidersKrb5ByIDParams) WithScope(scope *string) *GetProvidersKrb5ByIDParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the get providers krb5 by Id params
func (o *GetProvidersKrb5ByIDParams) SetScope(scope *string) {
	o.Scope = scope
}

// WriteToRequest writes these params to a swagger request
func (o *GetProvidersKrb5ByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ProvidersKrb5Id
	if err := r.SetPathParam("ProvidersKrb5Id", o.ProvidersKrb5ID); err != nil {
		return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
