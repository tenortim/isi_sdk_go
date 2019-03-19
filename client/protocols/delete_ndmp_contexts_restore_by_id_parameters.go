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

// NewDeleteNdmpContextsRestoreByIDParams creates a new DeleteNdmpContextsRestoreByIDParams object
// with the default values initialized.
func NewDeleteNdmpContextsRestoreByIDParams() *DeleteNdmpContextsRestoreByIDParams {
	var ()
	return &DeleteNdmpContextsRestoreByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNdmpContextsRestoreByIDParamsWithTimeout creates a new DeleteNdmpContextsRestoreByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteNdmpContextsRestoreByIDParamsWithTimeout(timeout time.Duration) *DeleteNdmpContextsRestoreByIDParams {
	var ()
	return &DeleteNdmpContextsRestoreByIDParams{

		timeout: timeout,
	}
}

// NewDeleteNdmpContextsRestoreByIDParamsWithContext creates a new DeleteNdmpContextsRestoreByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteNdmpContextsRestoreByIDParamsWithContext(ctx context.Context) *DeleteNdmpContextsRestoreByIDParams {
	var ()
	return &DeleteNdmpContextsRestoreByIDParams{

		Context: ctx,
	}
}

// NewDeleteNdmpContextsRestoreByIDParamsWithHTTPClient creates a new DeleteNdmpContextsRestoreByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteNdmpContextsRestoreByIDParamsWithHTTPClient(client *http.Client) *DeleteNdmpContextsRestoreByIDParams {
	var ()
	return &DeleteNdmpContextsRestoreByIDParams{
		HTTPClient: client,
	}
}

/*DeleteNdmpContextsRestoreByIDParams contains all the parameters to send to the API endpoint
for the delete ndmp contexts restore by Id operation typically these are written to a http.Request
*/
type DeleteNdmpContextsRestoreByIDParams struct {

	/*NdmpContextsRestoreID
	  Delete a restore context

	*/
	NdmpContextsRestoreID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete ndmp contexts restore by Id params
func (o *DeleteNdmpContextsRestoreByIDParams) WithTimeout(timeout time.Duration) *DeleteNdmpContextsRestoreByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete ndmp contexts restore by Id params
func (o *DeleteNdmpContextsRestoreByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete ndmp contexts restore by Id params
func (o *DeleteNdmpContextsRestoreByIDParams) WithContext(ctx context.Context) *DeleteNdmpContextsRestoreByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete ndmp contexts restore by Id params
func (o *DeleteNdmpContextsRestoreByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete ndmp contexts restore by Id params
func (o *DeleteNdmpContextsRestoreByIDParams) WithHTTPClient(client *http.Client) *DeleteNdmpContextsRestoreByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete ndmp contexts restore by Id params
func (o *DeleteNdmpContextsRestoreByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNdmpContextsRestoreID adds the ndmpContextsRestoreID to the delete ndmp contexts restore by Id params
func (o *DeleteNdmpContextsRestoreByIDParams) WithNdmpContextsRestoreID(ndmpContextsRestoreID string) *DeleteNdmpContextsRestoreByIDParams {
	o.SetNdmpContextsRestoreID(ndmpContextsRestoreID)
	return o
}

// SetNdmpContextsRestoreID adds the ndmpContextsRestoreId to the delete ndmp contexts restore by Id params
func (o *DeleteNdmpContextsRestoreByIDParams) SetNdmpContextsRestoreID(ndmpContextsRestoreID string) {
	o.NdmpContextsRestoreID = ndmpContextsRestoreID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNdmpContextsRestoreByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param NdmpContextsRestoreId
	if err := r.SetPathParam("NdmpContextsRestoreId", o.NdmpContextsRestoreID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}