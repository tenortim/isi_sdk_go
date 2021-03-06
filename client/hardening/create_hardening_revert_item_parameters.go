// Code generated by go-swagger; DO NOT EDIT.

package hardening

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// NewCreateHardeningRevertItemParams creates a new CreateHardeningRevertItemParams object
// with the default values initialized.
func NewCreateHardeningRevertItemParams() *CreateHardeningRevertItemParams {
	var ()
	return &CreateHardeningRevertItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateHardeningRevertItemParamsWithTimeout creates a new CreateHardeningRevertItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateHardeningRevertItemParamsWithTimeout(timeout time.Duration) *CreateHardeningRevertItemParams {
	var ()
	return &CreateHardeningRevertItemParams{

		timeout: timeout,
	}
}

// NewCreateHardeningRevertItemParamsWithContext creates a new CreateHardeningRevertItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateHardeningRevertItemParamsWithContext(ctx context.Context) *CreateHardeningRevertItemParams {
	var ()
	return &CreateHardeningRevertItemParams{

		Context: ctx,
	}
}

// NewCreateHardeningRevertItemParamsWithHTTPClient creates a new CreateHardeningRevertItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateHardeningRevertItemParamsWithHTTPClient(client *http.Client) *CreateHardeningRevertItemParams {
	var ()
	return &CreateHardeningRevertItemParams{
		HTTPClient: client,
	}
}

/*CreateHardeningRevertItemParams contains all the parameters to send to the API endpoint
for the create hardening revert item operation typically these are written to a http.Request
*/
type CreateHardeningRevertItemParams struct {

	/*HardeningRevertItem*/
	HardeningRevertItem models.Empty
	/*Force
	  If specified, revert operation continues even in case of a failure. Default is false in which case revert stops at the first failure.

	*/
	Force *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create hardening revert item params
func (o *CreateHardeningRevertItemParams) WithTimeout(timeout time.Duration) *CreateHardeningRevertItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create hardening revert item params
func (o *CreateHardeningRevertItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create hardening revert item params
func (o *CreateHardeningRevertItemParams) WithContext(ctx context.Context) *CreateHardeningRevertItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create hardening revert item params
func (o *CreateHardeningRevertItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create hardening revert item params
func (o *CreateHardeningRevertItemParams) WithHTTPClient(client *http.Client) *CreateHardeningRevertItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create hardening revert item params
func (o *CreateHardeningRevertItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHardeningRevertItem adds the hardeningRevertItem to the create hardening revert item params
func (o *CreateHardeningRevertItemParams) WithHardeningRevertItem(hardeningRevertItem models.Empty) *CreateHardeningRevertItemParams {
	o.SetHardeningRevertItem(hardeningRevertItem)
	return o
}

// SetHardeningRevertItem adds the hardeningRevertItem to the create hardening revert item params
func (o *CreateHardeningRevertItemParams) SetHardeningRevertItem(hardeningRevertItem models.Empty) {
	o.HardeningRevertItem = hardeningRevertItem
}

// WithForce adds the force to the create hardening revert item params
func (o *CreateHardeningRevertItemParams) WithForce(force *bool) *CreateHardeningRevertItemParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the create hardening revert item params
func (o *CreateHardeningRevertItemParams) SetForce(force *bool) {
	o.Force = force
}

// WriteToRequest writes these params to a swagger request
func (o *CreateHardeningRevertItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.HardeningRevertItem != nil {
		if err := r.SetBodyParam(o.HardeningRevertItem); err != nil {
			return err
		}
	}

	if o.Force != nil {

		// query param force
		var qrForce bool
		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {
			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
