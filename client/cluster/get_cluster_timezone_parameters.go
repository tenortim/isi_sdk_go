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

// NewGetClusterTimezoneParams creates a new GetClusterTimezoneParams object
// with the default values initialized.
func NewGetClusterTimezoneParams() *GetClusterTimezoneParams {

	return &GetClusterTimezoneParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterTimezoneParamsWithTimeout creates a new GetClusterTimezoneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClusterTimezoneParamsWithTimeout(timeout time.Duration) *GetClusterTimezoneParams {

	return &GetClusterTimezoneParams{

		timeout: timeout,
	}
}

// NewGetClusterTimezoneParamsWithContext creates a new GetClusterTimezoneParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClusterTimezoneParamsWithContext(ctx context.Context) *GetClusterTimezoneParams {

	return &GetClusterTimezoneParams{

		Context: ctx,
	}
}

// NewGetClusterTimezoneParamsWithHTTPClient creates a new GetClusterTimezoneParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClusterTimezoneParamsWithHTTPClient(client *http.Client) *GetClusterTimezoneParams {

	return &GetClusterTimezoneParams{
		HTTPClient: client,
	}
}

/*GetClusterTimezoneParams contains all the parameters to send to the API endpoint
for the get cluster timezone operation typically these are written to a http.Request
*/
type GetClusterTimezoneParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cluster timezone params
func (o *GetClusterTimezoneParams) WithTimeout(timeout time.Duration) *GetClusterTimezoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster timezone params
func (o *GetClusterTimezoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster timezone params
func (o *GetClusterTimezoneParams) WithContext(ctx context.Context) *GetClusterTimezoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster timezone params
func (o *GetClusterTimezoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster timezone params
func (o *GetClusterTimezoneParams) WithHTTPClient(client *http.Client) *GetClusterTimezoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster timezone params
func (o *GetClusterTimezoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterTimezoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
