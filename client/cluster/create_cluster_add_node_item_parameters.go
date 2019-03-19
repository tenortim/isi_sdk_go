// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// NewCreateClusterAddNodeItemParams creates a new CreateClusterAddNodeItemParams object
// with the default values initialized.
func NewCreateClusterAddNodeItemParams() *CreateClusterAddNodeItemParams {
	var ()
	return &CreateClusterAddNodeItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateClusterAddNodeItemParamsWithTimeout creates a new CreateClusterAddNodeItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateClusterAddNodeItemParamsWithTimeout(timeout time.Duration) *CreateClusterAddNodeItemParams {
	var ()
	return &CreateClusterAddNodeItemParams{

		timeout: timeout,
	}
}

// NewCreateClusterAddNodeItemParamsWithContext creates a new CreateClusterAddNodeItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateClusterAddNodeItemParamsWithContext(ctx context.Context) *CreateClusterAddNodeItemParams {
	var ()
	return &CreateClusterAddNodeItemParams{

		Context: ctx,
	}
}

// NewCreateClusterAddNodeItemParamsWithHTTPClient creates a new CreateClusterAddNodeItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateClusterAddNodeItemParamsWithHTTPClient(client *http.Client) *CreateClusterAddNodeItemParams {
	var ()
	return &CreateClusterAddNodeItemParams{
		HTTPClient: client,
	}
}

/*CreateClusterAddNodeItemParams contains all the parameters to send to the API endpoint
for the create cluster add node item operation typically these are written to a http.Request
*/
type CreateClusterAddNodeItemParams struct {

	/*ClusterAddNodeItem*/
	ClusterAddNodeItem *models.ClusterAddNodeItem

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create cluster add node item params
func (o *CreateClusterAddNodeItemParams) WithTimeout(timeout time.Duration) *CreateClusterAddNodeItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cluster add node item params
func (o *CreateClusterAddNodeItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cluster add node item params
func (o *CreateClusterAddNodeItemParams) WithContext(ctx context.Context) *CreateClusterAddNodeItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cluster add node item params
func (o *CreateClusterAddNodeItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cluster add node item params
func (o *CreateClusterAddNodeItemParams) WithHTTPClient(client *http.Client) *CreateClusterAddNodeItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cluster add node item params
func (o *CreateClusterAddNodeItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterAddNodeItem adds the clusterAddNodeItem to the create cluster add node item params
func (o *CreateClusterAddNodeItemParams) WithClusterAddNodeItem(clusterAddNodeItem *models.ClusterAddNodeItem) *CreateClusterAddNodeItemParams {
	o.SetClusterAddNodeItem(clusterAddNodeItem)
	return o
}

// SetClusterAddNodeItem adds the clusterAddNodeItem to the create cluster add node item params
func (o *CreateClusterAddNodeItemParams) SetClusterAddNodeItem(clusterAddNodeItem *models.ClusterAddNodeItem) {
	o.ClusterAddNodeItem = clusterAddNodeItem
}

// WriteToRequest writes these params to a swagger request
func (o *CreateClusterAddNodeItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClusterAddNodeItem != nil {
		if err := r.SetBodyParam(o.ClusterAddNodeItem); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}