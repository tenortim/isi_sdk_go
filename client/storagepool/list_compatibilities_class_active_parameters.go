// Code generated by go-swagger; DO NOT EDIT.

package storagepool

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

// NewListCompatibilitiesClassActiveParams creates a new ListCompatibilitiesClassActiveParams object
// with the default values initialized.
func NewListCompatibilitiesClassActiveParams() *ListCompatibilitiesClassActiveParams {

	return &ListCompatibilitiesClassActiveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListCompatibilitiesClassActiveParamsWithTimeout creates a new ListCompatibilitiesClassActiveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListCompatibilitiesClassActiveParamsWithTimeout(timeout time.Duration) *ListCompatibilitiesClassActiveParams {

	return &ListCompatibilitiesClassActiveParams{

		timeout: timeout,
	}
}

// NewListCompatibilitiesClassActiveParamsWithContext creates a new ListCompatibilitiesClassActiveParams object
// with the default values initialized, and the ability to set a context for a request
func NewListCompatibilitiesClassActiveParamsWithContext(ctx context.Context) *ListCompatibilitiesClassActiveParams {

	return &ListCompatibilitiesClassActiveParams{

		Context: ctx,
	}
}

// NewListCompatibilitiesClassActiveParamsWithHTTPClient creates a new ListCompatibilitiesClassActiveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListCompatibilitiesClassActiveParamsWithHTTPClient(client *http.Client) *ListCompatibilitiesClassActiveParams {

	return &ListCompatibilitiesClassActiveParams{
		HTTPClient: client,
	}
}

/*ListCompatibilitiesClassActiveParams contains all the parameters to send to the API endpoint
for the list compatibilities class active operation typically these are written to a http.Request
*/
type ListCompatibilitiesClassActiveParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list compatibilities class active params
func (o *ListCompatibilitiesClassActiveParams) WithTimeout(timeout time.Duration) *ListCompatibilitiesClassActiveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list compatibilities class active params
func (o *ListCompatibilitiesClassActiveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list compatibilities class active params
func (o *ListCompatibilitiesClassActiveParams) WithContext(ctx context.Context) *ListCompatibilitiesClassActiveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list compatibilities class active params
func (o *ListCompatibilitiesClassActiveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list compatibilities class active params
func (o *ListCompatibilitiesClassActiveParams) WithHTTPClient(client *http.Client) *ListCompatibilitiesClassActiveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list compatibilities class active params
func (o *ListCompatibilitiesClassActiveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListCompatibilitiesClassActiveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}