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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// NewCreateHardeningApplyItemParams creates a new CreateHardeningApplyItemParams object
// with the default values initialized.
func NewCreateHardeningApplyItemParams() *CreateHardeningApplyItemParams {
	var ()
	return &CreateHardeningApplyItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateHardeningApplyItemParamsWithTimeout creates a new CreateHardeningApplyItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateHardeningApplyItemParamsWithTimeout(timeout time.Duration) *CreateHardeningApplyItemParams {
	var ()
	return &CreateHardeningApplyItemParams{

		timeout: timeout,
	}
}

// NewCreateHardeningApplyItemParamsWithContext creates a new CreateHardeningApplyItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateHardeningApplyItemParamsWithContext(ctx context.Context) *CreateHardeningApplyItemParams {
	var ()
	return &CreateHardeningApplyItemParams{

		Context: ctx,
	}
}

// NewCreateHardeningApplyItemParamsWithHTTPClient creates a new CreateHardeningApplyItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateHardeningApplyItemParamsWithHTTPClient(client *http.Client) *CreateHardeningApplyItemParams {
	var ()
	return &CreateHardeningApplyItemParams{
		HTTPClient: client,
	}
}

/*CreateHardeningApplyItemParams contains all the parameters to send to the API endpoint
for the create hardening apply item operation typically these are written to a http.Request
*/
type CreateHardeningApplyItemParams struct {

	/*HardeningApplyItem*/
	HardeningApplyItem *models.HardeningApplyItem

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create hardening apply item params
func (o *CreateHardeningApplyItemParams) WithTimeout(timeout time.Duration) *CreateHardeningApplyItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create hardening apply item params
func (o *CreateHardeningApplyItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create hardening apply item params
func (o *CreateHardeningApplyItemParams) WithContext(ctx context.Context) *CreateHardeningApplyItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create hardening apply item params
func (o *CreateHardeningApplyItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create hardening apply item params
func (o *CreateHardeningApplyItemParams) WithHTTPClient(client *http.Client) *CreateHardeningApplyItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create hardening apply item params
func (o *CreateHardeningApplyItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHardeningApplyItem adds the hardeningApplyItem to the create hardening apply item params
func (o *CreateHardeningApplyItemParams) WithHardeningApplyItem(hardeningApplyItem *models.HardeningApplyItem) *CreateHardeningApplyItemParams {
	o.SetHardeningApplyItem(hardeningApplyItem)
	return o
}

// SetHardeningApplyItem adds the hardeningApplyItem to the create hardening apply item params
func (o *CreateHardeningApplyItemParams) SetHardeningApplyItem(hardeningApplyItem *models.HardeningApplyItem) {
	o.HardeningApplyItem = hardeningApplyItem
}

// WriteToRequest writes these params to a swagger request
func (o *CreateHardeningApplyItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.HardeningApplyItem != nil {
		if err := r.SetBodyParam(o.HardeningApplyItem); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}