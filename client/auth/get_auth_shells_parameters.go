// Code generated by go-swagger; DO NOT EDIT.

package auth

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

// NewGetAuthShellsParams creates a new GetAuthShellsParams object
// with the default values initialized.
func NewGetAuthShellsParams() *GetAuthShellsParams {

	return &GetAuthShellsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthShellsParamsWithTimeout creates a new GetAuthShellsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAuthShellsParamsWithTimeout(timeout time.Duration) *GetAuthShellsParams {

	return &GetAuthShellsParams{

		timeout: timeout,
	}
}

// NewGetAuthShellsParamsWithContext creates a new GetAuthShellsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAuthShellsParamsWithContext(ctx context.Context) *GetAuthShellsParams {

	return &GetAuthShellsParams{

		Context: ctx,
	}
}

// NewGetAuthShellsParamsWithHTTPClient creates a new GetAuthShellsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAuthShellsParamsWithHTTPClient(client *http.Client) *GetAuthShellsParams {

	return &GetAuthShellsParams{
		HTTPClient: client,
	}
}

/*GetAuthShellsParams contains all the parameters to send to the API endpoint
for the get auth shells operation typically these are written to a http.Request
*/
type GetAuthShellsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get auth shells params
func (o *GetAuthShellsParams) WithTimeout(timeout time.Duration) *GetAuthShellsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get auth shells params
func (o *GetAuthShellsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get auth shells params
func (o *GetAuthShellsParams) WithContext(ctx context.Context) *GetAuthShellsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get auth shells params
func (o *GetAuthShellsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get auth shells params
func (o *GetAuthShellsParams) WithHTTPClient(client *http.Client) *GetAuthShellsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get auth shells params
func (o *GetAuthShellsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthShellsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}