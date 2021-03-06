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

// NewGetSmbSessionsParams creates a new GetSmbSessionsParams object
// with the default values initialized.
func NewGetSmbSessionsParams() *GetSmbSessionsParams {
	var ()
	return &GetSmbSessionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSmbSessionsParamsWithTimeout creates a new GetSmbSessionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSmbSessionsParamsWithTimeout(timeout time.Duration) *GetSmbSessionsParams {
	var ()
	return &GetSmbSessionsParams{

		timeout: timeout,
	}
}

// NewGetSmbSessionsParamsWithContext creates a new GetSmbSessionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSmbSessionsParamsWithContext(ctx context.Context) *GetSmbSessionsParams {
	var ()
	return &GetSmbSessionsParams{

		Context: ctx,
	}
}

// NewGetSmbSessionsParamsWithHTTPClient creates a new GetSmbSessionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSmbSessionsParamsWithHTTPClient(client *http.Client) *GetSmbSessionsParams {
	var ()
	return &GetSmbSessionsParams{
		HTTPClient: client,
	}
}

/*GetSmbSessionsParams contains all the parameters to send to the API endpoint
for the get smb sessions operation typically these are written to a http.Request
*/
type GetSmbSessionsParams struct {

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
	  Order results by this field.

	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get smb sessions params
func (o *GetSmbSessionsParams) WithTimeout(timeout time.Duration) *GetSmbSessionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get smb sessions params
func (o *GetSmbSessionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get smb sessions params
func (o *GetSmbSessionsParams) WithContext(ctx context.Context) *GetSmbSessionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get smb sessions params
func (o *GetSmbSessionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get smb sessions params
func (o *GetSmbSessionsParams) WithHTTPClient(client *http.Client) *GetSmbSessionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get smb sessions params
func (o *GetSmbSessionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDir adds the dir to the get smb sessions params
func (o *GetSmbSessionsParams) WithDir(dir *string) *GetSmbSessionsParams {
	o.SetDir(dir)
	return o
}

// SetDir adds the dir to the get smb sessions params
func (o *GetSmbSessionsParams) SetDir(dir *string) {
	o.Dir = dir
}

// WithLimit adds the limit to the get smb sessions params
func (o *GetSmbSessionsParams) WithLimit(limit *int64) *GetSmbSessionsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get smb sessions params
func (o *GetSmbSessionsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithResume adds the resume to the get smb sessions params
func (o *GetSmbSessionsParams) WithResume(resume *string) *GetSmbSessionsParams {
	o.SetResume(resume)
	return o
}

// SetResume adds the resume to the get smb sessions params
func (o *GetSmbSessionsParams) SetResume(resume *string) {
	o.Resume = resume
}

// WithSort adds the sort to the get smb sessions params
func (o *GetSmbSessionsParams) WithSort(sort *string) *GetSmbSessionsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get smb sessions params
func (o *GetSmbSessionsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetSmbSessionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
