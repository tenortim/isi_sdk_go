// Code generated by go-swagger; DO NOT EDIT.

package snapshot

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

// NewDeleteSnapshotScheduleParams creates a new DeleteSnapshotScheduleParams object
// with the default values initialized.
func NewDeleteSnapshotScheduleParams() *DeleteSnapshotScheduleParams {
	var ()
	return &DeleteSnapshotScheduleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSnapshotScheduleParamsWithTimeout creates a new DeleteSnapshotScheduleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSnapshotScheduleParamsWithTimeout(timeout time.Duration) *DeleteSnapshotScheduleParams {
	var ()
	return &DeleteSnapshotScheduleParams{

		timeout: timeout,
	}
}

// NewDeleteSnapshotScheduleParamsWithContext creates a new DeleteSnapshotScheduleParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSnapshotScheduleParamsWithContext(ctx context.Context) *DeleteSnapshotScheduleParams {
	var ()
	return &DeleteSnapshotScheduleParams{

		Context: ctx,
	}
}

// NewDeleteSnapshotScheduleParamsWithHTTPClient creates a new DeleteSnapshotScheduleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSnapshotScheduleParamsWithHTTPClient(client *http.Client) *DeleteSnapshotScheduleParams {
	var ()
	return &DeleteSnapshotScheduleParams{
		HTTPClient: client,
	}
}

/*DeleteSnapshotScheduleParams contains all the parameters to send to the API endpoint
for the delete snapshot schedule operation typically these are written to a http.Request
*/
type DeleteSnapshotScheduleParams struct {

	/*SnapshotScheduleID
	  Delete the schedule. This does not affect already created snapshots.

	*/
	SnapshotScheduleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete snapshot schedule params
func (o *DeleteSnapshotScheduleParams) WithTimeout(timeout time.Duration) *DeleteSnapshotScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete snapshot schedule params
func (o *DeleteSnapshotScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete snapshot schedule params
func (o *DeleteSnapshotScheduleParams) WithContext(ctx context.Context) *DeleteSnapshotScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete snapshot schedule params
func (o *DeleteSnapshotScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete snapshot schedule params
func (o *DeleteSnapshotScheduleParams) WithHTTPClient(client *http.Client) *DeleteSnapshotScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete snapshot schedule params
func (o *DeleteSnapshotScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSnapshotScheduleID adds the snapshotScheduleID to the delete snapshot schedule params
func (o *DeleteSnapshotScheduleParams) WithSnapshotScheduleID(snapshotScheduleID string) *DeleteSnapshotScheduleParams {
	o.SetSnapshotScheduleID(snapshotScheduleID)
	return o
}

// SetSnapshotScheduleID adds the snapshotScheduleId to the delete snapshot schedule params
func (o *DeleteSnapshotScheduleParams) SetSnapshotScheduleID(snapshotScheduleID string) {
	o.SnapshotScheduleID = snapshotScheduleID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSnapshotScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param SnapshotScheduleId
	if err := r.SetPathParam("SnapshotScheduleId", o.SnapshotScheduleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}