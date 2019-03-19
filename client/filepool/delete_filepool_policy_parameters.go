// Code generated by go-swagger; DO NOT EDIT.

package filepool

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

// NewDeleteFilepoolPolicyParams creates a new DeleteFilepoolPolicyParams object
// with the default values initialized.
func NewDeleteFilepoolPolicyParams() *DeleteFilepoolPolicyParams {
	var ()
	return &DeleteFilepoolPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFilepoolPolicyParamsWithTimeout creates a new DeleteFilepoolPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteFilepoolPolicyParamsWithTimeout(timeout time.Duration) *DeleteFilepoolPolicyParams {
	var ()
	return &DeleteFilepoolPolicyParams{

		timeout: timeout,
	}
}

// NewDeleteFilepoolPolicyParamsWithContext creates a new DeleteFilepoolPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteFilepoolPolicyParamsWithContext(ctx context.Context) *DeleteFilepoolPolicyParams {
	var ()
	return &DeleteFilepoolPolicyParams{

		Context: ctx,
	}
}

// NewDeleteFilepoolPolicyParamsWithHTTPClient creates a new DeleteFilepoolPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteFilepoolPolicyParamsWithHTTPClient(client *http.Client) *DeleteFilepoolPolicyParams {
	var ()
	return &DeleteFilepoolPolicyParams{
		HTTPClient: client,
	}
}

/*DeleteFilepoolPolicyParams contains all the parameters to send to the API endpoint
for the delete filepool policy operation typically these are written to a http.Request
*/
type DeleteFilepoolPolicyParams struct {

	/*FilepoolPolicyID
	  Delete file pool policy.

	*/
	FilepoolPolicyID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete filepool policy params
func (o *DeleteFilepoolPolicyParams) WithTimeout(timeout time.Duration) *DeleteFilepoolPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete filepool policy params
func (o *DeleteFilepoolPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete filepool policy params
func (o *DeleteFilepoolPolicyParams) WithContext(ctx context.Context) *DeleteFilepoolPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete filepool policy params
func (o *DeleteFilepoolPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete filepool policy params
func (o *DeleteFilepoolPolicyParams) WithHTTPClient(client *http.Client) *DeleteFilepoolPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete filepool policy params
func (o *DeleteFilepoolPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilepoolPolicyID adds the filepoolPolicyID to the delete filepool policy params
func (o *DeleteFilepoolPolicyParams) WithFilepoolPolicyID(filepoolPolicyID string) *DeleteFilepoolPolicyParams {
	o.SetFilepoolPolicyID(filepoolPolicyID)
	return o
}

// SetFilepoolPolicyID adds the filepoolPolicyId to the delete filepool policy params
func (o *DeleteFilepoolPolicyParams) SetFilepoolPolicyID(filepoolPolicyID string) {
	o.FilepoolPolicyID = filepoolPolicyID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFilepoolPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param FilepoolPolicyId
	if err := r.SetPathParam("FilepoolPolicyId", o.FilepoolPolicyID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
