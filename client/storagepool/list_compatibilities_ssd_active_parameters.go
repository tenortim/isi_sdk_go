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

// NewListCompatibilitiesSsdActiveParams creates a new ListCompatibilitiesSsdActiveParams object
// with the default values initialized.
func NewListCompatibilitiesSsdActiveParams() *ListCompatibilitiesSsdActiveParams {

	return &ListCompatibilitiesSsdActiveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListCompatibilitiesSsdActiveParamsWithTimeout creates a new ListCompatibilitiesSsdActiveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListCompatibilitiesSsdActiveParamsWithTimeout(timeout time.Duration) *ListCompatibilitiesSsdActiveParams {

	return &ListCompatibilitiesSsdActiveParams{

		timeout: timeout,
	}
}

// NewListCompatibilitiesSsdActiveParamsWithContext creates a new ListCompatibilitiesSsdActiveParams object
// with the default values initialized, and the ability to set a context for a request
func NewListCompatibilitiesSsdActiveParamsWithContext(ctx context.Context) *ListCompatibilitiesSsdActiveParams {

	return &ListCompatibilitiesSsdActiveParams{

		Context: ctx,
	}
}

// NewListCompatibilitiesSsdActiveParamsWithHTTPClient creates a new ListCompatibilitiesSsdActiveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListCompatibilitiesSsdActiveParamsWithHTTPClient(client *http.Client) *ListCompatibilitiesSsdActiveParams {

	return &ListCompatibilitiesSsdActiveParams{
		HTTPClient: client,
	}
}

/*ListCompatibilitiesSsdActiveParams contains all the parameters to send to the API endpoint
for the list compatibilities ssd active operation typically these are written to a http.Request
*/
type ListCompatibilitiesSsdActiveParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list compatibilities ssd active params
func (o *ListCompatibilitiesSsdActiveParams) WithTimeout(timeout time.Duration) *ListCompatibilitiesSsdActiveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list compatibilities ssd active params
func (o *ListCompatibilitiesSsdActiveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list compatibilities ssd active params
func (o *ListCompatibilitiesSsdActiveParams) WithContext(ctx context.Context) *ListCompatibilitiesSsdActiveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list compatibilities ssd active params
func (o *ListCompatibilitiesSsdActiveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list compatibilities ssd active params
func (o *ListCompatibilitiesSsdActiveParams) WithHTTPClient(client *http.Client) *ListCompatibilitiesSsdActiveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list compatibilities ssd active params
func (o *ListCompatibilitiesSsdActiveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListCompatibilitiesSsdActiveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
