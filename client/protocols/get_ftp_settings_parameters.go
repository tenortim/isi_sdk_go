// Code generated by go-swagger; DO NOT EDIT.

package protocols

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

// NewGetFtpSettingsParams creates a new GetFtpSettingsParams object
// with the default values initialized.
func NewGetFtpSettingsParams() *GetFtpSettingsParams {

	return &GetFtpSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFtpSettingsParamsWithTimeout creates a new GetFtpSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFtpSettingsParamsWithTimeout(timeout time.Duration) *GetFtpSettingsParams {

	return &GetFtpSettingsParams{

		timeout: timeout,
	}
}

// NewGetFtpSettingsParamsWithContext creates a new GetFtpSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFtpSettingsParamsWithContext(ctx context.Context) *GetFtpSettingsParams {

	return &GetFtpSettingsParams{

		Context: ctx,
	}
}

// NewGetFtpSettingsParamsWithHTTPClient creates a new GetFtpSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFtpSettingsParamsWithHTTPClient(client *http.Client) *GetFtpSettingsParams {

	return &GetFtpSettingsParams{
		HTTPClient: client,
	}
}

/*GetFtpSettingsParams contains all the parameters to send to the API endpoint
for the get ftp settings operation typically these are written to a http.Request
*/
type GetFtpSettingsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ftp settings params
func (o *GetFtpSettingsParams) WithTimeout(timeout time.Duration) *GetFtpSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ftp settings params
func (o *GetFtpSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ftp settings params
func (o *GetFtpSettingsParams) WithContext(ctx context.Context) *GetFtpSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ftp settings params
func (o *GetFtpSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ftp settings params
func (o *GetFtpSettingsParams) WithHTTPClient(client *http.Client) *GetFtpSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ftp settings params
func (o *GetFtpSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetFtpSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}