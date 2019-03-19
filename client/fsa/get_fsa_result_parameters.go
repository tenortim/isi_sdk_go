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

// NewGetFsaResultParams creates a new GetFsaResultParams object
// with the default values initialized.
func NewGetFsaResultParams() *GetFsaResultParams {
	var ()
	return &GetFsaResultParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFsaResultParamsWithTimeout creates a new GetFsaResultParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFsaResultParamsWithTimeout(timeout time.Duration) *GetFsaResultParams {
	var ()
	return &GetFsaResultParams{

		timeout: timeout,
	}
}

// NewGetFsaResultParamsWithContext creates a new GetFsaResultParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFsaResultParamsWithContext(ctx context.Context) *GetFsaResultParams {
	var ()
	return &GetFsaResultParams{

		Context: ctx,
	}
}

// NewGetFsaResultParamsWithHTTPClient creates a new GetFsaResultParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFsaResultParamsWithHTTPClient(client *http.Client) *GetFsaResultParams {
	var ()
	return &GetFsaResultParams{
		HTTPClient: client,
	}
}

/*GetFsaResultParams contains all the parameters to send to the API endpoint
for the get fsa result operation typically these are written to a http.Request
*/
type GetFsaResultParams struct {

	/*FsaResultID
	  Retrieve result set information.

	*/
	FsaResultID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get fsa result params
func (o *GetFsaResultParams) WithTimeout(timeout time.Duration) *GetFsaResultParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fsa result params
func (o *GetFsaResultParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fsa result params
func (o *GetFsaResultParams) WithContext(ctx context.Context) *GetFsaResultParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fsa result params
func (o *GetFsaResultParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fsa result params
func (o *GetFsaResultParams) WithHTTPClient(client *http.Client) *GetFsaResultParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fsa result params
func (o *GetFsaResultParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFsaResultID adds the fsaResultID to the get fsa result params
func (o *GetFsaResultParams) WithFsaResultID(fsaResultID string) *GetFsaResultParams {
	o.SetFsaResultID(fsaResultID)
	return o
}

// SetFsaResultID adds the fsaResultId to the get fsa result params
func (o *GetFsaResultParams) SetFsaResultID(fsaResultID string) {
	o.FsaResultID = fsaResultID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFsaResultParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param FsaResultId
	if err := r.SetPathParam("FsaResultId", o.FsaResultID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}