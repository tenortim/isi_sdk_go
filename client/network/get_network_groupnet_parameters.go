// Code generated by go-swagger; DO NOT EDIT.

package network

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

// NewGetNetworkGroupnetParams creates a new GetNetworkGroupnetParams object
// with the default values initialized.
func NewGetNetworkGroupnetParams() *GetNetworkGroupnetParams {
	var ()
	return &GetNetworkGroupnetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkGroupnetParamsWithTimeout creates a new GetNetworkGroupnetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkGroupnetParamsWithTimeout(timeout time.Duration) *GetNetworkGroupnetParams {
	var ()
	return &GetNetworkGroupnetParams{

		timeout: timeout,
	}
}

// NewGetNetworkGroupnetParamsWithContext creates a new GetNetworkGroupnetParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkGroupnetParamsWithContext(ctx context.Context) *GetNetworkGroupnetParams {
	var ()
	return &GetNetworkGroupnetParams{

		Context: ctx,
	}
}

// NewGetNetworkGroupnetParamsWithHTTPClient creates a new GetNetworkGroupnetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkGroupnetParamsWithHTTPClient(client *http.Client) *GetNetworkGroupnetParams {
	var ()
	return &GetNetworkGroupnetParams{
		HTTPClient: client,
	}
}

/*GetNetworkGroupnetParams contains all the parameters to send to the API endpoint
for the get network groupnet operation typically these are written to a http.Request
*/
type GetNetworkGroupnetParams struct {

	/*NetworkGroupnetID
	  View a network groupnet.

	*/
	NetworkGroupnetID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network groupnet params
func (o *GetNetworkGroupnetParams) WithTimeout(timeout time.Duration) *GetNetworkGroupnetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network groupnet params
func (o *GetNetworkGroupnetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network groupnet params
func (o *GetNetworkGroupnetParams) WithContext(ctx context.Context) *GetNetworkGroupnetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network groupnet params
func (o *GetNetworkGroupnetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network groupnet params
func (o *GetNetworkGroupnetParams) WithHTTPClient(client *http.Client) *GetNetworkGroupnetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network groupnet params
func (o *GetNetworkGroupnetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkGroupnetID adds the networkGroupnetID to the get network groupnet params
func (o *GetNetworkGroupnetParams) WithNetworkGroupnetID(networkGroupnetID string) *GetNetworkGroupnetParams {
	o.SetNetworkGroupnetID(networkGroupnetID)
	return o
}

// SetNetworkGroupnetID adds the networkGroupnetId to the get network groupnet params
func (o *GetNetworkGroupnetParams) SetNetworkGroupnetID(networkGroupnetID string) {
	o.NetworkGroupnetID = networkGroupnetID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkGroupnetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param NetworkGroupnetId
	if err := r.SetPathParam("NetworkGroupnetId", o.NetworkGroupnetID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
