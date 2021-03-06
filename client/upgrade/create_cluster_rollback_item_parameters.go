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

// NewCreateClusterRollbackItemParams creates a new CreateClusterRollbackItemParams object
// with the default values initialized.
func NewCreateClusterRollbackItemParams() *CreateClusterRollbackItemParams {
	var ()
	return &CreateClusterRollbackItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateClusterRollbackItemParamsWithTimeout creates a new CreateClusterRollbackItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateClusterRollbackItemParamsWithTimeout(timeout time.Duration) *CreateClusterRollbackItemParams {
	var ()
	return &CreateClusterRollbackItemParams{

		timeout: timeout,
	}
}

// NewCreateClusterRollbackItemParamsWithContext creates a new CreateClusterRollbackItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateClusterRollbackItemParamsWithContext(ctx context.Context) *CreateClusterRollbackItemParams {
	var ()
	return &CreateClusterRollbackItemParams{

		Context: ctx,
	}
}

// NewCreateClusterRollbackItemParamsWithHTTPClient creates a new CreateClusterRollbackItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateClusterRollbackItemParamsWithHTTPClient(client *http.Client) *CreateClusterRollbackItemParams {
	var ()
	return &CreateClusterRollbackItemParams{
		HTTPClient: client,
	}
}

/*CreateClusterRollbackItemParams contains all the parameters to send to the API endpoint
for the create cluster rollback item operation typically these are written to a http.Request
*/
type CreateClusterRollbackItemParams struct {

	/*ClusterRollbackItem*/
	ClusterRollbackItem models.Empty

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create cluster rollback item params
func (o *CreateClusterRollbackItemParams) WithTimeout(timeout time.Duration) *CreateClusterRollbackItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cluster rollback item params
func (o *CreateClusterRollbackItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cluster rollback item params
func (o *CreateClusterRollbackItemParams) WithContext(ctx context.Context) *CreateClusterRollbackItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cluster rollback item params
func (o *CreateClusterRollbackItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cluster rollback item params
func (o *CreateClusterRollbackItemParams) WithHTTPClient(client *http.Client) *CreateClusterRollbackItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cluster rollback item params
func (o *CreateClusterRollbackItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterRollbackItem adds the clusterRollbackItem to the create cluster rollback item params
func (o *CreateClusterRollbackItemParams) WithClusterRollbackItem(clusterRollbackItem models.Empty) *CreateClusterRollbackItemParams {
	o.SetClusterRollbackItem(clusterRollbackItem)
	return o
}

// SetClusterRollbackItem adds the clusterRollbackItem to the create cluster rollback item params
func (o *CreateClusterRollbackItemParams) SetClusterRollbackItem(clusterRollbackItem models.Empty) {
	o.ClusterRollbackItem = clusterRollbackItem
}

// WriteToRequest writes these params to a swagger request
func (o *CreateClusterRollbackItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClusterRollbackItem != nil {
		if err := r.SetBodyParam(o.ClusterRollbackItem); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
