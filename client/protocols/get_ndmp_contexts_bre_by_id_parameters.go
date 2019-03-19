// Code generated by go-swagger; DO NOT EDIT.

package protocols

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

// NewGetNdmpContextsBreByIDParams creates a new GetNdmpContextsBreByIDParams object
// with the default values initialized.
func NewGetNdmpContextsBreByIDParams() *GetNdmpContextsBreByIDParams {
	var ()
	return &GetNdmpContextsBreByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNdmpContextsBreByIDParamsWithTimeout creates a new GetNdmpContextsBreByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNdmpContextsBreByIDParamsWithTimeout(timeout time.Duration) *GetNdmpContextsBreByIDParams {
	var ()
	return &GetNdmpContextsBreByIDParams{

		timeout: timeout,
	}
}

// NewGetNdmpContextsBreByIDParamsWithContext creates a new GetNdmpContextsBreByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNdmpContextsBreByIDParamsWithContext(ctx context.Context) *GetNdmpContextsBreByIDParams {
	var ()
	return &GetNdmpContextsBreByIDParams{

		Context: ctx,
	}
}

// NewGetNdmpContextsBreByIDParamsWithHTTPClient creates a new GetNdmpContextsBreByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNdmpContextsBreByIDParamsWithHTTPClient(client *http.Client) *GetNdmpContextsBreByIDParams {
	var ()
	return &GetNdmpContextsBreByIDParams{
		HTTPClient: client,
	}
}

/*GetNdmpContextsBreByIDParams contains all the parameters to send to the API endpoint
for the get ndmp contexts bre by Id operation typically these are written to a http.Request
*/
type GetNdmpContextsBreByIDParams struct {

	/*NdmpContextsBreID
	  View a NDMP BRE context

	*/
	NdmpContextsBreID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ndmp contexts bre by Id params
func (o *GetNdmpContextsBreByIDParams) WithTimeout(timeout time.Duration) *GetNdmpContextsBreByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ndmp contexts bre by Id params
func (o *GetNdmpContextsBreByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ndmp contexts bre by Id params
func (o *GetNdmpContextsBreByIDParams) WithContext(ctx context.Context) *GetNdmpContextsBreByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ndmp contexts bre by Id params
func (o *GetNdmpContextsBreByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ndmp contexts bre by Id params
func (o *GetNdmpContextsBreByIDParams) WithHTTPClient(client *http.Client) *GetNdmpContextsBreByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ndmp contexts bre by Id params
func (o *GetNdmpContextsBreByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNdmpContextsBreID adds the ndmpContextsBreID to the get ndmp contexts bre by Id params
func (o *GetNdmpContextsBreByIDParams) WithNdmpContextsBreID(ndmpContextsBreID string) *GetNdmpContextsBreByIDParams {
	o.SetNdmpContextsBreID(ndmpContextsBreID)
	return o
}

// SetNdmpContextsBreID adds the ndmpContextsBreId to the get ndmp contexts bre by Id params
func (o *GetNdmpContextsBreByIDParams) SetNdmpContextsBreID(ndmpContextsBreID string) {
	o.NdmpContextsBreID = ndmpContextsBreID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNdmpContextsBreByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param NdmpContextsBreId
	if err := r.SetPathParam("NdmpContextsBreId", o.NdmpContextsBreID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
