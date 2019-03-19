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

// NewGetSmbSharesSummaryParams creates a new GetSmbSharesSummaryParams object
// with the default values initialized.
func NewGetSmbSharesSummaryParams() *GetSmbSharesSummaryParams {
	var ()
	return &GetSmbSharesSummaryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSmbSharesSummaryParamsWithTimeout creates a new GetSmbSharesSummaryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSmbSharesSummaryParamsWithTimeout(timeout time.Duration) *GetSmbSharesSummaryParams {
	var ()
	return &GetSmbSharesSummaryParams{

		timeout: timeout,
	}
}

// NewGetSmbSharesSummaryParamsWithContext creates a new GetSmbSharesSummaryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSmbSharesSummaryParamsWithContext(ctx context.Context) *GetSmbSharesSummaryParams {
	var ()
	return &GetSmbSharesSummaryParams{

		Context: ctx,
	}
}

// NewGetSmbSharesSummaryParamsWithHTTPClient creates a new GetSmbSharesSummaryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSmbSharesSummaryParamsWithHTTPClient(client *http.Client) *GetSmbSharesSummaryParams {
	var ()
	return &GetSmbSharesSummaryParams{
		HTTPClient: client,
	}
}

/*GetSmbSharesSummaryParams contains all the parameters to send to the API endpoint
for the get smb shares summary operation typically these are written to a http.Request
*/
type GetSmbSharesSummaryParams struct {

	/*Zone
	  Specifies which access zone or zones to use.

	*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get smb shares summary params
func (o *GetSmbSharesSummaryParams) WithTimeout(timeout time.Duration) *GetSmbSharesSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get smb shares summary params
func (o *GetSmbSharesSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get smb shares summary params
func (o *GetSmbSharesSummaryParams) WithContext(ctx context.Context) *GetSmbSharesSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get smb shares summary params
func (o *GetSmbSharesSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get smb shares summary params
func (o *GetSmbSharesSummaryParams) WithHTTPClient(client *http.Client) *GetSmbSharesSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get smb shares summary params
func (o *GetSmbSharesSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithZone adds the zone to the get smb shares summary params
func (o *GetSmbSharesSummaryParams) WithZone(zone *string) *GetSmbSharesSummaryParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the get smb shares summary params
func (o *GetSmbSharesSummaryParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *GetSmbSharesSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
