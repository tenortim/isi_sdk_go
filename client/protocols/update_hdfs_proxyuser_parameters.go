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

// NewUpdateHdfsProxyuserParams creates a new UpdateHdfsProxyuserParams object
// with the default values initialized.
func NewUpdateHdfsProxyuserParams() *UpdateHdfsProxyuserParams {
	var ()
	return &UpdateHdfsProxyuserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateHdfsProxyuserParamsWithTimeout creates a new UpdateHdfsProxyuserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateHdfsProxyuserParamsWithTimeout(timeout time.Duration) *UpdateHdfsProxyuserParams {
	var ()
	return &UpdateHdfsProxyuserParams{

		timeout: timeout,
	}
}

// NewUpdateHdfsProxyuserParamsWithContext creates a new UpdateHdfsProxyuserParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateHdfsProxyuserParamsWithContext(ctx context.Context) *UpdateHdfsProxyuserParams {
	var ()
	return &UpdateHdfsProxyuserParams{

		Context: ctx,
	}
}

// NewUpdateHdfsProxyuserParamsWithHTTPClient creates a new UpdateHdfsProxyuserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateHdfsProxyuserParamsWithHTTPClient(client *http.Client) *UpdateHdfsProxyuserParams {
	var ()
	return &UpdateHdfsProxyuserParams{
		HTTPClient: client,
	}
}

/*UpdateHdfsProxyuserParams contains all the parameters to send to the API endpoint
for the update hdfs proxyuser operation typically these are written to a http.Request
*/
type UpdateHdfsProxyuserParams struct {

	/*HdfsProxyuser*/
	HdfsProxyuser models.Empty
	/*HdfsProxyuserID
	  Modify an HDFS proxyuser.

	*/
	HdfsProxyuserID string
	/*Zone
	  Access zone which contains HDFS proxyuser.

	*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update hdfs proxyuser params
func (o *UpdateHdfsProxyuserParams) WithTimeout(timeout time.Duration) *UpdateHdfsProxyuserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update hdfs proxyuser params
func (o *UpdateHdfsProxyuserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update hdfs proxyuser params
func (o *UpdateHdfsProxyuserParams) WithContext(ctx context.Context) *UpdateHdfsProxyuserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update hdfs proxyuser params
func (o *UpdateHdfsProxyuserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update hdfs proxyuser params
func (o *UpdateHdfsProxyuserParams) WithHTTPClient(client *http.Client) *UpdateHdfsProxyuserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update hdfs proxyuser params
func (o *UpdateHdfsProxyuserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHdfsProxyuser adds the hdfsProxyuser to the update hdfs proxyuser params
func (o *UpdateHdfsProxyuserParams) WithHdfsProxyuser(hdfsProxyuser models.Empty) *UpdateHdfsProxyuserParams {
	o.SetHdfsProxyuser(hdfsProxyuser)
	return o
}

// SetHdfsProxyuser adds the hdfsProxyuser to the update hdfs proxyuser params
func (o *UpdateHdfsProxyuserParams) SetHdfsProxyuser(hdfsProxyuser models.Empty) {
	o.HdfsProxyuser = hdfsProxyuser
}

// WithHdfsProxyuserID adds the hdfsProxyuserID to the update hdfs proxyuser params
func (o *UpdateHdfsProxyuserParams) WithHdfsProxyuserID(hdfsProxyuserID string) *UpdateHdfsProxyuserParams {
	o.SetHdfsProxyuserID(hdfsProxyuserID)
	return o
}

// SetHdfsProxyuserID adds the hdfsProxyuserId to the update hdfs proxyuser params
func (o *UpdateHdfsProxyuserParams) SetHdfsProxyuserID(hdfsProxyuserID string) {
	o.HdfsProxyuserID = hdfsProxyuserID
}

// WithZone adds the zone to the update hdfs proxyuser params
func (o *UpdateHdfsProxyuserParams) WithZone(zone *string) *UpdateHdfsProxyuserParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the update hdfs proxyuser params
func (o *UpdateHdfsProxyuserParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateHdfsProxyuserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.HdfsProxyuser != nil {
		if err := r.SetBodyParam(o.HdfsProxyuser); err != nil {
			return err
		}
	}

	// path param HdfsProxyuserId
	if err := r.SetPathParam("HdfsProxyuserId", o.HdfsProxyuserID); err != nil {
		return err
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
