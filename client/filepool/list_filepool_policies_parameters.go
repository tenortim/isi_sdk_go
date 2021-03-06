// Code generated by go-swagger; DO NOT EDIT.

package filepool

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

// NewListFilepoolPoliciesParams creates a new ListFilepoolPoliciesParams object
// with the default values initialized.
func NewListFilepoolPoliciesParams() *ListFilepoolPoliciesParams {

	return &ListFilepoolPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListFilepoolPoliciesParamsWithTimeout creates a new ListFilepoolPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListFilepoolPoliciesParamsWithTimeout(timeout time.Duration) *ListFilepoolPoliciesParams {

	return &ListFilepoolPoliciesParams{

		timeout: timeout,
	}
}

// NewListFilepoolPoliciesParamsWithContext creates a new ListFilepoolPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListFilepoolPoliciesParamsWithContext(ctx context.Context) *ListFilepoolPoliciesParams {

	return &ListFilepoolPoliciesParams{

		Context: ctx,
	}
}

// NewListFilepoolPoliciesParamsWithHTTPClient creates a new ListFilepoolPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListFilepoolPoliciesParamsWithHTTPClient(client *http.Client) *ListFilepoolPoliciesParams {

	return &ListFilepoolPoliciesParams{
		HTTPClient: client,
	}
}

/*ListFilepoolPoliciesParams contains all the parameters to send to the API endpoint
for the list filepool policies operation typically these are written to a http.Request
*/
type ListFilepoolPoliciesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list filepool policies params
func (o *ListFilepoolPoliciesParams) WithTimeout(timeout time.Duration) *ListFilepoolPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list filepool policies params
func (o *ListFilepoolPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list filepool policies params
func (o *ListFilepoolPoliciesParams) WithContext(ctx context.Context) *ListFilepoolPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list filepool policies params
func (o *ListFilepoolPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list filepool policies params
func (o *ListFilepoolPoliciesParams) WithHTTPClient(client *http.Client) *ListFilepoolPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list filepool policies params
func (o *ListFilepoolPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListFilepoolPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
