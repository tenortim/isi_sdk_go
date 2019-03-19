// Code generated by go-swagger; DO NOT EDIT.

package network_groupnets_subnets

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

// NewUpdatePoolsPoolRuleParams creates a new UpdatePoolsPoolRuleParams object
// with the default values initialized.
func NewUpdatePoolsPoolRuleParams() *UpdatePoolsPoolRuleParams {
	var ()
	return &UpdatePoolsPoolRuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePoolsPoolRuleParamsWithTimeout creates a new UpdatePoolsPoolRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePoolsPoolRuleParamsWithTimeout(timeout time.Duration) *UpdatePoolsPoolRuleParams {
	var ()
	return &UpdatePoolsPoolRuleParams{

		timeout: timeout,
	}
}

// NewUpdatePoolsPoolRuleParamsWithContext creates a new UpdatePoolsPoolRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePoolsPoolRuleParamsWithContext(ctx context.Context) *UpdatePoolsPoolRuleParams {
	var ()
	return &UpdatePoolsPoolRuleParams{

		Context: ctx,
	}
}

// NewUpdatePoolsPoolRuleParamsWithHTTPClient creates a new UpdatePoolsPoolRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePoolsPoolRuleParamsWithHTTPClient(client *http.Client) *UpdatePoolsPoolRuleParams {
	var ()
	return &UpdatePoolsPoolRuleParams{
		HTTPClient: client,
	}
}

/*UpdatePoolsPoolRuleParams contains all the parameters to send to the API endpoint
for the update pools pool rule operation typically these are written to a http.Request
*/
type UpdatePoolsPoolRuleParams struct {

	/*Groupnet*/
	Groupnet string
	/*Pool*/
	Pool string
	/*PoolsPoolRule*/
	PoolsPoolRule *models.PoolsPoolRule
	/*PoolsPoolRuleID
	  Modify a network rule.

	*/
	PoolsPoolRuleID string
	/*Subnet*/
	Subnet string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update pools pool rule params
func (o *UpdatePoolsPoolRuleParams) WithTimeout(timeout time.Duration) *UpdatePoolsPoolRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update pools pool rule params
func (o *UpdatePoolsPoolRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update pools pool rule params
func (o *UpdatePoolsPoolRuleParams) WithContext(ctx context.Context) *UpdatePoolsPoolRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update pools pool rule params
func (o *UpdatePoolsPoolRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update pools pool rule params
func (o *UpdatePoolsPoolRuleParams) WithHTTPClient(client *http.Client) *UpdatePoolsPoolRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update pools pool rule params
func (o *UpdatePoolsPoolRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupnet adds the groupnet to the update pools pool rule params
func (o *UpdatePoolsPoolRuleParams) WithGroupnet(groupnet string) *UpdatePoolsPoolRuleParams {
	o.SetGroupnet(groupnet)
	return o
}

// SetGroupnet adds the groupnet to the update pools pool rule params
func (o *UpdatePoolsPoolRuleParams) SetGroupnet(groupnet string) {
	o.Groupnet = groupnet
}

// WithPool adds the pool to the update pools pool rule params
func (o *UpdatePoolsPoolRuleParams) WithPool(pool string) *UpdatePoolsPoolRuleParams {
	o.SetPool(pool)
	return o
}

// SetPool adds the pool to the update pools pool rule params
func (o *UpdatePoolsPoolRuleParams) SetPool(pool string) {
	o.Pool = pool
}

// WithPoolsPoolRule adds the poolsPoolRule to the update pools pool rule params
func (o *UpdatePoolsPoolRuleParams) WithPoolsPoolRule(poolsPoolRule *models.PoolsPoolRule) *UpdatePoolsPoolRuleParams {
	o.SetPoolsPoolRule(poolsPoolRule)
	return o
}

// SetPoolsPoolRule adds the poolsPoolRule to the update pools pool rule params
func (o *UpdatePoolsPoolRuleParams) SetPoolsPoolRule(poolsPoolRule *models.PoolsPoolRule) {
	o.PoolsPoolRule = poolsPoolRule
}

// WithPoolsPoolRuleID adds the poolsPoolRuleID to the update pools pool rule params
func (o *UpdatePoolsPoolRuleParams) WithPoolsPoolRuleID(poolsPoolRuleID string) *UpdatePoolsPoolRuleParams {
	o.SetPoolsPoolRuleID(poolsPoolRuleID)
	return o
}

// SetPoolsPoolRuleID adds the poolsPoolRuleId to the update pools pool rule params
func (o *UpdatePoolsPoolRuleParams) SetPoolsPoolRuleID(poolsPoolRuleID string) {
	o.PoolsPoolRuleID = poolsPoolRuleID
}

// WithSubnet adds the subnet to the update pools pool rule params
func (o *UpdatePoolsPoolRuleParams) WithSubnet(subnet string) *UpdatePoolsPoolRuleParams {
	o.SetSubnet(subnet)
	return o
}

// SetSubnet adds the subnet to the update pools pool rule params
func (o *UpdatePoolsPoolRuleParams) SetSubnet(subnet string) {
	o.Subnet = subnet
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePoolsPoolRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Groupnet
	if err := r.SetPathParam("Groupnet", o.Groupnet); err != nil {
		return err
	}

	// path param Pool
	if err := r.SetPathParam("Pool", o.Pool); err != nil {
		return err
	}

	if o.PoolsPoolRule != nil {
		if err := r.SetBodyParam(o.PoolsPoolRule); err != nil {
			return err
		}
	}

	// path param PoolsPoolRuleId
	if err := r.SetPathParam("PoolsPoolRuleId", o.PoolsPoolRuleID); err != nil {
		return err
	}

	// path param Subnet
	if err := r.SetPathParam("Subnet", o.Subnet); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}