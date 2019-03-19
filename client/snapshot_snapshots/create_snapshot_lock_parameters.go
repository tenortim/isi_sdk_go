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

	models "github.com/tenortim/isi_sdk_go/models"
)

// NewCreateSnapshotLockParams creates a new CreateSnapshotLockParams object
// with the default values initialized.
func NewCreateSnapshotLockParams() *CreateSnapshotLockParams {
	var ()
	return &CreateSnapshotLockParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSnapshotLockParamsWithTimeout creates a new CreateSnapshotLockParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSnapshotLockParamsWithTimeout(timeout time.Duration) *CreateSnapshotLockParams {
	var ()
	return &CreateSnapshotLockParams{

		timeout: timeout,
	}
}

// NewCreateSnapshotLockParamsWithContext creates a new CreateSnapshotLockParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSnapshotLockParamsWithContext(ctx context.Context) *CreateSnapshotLockParams {
	var ()
	return &CreateSnapshotLockParams{

		Context: ctx,
	}
}

// NewCreateSnapshotLockParamsWithHTTPClient creates a new CreateSnapshotLockParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSnapshotLockParamsWithHTTPClient(client *http.Client) *CreateSnapshotLockParams {
	var ()
	return &CreateSnapshotLockParams{
		HTTPClient: client,
	}
}

/*CreateSnapshotLockParams contains all the parameters to send to the API endpoint
for the create snapshot lock operation typically these are written to a http.Request
*/
type CreateSnapshotLockParams struct {

	/*Sid*/
	Sid string
	/*SnapshotLock*/
	SnapshotLock *models.SnapshotLockCreateParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create snapshot lock params
func (o *CreateSnapshotLockParams) WithTimeout(timeout time.Duration) *CreateSnapshotLockParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create snapshot lock params
func (o *CreateSnapshotLockParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create snapshot lock params
func (o *CreateSnapshotLockParams) WithContext(ctx context.Context) *CreateSnapshotLockParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create snapshot lock params
func (o *CreateSnapshotLockParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create snapshot lock params
func (o *CreateSnapshotLockParams) WithHTTPClient(client *http.Client) *CreateSnapshotLockParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create snapshot lock params
func (o *CreateSnapshotLockParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSid adds the sid to the create snapshot lock params
func (o *CreateSnapshotLockParams) WithSid(sid string) *CreateSnapshotLockParams {
	o.SetSid(sid)
	return o
}

// SetSid adds the sid to the create snapshot lock params
func (o *CreateSnapshotLockParams) SetSid(sid string) {
	o.Sid = sid
}

// WithSnapshotLock adds the snapshotLock to the create snapshot lock params
func (o *CreateSnapshotLockParams) WithSnapshotLock(snapshotLock *models.SnapshotLockCreateParams) *CreateSnapshotLockParams {
	o.SetSnapshotLock(snapshotLock)
	return o
}

// SetSnapshotLock adds the snapshotLock to the create snapshot lock params
func (o *CreateSnapshotLockParams) SetSnapshotLock(snapshotLock *models.SnapshotLockCreateParams) {
	o.SnapshotLock = snapshotLock
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSnapshotLockParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Sid
	if err := r.SetPathParam("Sid", o.Sid); err != nil {
		return err
	}

	if o.SnapshotLock != nil {
		if err := r.SetBodyParam(o.SnapshotLock); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
