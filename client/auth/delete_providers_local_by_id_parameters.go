// Code generated by go-swagger; DO NOT EDIT.

package auth

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

// NewDeleteProvidersLocalByIDParams creates a new DeleteProvidersLocalByIDParams object
// with the default values initialized.
func NewDeleteProvidersLocalByIDParams() *DeleteProvidersLocalByIDParams {
	var ()
	return &DeleteProvidersLocalByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProvidersLocalByIDParamsWithTimeout creates a new DeleteProvidersLocalByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteProvidersLocalByIDParamsWithTimeout(timeout time.Duration) *DeleteProvidersLocalByIDParams {
	var ()
	return &DeleteProvidersLocalByIDParams{

		timeout: timeout,
	}
}

// NewDeleteProvidersLocalByIDParamsWithContext creates a new DeleteProvidersLocalByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteProvidersLocalByIDParamsWithContext(ctx context.Context) *DeleteProvidersLocalByIDParams {
	var ()
	return &DeleteProvidersLocalByIDParams{

		Context: ctx,
	}
}

// NewDeleteProvidersLocalByIDParamsWithHTTPClient creates a new DeleteProvidersLocalByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteProvidersLocalByIDParamsWithHTTPClient(client *http.Client) *DeleteProvidersLocalByIDParams {
	var ()
	return &DeleteProvidersLocalByIDParams{
		HTTPClient: client,
	}
}

/*DeleteProvidersLocalByIDParams contains all the parameters to send to the API endpoint
for the delete providers local by Id operation typically these are written to a http.Request
*/
type DeleteProvidersLocalByIDParams struct {

	/*ProvidersLocalID
	  Delete the local provider.

	*/
	ProvidersLocalID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete providers local by Id params
func (o *DeleteProvidersLocalByIDParams) WithTimeout(timeout time.Duration) *DeleteProvidersLocalByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete providers local by Id params
func (o *DeleteProvidersLocalByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete providers local by Id params
func (o *DeleteProvidersLocalByIDParams) WithContext(ctx context.Context) *DeleteProvidersLocalByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete providers local by Id params
func (o *DeleteProvidersLocalByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete providers local by Id params
func (o *DeleteProvidersLocalByIDParams) WithHTTPClient(client *http.Client) *DeleteProvidersLocalByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete providers local by Id params
func (o *DeleteProvidersLocalByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProvidersLocalID adds the providersLocalID to the delete providers local by Id params
func (o *DeleteProvidersLocalByIDParams) WithProvidersLocalID(providersLocalID string) *DeleteProvidersLocalByIDParams {
	o.SetProvidersLocalID(providersLocalID)
	return o
}

// SetProvidersLocalID adds the providersLocalId to the delete providers local by Id params
func (o *DeleteProvidersLocalByIDParams) SetProvidersLocalID(providersLocalID string) {
	o.ProvidersLocalID = providersLocalID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProvidersLocalByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ProvidersLocalId
	if err := r.SetPathParam("ProvidersLocalId", o.ProvidersLocalID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
