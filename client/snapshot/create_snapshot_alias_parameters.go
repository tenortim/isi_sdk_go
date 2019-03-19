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

	models "github.com/tenortim/isi_sdk_go/models"
)

// NewCreateSnapshotAliasParams creates a new CreateSnapshotAliasParams object
// with the default values initialized.
func NewCreateSnapshotAliasParams() *CreateSnapshotAliasParams {
	var ()
	return &CreateSnapshotAliasParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSnapshotAliasParamsWithTimeout creates a new CreateSnapshotAliasParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSnapshotAliasParamsWithTimeout(timeout time.Duration) *CreateSnapshotAliasParams {
	var ()
	return &CreateSnapshotAliasParams{

		timeout: timeout,
	}
}

// NewCreateSnapshotAliasParamsWithContext creates a new CreateSnapshotAliasParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSnapshotAliasParamsWithContext(ctx context.Context) *CreateSnapshotAliasParams {
	var ()
	return &CreateSnapshotAliasParams{

		Context: ctx,
	}
}

// NewCreateSnapshotAliasParamsWithHTTPClient creates a new CreateSnapshotAliasParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSnapshotAliasParamsWithHTTPClient(client *http.Client) *CreateSnapshotAliasParams {
	var ()
	return &CreateSnapshotAliasParams{
		HTTPClient: client,
	}
}

/*CreateSnapshotAliasParams contains all the parameters to send to the API endpoint
for the create snapshot alias operation typically these are written to a http.Request
*/
type CreateSnapshotAliasParams struct {

	/*SnapshotAlias*/
	SnapshotAlias *models.SnapshotAliasCreateParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create snapshot alias params
func (o *CreateSnapshotAliasParams) WithTimeout(timeout time.Duration) *CreateSnapshotAliasParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create snapshot alias params
func (o *CreateSnapshotAliasParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create snapshot alias params
func (o *CreateSnapshotAliasParams) WithContext(ctx context.Context) *CreateSnapshotAliasParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create snapshot alias params
func (o *CreateSnapshotAliasParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create snapshot alias params
func (o *CreateSnapshotAliasParams) WithHTTPClient(client *http.Client) *CreateSnapshotAliasParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create snapshot alias params
func (o *CreateSnapshotAliasParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSnapshotAlias adds the snapshotAlias to the create snapshot alias params
func (o *CreateSnapshotAliasParams) WithSnapshotAlias(snapshotAlias *models.SnapshotAliasCreateParams) *CreateSnapshotAliasParams {
	o.SetSnapshotAlias(snapshotAlias)
	return o
}

// SetSnapshotAlias adds the snapshotAlias to the create snapshot alias params
func (o *CreateSnapshotAliasParams) SetSnapshotAlias(snapshotAlias *models.SnapshotAliasCreateParams) {
	o.SnapshotAlias = snapshotAlias
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSnapshotAliasParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SnapshotAlias != nil {
		if err := r.SetBodyParam(o.SnapshotAlias); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}