// Code generated by go-swagger; DO NOT EDIT.

package filepool

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

// NewGetFilepoolTemplatesParams creates a new GetFilepoolTemplatesParams object
// with the default values initialized.
func NewGetFilepoolTemplatesParams() *GetFilepoolTemplatesParams {

	return &GetFilepoolTemplatesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFilepoolTemplatesParamsWithTimeout creates a new GetFilepoolTemplatesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFilepoolTemplatesParamsWithTimeout(timeout time.Duration) *GetFilepoolTemplatesParams {

	return &GetFilepoolTemplatesParams{

		timeout: timeout,
	}
}

// NewGetFilepoolTemplatesParamsWithContext creates a new GetFilepoolTemplatesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFilepoolTemplatesParamsWithContext(ctx context.Context) *GetFilepoolTemplatesParams {

	return &GetFilepoolTemplatesParams{

		Context: ctx,
	}
}

// NewGetFilepoolTemplatesParamsWithHTTPClient creates a new GetFilepoolTemplatesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFilepoolTemplatesParamsWithHTTPClient(client *http.Client) *GetFilepoolTemplatesParams {

	return &GetFilepoolTemplatesParams{
		HTTPClient: client,
	}
}

/*GetFilepoolTemplatesParams contains all the parameters to send to the API endpoint
for the get filepool templates operation typically these are written to a http.Request
*/
type GetFilepoolTemplatesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get filepool templates params
func (o *GetFilepoolTemplatesParams) WithTimeout(timeout time.Duration) *GetFilepoolTemplatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get filepool templates params
func (o *GetFilepoolTemplatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get filepool templates params
func (o *GetFilepoolTemplatesParams) WithContext(ctx context.Context) *GetFilepoolTemplatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get filepool templates params
func (o *GetFilepoolTemplatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get filepool templates params
func (o *GetFilepoolTemplatesParams) WithHTTPClient(client *http.Client) *GetFilepoolTemplatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get filepool templates params
func (o *GetFilepoolTemplatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetFilepoolTemplatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
