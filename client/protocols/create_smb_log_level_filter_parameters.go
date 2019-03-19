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

	models "github.com/tenortim/isi_sdk_go/models"
)

// NewCreateSmbLogLevelFilterParams creates a new CreateSmbLogLevelFilterParams object
// with the default values initialized.
func NewCreateSmbLogLevelFilterParams() *CreateSmbLogLevelFilterParams {
	var ()
	return &CreateSmbLogLevelFilterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSmbLogLevelFilterParamsWithTimeout creates a new CreateSmbLogLevelFilterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSmbLogLevelFilterParamsWithTimeout(timeout time.Duration) *CreateSmbLogLevelFilterParams {
	var ()
	return &CreateSmbLogLevelFilterParams{

		timeout: timeout,
	}
}

// NewCreateSmbLogLevelFilterParamsWithContext creates a new CreateSmbLogLevelFilterParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSmbLogLevelFilterParamsWithContext(ctx context.Context) *CreateSmbLogLevelFilterParams {
	var ()
	return &CreateSmbLogLevelFilterParams{

		Context: ctx,
	}
}

// NewCreateSmbLogLevelFilterParamsWithHTTPClient creates a new CreateSmbLogLevelFilterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSmbLogLevelFilterParamsWithHTTPClient(client *http.Client) *CreateSmbLogLevelFilterParams {
	var ()
	return &CreateSmbLogLevelFilterParams{
		HTTPClient: client,
	}
}

/*CreateSmbLogLevelFilterParams contains all the parameters to send to the API endpoint
for the create smb log level filter operation typically these are written to a http.Request
*/
type CreateSmbLogLevelFilterParams struct {

	/*SmbLogLevelFilter*/
	SmbLogLevelFilter *models.SmbLogLevelFilter

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create smb log level filter params
func (o *CreateSmbLogLevelFilterParams) WithTimeout(timeout time.Duration) *CreateSmbLogLevelFilterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create smb log level filter params
func (o *CreateSmbLogLevelFilterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create smb log level filter params
func (o *CreateSmbLogLevelFilterParams) WithContext(ctx context.Context) *CreateSmbLogLevelFilterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create smb log level filter params
func (o *CreateSmbLogLevelFilterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create smb log level filter params
func (o *CreateSmbLogLevelFilterParams) WithHTTPClient(client *http.Client) *CreateSmbLogLevelFilterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create smb log level filter params
func (o *CreateSmbLogLevelFilterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSmbLogLevelFilter adds the smbLogLevelFilter to the create smb log level filter params
func (o *CreateSmbLogLevelFilterParams) WithSmbLogLevelFilter(smbLogLevelFilter *models.SmbLogLevelFilter) *CreateSmbLogLevelFilterParams {
	o.SetSmbLogLevelFilter(smbLogLevelFilter)
	return o
}

// SetSmbLogLevelFilter adds the smbLogLevelFilter to the create smb log level filter params
func (o *CreateSmbLogLevelFilterParams) SetSmbLogLevelFilter(smbLogLevelFilter *models.SmbLogLevelFilter) {
	o.SmbLogLevelFilter = smbLogLevelFilter
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSmbLogLevelFilterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SmbLogLevelFilter != nil {
		if err := r.SetBodyParam(o.SmbLogLevelFilter); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}