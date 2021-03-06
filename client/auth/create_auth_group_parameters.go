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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// NewCreateAuthGroupParams creates a new CreateAuthGroupParams object
// with the default values initialized.
func NewCreateAuthGroupParams() *CreateAuthGroupParams {
	var ()
	return &CreateAuthGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAuthGroupParamsWithTimeout creates a new CreateAuthGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAuthGroupParamsWithTimeout(timeout time.Duration) *CreateAuthGroupParams {
	var ()
	return &CreateAuthGroupParams{

		timeout: timeout,
	}
}

// NewCreateAuthGroupParamsWithContext creates a new CreateAuthGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAuthGroupParamsWithContext(ctx context.Context) *CreateAuthGroupParams {
	var ()
	return &CreateAuthGroupParams{

		Context: ctx,
	}
}

// NewCreateAuthGroupParamsWithHTTPClient creates a new CreateAuthGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAuthGroupParamsWithHTTPClient(client *http.Client) *CreateAuthGroupParams {
	var ()
	return &CreateAuthGroupParams{
		HTTPClient: client,
	}
}

/*CreateAuthGroupParams contains all the parameters to send to the API endpoint
for the create auth group operation typically these are written to a http.Request
*/
type CreateAuthGroupParams struct {

	/*AuthGroup*/
	AuthGroup *models.AuthGroupCreateParams
	/*Force
	  Skip validation checks when creating a group.

	*/
	Force *bool
	/*Provider
	  Optional provider type.

	*/
	Provider *string
	/*Zone
	  Optional zone.

	*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create auth group params
func (o *CreateAuthGroupParams) WithTimeout(timeout time.Duration) *CreateAuthGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create auth group params
func (o *CreateAuthGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create auth group params
func (o *CreateAuthGroupParams) WithContext(ctx context.Context) *CreateAuthGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create auth group params
func (o *CreateAuthGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create auth group params
func (o *CreateAuthGroupParams) WithHTTPClient(client *http.Client) *CreateAuthGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create auth group params
func (o *CreateAuthGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthGroup adds the authGroup to the create auth group params
func (o *CreateAuthGroupParams) WithAuthGroup(authGroup *models.AuthGroupCreateParams) *CreateAuthGroupParams {
	o.SetAuthGroup(authGroup)
	return o
}

// SetAuthGroup adds the authGroup to the create auth group params
func (o *CreateAuthGroupParams) SetAuthGroup(authGroup *models.AuthGroupCreateParams) {
	o.AuthGroup = authGroup
}

// WithForce adds the force to the create auth group params
func (o *CreateAuthGroupParams) WithForce(force *bool) *CreateAuthGroupParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the create auth group params
func (o *CreateAuthGroupParams) SetForce(force *bool) {
	o.Force = force
}

// WithProvider adds the provider to the create auth group params
func (o *CreateAuthGroupParams) WithProvider(provider *string) *CreateAuthGroupParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the create auth group params
func (o *CreateAuthGroupParams) SetProvider(provider *string) {
	o.Provider = provider
}

// WithZone adds the zone to the create auth group params
func (o *CreateAuthGroupParams) WithZone(zone *string) *CreateAuthGroupParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the create auth group params
func (o *CreateAuthGroupParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAuthGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AuthGroup != nil {
		if err := r.SetBodyParam(o.AuthGroup); err != nil {
			return err
		}
	}

	if o.Force != nil {

		// query param force
		var qrForce bool
		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {
			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}

	}

	if o.Provider != nil {

		// query param provider
		var qrProvider string
		if o.Provider != nil {
			qrProvider = *o.Provider
		}
		qProvider := qrProvider
		if qProvider != "" {
			if err := r.SetQueryParam("provider", qProvider); err != nil {
				return err
			}
		}

	}

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
