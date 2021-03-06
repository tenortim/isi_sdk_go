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

// NewDeleteSnapshotAliasParams creates a new DeleteSnapshotAliasParams object
// with the default values initialized.
func NewDeleteSnapshotAliasParams() *DeleteSnapshotAliasParams {
	var ()
	return &DeleteSnapshotAliasParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSnapshotAliasParamsWithTimeout creates a new DeleteSnapshotAliasParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSnapshotAliasParamsWithTimeout(timeout time.Duration) *DeleteSnapshotAliasParams {
	var ()
	return &DeleteSnapshotAliasParams{

		timeout: timeout,
	}
}

// NewDeleteSnapshotAliasParamsWithContext creates a new DeleteSnapshotAliasParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSnapshotAliasParamsWithContext(ctx context.Context) *DeleteSnapshotAliasParams {
	var ()
	return &DeleteSnapshotAliasParams{

		Context: ctx,
	}
}

// NewDeleteSnapshotAliasParamsWithHTTPClient creates a new DeleteSnapshotAliasParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSnapshotAliasParamsWithHTTPClient(client *http.Client) *DeleteSnapshotAliasParams {
	var ()
	return &DeleteSnapshotAliasParams{
		HTTPClient: client,
	}
}

/*DeleteSnapshotAliasParams contains all the parameters to send to the API endpoint
for the delete snapshot alias operation typically these are written to a http.Request
*/
type DeleteSnapshotAliasParams struct {

	/*SnapshotAliasID
	  Delete the snapshot alias

	*/
	SnapshotAliasID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete snapshot alias params
func (o *DeleteSnapshotAliasParams) WithTimeout(timeout time.Duration) *DeleteSnapshotAliasParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete snapshot alias params
func (o *DeleteSnapshotAliasParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete snapshot alias params
func (o *DeleteSnapshotAliasParams) WithContext(ctx context.Context) *DeleteSnapshotAliasParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete snapshot alias params
func (o *DeleteSnapshotAliasParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete snapshot alias params
func (o *DeleteSnapshotAliasParams) WithHTTPClient(client *http.Client) *DeleteSnapshotAliasParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete snapshot alias params
func (o *DeleteSnapshotAliasParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSnapshotAliasID adds the snapshotAliasID to the delete snapshot alias params
func (o *DeleteSnapshotAliasParams) WithSnapshotAliasID(snapshotAliasID string) *DeleteSnapshotAliasParams {
	o.SetSnapshotAliasID(snapshotAliasID)
	return o
}

// SetSnapshotAliasID adds the snapshotAliasId to the delete snapshot alias params
func (o *DeleteSnapshotAliasParams) SetSnapshotAliasID(snapshotAliasID string) {
	o.SnapshotAliasID = snapshotAliasID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSnapshotAliasParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
