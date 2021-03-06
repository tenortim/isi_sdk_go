// Code generated by go-swagger; DO NOT EDIT.

package cloud

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

// NewDeleteSettingsReportingEulaParams creates a new DeleteSettingsReportingEulaParams object
// with the default values initialized.
func NewDeleteSettingsReportingEulaParams() *DeleteSettingsReportingEulaParams {

	return &DeleteSettingsReportingEulaParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSettingsReportingEulaParamsWithTimeout creates a new DeleteSettingsReportingEulaParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSettingsReportingEulaParamsWithTimeout(timeout time.Duration) *DeleteSettingsReportingEulaParams {

	return &DeleteSettingsReportingEulaParams{

		timeout: timeout,
	}
}

// NewDeleteSettingsReportingEulaParamsWithContext creates a new DeleteSettingsReportingEulaParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSettingsReportingEulaParamsWithContext(ctx context.Context) *DeleteSettingsReportingEulaParams {

	return &DeleteSettingsReportingEulaParams{

		Context: ctx,
	}
}

// NewDeleteSettingsReportingEulaParamsWithHTTPClient creates a new DeleteSettingsReportingEulaParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSettingsReportingEulaParamsWithHTTPClient(client *http.Client) *DeleteSettingsReportingEulaParams {

	return &DeleteSettingsReportingEulaParams{
		HTTPClient: client,
	}
}

/*DeleteSettingsReportingEulaParams contains all the parameters to send to the API endpoint
for the delete settings reporting eula operation typically these are written to a http.Request
*/
type DeleteSettingsReportingEulaParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete settings reporting eula params
func (o *DeleteSettingsReportingEulaParams) WithTimeout(timeout time.Duration) *DeleteSettingsReportingEulaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete settings reporting eula params
func (o *DeleteSettingsReportingEulaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete settings reporting eula params
func (o *DeleteSettingsReportingEulaParams) WithContext(ctx context.Context) *DeleteSettingsReportingEulaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete settings reporting eula params
func (o *DeleteSettingsReportingEulaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete settings reporting eula params
func (o *DeleteSettingsReportingEulaParams) WithHTTPClient(client *http.Client) *DeleteSettingsReportingEulaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete settings reporting eula params
func (o *DeleteSettingsReportingEulaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSettingsReportingEulaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
