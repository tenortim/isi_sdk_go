// Code generated by go-swagger; DO NOT EDIT.

package sync

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

// NewListSyncPoliciesParams creates a new ListSyncPoliciesParams object
// with the default values initialized.
func NewListSyncPoliciesParams() *ListSyncPoliciesParams {
	var ()
	return &ListSyncPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSyncPoliciesParamsWithTimeout creates a new ListSyncPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSyncPoliciesParamsWithTimeout(timeout time.Duration) *ListSyncPoliciesParams {
	var ()
	return &ListSyncPoliciesParams{

		timeout: timeout,
	}
}

// NewListSyncPoliciesParamsWithContext creates a new ListSyncPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSyncPoliciesParamsWithContext(ctx context.Context) *ListSyncPoliciesParams {
	var ()
	return &ListSyncPoliciesParams{

		Context: ctx,
	}
}

// NewListSyncPoliciesParamsWithHTTPClient creates a new ListSyncPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSyncPoliciesParamsWithHTTPClient(client *http.Client) *ListSyncPoliciesParams {
	var ()
	return &ListSyncPoliciesParams{
		HTTPClient: client,
	}
}

/*ListSyncPoliciesParams contains all the parameters to send to the API endpoint
for the list sync policies operation typically these are written to a http.Request
*/
type ListSyncPoliciesParams struct {

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
	/*Scope
	  If specified as "effective" or not specified, all fields are returned.  If specified as "user", only fields with non-default values are shown.  If specified as "default", the original values are returned.

	*/
	Scope *string
	/*Sort
	  The field that will be used for sorting.

	*/
	Sort *string
	/*Summary
	  Show only summary properties

	*/
	Summary *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list sync policies params
func (o *ListSyncPoliciesParams) WithTimeout(timeout time.Duration) *ListSyncPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list sync policies params
func (o *ListSyncPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list sync policies params
func (o *ListSyncPoliciesParams) WithContext(ctx context.Context) *ListSyncPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list sync policies params
func (o *ListSyncPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list sync policies params
func (o *ListSyncPoliciesParams) WithHTTPClient(client *http.Client) *ListSyncPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list sync policies params
func (o *ListSyncPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDir adds the dir to the list sync policies params
func (o *ListSyncPoliciesParams) WithDir(dir *string) *ListSyncPoliciesParams {
	o.SetDir(dir)
	return o
}

// SetDir adds the dir to the list sync policies params
func (o *ListSyncPoliciesParams) SetDir(dir *string) {
	o.Dir = dir
}

// WithLimit adds the limit to the list sync policies params
func (o *ListSyncPoliciesParams) WithLimit(limit *int64) *ListSyncPoliciesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list sync policies params
func (o *ListSyncPoliciesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithResume adds the resume to the list sync policies params
func (o *ListSyncPoliciesParams) WithResume(resume *string) *ListSyncPoliciesParams {
	o.SetResume(resume)
	return o
}

// SetResume adds the resume to the list sync policies params
func (o *ListSyncPoliciesParams) SetResume(resume *string) {
	o.Resume = resume
}

// WithScope adds the scope to the list sync policies params
func (o *ListSyncPoliciesParams) WithScope(scope *string) *ListSyncPoliciesParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the list sync policies params
func (o *ListSyncPoliciesParams) SetScope(scope *string) {
	o.Scope = scope
}

// WithSort adds the sort to the list sync policies params
func (o *ListSyncPoliciesParams) WithSort(sort *string) *ListSyncPoliciesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the list sync policies params
func (o *ListSyncPoliciesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithSummary adds the summary to the list sync policies params
func (o *ListSyncPoliciesParams) WithSummary(summary *bool) *ListSyncPoliciesParams {
	o.SetSummary(summary)
	return o
}

// SetSummary adds the summary to the list sync policies params
func (o *ListSyncPoliciesParams) SetSummary(summary *bool) {
	o.Summary = summary
}

// WriteToRequest writes these params to a swagger request
func (o *ListSyncPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Scope != nil {

		// query param scope
		var qrScope string
		if o.Scope != nil {
			qrScope = *o.Scope
		}
		qScope := qrScope
		if qScope != "" {
			if err := r.SetQueryParam("scope", qScope); err != nil {
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

	if o.Summary != nil {

		// query param summary
		var qrSummary bool
		if o.Summary != nil {
			qrSummary = *o.Summary
		}
		qSummary := swag.FormatBool(qrSummary)
		if qSummary != "" {
			if err := r.SetQueryParam("summary", qSummary); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
