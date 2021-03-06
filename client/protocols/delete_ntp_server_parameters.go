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

// NewDeleteNtpServerParams creates a new DeleteNtpServerParams object
// with the default values initialized.
func NewDeleteNtpServerParams() *DeleteNtpServerParams {
	var ()
	return &DeleteNtpServerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNtpServerParamsWithTimeout creates a new DeleteNtpServerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteNtpServerParamsWithTimeout(timeout time.Duration) *DeleteNtpServerParams {
	var ()
	return &DeleteNtpServerParams{

		timeout: timeout,
	}
}

// NewDeleteNtpServerParamsWithContext creates a new DeleteNtpServerParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteNtpServerParamsWithContext(ctx context.Context) *DeleteNtpServerParams {
	var ()
	return &DeleteNtpServerParams{

		Context: ctx,
	}
}

// NewDeleteNtpServerParamsWithHTTPClient creates a new DeleteNtpServerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteNtpServerParamsWithHTTPClient(client *http.Client) *DeleteNtpServerParams {
	var ()
	return &DeleteNtpServerParams{
		HTTPClient: client,
	}
}

/*DeleteNtpServerParams contains all the parameters to send to the API endpoint
for the delete ntp server operation typically these are written to a http.Request
*/
type DeleteNtpServerParams struct {

	/*NtpServerID
	  Delete an NTP server entry.

	*/
	NtpServerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete ntp server params
func (o *DeleteNtpServerParams) WithTimeout(timeout time.Duration) *DeleteNtpServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete ntp server params
func (o *DeleteNtpServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete ntp server params
func (o *DeleteNtpServerParams) WithContext(ctx context.Context) *DeleteNtpServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete ntp server params
func (o *DeleteNtpServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete ntp server params
func (o *DeleteNtpServerParams) WithHTTPClient(client *http.Client) *DeleteNtpServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete ntp server params
func (o *DeleteNtpServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNtpServerID adds the ntpServerID to the delete ntp server params
func (o *DeleteNtpServerParams) WithNtpServerID(ntpServerID string) *DeleteNtpServerParams {
	o.SetNtpServerID(ntpServerID)
	return o
}

// SetNtpServerID adds the ntpServerId to the delete ntp server params
func (o *DeleteNtpServerParams) SetNtpServerID(ntpServerID string) {
	o.NtpServerID = ntpServerID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNtpServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param NtpServerId
	if err := r.SetPathParam("NtpServerId", o.NtpServerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
