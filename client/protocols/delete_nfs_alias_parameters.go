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

// NewDeleteNfsAliasParams creates a new DeleteNfsAliasParams object
// with the default values initialized.
func NewDeleteNfsAliasParams() *DeleteNfsAliasParams {
	var ()
	return &DeleteNfsAliasParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNfsAliasParamsWithTimeout creates a new DeleteNfsAliasParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteNfsAliasParamsWithTimeout(timeout time.Duration) *DeleteNfsAliasParams {
	var ()
	return &DeleteNfsAliasParams{

		timeout: timeout,
	}
}

// NewDeleteNfsAliasParamsWithContext creates a new DeleteNfsAliasParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteNfsAliasParamsWithContext(ctx context.Context) *DeleteNfsAliasParams {
	var ()
	return &DeleteNfsAliasParams{

		Context: ctx,
	}
}

// NewDeleteNfsAliasParamsWithHTTPClient creates a new DeleteNfsAliasParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteNfsAliasParamsWithHTTPClient(client *http.Client) *DeleteNfsAliasParams {
	var ()
	return &DeleteNfsAliasParams{
		HTTPClient: client,
	}
}

/*DeleteNfsAliasParams contains all the parameters to send to the API endpoint
for the delete nfs alias operation typically these are written to a http.Request
*/
type DeleteNfsAliasParams struct {

	/*NfsAliasID
	  Delete the export.

	*/
	NfsAliasID string
	/*Zone
	  Access zone

	*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete nfs alias params
func (o *DeleteNfsAliasParams) WithTimeout(timeout time.Duration) *DeleteNfsAliasParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete nfs alias params
func (o *DeleteNfsAliasParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete nfs alias params
func (o *DeleteNfsAliasParams) WithContext(ctx context.Context) *DeleteNfsAliasParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete nfs alias params
func (o *DeleteNfsAliasParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete nfs alias params
func (o *DeleteNfsAliasParams) WithHTTPClient(client *http.Client) *DeleteNfsAliasParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete nfs alias params
func (o *DeleteNfsAliasParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNfsAliasID adds the nfsAliasID to the delete nfs alias params
func (o *DeleteNfsAliasParams) WithNfsAliasID(nfsAliasID string) *DeleteNfsAliasParams {
	o.SetNfsAliasID(nfsAliasID)
	return o
}

// SetNfsAliasID adds the nfsAliasId to the delete nfs alias params
func (o *DeleteNfsAliasParams) SetNfsAliasID(nfsAliasID string) {
	o.NfsAliasID = nfsAliasID
}

// WithZone adds the zone to the delete nfs alias params
func (o *DeleteNfsAliasParams) WithZone(zone *string) *DeleteNfsAliasParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the delete nfs alias params
func (o *DeleteNfsAliasParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNfsAliasParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param NfsAliasId
	if err := r.SetPathParam("NfsAliasId", o.NfsAliasID); err != nil {
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
