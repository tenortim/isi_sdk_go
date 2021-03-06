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

// NewGetFilepoolDefaultPolicyParams creates a new GetFilepoolDefaultPolicyParams object
// with the default values initialized.
func NewGetFilepoolDefaultPolicyParams() *GetFilepoolDefaultPolicyParams {

	return &GetFilepoolDefaultPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFilepoolDefaultPolicyParamsWithTimeout creates a new GetFilepoolDefaultPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFilepoolDefaultPolicyParamsWithTimeout(timeout time.Duration) *GetFilepoolDefaultPolicyParams {

	return &GetFilepoolDefaultPolicyParams{

		timeout: timeout,
	}
}

// NewGetFilepoolDefaultPolicyParamsWithContext creates a new GetFilepoolDefaultPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFilepoolDefaultPolicyParamsWithContext(ctx context.Context) *GetFilepoolDefaultPolicyParams {

	return &GetFilepoolDefaultPolicyParams{

		Context: ctx,
	}
}

// NewGetFilepoolDefaultPolicyParamsWithHTTPClient creates a new GetFilepoolDefaultPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFilepoolDefaultPolicyParamsWithHTTPClient(client *http.Client) *GetFilepoolDefaultPolicyParams {

	return &GetFilepoolDefaultPolicyParams{
		HTTPClient: client,
	}
}

/*GetFilepoolDefaultPolicyParams contains all the parameters to send to the API endpoint
for the get filepool default policy operation typically these are written to a http.Request
*/
type GetFilepoolDefaultPolicyParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get filepool default policy params
func (o *GetFilepoolDefaultPolicyParams) WithTimeout(timeout time.Duration) *GetFilepoolDefaultPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get filepool default policy params
func (o *GetFilepoolDefaultPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get filepool default policy params
func (o *GetFilepoolDefaultPolicyParams) WithContext(ctx context.Context) *GetFilepoolDefaultPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get filepool default policy params
func (o *GetFilepoolDefaultPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get filepool default policy params
func (o *GetFilepoolDefaultPolicyParams) WithHTTPClient(client *http.Client) *GetFilepoolDefaultPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get filepool default policy params
func (o *GetFilepoolDefaultPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetFilepoolDefaultPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
