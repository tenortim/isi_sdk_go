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

	models "github.com/tenortim/isi_sdk_go/models"
)

// NewCreateSubnetsSubnetPoolParams creates a new CreateSubnetsSubnetPoolParams object
// with the default values initialized.
func NewCreateSubnetsSubnetPoolParams() *CreateSubnetsSubnetPoolParams {
	var ()
	return &CreateSubnetsSubnetPoolParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSubnetsSubnetPoolParamsWithTimeout creates a new CreateSubnetsSubnetPoolParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSubnetsSubnetPoolParamsWithTimeout(timeout time.Duration) *CreateSubnetsSubnetPoolParams {
	var ()
	return &CreateSubnetsSubnetPoolParams{

		timeout: timeout,
	}
}

// NewCreateSubnetsSubnetPoolParamsWithContext creates a new CreateSubnetsSubnetPoolParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSubnetsSubnetPoolParamsWithContext(ctx context.Context) *CreateSubnetsSubnetPoolParams {
	var ()
	return &CreateSubnetsSubnetPoolParams{

		Context: ctx,
	}
}

// NewCreateSubnetsSubnetPoolParamsWithHTTPClient creates a new CreateSubnetsSubnetPoolParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSubnetsSubnetPoolParamsWithHTTPClient(client *http.Client) *CreateSubnetsSubnetPoolParams {
	var ()
	return &CreateSubnetsSubnetPoolParams{
		HTTPClient: client,
	}
}

/*CreateSubnetsSubnetPoolParams contains all the parameters to send to the API endpoint
for the create subnets subnet pool operation typically these are written to a http.Request
*/
type CreateSubnetsSubnetPoolParams struct {

	/*Groupnet*/
	Groupnet string
	/*Subnet*/
	Subnet string
	/*SubnetsSubnetPool*/
	SubnetsSubnetPool *models.SubnetsSubnetPoolCreateParams
	/*Force
	  Force creating this pool even if it causes an MTU conflict.

	*/
	Force *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create subnets subnet pool params
func (o *CreateSubnetsSubnetPoolParams) WithTimeout(timeout time.Duration) *CreateSubnetsSubnetPoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create subnets subnet pool params
func (o *CreateSubnetsSubnetPoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create subnets subnet pool params
func (o *CreateSubnetsSubnetPoolParams) WithContext(ctx context.Context) *CreateSubnetsSubnetPoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create subnets subnet pool params
func (o *CreateSubnetsSubnetPoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create subnets subnet pool params
func (o *CreateSubnetsSubnetPoolParams) WithHTTPClient(client *http.Client) *CreateSubnetsSubnetPoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create subnets subnet pool params
func (o *CreateSubnetsSubnetPoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupnet adds the groupnet to the create subnets subnet pool params
func (o *CreateSubnetsSubnetPoolParams) WithGroupnet(groupnet string) *CreateSubnetsSubnetPoolParams {
	o.SetGroupnet(groupnet)
	return o
}

// SetGroupnet adds the groupnet to the create subnets subnet pool params
func (o *CreateSubnetsSubnetPoolParams) SetGroupnet(groupnet string) {
	o.Groupnet = groupnet
}

// WithSubnet adds the subnet to the create subnets subnet pool params
func (o *CreateSubnetsSubnetPoolParams) WithSubnet(subnet string) *CreateSubnetsSubnetPoolParams {
	o.SetSubnet(subnet)
	return o
}

// SetSubnet adds the subnet to the create subnets subnet pool params
func (o *CreateSubnetsSubnetPoolParams) SetSubnet(subnet string) {
	o.Subnet = subnet
}

// WithSubnetsSubnetPool adds the subnetsSubnetPool to the create subnets subnet pool params
func (o *CreateSubnetsSubnetPoolParams) WithSubnetsSubnetPool(subnetsSubnetPool *models.SubnetsSubnetPoolCreateParams) *CreateSubnetsSubnetPoolParams {
	o.SetSubnetsSubnetPool(subnetsSubnetPool)
	return o
}

// SetSubnetsSubnetPool adds the subnetsSubnetPool to the create subnets subnet pool params
func (o *CreateSubnetsSubnetPoolParams) SetSubnetsSubnetPool(subnetsSubnetPool *models.SubnetsSubnetPoolCreateParams) {
	o.SubnetsSubnetPool = subnetsSubnetPool
}

// WithForce adds the force to the create subnets subnet pool params
func (o *CreateSubnetsSubnetPoolParams) WithForce(force *bool) *CreateSubnetsSubnetPoolParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the create subnets subnet pool params
func (o *CreateSubnetsSubnetPoolParams) SetForce(force *bool) {
	o.Force = force
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSubnetsSubnetPoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Groupnet
	if err := r.SetPathParam("Groupnet", o.Groupnet); err != nil {
		return err
	}

	// path param Subnet
	if err := r.SetPathParam("Subnet", o.Subnet); err != nil {
		return err
	}

	if o.SubnetsSubnetPool != nil {
		if err := r.SetBodyParam(o.SubnetsSubnetPool); err != nil {
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
