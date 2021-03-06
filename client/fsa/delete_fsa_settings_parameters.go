// Code generated by go-swagger; DO NOT EDIT.

package fsa

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

// NewDeleteFsaSettingsParams creates a new DeleteFsaSettingsParams object
// with the default values initialized.
func NewDeleteFsaSettingsParams() *DeleteFsaSettingsParams {

	return &DeleteFsaSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFsaSettingsParamsWithTimeout creates a new DeleteFsaSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteFsaSettingsParamsWithTimeout(timeout time.Duration) *DeleteFsaSettingsParams {

	return &DeleteFsaSettingsParams{

		timeout: timeout,
	}
}

// NewDeleteFsaSettingsParamsWithContext creates a new DeleteFsaSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteFsaSettingsParamsWithContext(ctx context.Context) *DeleteFsaSettingsParams {

	return &DeleteFsaSettingsParams{

		Context: ctx,
	}
}

// NewDeleteFsaSettingsParamsWithHTTPClient creates a new DeleteFsaSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteFsaSettingsParamsWithHTTPClient(client *http.Client) *DeleteFsaSettingsParams {

	return &DeleteFsaSettingsParams{
		HTTPClient: client,
	}
}

/*DeleteFsaSettingsParams contains all the parameters to send to the API endpoint
for the delete fsa settings operation typically these are written to a http.Request
*/
type DeleteFsaSettingsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete fsa settings params
func (o *DeleteFsaSettingsParams) WithTimeout(timeout time.Duration) *DeleteFsaSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete fsa settings params
func (o *DeleteFsaSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete fsa settings params
func (o *DeleteFsaSettingsParams) WithContext(ctx context.Context) *DeleteFsaSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete fsa settings params
func (o *DeleteFsaSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete fsa settings params
func (o *DeleteFsaSettingsParams) WithHTTPClient(client *http.Client) *DeleteFsaSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete fsa settings params
func (o *DeleteFsaSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFsaSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
