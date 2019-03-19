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

// NewGetNdmpContextsBackupByIDParams creates a new GetNdmpContextsBackupByIDParams object
// with the default values initialized.
func NewGetNdmpContextsBackupByIDParams() *GetNdmpContextsBackupByIDParams {
	var ()
	return &GetNdmpContextsBackupByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNdmpContextsBackupByIDParamsWithTimeout creates a new GetNdmpContextsBackupByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNdmpContextsBackupByIDParamsWithTimeout(timeout time.Duration) *GetNdmpContextsBackupByIDParams {
	var ()
	return &GetNdmpContextsBackupByIDParams{

		timeout: timeout,
	}
}

// NewGetNdmpContextsBackupByIDParamsWithContext creates a new GetNdmpContextsBackupByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNdmpContextsBackupByIDParamsWithContext(ctx context.Context) *GetNdmpContextsBackupByIDParams {
	var ()
	return &GetNdmpContextsBackupByIDParams{

		Context: ctx,
	}
}

// NewGetNdmpContextsBackupByIDParamsWithHTTPClient creates a new GetNdmpContextsBackupByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNdmpContextsBackupByIDParamsWithHTTPClient(client *http.Client) *GetNdmpContextsBackupByIDParams {
	var ()
	return &GetNdmpContextsBackupByIDParams{
		HTTPClient: client,
	}
}

/*GetNdmpContextsBackupByIDParams contains all the parameters to send to the API endpoint
for the get ndmp contexts backup by Id operation typically these are written to a http.Request
*/
type GetNdmpContextsBackupByIDParams struct {

	/*NdmpContextsBackupID
	  View a NDMP backup context

	*/
	NdmpContextsBackupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ndmp contexts backup by Id params
func (o *GetNdmpContextsBackupByIDParams) WithTimeout(timeout time.Duration) *GetNdmpContextsBackupByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ndmp contexts backup by Id params
func (o *GetNdmpContextsBackupByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ndmp contexts backup by Id params
func (o *GetNdmpContextsBackupByIDParams) WithContext(ctx context.Context) *GetNdmpContextsBackupByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ndmp contexts backup by Id params
func (o *GetNdmpContextsBackupByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ndmp contexts backup by Id params
func (o *GetNdmpContextsBackupByIDParams) WithHTTPClient(client *http.Client) *GetNdmpContextsBackupByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ndmp contexts backup by Id params
func (o *GetNdmpContextsBackupByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNdmpContextsBackupID adds the ndmpContextsBackupID to the get ndmp contexts backup by Id params
func (o *GetNdmpContextsBackupByIDParams) WithNdmpContextsBackupID(ndmpContextsBackupID string) *GetNdmpContextsBackupByIDParams {
	o.SetNdmpContextsBackupID(ndmpContextsBackupID)
	return o
}

// SetNdmpContextsBackupID adds the ndmpContextsBackupId to the get ndmp contexts backup by Id params
func (o *GetNdmpContextsBackupByIDParams) SetNdmpContextsBackupID(ndmpContextsBackupID string) {
	o.NdmpContextsBackupID = ndmpContextsBackupID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNdmpContextsBackupByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param NdmpContextsBackupId
	if err := r.SetPathParam("NdmpContextsBackupId", o.NdmpContextsBackupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}