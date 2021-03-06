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

// NewUpdateNtpServerParams creates a new UpdateNtpServerParams object
// with the default values initialized.
func NewUpdateNtpServerParams() *UpdateNtpServerParams {
	var ()
	return &UpdateNtpServerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNtpServerParamsWithTimeout creates a new UpdateNtpServerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateNtpServerParamsWithTimeout(timeout time.Duration) *UpdateNtpServerParams {
	var ()
	return &UpdateNtpServerParams{

		timeout: timeout,
	}
}

// NewUpdateNtpServerParamsWithContext creates a new UpdateNtpServerParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateNtpServerParamsWithContext(ctx context.Context) *UpdateNtpServerParams {
	var ()
	return &UpdateNtpServerParams{

		Context: ctx,
	}
}

// NewUpdateNtpServerParamsWithHTTPClient creates a new UpdateNtpServerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateNtpServerParamsWithHTTPClient(client *http.Client) *UpdateNtpServerParams {
	var ()
	return &UpdateNtpServerParams{
		HTTPClient: client,
	}
}

/*UpdateNtpServerParams contains all the parameters to send to the API endpoint
for the update ntp server operation typically these are written to a http.Request
*/
type UpdateNtpServerParams struct {

	/*NtpServer*/
	NtpServer *models.NtpServer
	/*NtpServerID
	  Modify the key value for an NTP server.

	*/
	NtpServerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update ntp server params
func (o *UpdateNtpServerParams) WithTimeout(timeout time.Duration) *UpdateNtpServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update ntp server params
func (o *UpdateNtpServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update ntp server params
func (o *UpdateNtpServerParams) WithContext(ctx context.Context) *UpdateNtpServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update ntp server params
func (o *UpdateNtpServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update ntp server params
func (o *UpdateNtpServerParams) WithHTTPClient(client *http.Client) *UpdateNtpServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update ntp server params
func (o *UpdateNtpServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNtpServer adds the ntpServer to the update ntp server params
func (o *UpdateNtpServerParams) WithNtpServer(ntpServer *models.NtpServer) *UpdateNtpServerParams {
	o.SetNtpServer(ntpServer)
	return o
}

// SetNtpServer adds the ntpServer to the update ntp server params
func (o *UpdateNtpServerParams) SetNtpServer(ntpServer *models.NtpServer) {
	o.NtpServer = ntpServer
}

// WithNtpServerID adds the ntpServerID to the update ntp server params
func (o *UpdateNtpServerParams) WithNtpServerID(ntpServerID string) *UpdateNtpServerParams {
	o.SetNtpServerID(ntpServerID)
	return o
}

// SetNtpServerID adds the ntpServerId to the update ntp server params
func (o *UpdateNtpServerParams) SetNtpServerID(ntpServerID string) {
	o.NtpServerID = ntpServerID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNtpServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.NtpServer != nil {
		if err := r.SetBodyParam(o.NtpServer); err != nil {
			return err
		}
	}

	// path param NtpServerId
	if err := r.SetPathParam("NtpServerId", o.NtpServerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
