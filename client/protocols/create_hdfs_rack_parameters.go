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

	models "github.com/tenortim/isi_sdk_go/models"
)

// NewCreateHdfsRackParams creates a new CreateHdfsRackParams object
// with the default values initialized.
func NewCreateHdfsRackParams() *CreateHdfsRackParams {
	var ()
	return &CreateHdfsRackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateHdfsRackParamsWithTimeout creates a new CreateHdfsRackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateHdfsRackParamsWithTimeout(timeout time.Duration) *CreateHdfsRackParams {
	var ()
	return &CreateHdfsRackParams{

		timeout: timeout,
	}
}

// NewCreateHdfsRackParamsWithContext creates a new CreateHdfsRackParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateHdfsRackParamsWithContext(ctx context.Context) *CreateHdfsRackParams {
	var ()
	return &CreateHdfsRackParams{

		Context: ctx,
	}
}

// NewCreateHdfsRackParamsWithHTTPClient creates a new CreateHdfsRackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateHdfsRackParamsWithHTTPClient(client *http.Client) *CreateHdfsRackParams {
	var ()
	return &CreateHdfsRackParams{
		HTTPClient: client,
	}
}

/*CreateHdfsRackParams contains all the parameters to send to the API endpoint
for the create hdfs rack operation typically these are written to a http.Request
*/
type CreateHdfsRackParams struct {

	/*HdfsRack*/
	HdfsRack *models.HdfsRackCreateParams
	/*Zone
	  Access zone which contains HDFS rack.

	*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create hdfs rack params
func (o *CreateHdfsRackParams) WithTimeout(timeout time.Duration) *CreateHdfsRackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create hdfs rack params
func (o *CreateHdfsRackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create hdfs rack params
func (o *CreateHdfsRackParams) WithContext(ctx context.Context) *CreateHdfsRackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create hdfs rack params
func (o *CreateHdfsRackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create hdfs rack params
func (o *CreateHdfsRackParams) WithHTTPClient(client *http.Client) *CreateHdfsRackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create hdfs rack params
func (o *CreateHdfsRackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHdfsRack adds the hdfsRack to the create hdfs rack params
func (o *CreateHdfsRackParams) WithHdfsRack(hdfsRack *models.HdfsRackCreateParams) *CreateHdfsRackParams {
	o.SetHdfsRack(hdfsRack)
	return o
}

// SetHdfsRack adds the hdfsRack to the create hdfs rack params
func (o *CreateHdfsRackParams) SetHdfsRack(hdfsRack *models.HdfsRackCreateParams) {
	o.HdfsRack = hdfsRack
}

// WithZone adds the zone to the create hdfs rack params
func (o *CreateHdfsRackParams) WithZone(zone *string) *CreateHdfsRackParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the create hdfs rack params
func (o *CreateHdfsRackParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *CreateHdfsRackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.HdfsRack != nil {
		if err := r.SetBodyParam(o.HdfsRack); err != nil {
			return err
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
