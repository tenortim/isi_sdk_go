// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// NewGetClusterExternalIpsParams creates a new GetClusterExternalIpsParams object
// with the default values initialized.
func NewGetClusterExternalIpsParams() *GetClusterExternalIpsParams {

	return &GetClusterExternalIpsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterExternalIpsParamsWithTimeout creates a new GetClusterExternalIpsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClusterExternalIpsParamsWithTimeout(timeout time.Duration) *GetClusterExternalIpsParams {

	return &GetClusterExternalIpsParams{

		timeout: timeout,
	}
}

// NewGetClusterExternalIpsParamsWithContext creates a new GetClusterExternalIpsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClusterExternalIpsParamsWithContext(ctx context.Context) *GetClusterExternalIpsParams {

	return &GetClusterExternalIpsParams{

		Context: ctx,
	}
}

// NewGetClusterExternalIpsParamsWithHTTPClient creates a new GetClusterExternalIpsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClusterExternalIpsParamsWithHTTPClient(client *http.Client) *GetClusterExternalIpsParams {

	return &GetClusterExternalIpsParams{
		HTTPClient: client,
	}
}

/*GetClusterExternalIpsParams contains all the parameters to send to the API endpoint
for the get cluster external ips operation typically these are written to a http.Request
*/
type GetClusterExternalIpsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cluster external ips params
func (o *GetClusterExternalIpsParams) WithTimeout(timeout time.Duration) *GetClusterExternalIpsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster external ips params
func (o *GetClusterExternalIpsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster external ips params
func (o *GetClusterExternalIpsParams) WithContext(ctx context.Context) *GetClusterExternalIpsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster external ips params
func (o *GetClusterExternalIpsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster external ips params
func (o *GetClusterExternalIpsParams) WithHTTPClient(client *http.Client) *GetClusterExternalIpsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster external ips params
func (o *GetClusterExternalIpsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterExternalIpsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}