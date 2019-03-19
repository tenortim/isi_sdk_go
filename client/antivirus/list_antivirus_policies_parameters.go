// Code generated by go-swagger; DO NOT EDIT.

package antivirus

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

// NewListAntivirusPoliciesParams creates a new ListAntivirusPoliciesParams object
// with the default values initialized.
func NewListAntivirusPoliciesParams() *ListAntivirusPoliciesParams {
	var ()
	return &ListAntivirusPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAntivirusPoliciesParamsWithTimeout creates a new ListAntivirusPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAntivirusPoliciesParamsWithTimeout(timeout time.Duration) *ListAntivirusPoliciesParams {
	var ()
	return &ListAntivirusPoliciesParams{

		timeout: timeout,
	}
}

// NewListAntivirusPoliciesParamsWithContext creates a new ListAntivirusPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAntivirusPoliciesParamsWithContext(ctx context.Context) *ListAntivirusPoliciesParams {
	var ()
	return &ListAntivirusPoliciesParams{

		Context: ctx,
	}
}

// NewListAntivirusPoliciesParamsWithHTTPClient creates a new ListAntivirusPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAntivirusPoliciesParamsWithHTTPClient(client *http.Client) *ListAntivirusPoliciesParams {
	var ()
	return &ListAntivirusPoliciesParams{
		HTTPClient: client,
	}
}

/*ListAntivirusPoliciesParams contains all the parameters to send to the API endpoint
for the list antivirus policies operation typically these are written to a http.Request
*/
type ListAntivirusPoliciesParams struct {

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

// WithTimeout adds the timeout to the list antivirus policies params
func (o *ListAntivirusPoliciesParams) WithTimeout(timeout time.Duration) *ListAntivirusPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list antivirus policies params
func (o *ListAntivirusPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list antivirus policies params
func (o *ListAntivirusPoliciesParams) WithContext(ctx context.Context) *ListAntivirusPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list antivirus policies params
func (o *ListAntivirusPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list antivirus policies params
func (o *ListAntivirusPoliciesParams) WithHTTPClient(client *http.Client) *ListAntivirusPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list antivirus policies params
func (o *ListAntivirusPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDir adds the dir to the list antivirus policies params
func (o *ListAntivirusPoliciesParams) WithDir(dir *string) *ListAntivirusPoliciesParams {
	o.SetDir(dir)
	return o
}

// SetDir adds the dir to the list antivirus policies params
func (o *ListAntivirusPoliciesParams) SetDir(dir *string) {
	o.Dir = dir
}

// WithLimit adds the limit to the list antivirus policies params
func (o *ListAntivirusPoliciesParams) WithLimit(limit *int64) *ListAntivirusPoliciesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list antivirus policies params
func (o *ListAntivirusPoliciesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithResume adds the resume to the list antivirus policies params
func (o *ListAntivirusPoliciesParams) WithResume(resume *string) *ListAntivirusPoliciesParams {
	o.SetResume(resume)
	return o
}

// SetResume adds the resume to the list antivirus policies params
func (o *ListAntivirusPoliciesParams) SetResume(resume *string) {
	o.Resume = resume
}

// WithSort adds the sort to the list antivirus policies params
func (o *ListAntivirusPoliciesParams) WithSort(sort *string) *ListAntivirusPoliciesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the list antivirus policies params
func (o *ListAntivirusPoliciesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *ListAntivirusPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
