// Code generated by go-swagger; DO NOT EDIT.

package snapshot

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

// NewGetSnapshotSnapshotsSummaryParams creates a new GetSnapshotSnapshotsSummaryParams object
// with the default values initialized.
func NewGetSnapshotSnapshotsSummaryParams() *GetSnapshotSnapshotsSummaryParams {

	return &GetSnapshotSnapshotsSummaryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSnapshotSnapshotsSummaryParamsWithTimeout creates a new GetSnapshotSnapshotsSummaryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSnapshotSnapshotsSummaryParamsWithTimeout(timeout time.Duration) *GetSnapshotSnapshotsSummaryParams {

	return &GetSnapshotSnapshotsSummaryParams{

		timeout: timeout,
	}
}

// NewGetSnapshotSnapshotsSummaryParamsWithContext creates a new GetSnapshotSnapshotsSummaryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSnapshotSnapshotsSummaryParamsWithContext(ctx context.Context) *GetSnapshotSnapshotsSummaryParams {

	return &GetSnapshotSnapshotsSummaryParams{

		Context: ctx,
	}
}

// NewGetSnapshotSnapshotsSummaryParamsWithHTTPClient creates a new GetSnapshotSnapshotsSummaryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSnapshotSnapshotsSummaryParamsWithHTTPClient(client *http.Client) *GetSnapshotSnapshotsSummaryParams {

	return &GetSnapshotSnapshotsSummaryParams{
		HTTPClient: client,
	}
}

/*GetSnapshotSnapshotsSummaryParams contains all the parameters to send to the API endpoint
for the get snapshot snapshots summary operation typically these are written to a http.Request
*/
type GetSnapshotSnapshotsSummaryParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get snapshot snapshots summary params
func (o *GetSnapshotSnapshotsSummaryParams) WithTimeout(timeout time.Duration) *GetSnapshotSnapshotsSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get snapshot snapshots summary params
func (o *GetSnapshotSnapshotsSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get snapshot snapshots summary params
func (o *GetSnapshotSnapshotsSummaryParams) WithContext(ctx context.Context) *GetSnapshotSnapshotsSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get snapshot snapshots summary params
func (o *GetSnapshotSnapshotsSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get snapshot snapshots summary params
func (o *GetSnapshotSnapshotsSummaryParams) WithHTTPClient(client *http.Client) *GetSnapshotSnapshotsSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get snapshot snapshots summary params
func (o *GetSnapshotSnapshotsSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSnapshotSnapshotsSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}