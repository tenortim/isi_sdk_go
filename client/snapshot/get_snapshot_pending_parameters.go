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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSnapshotPendingParams creates a new GetSnapshotPendingParams object
// with the default values initialized.
func NewGetSnapshotPendingParams() *GetSnapshotPendingParams {
	var ()
	return &GetSnapshotPendingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSnapshotPendingParamsWithTimeout creates a new GetSnapshotPendingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSnapshotPendingParamsWithTimeout(timeout time.Duration) *GetSnapshotPendingParams {
	var ()
	return &GetSnapshotPendingParams{

		timeout: timeout,
	}
}

// NewGetSnapshotPendingParamsWithContext creates a new GetSnapshotPendingParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSnapshotPendingParamsWithContext(ctx context.Context) *GetSnapshotPendingParams {
	var ()
	return &GetSnapshotPendingParams{

		Context: ctx,
	}
}

// NewGetSnapshotPendingParamsWithHTTPClient creates a new GetSnapshotPendingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSnapshotPendingParamsWithHTTPClient(client *http.Client) *GetSnapshotPendingParams {
	var ()
	return &GetSnapshotPendingParams{
		HTTPClient: client,
	}
}

/*GetSnapshotPendingParams contains all the parameters to send to the API endpoint
for the get snapshot pending operation typically these are written to a http.Request
*/
type GetSnapshotPendingParams struct {

	/*Begin
	  Unix Epoch time to start generating matches. Default is now.

	*/
	Begin *int64
	/*End
	  Unix Epoch time to end generating matches. Default is forever.

	*/
	End *int64
	/*Limit
	  Return no more than this many result at once (see resume).

	*/
	Limit *int64
	/*Resume
	  Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options).

	*/
	Resume *string
	/*Schedule
	  Limit output only to the named schedule.

	*/
	Schedule *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get snapshot pending params
func (o *GetSnapshotPendingParams) WithTimeout(timeout time.Duration) *GetSnapshotPendingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get snapshot pending params
func (o *GetSnapshotPendingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get snapshot pending params
func (o *GetSnapshotPendingParams) WithContext(ctx context.Context) *GetSnapshotPendingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get snapshot pending params
func (o *GetSnapshotPendingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get snapshot pending params
func (o *GetSnapshotPendingParams) WithHTTPClient(client *http.Client) *GetSnapshotPendingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get snapshot pending params
func (o *GetSnapshotPendingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBegin adds the begin to the get snapshot pending params
func (o *GetSnapshotPendingParams) WithBegin(begin *int64) *GetSnapshotPendingParams {
	o.SetBegin(begin)
	return o
}

// SetBegin adds the begin to the get snapshot pending params
func (o *GetSnapshotPendingParams) SetBegin(begin *int64) {
	o.Begin = begin
}

// WithEnd adds the end to the get snapshot pending params
func (o *GetSnapshotPendingParams) WithEnd(end *int64) *GetSnapshotPendingParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get snapshot pending params
func (o *GetSnapshotPendingParams) SetEnd(end *int64) {
	o.End = end
}

// WithLimit adds the limit to the get snapshot pending params
func (o *GetSnapshotPendingParams) WithLimit(limit *int64) *GetSnapshotPendingParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get snapshot pending params
func (o *GetSnapshotPendingParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithResume adds the resume to the get snapshot pending params
func (o *GetSnapshotPendingParams) WithResume(resume *string) *GetSnapshotPendingParams {
	o.SetResume(resume)
	return o
}

// SetResume adds the resume to the get snapshot pending params
func (o *GetSnapshotPendingParams) SetResume(resume *string) {
	o.Resume = resume
}

// WithSchedule adds the schedule to the get snapshot pending params
func (o *GetSnapshotPendingParams) WithSchedule(schedule *string) *GetSnapshotPendingParams {
	o.SetSchedule(schedule)
	return o
}

// SetSchedule adds the schedule to the get snapshot pending params
func (o *GetSnapshotPendingParams) SetSchedule(schedule *string) {
	o.Schedule = schedule
}

// WriteToRequest writes these params to a swagger request
func (o *GetSnapshotPendingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Begin != nil {

		// query param begin
		var qrBegin int64
		if o.Begin != nil {
			qrBegin = *o.Begin
		}
		qBegin := swag.FormatInt64(qrBegin)
		if qBegin != "" {
			if err := r.SetQueryParam("begin", qBegin); err != nil {
				return err
			}
		}

	}

	if o.End != nil {

		// query param end
		var qrEnd int64
		if o.End != nil {
			qrEnd = *o.End
		}
		qEnd := swag.FormatInt64(qrEnd)
		if qEnd != "" {
			if err := r.SetQueryParam("end", qEnd); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Resume != nil {

		// query param resume
		var qrResume string
		if o.Resume != nil {
			qrResume = *o.Resume
		}
		qResume := qrResume
		if qResume != "" {
			if err := r.SetQueryParam("resume", qResume); err != nil {
				return err
			}
		}

	}

	if o.Schedule != nil {

		// query param schedule
		var qrSchedule string
		if o.Schedule != nil {
			qrSchedule = *o.Schedule
		}
		qSchedule := qrSchedule
		if qSchedule != "" {
			if err := r.SetQueryParam("schedule", qSchedule); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}