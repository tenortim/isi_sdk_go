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

// NewCreateHardeningResolveItemParams creates a new CreateHardeningResolveItemParams object
// with the default values initialized.
func NewCreateHardeningResolveItemParams() *CreateHardeningResolveItemParams {
	var ()
	return &CreateHardeningResolveItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateHardeningResolveItemParamsWithTimeout creates a new CreateHardeningResolveItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateHardeningResolveItemParamsWithTimeout(timeout time.Duration) *CreateHardeningResolveItemParams {
	var ()
	return &CreateHardeningResolveItemParams{

		timeout: timeout,
	}
}

// NewCreateHardeningResolveItemParamsWithContext creates a new CreateHardeningResolveItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateHardeningResolveItemParamsWithContext(ctx context.Context) *CreateHardeningResolveItemParams {
	var ()
	return &CreateHardeningResolveItemParams{

		Context: ctx,
	}
}

// NewCreateHardeningResolveItemParamsWithHTTPClient creates a new CreateHardeningResolveItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateHardeningResolveItemParamsWithHTTPClient(client *http.Client) *CreateHardeningResolveItemParams {
	var ()
	return &CreateHardeningResolveItemParams{
		HTTPClient: client,
	}
}

/*CreateHardeningResolveItemParams contains all the parameters to send to the API endpoint
for the create hardening resolve item operation typically these are written to a http.Request
*/
type CreateHardeningResolveItemParams struct {

	/*HardeningResolveItem*/
	HardeningResolveItem *models.HardeningResolveItem
	/*Accept
	  If true, execution proceeds to resolve all issues. If false, executrion aborts. This is a required argument.

	*/
	Accept *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create hardening resolve item params
func (o *CreateHardeningResolveItemParams) WithTimeout(timeout time.Duration) *CreateHardeningResolveItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create hardening resolve item params
func (o *CreateHardeningResolveItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create hardening resolve item params
func (o *CreateHardeningResolveItemParams) WithContext(ctx context.Context) *CreateHardeningResolveItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create hardening resolve item params
func (o *CreateHardeningResolveItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create hardening resolve item params
func (o *CreateHardeningResolveItemParams) WithHTTPClient(client *http.Client) *CreateHardeningResolveItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create hardening resolve item params
func (o *CreateHardeningResolveItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHardeningResolveItem adds the hardeningResolveItem to the create hardening resolve item params
func (o *CreateHardeningResolveItemParams) WithHardeningResolveItem(hardeningResolveItem *models.HardeningResolveItem) *CreateHardeningResolveItemParams {
	o.SetHardeningResolveItem(hardeningResolveItem)
	return o
}

// SetHardeningResolveItem adds the hardeningResolveItem to the create hardening resolve item params
func (o *CreateHardeningResolveItemParams) SetHardeningResolveItem(hardeningResolveItem *models.HardeningResolveItem) {
	o.HardeningResolveItem = hardeningResolveItem
}

// WithAccept adds the accept to the create hardening resolve item params
func (o *CreateHardeningResolveItemParams) WithAccept(accept *bool) *CreateHardeningResolveItemParams {
	o.SetAccept(accept)
	return o
}

// SetAccept adds the accept to the create hardening resolve item params
func (o *CreateHardeningResolveItemParams) SetAccept(accept *bool) {
	o.Accept = accept
}

// WriteToRequest writes these params to a swagger request
func (o *CreateHardeningResolveItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.HardeningResolveItem != nil {
		if err := r.SetBodyParam(o.HardeningResolveItem); err != nil {
			return err
		}
	}

	if o.Accept != nil {

		// query param accept
		var qrAccept bool
		if o.Accept != nil {
			qrAccept = *o.Accept
		}
		qAccept := swag.FormatBool(qrAccept)
		if qAccept != "" {
			if err := r.SetQueryParam("accept", qAccept); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}