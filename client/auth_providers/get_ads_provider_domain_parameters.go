// Code generated by go-swagger; DO NOT EDIT.

package auth_providers

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

// NewGetAdsProviderDomainParams creates a new GetAdsProviderDomainParams object
// with the default values initialized.
func NewGetAdsProviderDomainParams() *GetAdsProviderDomainParams {
	var ()
	return &GetAdsProviderDomainParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAdsProviderDomainParamsWithTimeout creates a new GetAdsProviderDomainParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAdsProviderDomainParamsWithTimeout(timeout time.Duration) *GetAdsProviderDomainParams {
	var ()
	return &GetAdsProviderDomainParams{

		timeout: timeout,
	}
}

// NewGetAdsProviderDomainParamsWithContext creates a new GetAdsProviderDomainParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAdsProviderDomainParamsWithContext(ctx context.Context) *GetAdsProviderDomainParams {
	var ()
	return &GetAdsProviderDomainParams{

		Context: ctx,
	}
}

// NewGetAdsProviderDomainParamsWithHTTPClient creates a new GetAdsProviderDomainParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAdsProviderDomainParamsWithHTTPClient(client *http.Client) *GetAdsProviderDomainParams {
	var ()
	return &GetAdsProviderDomainParams{
		HTTPClient: client,
	}
}

/*GetAdsProviderDomainParams contains all the parameters to send to the API endpoint
for the get ads provider domain operation typically these are written to a http.Request
*/
type GetAdsProviderDomainParams struct {

	/*AdsProviderDomainID
	  Retrieve the ADS domain information.

	*/
	AdsProviderDomainID string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ads provider domain params
func (o *GetAdsProviderDomainParams) WithTimeout(timeout time.Duration) *GetAdsProviderDomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ads provider domain params
func (o *GetAdsProviderDomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ads provider domain params
func (o *GetAdsProviderDomainParams) WithContext(ctx context.Context) *GetAdsProviderDomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ads provider domain params
func (o *GetAdsProviderDomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ads provider domain params
func (o *GetAdsProviderDomainParams) WithHTTPClient(client *http.Client) *GetAdsProviderDomainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ads provider domain params
func (o *GetAdsProviderDomainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdsProviderDomainID adds the adsProviderDomainID to the get ads provider domain params
func (o *GetAdsProviderDomainParams) WithAdsProviderDomainID(adsProviderDomainID string) *GetAdsProviderDomainParams {
	o.SetAdsProviderDomainID(adsProviderDomainID)
	return o
}

// SetAdsProviderDomainID adds the adsProviderDomainId to the get ads provider domain params
func (o *GetAdsProviderDomainParams) SetAdsProviderDomainID(adsProviderDomainID string) {
	o.AdsProviderDomainID = adsProviderDomainID
}

// WithID adds the id to the get ads provider domain params
func (o *GetAdsProviderDomainParams) WithID(id string) *GetAdsProviderDomainParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get ads provider domain params
func (o *GetAdsProviderDomainParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAdsProviderDomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param AdsProviderDomainId
	if err := r.SetPathParam("AdsProviderDomainId", o.AdsProviderDomainID); err != nil {
		return err
	}

	// path param Id
	if err := r.SetPathParam("Id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
