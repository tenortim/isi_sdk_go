// Code generated by go-swagger; DO NOT EDIT.

package protocols

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

// NewGetNfsNetgroupParams creates a new GetNfsNetgroupParams object
// with the default values initialized.
func NewGetNfsNetgroupParams() *GetNfsNetgroupParams {
	var ()
	return &GetNfsNetgroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNfsNetgroupParamsWithTimeout creates a new GetNfsNetgroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNfsNetgroupParamsWithTimeout(timeout time.Duration) *GetNfsNetgroupParams {
	var ()
	return &GetNfsNetgroupParams{

		timeout: timeout,
	}
}

// NewGetNfsNetgroupParamsWithContext creates a new GetNfsNetgroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNfsNetgroupParamsWithContext(ctx context.Context) *GetNfsNetgroupParams {
	var ()
	return &GetNfsNetgroupParams{

		Context: ctx,
	}
}

// NewGetNfsNetgroupParamsWithHTTPClient creates a new GetNfsNetgroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNfsNetgroupParamsWithHTTPClient(client *http.Client) *GetNfsNetgroupParams {
	var ()
	return &GetNfsNetgroupParams{
		HTTPClient: client,
	}
}

/*GetNfsNetgroupParams contains all the parameters to send to the API endpoint
for the get nfs netgroup operation typically these are written to a http.Request
*/
type GetNfsNetgroupParams struct {

	/*Host
	  Host to retrieve netgroup cache settings from.

	*/
	Host *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get nfs netgroup params
func (o *GetNfsNetgroupParams) WithTimeout(timeout time.Duration) *GetNfsNetgroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nfs netgroup params
func (o *GetNfsNetgroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nfs netgroup params
func (o *GetNfsNetgroupParams) WithContext(ctx context.Context) *GetNfsNetgroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nfs netgroup params
func (o *GetNfsNetgroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nfs netgroup params
func (o *GetNfsNetgroupParams) WithHTTPClient(client *http.Client) *GetNfsNetgroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nfs netgroup params
func (o *GetNfsNetgroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHost adds the host to the get nfs netgroup params
func (o *GetNfsNetgroupParams) WithHost(host *string) *GetNfsNetgroupParams {
	o.SetHost(host)
	return o
}

// SetHost adds the host to the get nfs netgroup params
func (o *GetNfsNetgroupParams) SetHost(host *string) {
	o.Host = host
}

// WriteToRequest writes these params to a swagger request
func (o *GetNfsNetgroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Host != nil {

		// query param host
		var qrHost string
		if o.Host != nil {
			qrHost = *o.Host
		}
		qHost := qrHost
		if qHost != "" {
			if err := r.SetQueryParam("host", qHost); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
