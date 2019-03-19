// Code generated by go-swagger; DO NOT EDIT.

package sync

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

// NewGetSyncLicenseParams creates a new GetSyncLicenseParams object
// with the default values initialized.
func NewGetSyncLicenseParams() *GetSyncLicenseParams {

	return &GetSyncLicenseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSyncLicenseParamsWithTimeout creates a new GetSyncLicenseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSyncLicenseParamsWithTimeout(timeout time.Duration) *GetSyncLicenseParams {

	return &GetSyncLicenseParams{

		timeout: timeout,
	}
}

// NewGetSyncLicenseParamsWithContext creates a new GetSyncLicenseParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSyncLicenseParamsWithContext(ctx context.Context) *GetSyncLicenseParams {

	return &GetSyncLicenseParams{

		Context: ctx,
	}
}

// NewGetSyncLicenseParamsWithHTTPClient creates a new GetSyncLicenseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSyncLicenseParamsWithHTTPClient(client *http.Client) *GetSyncLicenseParams {

	return &GetSyncLicenseParams{
		HTTPClient: client,
	}
}

/*GetSyncLicenseParams contains all the parameters to send to the API endpoint
for the get sync license operation typically these are written to a http.Request
*/
type GetSyncLicenseParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sync license params
func (o *GetSyncLicenseParams) WithTimeout(timeout time.Duration) *GetSyncLicenseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sync license params
func (o *GetSyncLicenseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sync license params
func (o *GetSyncLicenseParams) WithContext(ctx context.Context) *GetSyncLicenseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sync license params
func (o *GetSyncLicenseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sync license params
func (o *GetSyncLicenseParams) WithHTTPClient(client *http.Client) *GetSyncLicenseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sync license params
func (o *GetSyncLicenseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSyncLicenseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
