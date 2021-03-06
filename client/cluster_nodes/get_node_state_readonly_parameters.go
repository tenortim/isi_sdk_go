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

// NewGetNodeStateReadonlyParams creates a new GetNodeStateReadonlyParams object
// with the default values initialized.
func NewGetNodeStateReadonlyParams() *GetNodeStateReadonlyParams {
	var ()
	return &GetNodeStateReadonlyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNodeStateReadonlyParamsWithTimeout creates a new GetNodeStateReadonlyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNodeStateReadonlyParamsWithTimeout(timeout time.Duration) *GetNodeStateReadonlyParams {
	var ()
	return &GetNodeStateReadonlyParams{

		timeout: timeout,
	}
}

// NewGetNodeStateReadonlyParamsWithContext creates a new GetNodeStateReadonlyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNodeStateReadonlyParamsWithContext(ctx context.Context) *GetNodeStateReadonlyParams {
	var ()
	return &GetNodeStateReadonlyParams{

		Context: ctx,
	}
}

// NewGetNodeStateReadonlyParamsWithHTTPClient creates a new GetNodeStateReadonlyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNodeStateReadonlyParamsWithHTTPClient(client *http.Client) *GetNodeStateReadonlyParams {
	var ()
	return &GetNodeStateReadonlyParams{
		HTTPClient: client,
	}
}

/*GetNodeStateReadonlyParams contains all the parameters to send to the API endpoint
for the get node state readonly operation typically these are written to a http.Request
*/
type GetNodeStateReadonlyParams struct {

	/*Lnn*/
	Lnn int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get node state readonly params
func (o *GetNodeStateReadonlyParams) WithTimeout(timeout time.Duration) *GetNodeStateReadonlyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get node state readonly params
func (o *GetNodeStateReadonlyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get node state readonly params
func (o *GetNodeStateReadonlyParams) WithContext(ctx context.Context) *GetNodeStateReadonlyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get node state readonly params
func (o *GetNodeStateReadonlyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get node state readonly params
func (o *GetNodeStateReadonlyParams) WithHTTPClient(client *http.Client) *GetNodeStateReadonlyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get node state readonly params
func (o *GetNodeStateReadonlyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLnn adds the lnn to the get node state readonly params
func (o *GetNodeStateReadonlyParams) WithLnn(lnn int64) *GetNodeStateReadonlyParams {
	o.SetLnn(lnn)
	return o
}

// SetLnn adds the lnn to the get node state readonly params
func (o *GetNodeStateReadonlyParams) SetLnn(lnn int64) {
	o.Lnn = lnn
}

// WriteToRequest writes these params to a swagger request
func (o *GetNodeStateReadonlyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
