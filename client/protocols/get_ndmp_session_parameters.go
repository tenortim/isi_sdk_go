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

// NewGetNdmpSessionParams creates a new GetNdmpSessionParams object
// with the default values initialized.
func NewGetNdmpSessionParams() *GetNdmpSessionParams {
	var ()
	return &GetNdmpSessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNdmpSessionParamsWithTimeout creates a new GetNdmpSessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNdmpSessionParamsWithTimeout(timeout time.Duration) *GetNdmpSessionParams {
	var ()
	return &GetNdmpSessionParams{

		timeout: timeout,
	}
}

// NewGetNdmpSessionParamsWithContext creates a new GetNdmpSessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNdmpSessionParamsWithContext(ctx context.Context) *GetNdmpSessionParams {
	var ()
	return &GetNdmpSessionParams{

		Context: ctx,
	}
}

// NewGetNdmpSessionParamsWithHTTPClient creates a new GetNdmpSessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNdmpSessionParamsWithHTTPClient(client *http.Client) *GetNdmpSessionParams {
	var ()
	return &GetNdmpSessionParams{
		HTTPClient: client,
	}
}

/*GetNdmpSessionParams contains all the parameters to send to the API endpoint
for the get ndmp session operation typically these are written to a http.Request
*/
type GetNdmpSessionParams struct {

	/*NdmpSessionID
	  Retrieve the ndmp session.

	*/
	NdmpSessionID string
	/*Lnn
	  Logical node number.

	*/
	Lnn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ndmp session params
func (o *GetNdmpSessionParams) WithTimeout(timeout time.Duration) *GetNdmpSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ndmp session params
func (o *GetNdmpSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ndmp session params
func (o *GetNdmpSessionParams) WithContext(ctx context.Context) *GetNdmpSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ndmp session params
func (o *GetNdmpSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ndmp session params
func (o *GetNdmpSessionParams) WithHTTPClient(client *http.Client) *GetNdmpSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ndmp session params
func (o *GetNdmpSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNdmpSessionID adds the ndmpSessionID to the get ndmp session params
func (o *GetNdmpSessionParams) WithNdmpSessionID(ndmpSessionID string) *GetNdmpSessionParams {
	o.SetNdmpSessionID(ndmpSessionID)
	return o
}

// SetNdmpSessionID adds the ndmpSessionId to the get ndmp session params
func (o *GetNdmpSessionParams) SetNdmpSessionID(ndmpSessionID string) {
	o.NdmpSessionID = ndmpSessionID
}

// WithLnn adds the lnn to the get ndmp session params
func (o *GetNdmpSessionParams) WithLnn(lnn *string) *GetNdmpSessionParams {
	o.SetLnn(lnn)
	return o
}

// SetLnn adds the lnn to the get ndmp session params
func (o *GetNdmpSessionParams) SetLnn(lnn *string) {
	o.Lnn = lnn
}

// WriteToRequest writes these params to a swagger request
func (o *GetNdmpSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param NdmpSessionId
	if err := r.SetPathParam("NdmpSessionId", o.NdmpSessionID); err != nil {
		return err
	}

	if o.Lnn != nil {

		// query param lnn
		var qrLnn string
		if o.Lnn != nil {
			qrLnn = *o.Lnn
		}
		qLnn := qrLnn
		if qLnn != "" {
			if err := r.SetQueryParam("lnn", qLnn); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}