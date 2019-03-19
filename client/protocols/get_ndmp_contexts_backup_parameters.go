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

// NewGetNdmpContextsBackupParams creates a new GetNdmpContextsBackupParams object
// with the default values initialized.
func NewGetNdmpContextsBackupParams() *GetNdmpContextsBackupParams {
	var ()
	return &GetNdmpContextsBackupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNdmpContextsBackupParamsWithTimeout creates a new GetNdmpContextsBackupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNdmpContextsBackupParamsWithTimeout(timeout time.Duration) *GetNdmpContextsBackupParams {
	var ()
	return &GetNdmpContextsBackupParams{

		timeout: timeout,
	}
}

// NewGetNdmpContextsBackupParamsWithContext creates a new GetNdmpContextsBackupParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNdmpContextsBackupParamsWithContext(ctx context.Context) *GetNdmpContextsBackupParams {
	var ()
	return &GetNdmpContextsBackupParams{

		Context: ctx,
	}
}

// NewGetNdmpContextsBackupParamsWithHTTPClient creates a new GetNdmpContextsBackupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNdmpContextsBackupParamsWithHTTPClient(client *http.Client) *GetNdmpContextsBackupParams {
	var ()
	return &GetNdmpContextsBackupParams{
		HTTPClient: client,
	}
}

/*GetNdmpContextsBackupParams contains all the parameters to send to the API endpoint
for the get ndmp contexts backup operation typically these are written to a http.Request
*/
type GetNdmpContextsBackupParams struct {

	/*Limit
	  Return no more than this many results at once (see resume).

	*/
	Limit *int64
	/*Resume
	  Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options).

	*/
	Resume *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ndmp contexts backup params
func (o *GetNdmpContextsBackupParams) WithTimeout(timeout time.Duration) *GetNdmpContextsBackupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ndmp contexts backup params
func (o *GetNdmpContextsBackupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ndmp contexts backup params
func (o *GetNdmpContextsBackupParams) WithContext(ctx context.Context) *GetNdmpContextsBackupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ndmp contexts backup params
func (o *GetNdmpContextsBackupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ndmp contexts backup params
func (o *GetNdmpContextsBackupParams) WithHTTPClient(client *http.Client) *GetNdmpContextsBackupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ndmp contexts backup params
func (o *GetNdmpContextsBackupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get ndmp contexts backup params
func (o *GetNdmpContextsBackupParams) WithLimit(limit *int64) *GetNdmpContextsBackupParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get ndmp contexts backup params
func (o *GetNdmpContextsBackupParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithResume adds the resume to the get ndmp contexts backup params
func (o *GetNdmpContextsBackupParams) WithResume(resume *string) *GetNdmpContextsBackupParams {
	o.SetResume(resume)
	return o
}

// SetResume adds the resume to the get ndmp contexts backup params
func (o *GetNdmpContextsBackupParams) SetResume(resume *string) {
	o.Resume = resume
}

// WriteToRequest writes these params to a swagger request
func (o *GetNdmpContextsBackupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}