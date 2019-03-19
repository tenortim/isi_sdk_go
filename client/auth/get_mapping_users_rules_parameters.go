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

// NewGetMappingUsersRulesParams creates a new GetMappingUsersRulesParams object
// with the default values initialized.
func NewGetMappingUsersRulesParams() *GetMappingUsersRulesParams {
	var ()
	return &GetMappingUsersRulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMappingUsersRulesParamsWithTimeout creates a new GetMappingUsersRulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMappingUsersRulesParamsWithTimeout(timeout time.Duration) *GetMappingUsersRulesParams {
	var ()
	return &GetMappingUsersRulesParams{

		timeout: timeout,
	}
}

// NewGetMappingUsersRulesParamsWithContext creates a new GetMappingUsersRulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMappingUsersRulesParamsWithContext(ctx context.Context) *GetMappingUsersRulesParams {
	var ()
	return &GetMappingUsersRulesParams{

		Context: ctx,
	}
}

// NewGetMappingUsersRulesParamsWithHTTPClient creates a new GetMappingUsersRulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMappingUsersRulesParamsWithHTTPClient(client *http.Client) *GetMappingUsersRulesParams {
	var ()
	return &GetMappingUsersRulesParams{
		HTTPClient: client,
	}
}

/*GetMappingUsersRulesParams contains all the parameters to send to the API endpoint
for the get mapping users rules operation typically these are written to a http.Request
*/
type GetMappingUsersRulesParams struct {

	/*Zone
	  The zone to which the user mapping applies.

	*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get mapping users rules params
func (o *GetMappingUsersRulesParams) WithTimeout(timeout time.Duration) *GetMappingUsersRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get mapping users rules params
func (o *GetMappingUsersRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get mapping users rules params
func (o *GetMappingUsersRulesParams) WithContext(ctx context.Context) *GetMappingUsersRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get mapping users rules params
func (o *GetMappingUsersRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get mapping users rules params
func (o *GetMappingUsersRulesParams) WithHTTPClient(client *http.Client) *GetMappingUsersRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get mapping users rules params
func (o *GetMappingUsersRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithZone adds the zone to the get mapping users rules params
func (o *GetMappingUsersRulesParams) WithZone(zone *string) *GetMappingUsersRulesParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the get mapping users rules params
func (o *GetMappingUsersRulesParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *GetMappingUsersRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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