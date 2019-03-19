// Code generated by go-swagger; DO NOT EDIT.

package statistics

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

// NewGetStatisticsKeysParams creates a new GetStatisticsKeysParams object
// with the default values initialized.
func NewGetStatisticsKeysParams() *GetStatisticsKeysParams {
	var ()
	return &GetStatisticsKeysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStatisticsKeysParamsWithTimeout creates a new GetStatisticsKeysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStatisticsKeysParamsWithTimeout(timeout time.Duration) *GetStatisticsKeysParams {
	var ()
	return &GetStatisticsKeysParams{

		timeout: timeout,
	}
}

// NewGetStatisticsKeysParamsWithContext creates a new GetStatisticsKeysParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStatisticsKeysParamsWithContext(ctx context.Context) *GetStatisticsKeysParams {
	var ()
	return &GetStatisticsKeysParams{

		Context: ctx,
	}
}

// NewGetStatisticsKeysParamsWithHTTPClient creates a new GetStatisticsKeysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStatisticsKeysParamsWithHTTPClient(client *http.Client) *GetStatisticsKeysParams {
	var ()
	return &GetStatisticsKeysParams{
		HTTPClient: client,
	}
}

/*GetStatisticsKeysParams contains all the parameters to send to the API endpoint
for the get statistics keys operation typically these are written to a http.Request
*/
type GetStatisticsKeysParams struct {

	/*Count
	  Only count matching keys, do not return meta-data.

	*/
	Count *bool
	/*Limit
	  Return no more than this many results at once (see resume).

	*/
	Limit *int64
	/*Queryable
	  Only list keys that can/cannot be queries. Default is true.

	*/
	Queryable *bool
	/*Resume
	  Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options).

	*/
	Resume *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get statistics keys params
func (o *GetStatisticsKeysParams) WithTimeout(timeout time.Duration) *GetStatisticsKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get statistics keys params
func (o *GetStatisticsKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get statistics keys params
func (o *GetStatisticsKeysParams) WithContext(ctx context.Context) *GetStatisticsKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get statistics keys params
func (o *GetStatisticsKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get statistics keys params
func (o *GetStatisticsKeysParams) WithHTTPClient(client *http.Client) *GetStatisticsKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get statistics keys params
func (o *GetStatisticsKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get statistics keys params
func (o *GetStatisticsKeysParams) WithCount(count *bool) *GetStatisticsKeysParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get statistics keys params
func (o *GetStatisticsKeysParams) SetCount(count *bool) {
	o.Count = count
}

// WithLimit adds the limit to the get statistics keys params
func (o *GetStatisticsKeysParams) WithLimit(limit *int64) *GetStatisticsKeysParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get statistics keys params
func (o *GetStatisticsKeysParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithQueryable adds the queryable to the get statistics keys params
func (o *GetStatisticsKeysParams) WithQueryable(queryable *bool) *GetStatisticsKeysParams {
	o.SetQueryable(queryable)
	return o
}

// SetQueryable adds the queryable to the get statistics keys params
func (o *GetStatisticsKeysParams) SetQueryable(queryable *bool) {
	o.Queryable = queryable
}

// WithResume adds the resume to the get statistics keys params
func (o *GetStatisticsKeysParams) WithResume(resume *string) *GetStatisticsKeysParams {
	o.SetResume(resume)
	return o
}

// SetResume adds the resume to the get statistics keys params
func (o *GetStatisticsKeysParams) SetResume(resume *string) {
	o.Resume = resume
}

// WriteToRequest writes these params to a swagger request
func (o *GetStatisticsKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Count != nil {

		// query param count
		var qrCount bool
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatBool(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("count", qCount); err != nil {
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

	if o.Queryable != nil {

		// query param queryable
		var qrQueryable bool
		if o.Queryable != nil {
			qrQueryable = *o.Queryable
		}
		qQueryable := swag.FormatBool(qrQueryable)
		if qQueryable != "" {
			if err := r.SetQueryParam("queryable", qQueryable); err != nil {
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