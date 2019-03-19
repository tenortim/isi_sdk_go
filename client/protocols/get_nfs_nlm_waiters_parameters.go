// Code generated by go-swagger; DO NOT EDIT.

package protocols

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

// NewGetNfsNlmWaitersParams creates a new GetNfsNlmWaitersParams object
// with the default values initialized.
func NewGetNfsNlmWaitersParams() *GetNfsNlmWaitersParams {
	var ()
	return &GetNfsNlmWaitersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNfsNlmWaitersParamsWithTimeout creates a new GetNfsNlmWaitersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNfsNlmWaitersParamsWithTimeout(timeout time.Duration) *GetNfsNlmWaitersParams {
	var ()
	return &GetNfsNlmWaitersParams{

		timeout: timeout,
	}
}

// NewGetNfsNlmWaitersParamsWithContext creates a new GetNfsNlmWaitersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNfsNlmWaitersParamsWithContext(ctx context.Context) *GetNfsNlmWaitersParams {
	var ()
	return &GetNfsNlmWaitersParams{

		Context: ctx,
	}
}

// NewGetNfsNlmWaitersParamsWithHTTPClient creates a new GetNfsNlmWaitersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNfsNlmWaitersParamsWithHTTPClient(client *http.Client) *GetNfsNlmWaitersParams {
	var ()
	return &GetNfsNlmWaitersParams{
		HTTPClient: client,
	}
}

/*GetNfsNlmWaitersParams contains all the parameters to send to the API endpoint
for the get nfs nlm waiters operation typically these are written to a http.Request
*/
type GetNfsNlmWaitersParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get nfs nlm waiters params
func (o *GetNfsNlmWaitersParams) WithTimeout(timeout time.Duration) *GetNfsNlmWaitersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nfs nlm waiters params
func (o *GetNfsNlmWaitersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nfs nlm waiters params
func (o *GetNfsNlmWaitersParams) WithContext(ctx context.Context) *GetNfsNlmWaitersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nfs nlm waiters params
func (o *GetNfsNlmWaitersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nfs nlm waiters params
func (o *GetNfsNlmWaitersParams) WithHTTPClient(client *http.Client) *GetNfsNlmWaitersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nfs nlm waiters params
func (o *GetNfsNlmWaitersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDir adds the dir to the get nfs nlm waiters params
func (o *GetNfsNlmWaitersParams) WithDir(dir *string) *GetNfsNlmWaitersParams {
	o.SetDir(dir)
	return o
}

// SetDir adds the dir to the get nfs nlm waiters params
func (o *GetNfsNlmWaitersParams) SetDir(dir *string) {
	o.Dir = dir
}

// WithLimit adds the limit to the get nfs nlm waiters params
func (o *GetNfsNlmWaitersParams) WithLimit(limit *int64) *GetNfsNlmWaitersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get nfs nlm waiters params
func (o *GetNfsNlmWaitersParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithResume adds the resume to the get nfs nlm waiters params
func (o *GetNfsNlmWaitersParams) WithResume(resume *string) *GetNfsNlmWaitersParams {
	o.SetResume(resume)
	return o
}

// SetResume adds the resume to the get nfs nlm waiters params
func (o *GetNfsNlmWaitersParams) SetResume(resume *string) {
	o.Resume = resume
}

// WithSort adds the sort to the get nfs nlm waiters params
func (o *GetNfsNlmWaitersParams) WithSort(sort *string) *GetNfsNlmWaitersParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get nfs nlm waiters params
func (o *GetNfsNlmWaitersParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetNfsNlmWaitersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
