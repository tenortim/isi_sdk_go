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

	models "github.com/tenortim/isi_sdk_go/models"
)

// NewUpdateSnapshotScheduleParams creates a new UpdateSnapshotScheduleParams object
// with the default values initialized.
func NewUpdateSnapshotScheduleParams() *UpdateSnapshotScheduleParams {
	var ()
	return &UpdateSnapshotScheduleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSnapshotScheduleParamsWithTimeout creates a new UpdateSnapshotScheduleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSnapshotScheduleParamsWithTimeout(timeout time.Duration) *UpdateSnapshotScheduleParams {
	var ()
	return &UpdateSnapshotScheduleParams{

		timeout: timeout,
	}
}

// NewUpdateSnapshotScheduleParamsWithContext creates a new UpdateSnapshotScheduleParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSnapshotScheduleParamsWithContext(ctx context.Context) *UpdateSnapshotScheduleParams {
	var ()
	return &UpdateSnapshotScheduleParams{

		Context: ctx,
	}
}

// NewUpdateSnapshotScheduleParamsWithHTTPClient creates a new UpdateSnapshotScheduleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSnapshotScheduleParamsWithHTTPClient(client *http.Client) *UpdateSnapshotScheduleParams {
	var ()
	return &UpdateSnapshotScheduleParams{
		HTTPClient: client,
	}
}

/*UpdateSnapshotScheduleParams contains all the parameters to send to the API endpoint
for the update snapshot schedule operation typically these are written to a http.Request
*/
type UpdateSnapshotScheduleParams struct {

	/*SnapshotSchedule*/
	SnapshotSchedule *models.SnapshotSchedule
	/*SnapshotScheduleID
	  Modify the schedule. All input fields are optional, but one or more must be supplied.

	*/
	SnapshotScheduleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update snapshot schedule params
func (o *UpdateSnapshotScheduleParams) WithTimeout(timeout time.Duration) *UpdateSnapshotScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update snapshot schedule params
func (o *UpdateSnapshotScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update snapshot schedule params
func (o *UpdateSnapshotScheduleParams) WithContext(ctx context.Context) *UpdateSnapshotScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update snapshot schedule params
func (o *UpdateSnapshotScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update snapshot schedule params
func (o *UpdateSnapshotScheduleParams) WithHTTPClient(client *http.Client) *UpdateSnapshotScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update snapshot schedule params
func (o *UpdateSnapshotScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSnapshotSchedule adds the snapshotSchedule to the update snapshot schedule params
func (o *UpdateSnapshotScheduleParams) WithSnapshotSchedule(snapshotSchedule *models.SnapshotSchedule) *UpdateSnapshotScheduleParams {
	o.SetSnapshotSchedule(snapshotSchedule)
	return o
}

// SetSnapshotSchedule adds the snapshotSchedule to the update snapshot schedule params
func (o *UpdateSnapshotScheduleParams) SetSnapshotSchedule(snapshotSchedule *models.SnapshotSchedule) {
	o.SnapshotSchedule = snapshotSchedule
}

// WithSnapshotScheduleID adds the snapshotScheduleID to the update snapshot schedule params
func (o *UpdateSnapshotScheduleParams) WithSnapshotScheduleID(snapshotScheduleID string) *UpdateSnapshotScheduleParams {
	o.SetSnapshotScheduleID(snapshotScheduleID)
	return o
}

// SetSnapshotScheduleID adds the snapshotScheduleId to the update snapshot schedule params
func (o *UpdateSnapshotScheduleParams) SetSnapshotScheduleID(snapshotScheduleID string) {
	o.SnapshotScheduleID = snapshotScheduleID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSnapshotScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SnapshotSchedule != nil {
		if err := r.SetBodyParam(o.SnapshotSchedule); err != nil {
			return err
		}
	}

	// path param SnapshotScheduleId
	if err := r.SetPathParam("SnapshotScheduleId", o.SnapshotScheduleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}