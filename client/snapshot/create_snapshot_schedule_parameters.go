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

// NewCreateSnapshotScheduleParams creates a new CreateSnapshotScheduleParams object
// with the default values initialized.
func NewCreateSnapshotScheduleParams() *CreateSnapshotScheduleParams {
	var ()
	return &CreateSnapshotScheduleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSnapshotScheduleParamsWithTimeout creates a new CreateSnapshotScheduleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSnapshotScheduleParamsWithTimeout(timeout time.Duration) *CreateSnapshotScheduleParams {
	var ()
	return &CreateSnapshotScheduleParams{

		timeout: timeout,
	}
}

// NewCreateSnapshotScheduleParamsWithContext creates a new CreateSnapshotScheduleParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSnapshotScheduleParamsWithContext(ctx context.Context) *CreateSnapshotScheduleParams {
	var ()
	return &CreateSnapshotScheduleParams{

		Context: ctx,
	}
}

// NewCreateSnapshotScheduleParamsWithHTTPClient creates a new CreateSnapshotScheduleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSnapshotScheduleParamsWithHTTPClient(client *http.Client) *CreateSnapshotScheduleParams {
	var ()
	return &CreateSnapshotScheduleParams{
		HTTPClient: client,
	}
}

/*CreateSnapshotScheduleParams contains all the parameters to send to the API endpoint
for the create snapshot schedule operation typically these are written to a http.Request
*/
type CreateSnapshotScheduleParams struct {

	/*SnapshotSchedule*/
	SnapshotSchedule *models.SnapshotScheduleCreateParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create snapshot schedule params
func (o *CreateSnapshotScheduleParams) WithTimeout(timeout time.Duration) *CreateSnapshotScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create snapshot schedule params
func (o *CreateSnapshotScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create snapshot schedule params
func (o *CreateSnapshotScheduleParams) WithContext(ctx context.Context) *CreateSnapshotScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create snapshot schedule params
func (o *CreateSnapshotScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create snapshot schedule params
func (o *CreateSnapshotScheduleParams) WithHTTPClient(client *http.Client) *CreateSnapshotScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create snapshot schedule params
func (o *CreateSnapshotScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSnapshotSchedule adds the snapshotSchedule to the create snapshot schedule params
func (o *CreateSnapshotScheduleParams) WithSnapshotSchedule(snapshotSchedule *models.SnapshotScheduleCreateParams) *CreateSnapshotScheduleParams {
	o.SetSnapshotSchedule(snapshotSchedule)
	return o
}

// SetSnapshotSchedule adds the snapshotSchedule to the create snapshot schedule params
func (o *CreateSnapshotScheduleParams) SetSnapshotSchedule(snapshotSchedule *models.SnapshotScheduleCreateParams) {
	o.SnapshotSchedule = snapshotSchedule
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSnapshotScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SnapshotSchedule != nil {
		if err := r.SetBodyParam(o.SnapshotSchedule); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
