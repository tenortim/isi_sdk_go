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

// NewGetMappingIdentityParams creates a new GetMappingIdentityParams object
// with the default values initialized.
func NewGetMappingIdentityParams() *GetMappingIdentityParams {
	var ()
	return &GetMappingIdentityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMappingIdentityParamsWithTimeout creates a new GetMappingIdentityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMappingIdentityParamsWithTimeout(timeout time.Duration) *GetMappingIdentityParams {
	var ()
	return &GetMappingIdentityParams{

		timeout: timeout,
	}
}

// NewGetMappingIdentityParamsWithContext creates a new GetMappingIdentityParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMappingIdentityParamsWithContext(ctx context.Context) *GetMappingIdentityParams {
	var ()
	return &GetMappingIdentityParams{

		Context: ctx,
	}
}

// NewGetMappingIdentityParamsWithHTTPClient creates a new GetMappingIdentityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMappingIdentityParamsWithHTTPClient(client *http.Client) *GetMappingIdentityParams {
	var ()
	return &GetMappingIdentityParams{
		HTTPClient: client,
	}
}

/*GetMappingIdentityParams contains all the parameters to send to the API endpoint
for the get mapping identity operation typically these are written to a http.Request
*/
type GetMappingIdentityParams struct {

	/*MappingIdentityID
	  Retrieve all identity mappings (uid, gid, sid, and on-disk) for the supplied source persona.

	*/
	MappingIdentityID string
	/*Nocreate
	  Idmap should attempt to create missing identity mappings.

	*/
	Nocreate *bool
	/*Zone
	  Optional zone.

	*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get mapping identity params
func (o *GetMappingIdentityParams) WithTimeout(timeout time.Duration) *GetMappingIdentityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get mapping identity params
func (o *GetMappingIdentityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get mapping identity params
func (o *GetMappingIdentityParams) WithContext(ctx context.Context) *GetMappingIdentityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get mapping identity params
func (o *GetMappingIdentityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get mapping identity params
func (o *GetMappingIdentityParams) WithHTTPClient(client *http.Client) *GetMappingIdentityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get mapping identity params
func (o *GetMappingIdentityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMappingIdentityID adds the mappingIdentityID to the get mapping identity params
func (o *GetMappingIdentityParams) WithMappingIdentityID(mappingIdentityID string) *GetMappingIdentityParams {
	o.SetMappingIdentityID(mappingIdentityID)
	return o
}

// SetMappingIdentityID adds the mappingIdentityId to the get mapping identity params
func (o *GetMappingIdentityParams) SetMappingIdentityID(mappingIdentityID string) {
	o.MappingIdentityID = mappingIdentityID
}

// WithNocreate adds the nocreate to the get mapping identity params
func (o *GetMappingIdentityParams) WithNocreate(nocreate *bool) *GetMappingIdentityParams {
	o.SetNocreate(nocreate)
	return o
}

// SetNocreate adds the nocreate to the get mapping identity params
func (o *GetMappingIdentityParams) SetNocreate(nocreate *bool) {
	o.Nocreate = nocreate
}

// WithZone adds the zone to the get mapping identity params
func (o *GetMappingIdentityParams) WithZone(zone *string) *GetMappingIdentityParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the get mapping identity params
func (o *GetMappingIdentityParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *GetMappingIdentityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param MappingIdentityId
	if err := r.SetPathParam("MappingIdentityId", o.MappingIdentityID); err != nil {
		return err
	}

	if o.Nocreate != nil {

		// query param nocreate
		var qrNocreate bool
		if o.Nocreate != nil {
			qrNocreate = *o.Nocreate
		}
		qNocreate := swag.FormatBool(qrNocreate)
		if qNocreate != "" {
			if err := r.SetQueryParam("nocreate", qNocreate); err != nil {
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