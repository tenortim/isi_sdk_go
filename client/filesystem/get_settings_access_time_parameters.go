// Code generated by go-swagger; DO NOT EDIT.

package filesystem

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

// NewGetSettingsAccessTimeParams creates a new GetSettingsAccessTimeParams object
// with the default values initialized.
func NewGetSettingsAccessTimeParams() *GetSettingsAccessTimeParams {

	return &GetSettingsAccessTimeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSettingsAccessTimeParamsWithTimeout creates a new GetSettingsAccessTimeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSettingsAccessTimeParamsWithTimeout(timeout time.Duration) *GetSettingsAccessTimeParams {

	return &GetSettingsAccessTimeParams{

		timeout: timeout,
	}
}

// NewGetSettingsAccessTimeParamsWithContext creates a new GetSettingsAccessTimeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSettingsAccessTimeParamsWithContext(ctx context.Context) *GetSettingsAccessTimeParams {

	return &GetSettingsAccessTimeParams{

		Context: ctx,
	}
}

// NewGetSettingsAccessTimeParamsWithHTTPClient creates a new GetSettingsAccessTimeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSettingsAccessTimeParamsWithHTTPClient(client *http.Client) *GetSettingsAccessTimeParams {

	return &GetSettingsAccessTimeParams{
		HTTPClient: client,
	}
}

/*GetSettingsAccessTimeParams contains all the parameters to send to the API endpoint
for the get settings access time operation typically these are written to a http.Request
*/
type GetSettingsAccessTimeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get settings access time params
func (o *GetSettingsAccessTimeParams) WithTimeout(timeout time.Duration) *GetSettingsAccessTimeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get settings access time params
func (o *GetSettingsAccessTimeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get settings access time params
func (o *GetSettingsAccessTimeParams) WithContext(ctx context.Context) *GetSettingsAccessTimeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get settings access time params
func (o *GetSettingsAccessTimeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get settings access time params
func (o *GetSettingsAccessTimeParams) WithHTTPClient(client *http.Client) *GetSettingsAccessTimeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get settings access time params
func (o *GetSettingsAccessTimeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSettingsAccessTimeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
