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

// NewDeleteAuthUserParams creates a new DeleteAuthUserParams object
// with the default values initialized.
func NewDeleteAuthUserParams() *DeleteAuthUserParams {
	var ()
	return &DeleteAuthUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAuthUserParamsWithTimeout creates a new DeleteAuthUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAuthUserParamsWithTimeout(timeout time.Duration) *DeleteAuthUserParams {
	var ()
	return &DeleteAuthUserParams{

		timeout: timeout,
	}
}

// NewDeleteAuthUserParamsWithContext creates a new DeleteAuthUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAuthUserParamsWithContext(ctx context.Context) *DeleteAuthUserParams {
	var ()
	return &DeleteAuthUserParams{

		Context: ctx,
	}
}

// NewDeleteAuthUserParamsWithHTTPClient creates a new DeleteAuthUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAuthUserParamsWithHTTPClient(client *http.Client) *DeleteAuthUserParams {
	var ()
	return &DeleteAuthUserParams{
		HTTPClient: client,
	}
}

/*DeleteAuthUserParams contains all the parameters to send to the API endpoint
for the delete auth user operation typically these are written to a http.Request
*/
type DeleteAuthUserParams struct {

	/*AuthUserID
	  Delete the user.

	*/
	AuthUserID string
	/*Cached
	  If true, flush the user from the cache.

	*/
	Cached *bool
	/*Provider
	  Filter users by provider.

	*/
	Provider *string
	/*Zone
	  Filter users by zone.

	*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete auth user params
func (o *DeleteAuthUserParams) WithTimeout(timeout time.Duration) *DeleteAuthUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete auth user params
func (o *DeleteAuthUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete auth user params
func (o *DeleteAuthUserParams) WithContext(ctx context.Context) *DeleteAuthUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete auth user params
func (o *DeleteAuthUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete auth user params
func (o *DeleteAuthUserParams) WithHTTPClient(client *http.Client) *DeleteAuthUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete auth user params
func (o *DeleteAuthUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthUserID adds the authUserID to the delete auth user params
func (o *DeleteAuthUserParams) WithAuthUserID(authUserID string) *DeleteAuthUserParams {
	o.SetAuthUserID(authUserID)
	return o
}

// SetAuthUserID adds the authUserId to the delete auth user params
func (o *DeleteAuthUserParams) SetAuthUserID(authUserID string) {
	o.AuthUserID = authUserID
}

// WithCached adds the cached to the delete auth user params
func (o *DeleteAuthUserParams) WithCached(cached *bool) *DeleteAuthUserParams {
	o.SetCached(cached)
	return o
}

// SetCached adds the cached to the delete auth user params
func (o *DeleteAuthUserParams) SetCached(cached *bool) {
	o.Cached = cached
}

// WithProvider adds the provider to the delete auth user params
func (o *DeleteAuthUserParams) WithProvider(provider *string) *DeleteAuthUserParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the delete auth user params
func (o *DeleteAuthUserParams) SetProvider(provider *string) {
	o.Provider = provider
}

// WithZone adds the zone to the delete auth user params
func (o *DeleteAuthUserParams) WithZone(zone *string) *DeleteAuthUserParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the delete auth user params
func (o *DeleteAuthUserParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAuthUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param AuthUserId
	if err := r.SetPathParam("AuthUserId", o.AuthUserID); err != nil {
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