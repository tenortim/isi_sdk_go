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

// NewDeleteMappingIdentityParams creates a new DeleteMappingIdentityParams object
// with the default values initialized.
func NewDeleteMappingIdentityParams() *DeleteMappingIdentityParams {
	var ()
	return &DeleteMappingIdentityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMappingIdentityParamsWithTimeout creates a new DeleteMappingIdentityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMappingIdentityParamsWithTimeout(timeout time.Duration) *DeleteMappingIdentityParams {
	var ()
	return &DeleteMappingIdentityParams{

		timeout: timeout,
	}
}

// NewDeleteMappingIdentityParamsWithContext creates a new DeleteMappingIdentityParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMappingIdentityParamsWithContext(ctx context.Context) *DeleteMappingIdentityParams {
	var ()
	return &DeleteMappingIdentityParams{

		Context: ctx,
	}
}

// NewDeleteMappingIdentityParamsWithHTTPClient creates a new DeleteMappingIdentityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMappingIdentityParamsWithHTTPClient(client *http.Client) *DeleteMappingIdentityParams {
	var ()
	return &DeleteMappingIdentityParams{
		HTTPClient: client,
	}
}

/*DeleteMappingIdentityParams contains all the parameters to send to the API endpoint
for the delete mapping identity operation typically these are written to a http.Request
*/
type DeleteMappingIdentityParams struct {

	/*Nr2way
	  Delete the bi-directional mapping from source to target and target to source.

	*/
	X2way *bool
	/*MappingIdentityID
	  Flush the entire idmap cache.

	*/
	MappingIdentityID string
	/*Remove
	  Delete mapping instead of flush mapping from cache.

	*/
	Remove *bool
	/*Target
	  Target identity persona.

	*/
	Target *string
	/*Zone
	  Optional zone.

	*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete mapping identity params
func (o *DeleteMappingIdentityParams) WithTimeout(timeout time.Duration) *DeleteMappingIdentityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete mapping identity params
func (o *DeleteMappingIdentityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete mapping identity params
func (o *DeleteMappingIdentityParams) WithContext(ctx context.Context) *DeleteMappingIdentityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete mapping identity params
func (o *DeleteMappingIdentityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete mapping identity params
func (o *DeleteMappingIdentityParams) WithHTTPClient(client *http.Client) *DeleteMappingIdentityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete mapping identity params
func (o *DeleteMappingIdentityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithX2way adds the x2way to the delete mapping identity params
func (o *DeleteMappingIdentityParams) WithX2way(x2way *bool) *DeleteMappingIdentityParams {
	o.SetX2way(x2way)
	return o
}

// SetX2way adds the 2way to the delete mapping identity params
func (o *DeleteMappingIdentityParams) SetX2way(x2way *bool) {
	o.X2way = x2way
}

// WithMappingIdentityID adds the mappingIdentityID to the delete mapping identity params
func (o *DeleteMappingIdentityParams) WithMappingIdentityID(mappingIdentityID string) *DeleteMappingIdentityParams {
	o.SetMappingIdentityID(mappingIdentityID)
	return o
}

// SetMappingIdentityID adds the mappingIdentityId to the delete mapping identity params
func (o *DeleteMappingIdentityParams) SetMappingIdentityID(mappingIdentityID string) {
	o.MappingIdentityID = mappingIdentityID
}

// WithRemove adds the remove to the delete mapping identity params
func (o *DeleteMappingIdentityParams) WithRemove(remove *bool) *DeleteMappingIdentityParams {
	o.SetRemove(remove)
	return o
}

// SetRemove adds the remove to the delete mapping identity params
func (o *DeleteMappingIdentityParams) SetRemove(remove *bool) {
	o.Remove = remove
}

// WithTarget adds the target to the delete mapping identity params
func (o *DeleteMappingIdentityParams) WithTarget(target *string) *DeleteMappingIdentityParams {
	o.SetTarget(target)
	return o
}

// SetTarget adds the target to the delete mapping identity params
func (o *DeleteMappingIdentityParams) SetTarget(target *string) {
	o.Target = target
}

// WithZone adds the zone to the delete mapping identity params
func (o *DeleteMappingIdentityParams) WithZone(zone *string) *DeleteMappingIdentityParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the delete mapping identity params
func (o *DeleteMappingIdentityParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMappingIdentityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.X2way != nil {

		// query param 2way
		var qrNr2way bool
		if o.X2way != nil {
			qrNr2way = *o.X2way
		}
		qNr2way := swag.FormatBool(qrNr2way)
		if qNr2way != "" {
			if err := r.SetQueryParam("2way", qNr2way); err != nil {
				return err
			}
		}

	}

	// path param MappingIdentityId
	if err := r.SetPathParam("MappingIdentityId", o.MappingIdentityID); err != nil {
		return err
	}

	if o.Remove != nil {

		// query param remove
		var qrRemove bool
		if o.Remove != nil {
			qrRemove = *o.Remove
		}
		qRemove := swag.FormatBool(qrRemove)
		if qRemove != "" {
			if err := r.SetQueryParam("remove", qRemove); err != nil {
				return err
			}
		}

	}

	if o.Target != nil {

		// query param target
		var qrTarget string
		if o.Target != nil {
			qrTarget = *o.Target
		}
		qTarget := qrTarget
		if qTarget != "" {
			if err := r.SetQueryParam("target", qTarget); err != nil {
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
