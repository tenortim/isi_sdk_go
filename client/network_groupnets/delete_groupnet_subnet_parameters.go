// Code generated by go-swagger; DO NOT EDIT.

package network_groupnets

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

// NewDeleteGroupnetSubnetParams creates a new DeleteGroupnetSubnetParams object
// with the default values initialized.
func NewDeleteGroupnetSubnetParams() *DeleteGroupnetSubnetParams {
	var ()
	return &DeleteGroupnetSubnetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteGroupnetSubnetParamsWithTimeout creates a new DeleteGroupnetSubnetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteGroupnetSubnetParamsWithTimeout(timeout time.Duration) *DeleteGroupnetSubnetParams {
	var ()
	return &DeleteGroupnetSubnetParams{

		timeout: timeout,
	}
}

// NewDeleteGroupnetSubnetParamsWithContext creates a new DeleteGroupnetSubnetParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteGroupnetSubnetParamsWithContext(ctx context.Context) *DeleteGroupnetSubnetParams {
	var ()
	return &DeleteGroupnetSubnetParams{

		Context: ctx,
	}
}

// NewDeleteGroupnetSubnetParamsWithHTTPClient creates a new DeleteGroupnetSubnetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteGroupnetSubnetParamsWithHTTPClient(client *http.Client) *DeleteGroupnetSubnetParams {
	var ()
	return &DeleteGroupnetSubnetParams{
		HTTPClient: client,
	}
}

/*DeleteGroupnetSubnetParams contains all the parameters to send to the API endpoint
for the delete groupnet subnet operation typically these are written to a http.Request
*/
type DeleteGroupnetSubnetParams struct {

	/*Groupnet*/
	Groupnet string
	/*GroupnetSubnetID
	  Delete a network subnet..

	*/
	GroupnetSubnetID string
	/*Force
	  force deleting this subnet even if pools in other subnets rely on this subnet's SC VIP.

	*/
	Force *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete groupnet subnet params
func (o *DeleteGroupnetSubnetParams) WithTimeout(timeout time.Duration) *DeleteGroupnetSubnetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete groupnet subnet params
func (o *DeleteGroupnetSubnetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete groupnet subnet params
func (o *DeleteGroupnetSubnetParams) WithContext(ctx context.Context) *DeleteGroupnetSubnetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete groupnet subnet params
func (o *DeleteGroupnetSubnetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete groupnet subnet params
func (o *DeleteGroupnetSubnetParams) WithHTTPClient(client *http.Client) *DeleteGroupnetSubnetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete groupnet subnet params
func (o *DeleteGroupnetSubnetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupnet adds the groupnet to the delete groupnet subnet params
func (o *DeleteGroupnetSubnetParams) WithGroupnet(groupnet string) *DeleteGroupnetSubnetParams {
	o.SetGroupnet(groupnet)
	return o
}

// SetGroupnet adds the groupnet to the delete groupnet subnet params
func (o *DeleteGroupnetSubnetParams) SetGroupnet(groupnet string) {
	o.Groupnet = groupnet
}

// WithGroupnetSubnetID adds the groupnetSubnetID to the delete groupnet subnet params
func (o *DeleteGroupnetSubnetParams) WithGroupnetSubnetID(groupnetSubnetID string) *DeleteGroupnetSubnetParams {
	o.SetGroupnetSubnetID(groupnetSubnetID)
	return o
}

// SetGroupnetSubnetID adds the groupnetSubnetId to the delete groupnet subnet params
func (o *DeleteGroupnetSubnetParams) SetGroupnetSubnetID(groupnetSubnetID string) {
	o.GroupnetSubnetID = groupnetSubnetID
}

// WithForce adds the force to the delete groupnet subnet params
func (o *DeleteGroupnetSubnetParams) WithForce(force *bool) *DeleteGroupnetSubnetParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the delete groupnet subnet params
func (o *DeleteGroupnetSubnetParams) SetForce(force *bool) {
	o.Force = force
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteGroupnetSubnetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Groupnet
	if err := r.SetPathParam("Groupnet", o.Groupnet); err != nil {
		return err
	}

	// path param GroupnetSubnetId
	if err := r.SetPathParam("GroupnetSubnetId", o.GroupnetSubnetID); err != nil {
		return err
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
