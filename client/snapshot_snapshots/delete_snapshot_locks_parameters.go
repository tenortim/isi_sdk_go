// Code generated by go-swagger; DO NOT EDIT.

package snapshot_snapshots

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

// NewDeleteSnapshotLocksParams creates a new DeleteSnapshotLocksParams object
// with the default values initialized.
func NewDeleteSnapshotLocksParams() *DeleteSnapshotLocksParams {
	var ()
	return &DeleteSnapshotLocksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSnapshotLocksParamsWithTimeout creates a new DeleteSnapshotLocksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSnapshotLocksParamsWithTimeout(timeout time.Duration) *DeleteSnapshotLocksParams {
	var ()
	return &DeleteSnapshotLocksParams{

		timeout: timeout,
	}
}

// NewDeleteSnapshotLocksParamsWithContext creates a new DeleteSnapshotLocksParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSnapshotLocksParamsWithContext(ctx context.Context) *DeleteSnapshotLocksParams {
	var ()
	return &DeleteSnapshotLocksParams{

		Context: ctx,
	}
}

// NewDeleteSnapshotLocksParamsWithHTTPClient creates a new DeleteSnapshotLocksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSnapshotLocksParamsWithHTTPClient(client *http.Client) *DeleteSnapshotLocksParams {
	var ()
	return &DeleteSnapshotLocksParams{
		HTTPClient: client,
	}
}

/*DeleteSnapshotLocksParams contains all the parameters to send to the API endpoint
for the delete snapshot locks operation typically these are written to a http.Request
*/
type DeleteSnapshotLocksParams struct {

	/*Sid*/
	Sid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete snapshot locks params
func (o *DeleteSnapshotLocksParams) WithTimeout(timeout time.Duration) *DeleteSnapshotLocksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete snapshot locks params
func (o *DeleteSnapshotLocksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete snapshot locks params
func (o *DeleteSnapshotLocksParams) WithContext(ctx context.Context) *DeleteSnapshotLocksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete snapshot locks params
func (o *DeleteSnapshotLocksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete snapshot locks params
func (o *DeleteSnapshotLocksParams) WithHTTPClient(client *http.Client) *DeleteSnapshotLocksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete snapshot locks params
func (o *DeleteSnapshotLocksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSid adds the sid to the delete snapshot locks params
func (o *DeleteSnapshotLocksParams) WithSid(sid string) *DeleteSnapshotLocksParams {
	o.SetSid(sid)
	return o
}

// SetSid adds the sid to the delete snapshot locks params
func (o *DeleteSnapshotLocksParams) SetSid(sid string) {
	o.Sid = sid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSnapshotLocksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Sid
	if err := r.SetPathParam("Sid", o.Sid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
