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

// NewDeleteSmbShareParams creates a new DeleteSmbShareParams object
// with the default values initialized.
func NewDeleteSmbShareParams() *DeleteSmbShareParams {
	var ()
	return &DeleteSmbShareParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSmbShareParamsWithTimeout creates a new DeleteSmbShareParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSmbShareParamsWithTimeout(timeout time.Duration) *DeleteSmbShareParams {
	var ()
	return &DeleteSmbShareParams{

		timeout: timeout,
	}
}

// NewDeleteSmbShareParamsWithContext creates a new DeleteSmbShareParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSmbShareParamsWithContext(ctx context.Context) *DeleteSmbShareParams {
	var ()
	return &DeleteSmbShareParams{

		Context: ctx,
	}
}

// NewDeleteSmbShareParamsWithHTTPClient creates a new DeleteSmbShareParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSmbShareParamsWithHTTPClient(client *http.Client) *DeleteSmbShareParams {
	var ()
	return &DeleteSmbShareParams{
		HTTPClient: client,
	}
}

/*DeleteSmbShareParams contains all the parameters to send to the API endpoint
for the delete smb share operation typically these are written to a http.Request
*/
type DeleteSmbShareParams struct {

	/*SmbShareID
	  Delete the share.

	*/
	SmbShareID string
	/*Zone
	  Zone which contains this share.

	*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete smb share params
func (o *DeleteSmbShareParams) WithTimeout(timeout time.Duration) *DeleteSmbShareParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete smb share params
func (o *DeleteSmbShareParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete smb share params
func (o *DeleteSmbShareParams) WithContext(ctx context.Context) *DeleteSmbShareParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete smb share params
func (o *DeleteSmbShareParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete smb share params
func (o *DeleteSmbShareParams) WithHTTPClient(client *http.Client) *DeleteSmbShareParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete smb share params
func (o *DeleteSmbShareParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSmbShareID adds the smbShareID to the delete smb share params
func (o *DeleteSmbShareParams) WithSmbShareID(smbShareID string) *DeleteSmbShareParams {
	o.SetSmbShareID(smbShareID)
	return o
}

// SetSmbShareID adds the smbShareId to the delete smb share params
func (o *DeleteSmbShareParams) SetSmbShareID(smbShareID string) {
	o.SmbShareID = smbShareID
}

// WithZone adds the zone to the delete smb share params
func (o *DeleteSmbShareParams) WithZone(zone *string) *DeleteSmbShareParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the delete smb share params
func (o *DeleteSmbShareParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSmbShareParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param SmbShareId
	if err := r.SetPathParam("SmbShareId", o.SmbShareID); err != nil {
		return err
	}

	if o.Zone != nil {

		// query param zone
		var qrZone string
		if o.Zone != nil {
			qrZone = *o.Zone
		}
		qZone := qrZone
		if qZone != "" {
			if err := r.SetQueryParam("zone", qZone); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}