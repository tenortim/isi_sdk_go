// Code generated by go-swagger; DO NOT EDIT.

package network

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

// NewCreateNetworkScRebalanceAllItemParams creates a new CreateNetworkScRebalanceAllItemParams object
// with the default values initialized.
func NewCreateNetworkScRebalanceAllItemParams() *CreateNetworkScRebalanceAllItemParams {
	var ()
	return &CreateNetworkScRebalanceAllItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNetworkScRebalanceAllItemParamsWithTimeout creates a new CreateNetworkScRebalanceAllItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateNetworkScRebalanceAllItemParamsWithTimeout(timeout time.Duration) *CreateNetworkScRebalanceAllItemParams {
	var ()
	return &CreateNetworkScRebalanceAllItemParams{

		timeout: timeout,
	}
}

// NewCreateNetworkScRebalanceAllItemParamsWithContext creates a new CreateNetworkScRebalanceAllItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateNetworkScRebalanceAllItemParamsWithContext(ctx context.Context) *CreateNetworkScRebalanceAllItemParams {
	var ()
	return &CreateNetworkScRebalanceAllItemParams{

		Context: ctx,
	}
}

// NewCreateNetworkScRebalanceAllItemParamsWithHTTPClient creates a new CreateNetworkScRebalanceAllItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateNetworkScRebalanceAllItemParamsWithHTTPClient(client *http.Client) *CreateNetworkScRebalanceAllItemParams {
	var ()
	return &CreateNetworkScRebalanceAllItemParams{
		HTTPClient: client,
	}
}

/*CreateNetworkScRebalanceAllItemParams contains all the parameters to send to the API endpoint
for the create network sc rebalance all item operation typically these are written to a http.Request
*/
type CreateNetworkScRebalanceAllItemParams struct {

	/*NetworkScRebalanceAllItem*/
	NetworkScRebalanceAllItem models.Empty

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create network sc rebalance all item params
func (o *CreateNetworkScRebalanceAllItemParams) WithTimeout(timeout time.Duration) *CreateNetworkScRebalanceAllItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create network sc rebalance all item params
func (o *CreateNetworkScRebalanceAllItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create network sc rebalance all item params
func (o *CreateNetworkScRebalanceAllItemParams) WithContext(ctx context.Context) *CreateNetworkScRebalanceAllItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create network sc rebalance all item params
func (o *CreateNetworkScRebalanceAllItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create network sc rebalance all item params
func (o *CreateNetworkScRebalanceAllItemParams) WithHTTPClient(client *http.Client) *CreateNetworkScRebalanceAllItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create network sc rebalance all item params
func (o *CreateNetworkScRebalanceAllItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkScRebalanceAllItem adds the networkScRebalanceAllItem to the create network sc rebalance all item params
func (o *CreateNetworkScRebalanceAllItemParams) WithNetworkScRebalanceAllItem(networkScRebalanceAllItem models.Empty) *CreateNetworkScRebalanceAllItemParams {
	o.SetNetworkScRebalanceAllItem(networkScRebalanceAllItem)
	return o
}

// SetNetworkScRebalanceAllItem adds the networkScRebalanceAllItem to the create network sc rebalance all item params
func (o *CreateNetworkScRebalanceAllItemParams) SetNetworkScRebalanceAllItem(networkScRebalanceAllItem models.Empty) {
	o.NetworkScRebalanceAllItem = networkScRebalanceAllItem
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNetworkScRebalanceAllItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.NetworkScRebalanceAllItem != nil {
		if err := r.SetBodyParam(o.NetworkScRebalanceAllItem); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
