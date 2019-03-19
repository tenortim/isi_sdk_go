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

	models "github.com/tenortim/isi_sdk_go/models"
)

// NewUpdateFilepoolPolicyParams creates a new UpdateFilepoolPolicyParams object
// with the default values initialized.
func NewUpdateFilepoolPolicyParams() *UpdateFilepoolPolicyParams {
	var ()
	return &UpdateFilepoolPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateFilepoolPolicyParamsWithTimeout creates a new UpdateFilepoolPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateFilepoolPolicyParamsWithTimeout(timeout time.Duration) *UpdateFilepoolPolicyParams {
	var ()
	return &UpdateFilepoolPolicyParams{

		timeout: timeout,
	}
}

// NewUpdateFilepoolPolicyParamsWithContext creates a new UpdateFilepoolPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateFilepoolPolicyParamsWithContext(ctx context.Context) *UpdateFilepoolPolicyParams {
	var ()
	return &UpdateFilepoolPolicyParams{

		Context: ctx,
	}
}

// NewUpdateFilepoolPolicyParamsWithHTTPClient creates a new UpdateFilepoolPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateFilepoolPolicyParamsWithHTTPClient(client *http.Client) *UpdateFilepoolPolicyParams {
	var ()
	return &UpdateFilepoolPolicyParams{
		HTTPClient: client,
	}
}

/*UpdateFilepoolPolicyParams contains all the parameters to send to the API endpoint
for the update filepool policy operation typically these are written to a http.Request
*/
type UpdateFilepoolPolicyParams struct {

	/*FilepoolPolicy*/
	FilepoolPolicy *models.FilepoolPolicy
	/*FilepoolPolicyID
	  Modify file pool policy. All input fields are optional, but one or more must be supplied.

	*/
	FilepoolPolicyID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update filepool policy params
func (o *UpdateFilepoolPolicyParams) WithTimeout(timeout time.Duration) *UpdateFilepoolPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update filepool policy params
func (o *UpdateFilepoolPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update filepool policy params
func (o *UpdateFilepoolPolicyParams) WithContext(ctx context.Context) *UpdateFilepoolPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update filepool policy params
func (o *UpdateFilepoolPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update filepool policy params
func (o *UpdateFilepoolPolicyParams) WithHTTPClient(client *http.Client) *UpdateFilepoolPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update filepool policy params
func (o *UpdateFilepoolPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilepoolPolicy adds the filepoolPolicy to the update filepool policy params
func (o *UpdateFilepoolPolicyParams) WithFilepoolPolicy(filepoolPolicy *models.FilepoolPolicy) *UpdateFilepoolPolicyParams {
	o.SetFilepoolPolicy(filepoolPolicy)
	return o
}

// SetFilepoolPolicy adds the filepoolPolicy to the update filepool policy params
func (o *UpdateFilepoolPolicyParams) SetFilepoolPolicy(filepoolPolicy *models.FilepoolPolicy) {
	o.FilepoolPolicy = filepoolPolicy
}

// WithFilepoolPolicyID adds the filepoolPolicyID to the update filepool policy params
func (o *UpdateFilepoolPolicyParams) WithFilepoolPolicyID(filepoolPolicyID string) *UpdateFilepoolPolicyParams {
	o.SetFilepoolPolicyID(filepoolPolicyID)
	return o
}

// SetFilepoolPolicyID adds the filepoolPolicyId to the update filepool policy params
func (o *UpdateFilepoolPolicyParams) SetFilepoolPolicyID(filepoolPolicyID string) {
	o.FilepoolPolicyID = filepoolPolicyID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateFilepoolPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FilepoolPolicy != nil {
		if err := r.SetBodyParam(o.FilepoolPolicy); err != nil {
			return err
		}
	}

	// path param FilepoolPolicyId
	if err := r.SetPathParam("FilepoolPolicyId", o.FilepoolPolicyID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
