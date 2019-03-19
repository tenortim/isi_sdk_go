// Code generated by go-swagger; DO NOT EDIT.

package job

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

// NewListJobJobsParams creates a new ListJobJobsParams object
// with the default values initialized.
func NewListJobJobsParams() *ListJobJobsParams {
	var ()
	return &ListJobJobsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListJobJobsParamsWithTimeout creates a new ListJobJobsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListJobJobsParamsWithTimeout(timeout time.Duration) *ListJobJobsParams {
	var ()
	return &ListJobJobsParams{

		timeout: timeout,
	}
}

// NewListJobJobsParamsWithContext creates a new ListJobJobsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListJobJobsParamsWithContext(ctx context.Context) *ListJobJobsParams {
	var ()
	return &ListJobJobsParams{

		Context: ctx,
	}
}

// NewListJobJobsParamsWithHTTPClient creates a new ListJobJobsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListJobJobsParamsWithHTTPClient(client *http.Client) *ListJobJobsParams {
	var ()
	return &ListJobJobsParams{
		HTTPClient: client,
	}
}

/*ListJobJobsParams contains all the parameters to send to the API endpoint
for the list job jobs operation typically these are written to a http.Request
*/
type ListJobJobsParams struct {

	/*Batch
	  If true, other arguments are ignored, and the query will return all results, unsorted, as quickly as possible.

	*/
	Batch *bool
	/*Dir
	  The direction of the sort.

	*/
	Dir *string
	/*Limit
	  Return no more than this many results at once (see resume).

	*/
	Limit *int64
	/*Resume
	  Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options).

	*/
	Resume *string
	/*Sort
	  The field that will be used for sorting.

	*/
	Sort *string
	/*State
	  Limit the results to jobs in the specified state.

	*/
	State *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list job jobs params
func (o *ListJobJobsParams) WithTimeout(timeout time.Duration) *ListJobJobsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list job jobs params
func (o *ListJobJobsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list job jobs params
func (o *ListJobJobsParams) WithContext(ctx context.Context) *ListJobJobsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list job jobs params
func (o *ListJobJobsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list job jobs params
func (o *ListJobJobsParams) WithHTTPClient(client *http.Client) *ListJobJobsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list job jobs params
func (o *ListJobJobsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBatch adds the batch to the list job jobs params
func (o *ListJobJobsParams) WithBatch(batch *bool) *ListJobJobsParams {
	o.SetBatch(batch)
	return o
}

// SetBatch adds the batch to the list job jobs params
func (o *ListJobJobsParams) SetBatch(batch *bool) {
	o.Batch = batch
}

// WithDir adds the dir to the list job jobs params
func (o *ListJobJobsParams) WithDir(dir *string) *ListJobJobsParams {
	o.SetDir(dir)
	return o
}

// SetDir adds the dir to the list job jobs params
func (o *ListJobJobsParams) SetDir(dir *string) {
	o.Dir = dir
}

// WithLimit adds the limit to the list job jobs params
func (o *ListJobJobsParams) WithLimit(limit *int64) *ListJobJobsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list job jobs params
func (o *ListJobJobsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithResume adds the resume to the list job jobs params
func (o *ListJobJobsParams) WithResume(resume *string) *ListJobJobsParams {
	o.SetResume(resume)
	return o
}

// SetResume adds the resume to the list job jobs params
func (o *ListJobJobsParams) SetResume(resume *string) {
	o.Resume = resume
}

// WithSort adds the sort to the list job jobs params
func (o *ListJobJobsParams) WithSort(sort *string) *ListJobJobsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the list job jobs params
func (o *ListJobJobsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithState adds the state to the list job jobs params
func (o *ListJobJobsParams) WithState(state *string) *ListJobJobsParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the list job jobs params
func (o *ListJobJobsParams) SetState(state *string) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *ListJobJobsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Batch != nil {

		// query param batch
		var qrBatch bool
		if o.Batch != nil {
			qrBatch = *o.Batch
		}
		qBatch := swag.FormatBool(qrBatch)
		if qBatch != "" {
			if err := r.SetQueryParam("batch", qBatch); err != nil {
				return err
			}
		}

	}

	if o.Dir != nil {

		// query param dir
		var qrDir string
		if o.Dir != nil {
			qrDir = *o.Dir
		}
		qDir := qrDir
		if qDir != "" {
			if err := r.SetQueryParam("dir", qDir); err != nil {
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

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

	}

	if o.State != nil {

		// query param state
		var qrState string
		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {
			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
