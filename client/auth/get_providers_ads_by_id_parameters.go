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

// NewGetProvidersAdsByIDParams creates a new GetProvidersAdsByIDParams object
// with the default values initialized.
func NewGetProvidersAdsByIDParams() *GetProvidersAdsByIDParams {
	var ()
	return &GetProvidersAdsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProvidersAdsByIDParamsWithTimeout creates a new GetProvidersAdsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProvidersAdsByIDParamsWithTimeout(timeout time.Duration) *GetProvidersAdsByIDParams {
	var ()
	return &GetProvidersAdsByIDParams{

		timeout: timeout,
	}
}

// NewGetProvidersAdsByIDParamsWithContext creates a new GetProvidersAdsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProvidersAdsByIDParamsWithContext(ctx context.Context) *GetProvidersAdsByIDParams {
	var ()
	return &GetProvidersAdsByIDParams{

		Context: ctx,
	}
}

// NewGetProvidersAdsByIDParamsWithHTTPClient creates a new GetProvidersAdsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProvidersAdsByIDParamsWithHTTPClient(client *http.Client) *GetProvidersAdsByIDParams {
	var ()
	return &GetProvidersAdsByIDParams{
		HTTPClient: client,
	}
}

/*GetProvidersAdsByIDParams contains all the parameters to send to the API endpoint
for the get providers ads by Id operation typically these are written to a http.Request
*/
type GetProvidersAdsByIDParams struct {

	/*ProvidersAdsID
	  Retrieve the ADS provider.

	*/
	ProvidersAdsID string
	/*Scope
	  If specified as "effective" or not specified, all fields are returned.  If specified as "user", only fields with non-default values are shown.  If specified as "default", the original values are returned.

	*/
	Scope *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get providers ads by Id params
func (o *GetProvidersAdsByIDParams) WithTimeout(timeout time.Duration) *GetProvidersAdsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get providers ads by Id params
func (o *GetProvidersAdsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get providers ads by Id params
func (o *GetProvidersAdsByIDParams) WithContext(ctx context.Context) *GetProvidersAdsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get providers ads by Id params
func (o *GetProvidersAdsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get providers ads by Id params
func (o *GetProvidersAdsByIDParams) WithHTTPClient(client *http.Client) *GetProvidersAdsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get providers ads by Id params
func (o *GetProvidersAdsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProvidersAdsID adds the providersAdsID to the get providers ads by Id params
func (o *GetProvidersAdsByIDParams) WithProvidersAdsID(providersAdsID string) *GetProvidersAdsByIDParams {
	o.SetProvidersAdsID(providersAdsID)
	return o
}

// SetProvidersAdsID adds the providersAdsId to the get providers ads by Id params
func (o *GetProvidersAdsByIDParams) SetProvidersAdsID(providersAdsID string) {
	o.ProvidersAdsID = providersAdsID
}

// WithScope adds the scope to the get providers ads by Id params
func (o *GetProvidersAdsByIDParams) WithScope(scope *string) *GetProvidersAdsByIDParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the get providers ads by Id params
func (o *GetProvidersAdsByIDParams) SetScope(scope *string) {
	o.Scope = scope
}

// WriteToRequest writes these params to a swagger request
func (o *GetProvidersAdsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ProvidersAdsId
	if err := r.SetPathParam("ProvidersAdsId", o.ProvidersAdsID); err != nil {
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
