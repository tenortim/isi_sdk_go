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

// NewUpdateSnapshotLockParams creates a new UpdateSnapshotLockParams object
// with the default values initialized.
func NewUpdateSnapshotLockParams() *UpdateSnapshotLockParams {
	var ()
	return &UpdateSnapshotLockParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSnapshotLockParamsWithTimeout creates a new UpdateSnapshotLockParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSnapshotLockParamsWithTimeout(timeout time.Duration) *UpdateSnapshotLockParams {
	var ()
	return &UpdateSnapshotLockParams{

		timeout: timeout,
	}
}

// NewUpdateSnapshotLockParamsWithContext creates a new UpdateSnapshotLockParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSnapshotLockParamsWithContext(ctx context.Context) *UpdateSnapshotLockParams {
	var ()
	return &UpdateSnapshotLockParams{

		Context: ctx,
	}
}

// NewUpdateSnapshotLockParamsWithHTTPClient creates a new UpdateSnapshotLockParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSnapshotLockParamsWithHTTPClient(client *http.Client) *UpdateSnapshotLockParams {
	var ()
	return &UpdateSnapshotLockParams{
		HTTPClient: client,
	}
}

/*UpdateSnapshotLockParams contains all the parameters to send to the API endpoint
for the update snapshot lock operation typically these are written to a http.Request
*/
type UpdateSnapshotLockParams struct {

	/*Sid*/
	Sid string
	/*SnapshotLock*/
	SnapshotLock *models.SnapshotLock
	/*SnapshotLockID
	  Modify lock. All input fields are optional, but one or more must be supplied.

	*/
	SnapshotLockID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update snapshot lock params
func (o *UpdateSnapshotLockParams) WithTimeout(timeout time.Duration) *UpdateSnapshotLockParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update snapshot lock params
func (o *UpdateSnapshotLockParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update snapshot lock params
func (o *UpdateSnapshotLockParams) WithContext(ctx context.Context) *UpdateSnapshotLockParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update snapshot lock params
func (o *UpdateSnapshotLockParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update snapshot lock params
func (o *UpdateSnapshotLockParams) WithHTTPClient(client *http.Client) *UpdateSnapshotLockParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update snapshot lock params
func (o *UpdateSnapshotLockParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSid adds the sid to the update snapshot lock params
func (o *UpdateSnapshotLockParams) WithSid(sid string) *UpdateSnapshotLockParams {
	o.SetSid(sid)
	return o
}

// SetSid adds the sid to the update snapshot lock params
func (o *UpdateSnapshotLockParams) SetSid(sid string) {
	o.Sid = sid
}

// WithSnapshotLock adds the snapshotLock to the update snapshot lock params
func (o *UpdateSnapshotLockParams) WithSnapshotLock(snapshotLock *models.SnapshotLock) *UpdateSnapshotLockParams {
	o.SetSnapshotLock(snapshotLock)
	return o
}

// SetSnapshotLock adds the snapshotLock to the update snapshot lock params
func (o *UpdateSnapshotLockParams) SetSnapshotLock(snapshotLock *models.SnapshotLock) {
	o.SnapshotLock = snapshotLock
}

// WithSnapshotLockID adds the snapshotLockID to the update snapshot lock params
func (o *UpdateSnapshotLockParams) WithSnapshotLockID(snapshotLockID string) *UpdateSnapshotLockParams {
	o.SetSnapshotLockID(snapshotLockID)
	return o
}

// SetSnapshotLockID adds the snapshotLockId to the update snapshot lock params
func (o *UpdateSnapshotLockParams) SetSnapshotLockID(snapshotLockID string) {
	o.SnapshotLockID = snapshotLockID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSnapshotLockParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param SnapshotLockId
	if err := r.SetPathParam("SnapshotLockId", o.SnapshotLockID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}