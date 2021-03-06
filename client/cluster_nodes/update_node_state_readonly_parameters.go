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

// NewUpdateNodeStateReadonlyParams creates a new UpdateNodeStateReadonlyParams object
// with the default values initialized.
func NewUpdateNodeStateReadonlyParams() *UpdateNodeStateReadonlyParams {
	var ()
	return &UpdateNodeStateReadonlyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNodeStateReadonlyParamsWithTimeout creates a new UpdateNodeStateReadonlyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateNodeStateReadonlyParamsWithTimeout(timeout time.Duration) *UpdateNodeStateReadonlyParams {
	var ()
	return &UpdateNodeStateReadonlyParams{

		timeout: timeout,
	}
}

// NewUpdateNodeStateReadonlyParamsWithContext creates a new UpdateNodeStateReadonlyParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateNodeStateReadonlyParamsWithContext(ctx context.Context) *UpdateNodeStateReadonlyParams {
	var ()
	return &UpdateNodeStateReadonlyParams{

		Context: ctx,
	}
}

// NewUpdateNodeStateReadonlyParamsWithHTTPClient creates a new UpdateNodeStateReadonlyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateNodeStateReadonlyParamsWithHTTPClient(client *http.Client) *UpdateNodeStateReadonlyParams {
	var ()
	return &UpdateNodeStateReadonlyParams{
		HTTPClient: client,
	}
}

/*UpdateNodeStateReadonlyParams contains all the parameters to send to the API endpoint
for the update node state readonly operation typically these are written to a http.Request
*/
type UpdateNodeStateReadonlyParams struct {

	/*Lnn*/
	Lnn int64
	/*NodeStateReadonly*/
	NodeStateReadonly *models.NodeStateReadonlyExtended

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update node state readonly params
func (o *UpdateNodeStateReadonlyParams) WithTimeout(timeout time.Duration) *UpdateNodeStateReadonlyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update node state readonly params
func (o *UpdateNodeStateReadonlyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update node state readonly params
func (o *UpdateNodeStateReadonlyParams) WithContext(ctx context.Context) *UpdateNodeStateReadonlyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update node state readonly params
func (o *UpdateNodeStateReadonlyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update node state readonly params
func (o *UpdateNodeStateReadonlyParams) WithHTTPClient(client *http.Client) *UpdateNodeStateReadonlyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update node state readonly params
func (o *UpdateNodeStateReadonlyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLnn adds the lnn to the update node state readonly params
func (o *UpdateNodeStateReadonlyParams) WithLnn(lnn int64) *UpdateNodeStateReadonlyParams {
	o.SetLnn(lnn)
	return o
}

// SetLnn adds the lnn to the update node state readonly params
func (o *UpdateNodeStateReadonlyParams) SetLnn(lnn int64) {
	o.Lnn = lnn
}

// WithNodeStateReadonly adds the nodeStateReadonly to the update node state readonly params
func (o *UpdateNodeStateReadonlyParams) WithNodeStateReadonly(nodeStateReadonly *models.NodeStateReadonlyExtended) *UpdateNodeStateReadonlyParams {
	o.SetNodeStateReadonly(nodeStateReadonly)
	return o
}

// SetNodeStateReadonly adds the nodeStateReadonly to the update node state readonly params
func (o *UpdateNodeStateReadonlyParams) SetNodeStateReadonly(nodeStateReadonly *models.NodeStateReadonlyExtended) {
	o.NodeStateReadonly = nodeStateReadonly
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNodeStateReadonlyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Lnn
	if err := r.SetPathParam("Lnn", swag.FormatInt64(o.Lnn)); err != nil {
		return err
	}

	if o.NodeStateReadonly != nil {
		if err := r.SetBodyParam(o.NodeStateReadonly); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
