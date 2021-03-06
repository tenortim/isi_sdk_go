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

// NewGetAntivirusQuarantinePathParams creates a new GetAntivirusQuarantinePathParams object
// with the default values initialized.
func NewGetAntivirusQuarantinePathParams() *GetAntivirusQuarantinePathParams {
	var ()
	return &GetAntivirusQuarantinePathParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAntivirusQuarantinePathParamsWithTimeout creates a new GetAntivirusQuarantinePathParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAntivirusQuarantinePathParamsWithTimeout(timeout time.Duration) *GetAntivirusQuarantinePathParams {
	var ()
	return &GetAntivirusQuarantinePathParams{

		timeout: timeout,
	}
}

// NewGetAntivirusQuarantinePathParamsWithContext creates a new GetAntivirusQuarantinePathParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAntivirusQuarantinePathParamsWithContext(ctx context.Context) *GetAntivirusQuarantinePathParams {
	var ()
	return &GetAntivirusQuarantinePathParams{

		Context: ctx,
	}
}

// NewGetAntivirusQuarantinePathParamsWithHTTPClient creates a new GetAntivirusQuarantinePathParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAntivirusQuarantinePathParamsWithHTTPClient(client *http.Client) *GetAntivirusQuarantinePathParams {
	var ()
	return &GetAntivirusQuarantinePathParams{
		HTTPClient: client,
	}
}

/*GetAntivirusQuarantinePathParams contains all the parameters to send to the API endpoint
for the get antivirus quarantine path operation typically these are written to a http.Request
*/
type GetAntivirusQuarantinePathParams struct {

	/*AntivirusQuarantinePath
	  Retrieve the quarantine status of the file at the specified path.

	*/
	AntivirusQuarantinePath string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get antivirus quarantine path params
func (o *GetAntivirusQuarantinePathParams) WithTimeout(timeout time.Duration) *GetAntivirusQuarantinePathParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get antivirus quarantine path params
func (o *GetAntivirusQuarantinePathParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get antivirus quarantine path params
func (o *GetAntivirusQuarantinePathParams) WithContext(ctx context.Context) *GetAntivirusQuarantinePathParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get antivirus quarantine path params
func (o *GetAntivirusQuarantinePathParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get antivirus quarantine path params
func (o *GetAntivirusQuarantinePathParams) WithHTTPClient(client *http.Client) *GetAntivirusQuarantinePathParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get antivirus quarantine path params
func (o *GetAntivirusQuarantinePathParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAntivirusQuarantinePath adds the antivirusQuarantinePath to the get antivirus quarantine path params
func (o *GetAntivirusQuarantinePathParams) WithAntivirusQuarantinePath(antivirusQuarantinePath string) *GetAntivirusQuarantinePathParams {
	o.SetAntivirusQuarantinePath(antivirusQuarantinePath)
	return o
}

// SetAntivirusQuarantinePath adds the antivirusQuarantinePath to the get antivirus quarantine path params
func (o *GetAntivirusQuarantinePathParams) SetAntivirusQuarantinePath(antivirusQuarantinePath string) {
	o.AntivirusQuarantinePath = antivirusQuarantinePath
}

// WriteToRequest writes these params to a swagger request
func (o *GetAntivirusQuarantinePathParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param AntivirusQuarantinePath
	if err := r.SetPathParam("AntivirusQuarantinePath", o.AntivirusQuarantinePath); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
