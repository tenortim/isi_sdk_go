// Code generated by go-swagger; DO NOT EDIT.

package cluster_nodes

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

// NewUpdateNodeStateServicelightParams creates a new UpdateNodeStateServicelightParams object
// with the default values initialized.
func NewUpdateNodeStateServicelightParams() *UpdateNodeStateServicelightParams {
	var ()
	return &UpdateNodeStateServicelightParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNodeStateServicelightParamsWithTimeout creates a new UpdateNodeStateServicelightParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateNodeStateServicelightParamsWithTimeout(timeout time.Duration) *UpdateNodeStateServicelightParams {
	var ()
	return &UpdateNodeStateServicelightParams{

		timeout: timeout,
	}
}

// NewUpdateNodeStateServicelightParamsWithContext creates a new UpdateNodeStateServicelightParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateNodeStateServicelightParamsWithContext(ctx context.Context) *UpdateNodeStateServicelightParams {
	var ()
	return &UpdateNodeStateServicelightParams{

		Context: ctx,
	}
}

// NewUpdateNodeStateServicelightParamsWithHTTPClient creates a new UpdateNodeStateServicelightParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateNodeStateServicelightParamsWithHTTPClient(client *http.Client) *UpdateNodeStateServicelightParams {
	var ()
	return &UpdateNodeStateServicelightParams{
		HTTPClient: client,
	}
}

/*UpdateNodeStateServicelightParams contains all the parameters to send to the API endpoint
for the update node state servicelight operation typically these are written to a http.Request
*/
type UpdateNodeStateServicelightParams struct {

	/*Lnn*/
	Lnn int64
	/*NodeStateServicelight*/
	NodeStateServicelight *models.NodeStateServicelightExtended

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update node state servicelight params
func (o *UpdateNodeStateServicelightParams) WithTimeout(timeout time.Duration) *UpdateNodeStateServicelightParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update node state servicelight params
func (o *UpdateNodeStateServicelightParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update node state servicelight params
func (o *UpdateNodeStateServicelightParams) WithContext(ctx context.Context) *UpdateNodeStateServicelightParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update node state servicelight params
func (o *UpdateNodeStateServicelightParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update node state servicelight params
func (o *UpdateNodeStateServicelightParams) WithHTTPClient(client *http.Client) *UpdateNodeStateServicelightParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update node state servicelight params
func (o *UpdateNodeStateServicelightParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLnn adds the lnn to the update node state servicelight params
func (o *UpdateNodeStateServicelightParams) WithLnn(lnn int64) *UpdateNodeStateServicelightParams {
	o.SetLnn(lnn)
	return o
}

// SetLnn adds the lnn to the update node state servicelight params
func (o *UpdateNodeStateServicelightParams) SetLnn(lnn int64) {
	o.Lnn = lnn
}

// WithNodeStateServicelight adds the nodeStateServicelight to the update node state servicelight params
func (o *UpdateNodeStateServicelightParams) WithNodeStateServicelight(nodeStateServicelight *models.NodeStateServicelightExtended) *UpdateNodeStateServicelightParams {
	o.SetNodeStateServicelight(nodeStateServicelight)
	return o
}

// SetNodeStateServicelight adds the nodeStateServicelight to the update node state servicelight params
func (o *UpdateNodeStateServicelightParams) SetNodeStateServicelight(nodeStateServicelight *models.NodeStateServicelightExtended) {
	o.NodeStateServicelight = nodeStateServicelight
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNodeStateServicelightParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Lnn
	if err := r.SetPathParam("Lnn", swag.FormatInt64(o.Lnn)); err != nil {
		return err
	}

	if o.NodeStateServicelight != nil {
		if err := r.SetBodyParam(o.NodeStateServicelight); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
