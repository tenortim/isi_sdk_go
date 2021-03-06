// Code generated by go-swagger; DO NOT EDIT.

package cluster_nodes

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

// NewGetNodeHardwareParams creates a new GetNodeHardwareParams object
// with the default values initialized.
func NewGetNodeHardwareParams() *GetNodeHardwareParams {
	var ()
	return &GetNodeHardwareParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNodeHardwareParamsWithTimeout creates a new GetNodeHardwareParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNodeHardwareParamsWithTimeout(timeout time.Duration) *GetNodeHardwareParams {
	var ()
	return &GetNodeHardwareParams{

		timeout: timeout,
	}
}

// NewGetNodeHardwareParamsWithContext creates a new GetNodeHardwareParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNodeHardwareParamsWithContext(ctx context.Context) *GetNodeHardwareParams {
	var ()
	return &GetNodeHardwareParams{

		Context: ctx,
	}
}

// NewGetNodeHardwareParamsWithHTTPClient creates a new GetNodeHardwareParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNodeHardwareParamsWithHTTPClient(client *http.Client) *GetNodeHardwareParams {
	var ()
	return &GetNodeHardwareParams{
		HTTPClient: client,
	}
}

/*GetNodeHardwareParams contains all the parameters to send to the API endpoint
for the get node hardware operation typically these are written to a http.Request
*/
type GetNodeHardwareParams struct {

	/*Lnn*/
	Lnn int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get node hardware params
func (o *GetNodeHardwareParams) WithTimeout(timeout time.Duration) *GetNodeHardwareParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get node hardware params
func (o *GetNodeHardwareParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get node hardware params
func (o *GetNodeHardwareParams) WithContext(ctx context.Context) *GetNodeHardwareParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get node hardware params
func (o *GetNodeHardwareParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get node hardware params
func (o *GetNodeHardwareParams) WithHTTPClient(client *http.Client) *GetNodeHardwareParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get node hardware params
func (o *GetNodeHardwareParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLnn adds the lnn to the get node hardware params
func (o *GetNodeHardwareParams) WithLnn(lnn int64) *GetNodeHardwareParams {
	o.SetLnn(lnn)
	return o
}

// SetLnn adds the lnn to the get node hardware params
func (o *GetNodeHardwareParams) SetLnn(lnn int64) {
	o.Lnn = lnn
}

// WriteToRequest writes these params to a swagger request
func (o *GetNodeHardwareParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Lnn
	if err := r.SetPathParam("Lnn", swag.FormatInt64(o.Lnn)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
