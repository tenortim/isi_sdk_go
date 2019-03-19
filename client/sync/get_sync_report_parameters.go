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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSyncReportParams creates a new GetSyncReportParams object
// with the default values initialized.
func NewGetSyncReportParams() *GetSyncReportParams {
	var ()
	return &GetSyncReportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSyncReportParamsWithTimeout creates a new GetSyncReportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSyncReportParamsWithTimeout(timeout time.Duration) *GetSyncReportParams {
	var ()
	return &GetSyncReportParams{

		timeout: timeout,
	}
}

// NewGetSyncReportParamsWithContext creates a new GetSyncReportParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSyncReportParamsWithContext(ctx context.Context) *GetSyncReportParams {
	var ()
	return &GetSyncReportParams{

		Context: ctx,
	}
}

// NewGetSyncReportParamsWithHTTPClient creates a new GetSyncReportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSyncReportParamsWithHTTPClient(client *http.Client) *GetSyncReportParams {
	var ()
	return &GetSyncReportParams{
		HTTPClient: client,
	}
}

/*GetSyncReportParams contains all the parameters to send to the API endpoint
for the get sync report operation typically these are written to a http.Request
*/
type GetSyncReportParams struct {

	/*SyncReportID
	  View a single SyncIQ report.

	*/
	SyncReportID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sync report params
func (o *GetSyncReportParams) WithTimeout(timeout time.Duration) *GetSyncReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sync report params
func (o *GetSyncReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sync report params
func (o *GetSyncReportParams) WithContext(ctx context.Context) *GetSyncReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sync report params
func (o *GetSyncReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sync report params
func (o *GetSyncReportParams) WithHTTPClient(client *http.Client) *GetSyncReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sync report params
func (o *GetSyncReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSyncReportID adds the syncReportID to the get sync report params
func (o *GetSyncReportParams) WithSyncReportID(syncReportID string) *GetSyncReportParams {
	o.SetSyncReportID(syncReportID)
	return o
}

// SetSyncReportID adds the syncReportId to the get sync report params
func (o *GetSyncReportParams) SetSyncReportID(syncReportID string) {
	o.SyncReportID = syncReportID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSyncReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param SyncReportId
	if err := r.SetPathParam("SyncReportId", o.SyncReportID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
