// Code generated by go-swagger; DO NOT EDIT.

package auth

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

// NewListSettingsKrb5DomainsParams creates a new ListSettingsKrb5DomainsParams object
// with the default values initialized.
func NewListSettingsKrb5DomainsParams() *ListSettingsKrb5DomainsParams {

	return &ListSettingsKrb5DomainsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSettingsKrb5DomainsParamsWithTimeout creates a new ListSettingsKrb5DomainsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSettingsKrb5DomainsParamsWithTimeout(timeout time.Duration) *ListSettingsKrb5DomainsParams {

	return &ListSettingsKrb5DomainsParams{

		timeout: timeout,
	}
}

// NewListSettingsKrb5DomainsParamsWithContext creates a new ListSettingsKrb5DomainsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSettingsKrb5DomainsParamsWithContext(ctx context.Context) *ListSettingsKrb5DomainsParams {

	return &ListSettingsKrb5DomainsParams{

		Context: ctx,
	}
}

// NewListSettingsKrb5DomainsParamsWithHTTPClient creates a new ListSettingsKrb5DomainsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSettingsKrb5DomainsParamsWithHTTPClient(client *http.Client) *ListSettingsKrb5DomainsParams {

	return &ListSettingsKrb5DomainsParams{
		HTTPClient: client,
	}
}

/*ListSettingsKrb5DomainsParams contains all the parameters to send to the API endpoint
for the list settings krb5 domains operation typically these are written to a http.Request
*/
type ListSettingsKrb5DomainsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list settings krb5 domains params
func (o *ListSettingsKrb5DomainsParams) WithTimeout(timeout time.Duration) *ListSettingsKrb5DomainsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list settings krb5 domains params
func (o *ListSettingsKrb5DomainsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list settings krb5 domains params
func (o *ListSettingsKrb5DomainsParams) WithContext(ctx context.Context) *ListSettingsKrb5DomainsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list settings krb5 domains params
func (o *ListSettingsKrb5DomainsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list settings krb5 domains params
func (o *ListSettingsKrb5DomainsParams) WithHTTPClient(client *http.Client) *ListSettingsKrb5DomainsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list settings krb5 domains params
func (o *ListSettingsKrb5DomainsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListSettingsKrb5DomainsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
