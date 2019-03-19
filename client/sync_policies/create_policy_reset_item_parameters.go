// Code generated by go-swagger; DO NOT EDIT.

package sync_policies

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

// NewCreatePolicyResetItemParams creates a new CreatePolicyResetItemParams object
// with the default values initialized.
func NewCreatePolicyResetItemParams() *CreatePolicyResetItemParams {
	var ()
	return &CreatePolicyResetItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePolicyResetItemParamsWithTimeout creates a new CreatePolicyResetItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreatePolicyResetItemParamsWithTimeout(timeout time.Duration) *CreatePolicyResetItemParams {
	var ()
	return &CreatePolicyResetItemParams{

		timeout: timeout,
	}
}

// NewCreatePolicyResetItemParamsWithContext creates a new CreatePolicyResetItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreatePolicyResetItemParamsWithContext(ctx context.Context) *CreatePolicyResetItemParams {
	var ()
	return &CreatePolicyResetItemParams{

		Context: ctx,
	}
}

// NewCreatePolicyResetItemParamsWithHTTPClient creates a new CreatePolicyResetItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreatePolicyResetItemParamsWithHTTPClient(client *http.Client) *CreatePolicyResetItemParams {
	var ()
	return &CreatePolicyResetItemParams{
		HTTPClient: client,
	}
}

/*CreatePolicyResetItemParams contains all the parameters to send to the API endpoint
for the create policy reset item operation typically these are written to a http.Request
*/
type CreatePolicyResetItemParams struct {

	/*Policy*/
	Policy string
	/*PolicyResetItem*/
	PolicyResetItem models.Empty

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create policy reset item params
func (o *CreatePolicyResetItemParams) WithTimeout(timeout time.Duration) *CreatePolicyResetItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create policy reset item params
func (o *CreatePolicyResetItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create policy reset item params
func (o *CreatePolicyResetItemParams) WithContext(ctx context.Context) *CreatePolicyResetItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create policy reset item params
func (o *CreatePolicyResetItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create policy reset item params
func (o *CreatePolicyResetItemParams) WithHTTPClient(client *http.Client) *CreatePolicyResetItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create policy reset item params
func (o *CreatePolicyResetItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPolicy adds the policy to the create policy reset item params
func (o *CreatePolicyResetItemParams) WithPolicy(policy string) *CreatePolicyResetItemParams {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the create policy reset item params
func (o *CreatePolicyResetItemParams) SetPolicy(policy string) {
	o.Policy = policy
}

// WithPolicyResetItem adds the policyResetItem to the create policy reset item params
func (o *CreatePolicyResetItemParams) WithPolicyResetItem(policyResetItem models.Empty) *CreatePolicyResetItemParams {
	o.SetPolicyResetItem(policyResetItem)
	return o
}

// SetPolicyResetItem adds the policyResetItem to the create policy reset item params
func (o *CreatePolicyResetItemParams) SetPolicyResetItem(policyResetItem models.Empty) {
	o.PolicyResetItem = policyResetItem
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePolicyResetItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Policy
	if err := r.SetPathParam("Policy", o.Policy); err != nil {
		return err
	}

	if o.PolicyResetItem != nil {
		if err := r.SetBodyParam(o.PolicyResetItem); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
