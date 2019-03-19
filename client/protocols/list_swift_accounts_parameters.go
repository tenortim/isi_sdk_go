// Code generated by go-swagger; DO NOT EDIT.

package protocols

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

// NewListSwiftAccountsParams creates a new ListSwiftAccountsParams object
// with the default values initialized.
func NewListSwiftAccountsParams() *ListSwiftAccountsParams {
	var ()
	return &ListSwiftAccountsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSwiftAccountsParamsWithTimeout creates a new ListSwiftAccountsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSwiftAccountsParamsWithTimeout(timeout time.Duration) *ListSwiftAccountsParams {
	var ()
	return &ListSwiftAccountsParams{

		timeout: timeout,
	}
}

// NewListSwiftAccountsParamsWithContext creates a new ListSwiftAccountsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSwiftAccountsParamsWithContext(ctx context.Context) *ListSwiftAccountsParams {
	var ()
	return &ListSwiftAccountsParams{

		Context: ctx,
	}
}

// NewListSwiftAccountsParamsWithHTTPClient creates a new ListSwiftAccountsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSwiftAccountsParamsWithHTTPClient(client *http.Client) *ListSwiftAccountsParams {
	var ()
	return &ListSwiftAccountsParams{
		HTTPClient: client,
	}
}

/*ListSwiftAccountsParams contains all the parameters to send to the API endpoint
for the list swift accounts operation typically these are written to a http.Request
*/
type ListSwiftAccountsParams struct {

	/*Zone
	  Access zone which contains Swift accounts.

	*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list swift accounts params
func (o *ListSwiftAccountsParams) WithTimeout(timeout time.Duration) *ListSwiftAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list swift accounts params
func (o *ListSwiftAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list swift accounts params
func (o *ListSwiftAccountsParams) WithContext(ctx context.Context) *ListSwiftAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list swift accounts params
func (o *ListSwiftAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list swift accounts params
func (o *ListSwiftAccountsParams) WithHTTPClient(client *http.Client) *ListSwiftAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list swift accounts params
func (o *ListSwiftAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithZone adds the zone to the list swift accounts params
func (o *ListSwiftAccountsParams) WithZone(zone *string) *ListSwiftAccountsParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the list swift accounts params
func (o *ListSwiftAccountsParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *ListSwiftAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Zone != nil {

		// query param zone
		var qrZone string
		if o.Zone != nil {
			qrZone = *o.Zone
		}
		qZone := qrZone
		if qZone != "" {
			if err := r.SetQueryParam("zone", qZone); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
