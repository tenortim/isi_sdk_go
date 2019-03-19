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
)

// NewGetCompatibilitiesClassAvailableParams creates a new GetCompatibilitiesClassAvailableParams object
// with the default values initialized.
func NewGetCompatibilitiesClassAvailableParams() *GetCompatibilitiesClassAvailableParams {

	return &GetCompatibilitiesClassAvailableParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCompatibilitiesClassAvailableParamsWithTimeout creates a new GetCompatibilitiesClassAvailableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCompatibilitiesClassAvailableParamsWithTimeout(timeout time.Duration) *GetCompatibilitiesClassAvailableParams {

	return &GetCompatibilitiesClassAvailableParams{

		timeout: timeout,
	}
}

// NewGetCompatibilitiesClassAvailableParamsWithContext creates a new GetCompatibilitiesClassAvailableParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCompatibilitiesClassAvailableParamsWithContext(ctx context.Context) *GetCompatibilitiesClassAvailableParams {

	return &GetCompatibilitiesClassAvailableParams{

		Context: ctx,
	}
}

// NewGetCompatibilitiesClassAvailableParamsWithHTTPClient creates a new GetCompatibilitiesClassAvailableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCompatibilitiesClassAvailableParamsWithHTTPClient(client *http.Client) *GetCompatibilitiesClassAvailableParams {

	return &GetCompatibilitiesClassAvailableParams{
		HTTPClient: client,
	}
}

/*GetCompatibilitiesClassAvailableParams contains all the parameters to send to the API endpoint
for the get compatibilities class available operation typically these are written to a http.Request
*/
type GetCompatibilitiesClassAvailableParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get compatibilities class available params
func (o *GetCompatibilitiesClassAvailableParams) WithTimeout(timeout time.Duration) *GetCompatibilitiesClassAvailableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get compatibilities class available params
func (o *GetCompatibilitiesClassAvailableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get compatibilities class available params
func (o *GetCompatibilitiesClassAvailableParams) WithContext(ctx context.Context) *GetCompatibilitiesClassAvailableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get compatibilities class available params
func (o *GetCompatibilitiesClassAvailableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get compatibilities class available params
func (o *GetCompatibilitiesClassAvailableParams) WithHTTPClient(client *http.Client) *GetCompatibilitiesClassAvailableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get compatibilities class available params
func (o *GetCompatibilitiesClassAvailableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetCompatibilitiesClassAvailableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
