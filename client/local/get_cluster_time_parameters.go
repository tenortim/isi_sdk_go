// Code generated by go-swagger; DO NOT EDIT.

package local

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

// NewGetClusterTimeParams creates a new GetClusterTimeParams object
// with the default values initialized.
func NewGetClusterTimeParams() *GetClusterTimeParams {

	return &GetClusterTimeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterTimeParamsWithTimeout creates a new GetClusterTimeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClusterTimeParamsWithTimeout(timeout time.Duration) *GetClusterTimeParams {

	return &GetClusterTimeParams{

		timeout: timeout,
	}
}

// NewGetClusterTimeParamsWithContext creates a new GetClusterTimeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClusterTimeParamsWithContext(ctx context.Context) *GetClusterTimeParams {

	return &GetClusterTimeParams{

		Context: ctx,
	}
}

// NewGetClusterTimeParamsWithHTTPClient creates a new GetClusterTimeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClusterTimeParamsWithHTTPClient(client *http.Client) *GetClusterTimeParams {

	return &GetClusterTimeParams{
		HTTPClient: client,
	}
}

/*GetClusterTimeParams contains all the parameters to send to the API endpoint
for the get cluster time operation typically these are written to a http.Request
*/
type GetClusterTimeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cluster time params
func (o *GetClusterTimeParams) WithTimeout(timeout time.Duration) *GetClusterTimeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster time params
func (o *GetClusterTimeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster time params
func (o *GetClusterTimeParams) WithContext(ctx context.Context) *GetClusterTimeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster time params
func (o *GetClusterTimeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster time params
func (o *GetClusterTimeParams) WithHTTPClient(client *http.Client) *GetClusterTimeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster time params
func (o *GetClusterTimeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterTimeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
