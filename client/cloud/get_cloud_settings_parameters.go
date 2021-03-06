// Code generated by go-swagger; DO NOT EDIT.

package cloud

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

// NewGetCloudSettingsParams creates a new GetCloudSettingsParams object
// with the default values initialized.
func NewGetCloudSettingsParams() *GetCloudSettingsParams {

	return &GetCloudSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCloudSettingsParamsWithTimeout creates a new GetCloudSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCloudSettingsParamsWithTimeout(timeout time.Duration) *GetCloudSettingsParams {

	return &GetCloudSettingsParams{

		timeout: timeout,
	}
}

// NewGetCloudSettingsParamsWithContext creates a new GetCloudSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCloudSettingsParamsWithContext(ctx context.Context) *GetCloudSettingsParams {

	return &GetCloudSettingsParams{

		Context: ctx,
	}
}

// NewGetCloudSettingsParamsWithHTTPClient creates a new GetCloudSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCloudSettingsParamsWithHTTPClient(client *http.Client) *GetCloudSettingsParams {

	return &GetCloudSettingsParams{
		HTTPClient: client,
	}
}

/*GetCloudSettingsParams contains all the parameters to send to the API endpoint
for the get cloud settings operation typically these are written to a http.Request
*/
type GetCloudSettingsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cloud settings params
func (o *GetCloudSettingsParams) WithTimeout(timeout time.Duration) *GetCloudSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cloud settings params
func (o *GetCloudSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cloud settings params
func (o *GetCloudSettingsParams) WithContext(ctx context.Context) *GetCloudSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cloud settings params
func (o *GetCloudSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cloud settings params
func (o *GetCloudSettingsParams) WithHTTPClient(client *http.Client) *GetCloudSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cloud settings params
func (o *GetCloudSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetCloudSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
