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

// NewListSyncRulesParams creates a new ListSyncRulesParams object
// with the default values initialized.
func NewListSyncRulesParams() *ListSyncRulesParams {
	var ()
	return &ListSyncRulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSyncRulesParamsWithTimeout creates a new ListSyncRulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSyncRulesParamsWithTimeout(timeout time.Duration) *ListSyncRulesParams {
	var ()
	return &ListSyncRulesParams{

		timeout: timeout,
	}
}

// NewListSyncRulesParamsWithContext creates a new ListSyncRulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSyncRulesParamsWithContext(ctx context.Context) *ListSyncRulesParams {
	var ()
	return &ListSyncRulesParams{

		Context: ctx,
	}
}

// NewListSyncRulesParamsWithHTTPClient creates a new ListSyncRulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSyncRulesParamsWithHTTPClient(client *http.Client) *ListSyncRulesParams {
	var ()
	return &ListSyncRulesParams{
		HTTPClient: client,
	}
}

/*ListSyncRulesParams contains all the parameters to send to the API endpoint
for the list sync rules operation typically these are written to a http.Request
*/
type ListSyncRulesParams struct {

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
	/*Type
	  Filter the returned rules to include only those with this rule type.

	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list sync rules params
func (o *ListSyncRulesParams) WithTimeout(timeout time.Duration) *ListSyncRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list sync rules params
func (o *ListSyncRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list sync rules params
func (o *ListSyncRulesParams) WithContext(ctx context.Context) *ListSyncRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list sync rules params
func (o *ListSyncRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list sync rules params
func (o *ListSyncRulesParams) WithHTTPClient(client *http.Client) *ListSyncRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list sync rules params
func (o *ListSyncRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDir adds the dir to the list sync rules params
func (o *ListSyncRulesParams) WithDir(dir *string) *ListSyncRulesParams {
	o.SetDir(dir)
	return o
}

// SetDir adds the dir to the list sync rules params
func (o *ListSyncRulesParams) SetDir(dir *string) {
	o.Dir = dir
}

// WithLimit adds the limit to the list sync rules params
func (o *ListSyncRulesParams) WithLimit(limit *int64) *ListSyncRulesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list sync rules params
func (o *ListSyncRulesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithResume adds the resume to the list sync rules params
func (o *ListSyncRulesParams) WithResume(resume *string) *ListSyncRulesParams {
	o.SetResume(resume)
	return o
}

// SetResume adds the resume to the list sync rules params
func (o *ListSyncRulesParams) SetResume(resume *string) {
	o.Resume = resume
}

// WithSort adds the sort to the list sync rules params
func (o *ListSyncRulesParams) WithSort(sort *string) *ListSyncRulesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the list sync rules params
func (o *ListSyncRulesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithType adds the typeVar to the list sync rules params
func (o *ListSyncRulesParams) WithType(typeVar *string) *ListSyncRulesParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the list sync rules params
func (o *ListSyncRulesParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *ListSyncRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
