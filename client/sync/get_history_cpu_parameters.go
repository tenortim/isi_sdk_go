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

// NewGetHistoryCPUParams creates a new GetHistoryCPUParams object
// with the default values initialized.
func NewGetHistoryCPUParams() *GetHistoryCPUParams {
	var ()
	return &GetHistoryCPUParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHistoryCPUParamsWithTimeout creates a new GetHistoryCPUParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHistoryCPUParamsWithTimeout(timeout time.Duration) *GetHistoryCPUParams {
	var ()
	return &GetHistoryCPUParams{

		timeout: timeout,
	}
}

// NewGetHistoryCPUParamsWithContext creates a new GetHistoryCPUParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHistoryCPUParamsWithContext(ctx context.Context) *GetHistoryCPUParams {
	var ()
	return &GetHistoryCPUParams{

		Context: ctx,
	}
}

// NewGetHistoryCPUParamsWithHTTPClient creates a new GetHistoryCPUParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHistoryCPUParamsWithHTTPClient(client *http.Client) *GetHistoryCPUParams {
	var ()
	return &GetHistoryCPUParams{
		HTTPClient: client,
	}
}

/*GetHistoryCPUParams contains all the parameters to send to the API endpoint
for the get history Cpu operation typically these are written to a http.Request
*/
type GetHistoryCPUParams struct {

	/*Begin
	  Begin timestamp for time-series report.

	*/
	Begin *int64
	/*End
	  End timestamp for time-series report.

	*/
	End *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get history Cpu params
func (o *GetHistoryCPUParams) WithTimeout(timeout time.Duration) *GetHistoryCPUParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get history Cpu params
func (o *GetHistoryCPUParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get history Cpu params
func (o *GetHistoryCPUParams) WithContext(ctx context.Context) *GetHistoryCPUParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get history Cpu params
func (o *GetHistoryCPUParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get history Cpu params
func (o *GetHistoryCPUParams) WithHTTPClient(client *http.Client) *GetHistoryCPUParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get history Cpu params
func (o *GetHistoryCPUParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBegin adds the begin to the get history Cpu params
func (o *GetHistoryCPUParams) WithBegin(begin *int64) *GetHistoryCPUParams {
	o.SetBegin(begin)
	return o
}

// SetBegin adds the begin to the get history Cpu params
func (o *GetHistoryCPUParams) SetBegin(begin *int64) {
	o.Begin = begin
}

// WithEnd adds the end to the get history Cpu params
func (o *GetHistoryCPUParams) WithEnd(end *int64) *GetHistoryCPUParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get history Cpu params
func (o *GetHistoryCPUParams) SetEnd(end *int64) {
	o.End = end
}

// WriteToRequest writes these params to a swagger request
func (o *GetHistoryCPUParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}