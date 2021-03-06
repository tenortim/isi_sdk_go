// Code generated by go-swagger; DO NOT EDIT.

package sync

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
)

// NewDeleteSyncPoliciesParams creates a new DeleteSyncPoliciesParams object
// with the default values initialized.
func NewDeleteSyncPoliciesParams() *DeleteSyncPoliciesParams {
	var ()
	return &DeleteSyncPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSyncPoliciesParamsWithTimeout creates a new DeleteSyncPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSyncPoliciesParamsWithTimeout(timeout time.Duration) *DeleteSyncPoliciesParams {
	var ()
	return &DeleteSyncPoliciesParams{

		timeout: timeout,
	}
}

// NewDeleteSyncPoliciesParamsWithContext creates a new DeleteSyncPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSyncPoliciesParamsWithContext(ctx context.Context) *DeleteSyncPoliciesParams {
	var ()
	return &DeleteSyncPoliciesParams{

		Context: ctx,
	}
}

// NewDeleteSyncPoliciesParamsWithHTTPClient creates a new DeleteSyncPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSyncPoliciesParamsWithHTTPClient(client *http.Client) *DeleteSyncPoliciesParams {
	var ()
	return &DeleteSyncPoliciesParams{
		HTTPClient: client,
	}
}

/*DeleteSyncPoliciesParams contains all the parameters to send to the API endpoint
for the delete sync policies operation typically these are written to a http.Request
*/
type DeleteSyncPoliciesParams struct {

	/*Force
	  Ignore any running jobs when preparing to delete a policy.

	*/
	Force *bool
	/*LocalOnly
	  Skip deleting the policy association on the target.

	*/
	LocalOnly *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete sync policies params
func (o *DeleteSyncPoliciesParams) WithTimeout(timeout time.Duration) *DeleteSyncPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete sync policies params
func (o *DeleteSyncPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete sync policies params
func (o *DeleteSyncPoliciesParams) WithContext(ctx context.Context) *DeleteSyncPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete sync policies params
func (o *DeleteSyncPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete sync policies params
func (o *DeleteSyncPoliciesParams) WithHTTPClient(client *http.Client) *DeleteSyncPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete sync policies params
func (o *DeleteSyncPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the delete sync policies params
func (o *DeleteSyncPoliciesParams) WithForce(force *bool) *DeleteSyncPoliciesParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the delete sync policies params
func (o *DeleteSyncPoliciesParams) SetForce(force *bool) {
	o.Force = force
}

// WithLocalOnly adds the localOnly to the delete sync policies params
func (o *DeleteSyncPoliciesParams) WithLocalOnly(localOnly *bool) *DeleteSyncPoliciesParams {
	o.SetLocalOnly(localOnly)
	return o
}

// SetLocalOnly adds the localOnly to the delete sync policies params
func (o *DeleteSyncPoliciesParams) SetLocalOnly(localOnly *bool) {
	o.LocalOnly = localOnly
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSyncPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.LocalOnly != nil {

		// query param local_only
		var qrLocalOnly bool
		if o.LocalOnly != nil {
			qrLocalOnly = *o.LocalOnly
		}
		qLocalOnly := swag.FormatBool(qrLocalOnly)
		if qLocalOnly != "" {
			if err := r.SetQueryParam("local_only", qLocalOnly); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
