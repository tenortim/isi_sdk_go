// Code generated by go-swagger; DO NOT EDIT.

package sync_target

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

// NewGetReportsReportSubreportsParams creates a new GetReportsReportSubreportsParams object
// with the default values initialized.
func NewGetReportsReportSubreportsParams() *GetReportsReportSubreportsParams {
	var ()
	return &GetReportsReportSubreportsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReportsReportSubreportsParamsWithTimeout creates a new GetReportsReportSubreportsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReportsReportSubreportsParamsWithTimeout(timeout time.Duration) *GetReportsReportSubreportsParams {
	var ()
	return &GetReportsReportSubreportsParams{

		timeout: timeout,
	}
}

// NewGetReportsReportSubreportsParamsWithContext creates a new GetReportsReportSubreportsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReportsReportSubreportsParamsWithContext(ctx context.Context) *GetReportsReportSubreportsParams {
	var ()
	return &GetReportsReportSubreportsParams{

		Context: ctx,
	}
}

// NewGetReportsReportSubreportsParamsWithHTTPClient creates a new GetReportsReportSubreportsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReportsReportSubreportsParamsWithHTTPClient(client *http.Client) *GetReportsReportSubreportsParams {
	var ()
	return &GetReportsReportSubreportsParams{
		HTTPClient: client,
	}
}

/*GetReportsReportSubreportsParams contains all the parameters to send to the API endpoint
for the get reports report subreports operation typically these are written to a http.Request
*/
type GetReportsReportSubreportsParams struct {

	/*Rid*/
	Rid string
	/*Dir
	  The direction of the sort.

	*/
	Dir *string
	/*Limit
	  Return no more than this many results at once (see resume).

	*/
	Limit *int64
	/*NewerThan
	  Filter the returned reports to include only those whose jobs started more recently than the specified number of days ago.

	*/
	NewerThan *int64
	/*Resume
	  Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options).

	*/
	Resume *string
	/*Sort
	  The field that will be used for sorting.

	*/
	Sort *string
	/*State
	  Filter the returned reports to include only those whose jobs are in this state.

	*/
	State *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get reports report subreports params
func (o *GetReportsReportSubreportsParams) WithTimeout(timeout time.Duration) *GetReportsReportSubreportsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get reports report subreports params
func (o *GetReportsReportSubreportsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get reports report subreports params
func (o *GetReportsReportSubreportsParams) WithContext(ctx context.Context) *GetReportsReportSubreportsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get reports report subreports params
func (o *GetReportsReportSubreportsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get reports report subreports params
func (o *GetReportsReportSubreportsParams) WithHTTPClient(client *http.Client) *GetReportsReportSubreportsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get reports report subreports params
func (o *GetReportsReportSubreportsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRid adds the rid to the get reports report subreports params
func (o *GetReportsReportSubreportsParams) WithRid(rid string) *GetReportsReportSubreportsParams {
	o.SetRid(rid)
	return o
}

// SetRid adds the rid to the get reports report subreports params
func (o *GetReportsReportSubreportsParams) SetRid(rid string) {
	o.Rid = rid
}

// WithDir adds the dir to the get reports report subreports params
func (o *GetReportsReportSubreportsParams) WithDir(dir *string) *GetReportsReportSubreportsParams {
	o.SetDir(dir)
	return o
}

// SetDir adds the dir to the get reports report subreports params
func (o *GetReportsReportSubreportsParams) SetDir(dir *string) {
	o.Dir = dir
}

// WithLimit adds the limit to the get reports report subreports params
func (o *GetReportsReportSubreportsParams) WithLimit(limit *int64) *GetReportsReportSubreportsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get reports report subreports params
func (o *GetReportsReportSubreportsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNewerThan adds the newerThan to the get reports report subreports params
func (o *GetReportsReportSubreportsParams) WithNewerThan(newerThan *int64) *GetReportsReportSubreportsParams {
	o.SetNewerThan(newerThan)
	return o
}

// SetNewerThan adds the newerThan to the get reports report subreports params
func (o *GetReportsReportSubreportsParams) SetNewerThan(newerThan *int64) {
	o.NewerThan = newerThan
}

// WithResume adds the resume to the get reports report subreports params
func (o *GetReportsReportSubreportsParams) WithResume(resume *string) *GetReportsReportSubreportsParams {
	o.SetResume(resume)
	return o
}

// SetResume adds the resume to the get reports report subreports params
func (o *GetReportsReportSubreportsParams) SetResume(resume *string) {
	o.Resume = resume
}

// WithSort adds the sort to the get reports report subreports params
func (o *GetReportsReportSubreportsParams) WithSort(sort *string) *GetReportsReportSubreportsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get reports report subreports params
func (o *GetReportsReportSubreportsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithState adds the state to the get reports report subreports params
func (o *GetReportsReportSubreportsParams) WithState(state *string) *GetReportsReportSubreportsParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the get reports report subreports params
func (o *GetReportsReportSubreportsParams) SetState(state *string) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *GetReportsReportSubreportsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Rid
	if err := r.SetPathParam("Rid", o.Rid); err != nil {
		return err
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

	if o.NewerThan != nil {

		// query param newer_than
		var qrNewerThan int64
		if o.NewerThan != nil {
			qrNewerThan = *o.NewerThan
		}
		qNewerThan := swag.FormatInt64(qrNewerThan)
		if qNewerThan != "" {
			if err := r.SetQueryParam("newer_than", qNewerThan); err != nil {
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
