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

// NewGetHistoryNetworkParams creates a new GetHistoryNetworkParams object
// with the default values initialized.
func NewGetHistoryNetworkParams() *GetHistoryNetworkParams {
	var ()
	return &GetHistoryNetworkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHistoryNetworkParamsWithTimeout creates a new GetHistoryNetworkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHistoryNetworkParamsWithTimeout(timeout time.Duration) *GetHistoryNetworkParams {
	var ()
	return &GetHistoryNetworkParams{

		timeout: timeout,
	}
}

// NewGetHistoryNetworkParamsWithContext creates a new GetHistoryNetworkParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHistoryNetworkParamsWithContext(ctx context.Context) *GetHistoryNetworkParams {
	var ()
	return &GetHistoryNetworkParams{

		Context: ctx,
	}
}

// NewGetHistoryNetworkParamsWithHTTPClient creates a new GetHistoryNetworkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHistoryNetworkParamsWithHTTPClient(client *http.Client) *GetHistoryNetworkParams {
	var ()
	return &GetHistoryNetworkParams{
		HTTPClient: client,
	}
}

/*GetHistoryNetworkParams contains all the parameters to send to the API endpoint
for the get history network operation typically these are written to a http.Request
*/
type GetHistoryNetworkParams struct {

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

// WithTimeout adds the timeout to the get history network params
func (o *GetHistoryNetworkParams) WithTimeout(timeout time.Duration) *GetHistoryNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get history network params
func (o *GetHistoryNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get history network params
func (o *GetHistoryNetworkParams) WithContext(ctx context.Context) *GetHistoryNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get history network params
func (o *GetHistoryNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get history network params
func (o *GetHistoryNetworkParams) WithHTTPClient(client *http.Client) *GetHistoryNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get history network params
func (o *GetHistoryNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBegin adds the begin to the get history network params
func (o *GetHistoryNetworkParams) WithBegin(begin *int64) *GetHistoryNetworkParams {
	o.SetBegin(begin)
	return o
}

// SetBegin adds the begin to the get history network params
func (o *GetHistoryNetworkParams) SetBegin(begin *int64) {
	o.Begin = begin
}

// WithEnd adds the end to the get history network params
func (o *GetHistoryNetworkParams) WithEnd(end *int64) *GetHistoryNetworkParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get history network params
func (o *GetHistoryNetworkParams) SetEnd(end *int64) {
	o.End = end
}

// WriteToRequest writes these params to a swagger request
func (o *GetHistoryNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
