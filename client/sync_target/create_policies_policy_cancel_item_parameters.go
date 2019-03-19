// Code generated by go-swagger; DO NOT EDIT.

package sync_target

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

// NewCreatePoliciesPolicyCancelItemParams creates a new CreatePoliciesPolicyCancelItemParams object
// with the default values initialized.
func NewCreatePoliciesPolicyCancelItemParams() *CreatePoliciesPolicyCancelItemParams {
	var ()
	return &CreatePoliciesPolicyCancelItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePoliciesPolicyCancelItemParamsWithTimeout creates a new CreatePoliciesPolicyCancelItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreatePoliciesPolicyCancelItemParamsWithTimeout(timeout time.Duration) *CreatePoliciesPolicyCancelItemParams {
	var ()
	return &CreatePoliciesPolicyCancelItemParams{

		timeout: timeout,
	}
}

// NewCreatePoliciesPolicyCancelItemParamsWithContext creates a new CreatePoliciesPolicyCancelItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreatePoliciesPolicyCancelItemParamsWithContext(ctx context.Context) *CreatePoliciesPolicyCancelItemParams {
	var ()
	return &CreatePoliciesPolicyCancelItemParams{

		Context: ctx,
	}
}

// NewCreatePoliciesPolicyCancelItemParamsWithHTTPClient creates a new CreatePoliciesPolicyCancelItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreatePoliciesPolicyCancelItemParamsWithHTTPClient(client *http.Client) *CreatePoliciesPolicyCancelItemParams {
	var ()
	return &CreatePoliciesPolicyCancelItemParams{
		HTTPClient: client,
	}
}

/*CreatePoliciesPolicyCancelItemParams contains all the parameters to send to the API endpoint
for the create policies policy cancel item operation typically these are written to a http.Request
*/
type CreatePoliciesPolicyCancelItemParams struct {

	/*PoliciesPolicyCancelItem*/
	PoliciesPolicyCancelItem models.Empty
	/*Policy*/
	Policy string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create policies policy cancel item params
func (o *CreatePoliciesPolicyCancelItemParams) WithTimeout(timeout time.Duration) *CreatePoliciesPolicyCancelItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create policies policy cancel item params
func (o *CreatePoliciesPolicyCancelItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create policies policy cancel item params
func (o *CreatePoliciesPolicyCancelItemParams) WithContext(ctx context.Context) *CreatePoliciesPolicyCancelItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create policies policy cancel item params
func (o *CreatePoliciesPolicyCancelItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create policies policy cancel item params
func (o *CreatePoliciesPolicyCancelItemParams) WithHTTPClient(client *http.Client) *CreatePoliciesPolicyCancelItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create policies policy cancel item params
func (o *CreatePoliciesPolicyCancelItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPoliciesPolicyCancelItem adds the policiesPolicyCancelItem to the create policies policy cancel item params
func (o *CreatePoliciesPolicyCancelItemParams) WithPoliciesPolicyCancelItem(policiesPolicyCancelItem models.Empty) *CreatePoliciesPolicyCancelItemParams {
	o.SetPoliciesPolicyCancelItem(policiesPolicyCancelItem)
	return o
}

// SetPoliciesPolicyCancelItem adds the policiesPolicyCancelItem to the create policies policy cancel item params
func (o *CreatePoliciesPolicyCancelItemParams) SetPoliciesPolicyCancelItem(policiesPolicyCancelItem models.Empty) {
	o.PoliciesPolicyCancelItem = policiesPolicyCancelItem
}

// WithPolicy adds the policy to the create policies policy cancel item params
func (o *CreatePoliciesPolicyCancelItemParams) WithPolicy(policy string) *CreatePoliciesPolicyCancelItemParams {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the create policies policy cancel item params
func (o *CreatePoliciesPolicyCancelItemParams) SetPolicy(policy string) {
	o.Policy = policy
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePoliciesPolicyCancelItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PoliciesPolicyCancelItem != nil {
		if err := r.SetBodyParam(o.PoliciesPolicyCancelItem); err != nil {
			return err
		}
	}

	// path param Policy
	if err := r.SetPathParam("Policy", o.Policy); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}