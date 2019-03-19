// Code generated by go-swagger; DO NOT EDIT.

package upgrade

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

// NewDeleteClusterPatchPatchParams creates a new DeleteClusterPatchPatchParams object
// with the default values initialized.
func NewDeleteClusterPatchPatchParams() *DeleteClusterPatchPatchParams {
	var ()
	return &DeleteClusterPatchPatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClusterPatchPatchParamsWithTimeout creates a new DeleteClusterPatchPatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteClusterPatchPatchParamsWithTimeout(timeout time.Duration) *DeleteClusterPatchPatchParams {
	var ()
	return &DeleteClusterPatchPatchParams{

		timeout: timeout,
	}
}

// NewDeleteClusterPatchPatchParamsWithContext creates a new DeleteClusterPatchPatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteClusterPatchPatchParamsWithContext(ctx context.Context) *DeleteClusterPatchPatchParams {
	var ()
	return &DeleteClusterPatchPatchParams{

		Context: ctx,
	}
}

// NewDeleteClusterPatchPatchParamsWithHTTPClient creates a new DeleteClusterPatchPatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteClusterPatchPatchParamsWithHTTPClient(client *http.Client) *DeleteClusterPatchPatchParams {
	var ()
	return &DeleteClusterPatchPatchParams{
		HTTPClient: client,
	}
}

/*DeleteClusterPatchPatchParams contains all the parameters to send to the API endpoint
for the delete cluster patch patch operation typically these are written to a http.Request
*/
type DeleteClusterPatchPatchParams struct {

	/*ClusterPatchPatchID
	  Uninstall a patch.

	*/
	ClusterPatchPatchID string
	/*Override
	  Whether to ignore patch system validation and force the uninstallation.

	*/
	Override *bool
	/*Rolling
	  Whether to uninstall the patch on one node at a time. Defaults to true.

	*/
	Rolling *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete cluster patch patch params
func (o *DeleteClusterPatchPatchParams) WithTimeout(timeout time.Duration) *DeleteClusterPatchPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cluster patch patch params
func (o *DeleteClusterPatchPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cluster patch patch params
func (o *DeleteClusterPatchPatchParams) WithContext(ctx context.Context) *DeleteClusterPatchPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cluster patch patch params
func (o *DeleteClusterPatchPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cluster patch patch params
func (o *DeleteClusterPatchPatchParams) WithHTTPClient(client *http.Client) *DeleteClusterPatchPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cluster patch patch params
func (o *DeleteClusterPatchPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterPatchPatchID adds the clusterPatchPatchID to the delete cluster patch patch params
func (o *DeleteClusterPatchPatchParams) WithClusterPatchPatchID(clusterPatchPatchID string) *DeleteClusterPatchPatchParams {
	o.SetClusterPatchPatchID(clusterPatchPatchID)
	return o
}

// SetClusterPatchPatchID adds the clusterPatchPatchId to the delete cluster patch patch params
func (o *DeleteClusterPatchPatchParams) SetClusterPatchPatchID(clusterPatchPatchID string) {
	o.ClusterPatchPatchID = clusterPatchPatchID
}

// WithOverride adds the override to the delete cluster patch patch params
func (o *DeleteClusterPatchPatchParams) WithOverride(override *bool) *DeleteClusterPatchPatchParams {
	o.SetOverride(override)
	return o
}

// SetOverride adds the override to the delete cluster patch patch params
func (o *DeleteClusterPatchPatchParams) SetOverride(override *bool) {
	o.Override = override
}

// WithRolling adds the rolling to the delete cluster patch patch params
func (o *DeleteClusterPatchPatchParams) WithRolling(rolling *bool) *DeleteClusterPatchPatchParams {
	o.SetRolling(rolling)
	return o
}

// SetRolling adds the rolling to the delete cluster patch patch params
func (o *DeleteClusterPatchPatchParams) SetRolling(rolling *bool) {
	o.Rolling = rolling
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClusterPatchPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ClusterPatchPatchId
	if err := r.SetPathParam("ClusterPatchPatchId", o.ClusterPatchPatchID); err != nil {
		return err
	}

	if o.Override != nil {

		// query param override
		var qrOverride bool
		if o.Override != nil {
			qrOverride = *o.Override
		}
		qOverride := swag.FormatBool(qrOverride)
		if qOverride != "" {
			if err := r.SetQueryParam("override", qOverride); err != nil {
				return err
			}
		}

	}

	if o.Rolling != nil {

		// query param rolling
		var qrRolling bool
		if o.Rolling != nil {
			qrRolling = *o.Rolling
		}
		qRolling := swag.FormatBool(qrRolling)
		if qRolling != "" {
			if err := r.SetQueryParam("rolling", qRolling); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}