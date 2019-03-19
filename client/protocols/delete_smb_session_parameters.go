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

// NewDeleteSmbSessionParams creates a new DeleteSmbSessionParams object
// with the default values initialized.
func NewDeleteSmbSessionParams() *DeleteSmbSessionParams {
	var ()
	return &DeleteSmbSessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSmbSessionParamsWithTimeout creates a new DeleteSmbSessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSmbSessionParamsWithTimeout(timeout time.Duration) *DeleteSmbSessionParams {
	var ()
	return &DeleteSmbSessionParams{

		timeout: timeout,
	}
}

// NewDeleteSmbSessionParamsWithContext creates a new DeleteSmbSessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSmbSessionParamsWithContext(ctx context.Context) *DeleteSmbSessionParams {
	var ()
	return &DeleteSmbSessionParams{

		Context: ctx,
	}
}

// NewDeleteSmbSessionParamsWithHTTPClient creates a new DeleteSmbSessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSmbSessionParamsWithHTTPClient(client *http.Client) *DeleteSmbSessionParams {
	var ()
	return &DeleteSmbSessionParams{
		HTTPClient: client,
	}
}

/*DeleteSmbSessionParams contains all the parameters to send to the API endpoint
for the delete smb session operation typically these are written to a http.Request
*/
type DeleteSmbSessionParams struct {

	/*SmbSessionID
	  Close the SMB session.

	*/
	SmbSessionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete smb session params
func (o *DeleteSmbSessionParams) WithTimeout(timeout time.Duration) *DeleteSmbSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete smb session params
func (o *DeleteSmbSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete smb session params
func (o *DeleteSmbSessionParams) WithContext(ctx context.Context) *DeleteSmbSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete smb session params
func (o *DeleteSmbSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete smb session params
func (o *DeleteSmbSessionParams) WithHTTPClient(client *http.Client) *DeleteSmbSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete smb session params
func (o *DeleteSmbSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSmbSessionID adds the smbSessionID to the delete smb session params
func (o *DeleteSmbSessionParams) WithSmbSessionID(smbSessionID string) *DeleteSmbSessionParams {
	o.SetSmbSessionID(smbSessionID)
	return o
}

// SetSmbSessionID adds the smbSessionId to the delete smb session params
func (o *DeleteSmbSessionParams) SetSmbSessionID(smbSessionID string) {
	o.SmbSessionID = smbSessionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSmbSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param SmbSessionId
	if err := r.SetPathParam("SmbSessionId", o.SmbSessionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
