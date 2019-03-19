// Code generated by go-swagger; DO NOT EDIT.

package storagepool

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

// NewCreateCompatibilitiesClassActiveItemParams creates a new CreateCompatibilitiesClassActiveItemParams object
// with the default values initialized.
func NewCreateCompatibilitiesClassActiveItemParams() *CreateCompatibilitiesClassActiveItemParams {
	var ()
	return &CreateCompatibilitiesClassActiveItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCompatibilitiesClassActiveItemParamsWithTimeout creates a new CreateCompatibilitiesClassActiveItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateCompatibilitiesClassActiveItemParamsWithTimeout(timeout time.Duration) *CreateCompatibilitiesClassActiveItemParams {
	var ()
	return &CreateCompatibilitiesClassActiveItemParams{

		timeout: timeout,
	}
}

// NewCreateCompatibilitiesClassActiveItemParamsWithContext creates a new CreateCompatibilitiesClassActiveItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateCompatibilitiesClassActiveItemParamsWithContext(ctx context.Context) *CreateCompatibilitiesClassActiveItemParams {
	var ()
	return &CreateCompatibilitiesClassActiveItemParams{

		Context: ctx,
	}
}

// NewCreateCompatibilitiesClassActiveItemParamsWithHTTPClient creates a new CreateCompatibilitiesClassActiveItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateCompatibilitiesClassActiveItemParamsWithHTTPClient(client *http.Client) *CreateCompatibilitiesClassActiveItemParams {
	var ()
	return &CreateCompatibilitiesClassActiveItemParams{
		HTTPClient: client,
	}
}

/*CreateCompatibilitiesClassActiveItemParams contains all the parameters to send to the API endpoint
for the create compatibilities class active item operation typically these are written to a http.Request
*/
type CreateCompatibilitiesClassActiveItemParams struct {

	/*CompatibilitiesClassActiveItem*/
	CompatibilitiesClassActiveItem *models.CompatibilitiesClassActiveItem

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create compatibilities class active item params
func (o *CreateCompatibilitiesClassActiveItemParams) WithTimeout(timeout time.Duration) *CreateCompatibilitiesClassActiveItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create compatibilities class active item params
func (o *CreateCompatibilitiesClassActiveItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create compatibilities class active item params
func (o *CreateCompatibilitiesClassActiveItemParams) WithContext(ctx context.Context) *CreateCompatibilitiesClassActiveItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create compatibilities class active item params
func (o *CreateCompatibilitiesClassActiveItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create compatibilities class active item params
func (o *CreateCompatibilitiesClassActiveItemParams) WithHTTPClient(client *http.Client) *CreateCompatibilitiesClassActiveItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create compatibilities class active item params
func (o *CreateCompatibilitiesClassActiveItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCompatibilitiesClassActiveItem adds the compatibilitiesClassActiveItem to the create compatibilities class active item params
func (o *CreateCompatibilitiesClassActiveItemParams) WithCompatibilitiesClassActiveItem(compatibilitiesClassActiveItem *models.CompatibilitiesClassActiveItem) *CreateCompatibilitiesClassActiveItemParams {
	o.SetCompatibilitiesClassActiveItem(compatibilitiesClassActiveItem)
	return o
}

// SetCompatibilitiesClassActiveItem adds the compatibilitiesClassActiveItem to the create compatibilities class active item params
func (o *CreateCompatibilitiesClassActiveItemParams) SetCompatibilitiesClassActiveItem(compatibilitiesClassActiveItem *models.CompatibilitiesClassActiveItem) {
	o.CompatibilitiesClassActiveItem = compatibilitiesClassActiveItem
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCompatibilitiesClassActiveItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CompatibilitiesClassActiveItem != nil {
		if err := r.SetBodyParam(o.CompatibilitiesClassActiveItem); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
