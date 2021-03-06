// Code generated by go-swagger; DO NOT EDIT.

package quota

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

// NewGetQuotaLicenseParams creates a new GetQuotaLicenseParams object
// with the default values initialized.
func NewGetQuotaLicenseParams() *GetQuotaLicenseParams {

	return &GetQuotaLicenseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetQuotaLicenseParamsWithTimeout creates a new GetQuotaLicenseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetQuotaLicenseParamsWithTimeout(timeout time.Duration) *GetQuotaLicenseParams {

	return &GetQuotaLicenseParams{

		timeout: timeout,
	}
}

// NewGetQuotaLicenseParamsWithContext creates a new GetQuotaLicenseParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetQuotaLicenseParamsWithContext(ctx context.Context) *GetQuotaLicenseParams {

	return &GetQuotaLicenseParams{

		Context: ctx,
	}
}

// NewGetQuotaLicenseParamsWithHTTPClient creates a new GetQuotaLicenseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetQuotaLicenseParamsWithHTTPClient(client *http.Client) *GetQuotaLicenseParams {

	return &GetQuotaLicenseParams{
		HTTPClient: client,
	}
}

/*GetQuotaLicenseParams contains all the parameters to send to the API endpoint
for the get quota license operation typically these are written to a http.Request
*/
type GetQuotaLicenseParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get quota license params
func (o *GetQuotaLicenseParams) WithTimeout(timeout time.Duration) *GetQuotaLicenseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get quota license params
func (o *GetQuotaLicenseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get quota license params
func (o *GetQuotaLicenseParams) WithContext(ctx context.Context) *GetQuotaLicenseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get quota license params
func (o *GetQuotaLicenseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get quota license params
func (o *GetQuotaLicenseParams) WithHTTPClient(client *http.Client) *GetQuotaLicenseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get quota license params
func (o *GetQuotaLicenseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetQuotaLicenseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
