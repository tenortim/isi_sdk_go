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
)

// NewDeleteAuthGroupParams creates a new DeleteAuthGroupParams object
// with the default values initialized.
func NewDeleteAuthGroupParams() *DeleteAuthGroupParams {
	var ()
	return &DeleteAuthGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAuthGroupParamsWithTimeout creates a new DeleteAuthGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAuthGroupParamsWithTimeout(timeout time.Duration) *DeleteAuthGroupParams {
	var ()
	return &DeleteAuthGroupParams{

		timeout: timeout,
	}
}

// NewDeleteAuthGroupParamsWithContext creates a new DeleteAuthGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAuthGroupParamsWithContext(ctx context.Context) *DeleteAuthGroupParams {
	var ()
	return &DeleteAuthGroupParams{

		Context: ctx,
	}
}

// NewDeleteAuthGroupParamsWithHTTPClient creates a new DeleteAuthGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAuthGroupParamsWithHTTPClient(client *http.Client) *DeleteAuthGroupParams {
	var ()
	return &DeleteAuthGroupParams{
		HTTPClient: client,
	}
}

/*DeleteAuthGroupParams contains all the parameters to send to the API endpoint
for the delete auth group operation typically these are written to a http.Request
*/
type DeleteAuthGroupParams struct {

	/*AuthGroupID
	  Delete the group.

	*/
	AuthGroupID string
	/*Cached
	  If true, flush the group from the cache.

	*/
	Cached *bool
	/*Provider
	  Filter groups by provider.

	*/
	Provider *string
	/*Zone
	  Filter groups by zone.

	*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete auth group params
func (o *DeleteAuthGroupParams) WithTimeout(timeout time.Duration) *DeleteAuthGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete auth group params
func (o *DeleteAuthGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete auth group params
func (o *DeleteAuthGroupParams) WithContext(ctx context.Context) *DeleteAuthGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete auth group params
func (o *DeleteAuthGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete auth group params
func (o *DeleteAuthGroupParams) WithHTTPClient(client *http.Client) *DeleteAuthGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete auth group params
func (o *DeleteAuthGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthGroupID adds the authGroupID to the delete auth group params
func (o *DeleteAuthGroupParams) WithAuthGroupID(authGroupID string) *DeleteAuthGroupParams {
	o.SetAuthGroupID(authGroupID)
	return o
}

// SetAuthGroupID adds the authGroupId to the delete auth group params
func (o *DeleteAuthGroupParams) SetAuthGroupID(authGroupID string) {
	o.AuthGroupID = authGroupID
}

// WithCached adds the cached to the delete auth group params
func (o *DeleteAuthGroupParams) WithCached(cached *bool) *DeleteAuthGroupParams {
	o.SetCached(cached)
	return o
}

// SetCached adds the cached to the delete auth group params
func (o *DeleteAuthGroupParams) SetCached(cached *bool) {
	o.Cached = cached
}

// WithProvider adds the provider to the delete auth group params
func (o *DeleteAuthGroupParams) WithProvider(provider *string) *DeleteAuthGroupParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the delete auth group params
func (o *DeleteAuthGroupParams) SetProvider(provider *string) {
	o.Provider = provider
}

// WithZone adds the zone to the delete auth group params
func (o *DeleteAuthGroupParams) WithZone(zone *string) *DeleteAuthGroupParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the delete auth group params
func (o *DeleteAuthGroupParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAuthGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param AuthGroupId
	if err := r.SetPathParam("AuthGroupId", o.AuthGroupID); err != nil {
		return err
	}

	if o.Cached != nil {

		// query param cached
		var qrCached bool
		if o.Cached != nil {
			qrCached = *o.Cached
		}
		qCached := swag.FormatBool(qrCached)
		if qCached != "" {
			if err := r.SetQueryParam("cached", qCached); err != nil {
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
