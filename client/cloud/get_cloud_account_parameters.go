// Code generated by go-swagger; DO NOT EDIT.

package cloud

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
)

// NewGetCloudAccountParams creates a new GetCloudAccountParams object
// with the default values initialized.
func NewGetCloudAccountParams() *GetCloudAccountParams {
	var ()
	return &GetCloudAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCloudAccountParamsWithTimeout creates a new GetCloudAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCloudAccountParamsWithTimeout(timeout time.Duration) *GetCloudAccountParams {
	var ()
	return &GetCloudAccountParams{

		timeout: timeout,
	}
}

// NewGetCloudAccountParamsWithContext creates a new GetCloudAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCloudAccountParamsWithContext(ctx context.Context) *GetCloudAccountParams {
	var ()
	return &GetCloudAccountParams{

		Context: ctx,
	}
}

// NewGetCloudAccountParamsWithHTTPClient creates a new GetCloudAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCloudAccountParamsWithHTTPClient(client *http.Client) *GetCloudAccountParams {
	var ()
	return &GetCloudAccountParams{
		HTTPClient: client,
	}
}

/*GetCloudAccountParams contains all the parameters to send to the API endpoint
for the get cloud account operation typically these are written to a http.Request
*/
type GetCloudAccountParams struct {

	/*CloudAccountID
	  Retrieve cloud account information.

	*/
	CloudAccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cloud account params
func (o *GetCloudAccountParams) WithTimeout(timeout time.Duration) *GetCloudAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cloud account params
func (o *GetCloudAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cloud account params
func (o *GetCloudAccountParams) WithContext(ctx context.Context) *GetCloudAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cloud account params
func (o *GetCloudAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cloud account params
func (o *GetCloudAccountParams) WithHTTPClient(client *http.Client) *GetCloudAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cloud account params
func (o *GetCloudAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudAccountID adds the cloudAccountID to the get cloud account params
func (o *GetCloudAccountParams) WithCloudAccountID(cloudAccountID string) *GetCloudAccountParams {
	o.SetCloudAccountID(cloudAccountID)
	return o
}

// SetCloudAccountID adds the cloudAccountId to the get cloud account params
func (o *GetCloudAccountParams) SetCloudAccountID(cloudAccountID string) {
	o.CloudAccountID = cloudAccountID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCloudAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param CloudAccountId
	if err := r.SetPathParam("CloudAccountId", o.CloudAccountID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
