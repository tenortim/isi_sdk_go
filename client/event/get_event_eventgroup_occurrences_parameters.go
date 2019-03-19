// Code generated by go-swagger; DO NOT EDIT.

package event

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

// NewGetEventEventgroupOccurrencesParams creates a new GetEventEventgroupOccurrencesParams object
// with the default values initialized.
func NewGetEventEventgroupOccurrencesParams() *GetEventEventgroupOccurrencesParams {
	var ()
	return &GetEventEventgroupOccurrencesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEventEventgroupOccurrencesParamsWithTimeout creates a new GetEventEventgroupOccurrencesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEventEventgroupOccurrencesParamsWithTimeout(timeout time.Duration) *GetEventEventgroupOccurrencesParams {
	var ()
	return &GetEventEventgroupOccurrencesParams{

		timeout: timeout,
	}
}

// NewGetEventEventgroupOccurrencesParamsWithContext creates a new GetEventEventgroupOccurrencesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEventEventgroupOccurrencesParamsWithContext(ctx context.Context) *GetEventEventgroupOccurrencesParams {
	var ()
	return &GetEventEventgroupOccurrencesParams{

		Context: ctx,
	}
}

// NewGetEventEventgroupOccurrencesParamsWithHTTPClient creates a new GetEventEventgroupOccurrencesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEventEventgroupOccurrencesParamsWithHTTPClient(client *http.Client) *GetEventEventgroupOccurrencesParams {
	var ()
	return &GetEventEventgroupOccurrencesParams{
		HTTPClient: client,
	}
}

/*GetEventEventgroupOccurrencesParams contains all the parameters to send to the API endpoint
for the get event eventgroup occurrences operation typically these are written to a http.Request
*/
type GetEventEventgroupOccurrencesParams struct {

	/*Begin
	  events that are in progress after this time

	*/
	Begin *int64
	/*Cause
	  Filter for cause

	*/
	Cause *string
	/*Dir
	  The direction of the sort.

	*/
	Dir *string
	/*End
	  events that were in progress before this time

	*/
	End *int64
	/*EventCount
	  events for which event_count > this

	*/
	EventCount *int64
	/*Fixer
	  Filter for eventgroup fixer

	*/
	Fixer *string
	/*Ignore
	  Filter for ignored eventgroups

	*/
	Ignore *bool
	/*Limit
	  Return no more than this many results at once (see resume).

	*/
	Limit *int64
	/*Resolved
	  Filter for resolved eventgroups

	*/
	Resolved *bool
	/*Resume
	  Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options).

	*/
	Resume *string
	/*Sort
	  The field that will be used for sorting.

	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) WithTimeout(timeout time.Duration) *GetEventEventgroupOccurrencesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) WithContext(ctx context.Context) *GetEventEventgroupOccurrencesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) WithHTTPClient(client *http.Client) *GetEventEventgroupOccurrencesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBegin adds the begin to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) WithBegin(begin *int64) *GetEventEventgroupOccurrencesParams {
	o.SetBegin(begin)
	return o
}

// SetBegin adds the begin to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) SetBegin(begin *int64) {
	o.Begin = begin
}

// WithCause adds the cause to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) WithCause(cause *string) *GetEventEventgroupOccurrencesParams {
	o.SetCause(cause)
	return o
}

// SetCause adds the cause to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) SetCause(cause *string) {
	o.Cause = cause
}

// WithDir adds the dir to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) WithDir(dir *string) *GetEventEventgroupOccurrencesParams {
	o.SetDir(dir)
	return o
}

// SetDir adds the dir to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) SetDir(dir *string) {
	o.Dir = dir
}

// WithEnd adds the end to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) WithEnd(end *int64) *GetEventEventgroupOccurrencesParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) SetEnd(end *int64) {
	o.End = end
}

// WithEventCount adds the eventCount to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) WithEventCount(eventCount *int64) *GetEventEventgroupOccurrencesParams {
	o.SetEventCount(eventCount)
	return o
}

// SetEventCount adds the eventCount to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) SetEventCount(eventCount *int64) {
	o.EventCount = eventCount
}

// WithFixer adds the fixer to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) WithFixer(fixer *string) *GetEventEventgroupOccurrencesParams {
	o.SetFixer(fixer)
	return o
}

// SetFixer adds the fixer to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) SetFixer(fixer *string) {
	o.Fixer = fixer
}

// WithIgnore adds the ignore to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) WithIgnore(ignore *bool) *GetEventEventgroupOccurrencesParams {
	o.SetIgnore(ignore)
	return o
}

// SetIgnore adds the ignore to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) SetIgnore(ignore *bool) {
	o.Ignore = ignore
}

// WithLimit adds the limit to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) WithLimit(limit *int64) *GetEventEventgroupOccurrencesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithResolved adds the resolved to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) WithResolved(resolved *bool) *GetEventEventgroupOccurrencesParams {
	o.SetResolved(resolved)
	return o
}

// SetResolved adds the resolved to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) SetResolved(resolved *bool) {
	o.Resolved = resolved
}

// WithResume adds the resume to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) WithResume(resume *string) *GetEventEventgroupOccurrencesParams {
	o.SetResume(resume)
	return o
}

// SetResume adds the resume to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) SetResume(resume *string) {
	o.Resume = resume
}

// WithSort adds the sort to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) WithSort(sort *string) *GetEventEventgroupOccurrencesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get event eventgroup occurrences params
func (o *GetEventEventgroupOccurrencesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetEventEventgroupOccurrencesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Cause != nil {

		// query param cause
		var qrCause string
		if o.Cause != nil {
			qrCause = *o.Cause
		}
		qCause := qrCause
		if qCause != "" {
			if err := r.SetQueryParam("cause", qCause); err != nil {
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

	if o.EventCount != nil {

		// query param event_count
		var qrEventCount int64
		if o.EventCount != nil {
			qrEventCount = *o.EventCount
		}
		qEventCount := swag.FormatInt64(qrEventCount)
		if qEventCount != "" {
			if err := r.SetQueryParam("event_count", qEventCount); err != nil {
				return err
			}
		}

	}

	if o.Fixer != nil {

		// query param fixer
		var qrFixer string
		if o.Fixer != nil {
			qrFixer = *o.Fixer
		}
		qFixer := qrFixer
		if qFixer != "" {
			if err := r.SetQueryParam("fixer", qFixer); err != nil {
				return err
			}
		}

	}

	if o.Ignore != nil {

		// query param ignore
		var qrIgnore bool
		if o.Ignore != nil {
			qrIgnore = *o.Ignore
		}
		qIgnore := swag.FormatBool(qrIgnore)
		if qIgnore != "" {
			if err := r.SetQueryParam("ignore", qIgnore); err != nil {
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

	if o.Resolved != nil {

		// query param resolved
		var qrResolved bool
		if o.Resolved != nil {
			qrResolved = *o.Resolved
		}
		qResolved := swag.FormatBool(qrResolved)
		if qResolved != "" {
			if err := r.SetQueryParam("resolved", qResolved); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
