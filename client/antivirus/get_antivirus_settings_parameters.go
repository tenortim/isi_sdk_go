// Code generated by go-swagger; DO NOT EDIT.

package antivirus

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

// NewGetAntivirusSettingsParams creates a new GetAntivirusSettingsParams object
// with the default values initialized.
func NewGetAntivirusSettingsParams() *GetAntivirusSettingsParams {

	return &GetAntivirusSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAntivirusSettingsParamsWithTimeout creates a new GetAntivirusSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAntivirusSettingsParamsWithTimeout(timeout time.Duration) *GetAntivirusSettingsParams {

	return &GetAntivirusSettingsParams{

		timeout: timeout,
	}
}

// NewGetAntivirusSettingsParamsWithContext creates a new GetAntivirusSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAntivirusSettingsParamsWithContext(ctx context.Context) *GetAntivirusSettingsParams {

	return &GetAntivirusSettingsParams{

		Context: ctx,
	}
}

// NewGetAntivirusSettingsParamsWithHTTPClient creates a new GetAntivirusSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAntivirusSettingsParamsWithHTTPClient(client *http.Client) *GetAntivirusSettingsParams {

	return &GetAntivirusSettingsParams{
		HTTPClient: client,
	}
}

/*GetAntivirusSettingsParams contains all the parameters to send to the API endpoint
for the get antivirus settings operation typically these are written to a http.Request
*/
type GetAntivirusSettingsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get antivirus settings params
func (o *GetAntivirusSettingsParams) WithTimeout(timeout time.Duration) *GetAntivirusSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get antivirus settings params
func (o *GetAntivirusSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get antivirus settings params
func (o *GetAntivirusSettingsParams) WithContext(ctx context.Context) *GetAntivirusSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get antivirus settings params
func (o *GetAntivirusSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get antivirus settings params
func (o *GetAntivirusSettingsParams) WithHTTPClient(client *http.Client) *GetAntivirusSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get antivirus settings params
func (o *GetAntivirusSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAntivirusSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
