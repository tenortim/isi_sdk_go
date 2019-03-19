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

// NewGetSnapshotAliasParams creates a new GetSnapshotAliasParams object
// with the default values initialized.
func NewGetSnapshotAliasParams() *GetSnapshotAliasParams {
	var ()
	return &GetSnapshotAliasParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSnapshotAliasParamsWithTimeout creates a new GetSnapshotAliasParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSnapshotAliasParamsWithTimeout(timeout time.Duration) *GetSnapshotAliasParams {
	var ()
	return &GetSnapshotAliasParams{

		timeout: timeout,
	}
}

// NewGetSnapshotAliasParamsWithContext creates a new GetSnapshotAliasParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSnapshotAliasParamsWithContext(ctx context.Context) *GetSnapshotAliasParams {
	var ()
	return &GetSnapshotAliasParams{

		Context: ctx,
	}
}

// NewGetSnapshotAliasParamsWithHTTPClient creates a new GetSnapshotAliasParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSnapshotAliasParamsWithHTTPClient(client *http.Client) *GetSnapshotAliasParams {
	var ()
	return &GetSnapshotAliasParams{
		HTTPClient: client,
	}
}

/*GetSnapshotAliasParams contains all the parameters to send to the API endpoint
for the get snapshot alias operation typically these are written to a http.Request
*/
type GetSnapshotAliasParams struct {

	/*SnapshotAliasID
	  Retrieve snapshot alias information.

	*/
	SnapshotAliasID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get snapshot alias params
func (o *GetSnapshotAliasParams) WithTimeout(timeout time.Duration) *GetSnapshotAliasParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get snapshot alias params
func (o *GetSnapshotAliasParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get snapshot alias params
func (o *GetSnapshotAliasParams) WithContext(ctx context.Context) *GetSnapshotAliasParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get snapshot alias params
func (o *GetSnapshotAliasParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get snapshot alias params
func (o *GetSnapshotAliasParams) WithHTTPClient(client *http.Client) *GetSnapshotAliasParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get snapshot alias params
func (o *GetSnapshotAliasParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSnapshotAliasID adds the snapshotAliasID to the get snapshot alias params
func (o *GetSnapshotAliasParams) WithSnapshotAliasID(snapshotAliasID string) *GetSnapshotAliasParams {
	o.SetSnapshotAliasID(snapshotAliasID)
	return o
}

// SetSnapshotAliasID adds the snapshotAliasId to the get snapshot alias params
func (o *GetSnapshotAliasParams) SetSnapshotAliasID(snapshotAliasID string) {
	o.SnapshotAliasID = snapshotAliasID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSnapshotAliasParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param SnapshotAliasId
	if err := r.SetPathParam("SnapshotAliasId", o.SnapshotAliasID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
