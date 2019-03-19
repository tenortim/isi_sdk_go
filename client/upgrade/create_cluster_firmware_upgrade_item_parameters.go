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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// NewCreateClusterFirmwareUpgradeItemParams creates a new CreateClusterFirmwareUpgradeItemParams object
// with the default values initialized.
func NewCreateClusterFirmwareUpgradeItemParams() *CreateClusterFirmwareUpgradeItemParams {
	var ()
	return &CreateClusterFirmwareUpgradeItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateClusterFirmwareUpgradeItemParamsWithTimeout creates a new CreateClusterFirmwareUpgradeItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateClusterFirmwareUpgradeItemParamsWithTimeout(timeout time.Duration) *CreateClusterFirmwareUpgradeItemParams {
	var ()
	return &CreateClusterFirmwareUpgradeItemParams{

		timeout: timeout,
	}
}

// NewCreateClusterFirmwareUpgradeItemParamsWithContext creates a new CreateClusterFirmwareUpgradeItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateClusterFirmwareUpgradeItemParamsWithContext(ctx context.Context) *CreateClusterFirmwareUpgradeItemParams {
	var ()
	return &CreateClusterFirmwareUpgradeItemParams{

		Context: ctx,
	}
}

// NewCreateClusterFirmwareUpgradeItemParamsWithHTTPClient creates a new CreateClusterFirmwareUpgradeItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateClusterFirmwareUpgradeItemParamsWithHTTPClient(client *http.Client) *CreateClusterFirmwareUpgradeItemParams {
	var ()
	return &CreateClusterFirmwareUpgradeItemParams{
		HTTPClient: client,
	}
}

/*CreateClusterFirmwareUpgradeItemParams contains all the parameters to send to the API endpoint
for the create cluster firmware upgrade item operation typically these are written to a http.Request
*/
type CreateClusterFirmwareUpgradeItemParams struct {

	/*ClusterFirmwareUpgradeItem*/
	ClusterFirmwareUpgradeItem *models.ClusterFirmwareUpgradeItem

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create cluster firmware upgrade item params
func (o *CreateClusterFirmwareUpgradeItemParams) WithTimeout(timeout time.Duration) *CreateClusterFirmwareUpgradeItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cluster firmware upgrade item params
func (o *CreateClusterFirmwareUpgradeItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cluster firmware upgrade item params
func (o *CreateClusterFirmwareUpgradeItemParams) WithContext(ctx context.Context) *CreateClusterFirmwareUpgradeItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cluster firmware upgrade item params
func (o *CreateClusterFirmwareUpgradeItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cluster firmware upgrade item params
func (o *CreateClusterFirmwareUpgradeItemParams) WithHTTPClient(client *http.Client) *CreateClusterFirmwareUpgradeItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cluster firmware upgrade item params
func (o *CreateClusterFirmwareUpgradeItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterFirmwareUpgradeItem adds the clusterFirmwareUpgradeItem to the create cluster firmware upgrade item params
func (o *CreateClusterFirmwareUpgradeItemParams) WithClusterFirmwareUpgradeItem(clusterFirmwareUpgradeItem *models.ClusterFirmwareUpgradeItem) *CreateClusterFirmwareUpgradeItemParams {
	o.SetClusterFirmwareUpgradeItem(clusterFirmwareUpgradeItem)
	return o
}

// SetClusterFirmwareUpgradeItem adds the clusterFirmwareUpgradeItem to the create cluster firmware upgrade item params
func (o *CreateClusterFirmwareUpgradeItemParams) SetClusterFirmwareUpgradeItem(clusterFirmwareUpgradeItem *models.ClusterFirmwareUpgradeItem) {
	o.ClusterFirmwareUpgradeItem = clusterFirmwareUpgradeItem
}

// WriteToRequest writes these params to a swagger request
func (o *CreateClusterFirmwareUpgradeItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClusterFirmwareUpgradeItem != nil {
		if err := r.SetBodyParam(o.ClusterFirmwareUpgradeItem); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
