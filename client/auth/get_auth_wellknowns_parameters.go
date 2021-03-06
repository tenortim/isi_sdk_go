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

// NewGetAuthWellknownsParams creates a new GetAuthWellknownsParams object
// with the default values initialized.
func NewGetAuthWellknownsParams() *GetAuthWellknownsParams {

	return &GetAuthWellknownsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthWellknownsParamsWithTimeout creates a new GetAuthWellknownsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAuthWellknownsParamsWithTimeout(timeout time.Duration) *GetAuthWellknownsParams {

	return &GetAuthWellknownsParams{

		timeout: timeout,
	}
}

// NewGetAuthWellknownsParamsWithContext creates a new GetAuthWellknownsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAuthWellknownsParamsWithContext(ctx context.Context) *GetAuthWellknownsParams {

	return &GetAuthWellknownsParams{

		Context: ctx,
	}
}

// NewGetAuthWellknownsParamsWithHTTPClient creates a new GetAuthWellknownsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAuthWellknownsParamsWithHTTPClient(client *http.Client) *GetAuthWellknownsParams {

	return &GetAuthWellknownsParams{
		HTTPClient: client,
	}
}

/*GetAuthWellknownsParams contains all the parameters to send to the API endpoint
for the get auth wellknowns operation typically these are written to a http.Request
*/
type GetAuthWellknownsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get auth wellknowns params
func (o *GetAuthWellknownsParams) WithTimeout(timeout time.Duration) *GetAuthWellknownsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get auth wellknowns params
func (o *GetAuthWellknownsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get auth wellknowns params
func (o *GetAuthWellknownsParams) WithContext(ctx context.Context) *GetAuthWellknownsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get auth wellknowns params
func (o *GetAuthWellknownsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get auth wellknowns params
func (o *GetAuthWellknownsParams) WithHTTPClient(client *http.Client) *GetAuthWellknownsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get auth wellknowns params
func (o *GetAuthWellknownsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthWellknownsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
