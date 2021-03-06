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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetReportsThreatParams creates a new GetReportsThreatParams object
// with the default values initialized.
func NewGetReportsThreatParams() *GetReportsThreatParams {
	var ()
	return &GetReportsThreatParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReportsThreatParamsWithTimeout creates a new GetReportsThreatParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReportsThreatParamsWithTimeout(timeout time.Duration) *GetReportsThreatParams {
	var ()
	return &GetReportsThreatParams{

		timeout: timeout,
	}
}

// NewGetReportsThreatParamsWithContext creates a new GetReportsThreatParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReportsThreatParamsWithContext(ctx context.Context) *GetReportsThreatParams {
	var ()
	return &GetReportsThreatParams{

		Context: ctx,
	}
}

// NewGetReportsThreatParamsWithHTTPClient creates a new GetReportsThreatParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReportsThreatParamsWithHTTPClient(client *http.Client) *GetReportsThreatParams {
	var ()
	return &GetReportsThreatParams{
		HTTPClient: client,
	}
}

/*GetReportsThreatParams contains all the parameters to send to the API endpoint
for the get reports threat operation typically these are written to a http.Request
*/
type GetReportsThreatParams struct {

	/*ReportsThreatID
	  Retrieve one antivirus threat report.

	*/
	ReportsThreatID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get reports threat params
func (o *GetReportsThreatParams) WithTimeout(timeout time.Duration) *GetReportsThreatParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get reports threat params
func (o *GetReportsThreatParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get reports threat params
func (o *GetReportsThreatParams) WithContext(ctx context.Context) *GetReportsThreatParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get reports threat params
func (o *GetReportsThreatParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get reports threat params
func (o *GetReportsThreatParams) WithHTTPClient(client *http.Client) *GetReportsThreatParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get reports threat params
func (o *GetReportsThreatParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReportsThreatID adds the reportsThreatID to the get reports threat params
func (o *GetReportsThreatParams) WithReportsThreatID(reportsThreatID string) *GetReportsThreatParams {
	o.SetReportsThreatID(reportsThreatID)
	return o
}

// SetReportsThreatID adds the reportsThreatId to the get reports threat params
func (o *GetReportsThreatParams) SetReportsThreatID(reportsThreatID string) {
	o.ReportsThreatID = reportsThreatID
}

// WriteToRequest writes these params to a swagger request
func (o *GetReportsThreatParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ReportsThreatId
	if err := r.SetPathParam("ReportsThreatId", o.ReportsThreatID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
