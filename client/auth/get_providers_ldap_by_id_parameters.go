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

// NewGetProvidersLdapByIDParams creates a new GetProvidersLdapByIDParams object
// with the default values initialized.
func NewGetProvidersLdapByIDParams() *GetProvidersLdapByIDParams {
	var ()
	return &GetProvidersLdapByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProvidersLdapByIDParamsWithTimeout creates a new GetProvidersLdapByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProvidersLdapByIDParamsWithTimeout(timeout time.Duration) *GetProvidersLdapByIDParams {
	var ()
	return &GetProvidersLdapByIDParams{

		timeout: timeout,
	}
}

// NewGetProvidersLdapByIDParamsWithContext creates a new GetProvidersLdapByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProvidersLdapByIDParamsWithContext(ctx context.Context) *GetProvidersLdapByIDParams {
	var ()
	return &GetProvidersLdapByIDParams{

		Context: ctx,
	}
}

// NewGetProvidersLdapByIDParamsWithHTTPClient creates a new GetProvidersLdapByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProvidersLdapByIDParamsWithHTTPClient(client *http.Client) *GetProvidersLdapByIDParams {
	var ()
	return &GetProvidersLdapByIDParams{
		HTTPClient: client,
	}
}

/*GetProvidersLdapByIDParams contains all the parameters to send to the API endpoint
for the get providers ldap by Id operation typically these are written to a http.Request
*/
type GetProvidersLdapByIDParams struct {

	/*ProvidersLdapID
	  Retrieve the LDAP provider.

	*/
	ProvidersLdapID string
	/*Scope
	  If specified as "effective" or not specified, all fields are returned.  If specified as "user", only fields with non-default values are shown.  If specified as "default", the original values are returned.

	*/
	Scope *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get providers ldap by Id params
func (o *GetProvidersLdapByIDParams) WithTimeout(timeout time.Duration) *GetProvidersLdapByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get providers ldap by Id params
func (o *GetProvidersLdapByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get providers ldap by Id params
func (o *GetProvidersLdapByIDParams) WithContext(ctx context.Context) *GetProvidersLdapByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get providers ldap by Id params
func (o *GetProvidersLdapByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get providers ldap by Id params
func (o *GetProvidersLdapByIDParams) WithHTTPClient(client *http.Client) *GetProvidersLdapByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get providers ldap by Id params
func (o *GetProvidersLdapByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProvidersLdapID adds the providersLdapID to the get providers ldap by Id params
func (o *GetProvidersLdapByIDParams) WithProvidersLdapID(providersLdapID string) *GetProvidersLdapByIDParams {
	o.SetProvidersLdapID(providersLdapID)
	return o
}

// SetProvidersLdapID adds the providersLdapId to the get providers ldap by Id params
func (o *GetProvidersLdapByIDParams) SetProvidersLdapID(providersLdapID string) {
	o.ProvidersLdapID = providersLdapID
}

// WithScope adds the scope to the get providers ldap by Id params
func (o *GetProvidersLdapByIDParams) WithScope(scope *string) *GetProvidersLdapByIDParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the get providers ldap by Id params
func (o *GetProvidersLdapByIDParams) SetScope(scope *string) {
	o.Scope = scope
}

// WriteToRequest writes these params to a swagger request
func (o *GetProvidersLdapByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ProvidersLdapId
	if err := r.SetPathParam("ProvidersLdapId", o.ProvidersLdapID); err != nil {
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
