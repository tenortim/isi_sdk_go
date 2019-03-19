// Code generated by go-swagger; DO NOT EDIT.

package upgrade_cluster

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

// NewGetNodesNodeFirmwareStatusParams creates a new GetNodesNodeFirmwareStatusParams object
// with the default values initialized.
func NewGetNodesNodeFirmwareStatusParams() *GetNodesNodeFirmwareStatusParams {
	var ()
	return &GetNodesNodeFirmwareStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNodesNodeFirmwareStatusParamsWithTimeout creates a new GetNodesNodeFirmwareStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNodesNodeFirmwareStatusParamsWithTimeout(timeout time.Duration) *GetNodesNodeFirmwareStatusParams {
	var ()
	return &GetNodesNodeFirmwareStatusParams{

		timeout: timeout,
	}
}

// NewGetNodesNodeFirmwareStatusParamsWithContext creates a new GetNodesNodeFirmwareStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNodesNodeFirmwareStatusParamsWithContext(ctx context.Context) *GetNodesNodeFirmwareStatusParams {
	var ()
	return &GetNodesNodeFirmwareStatusParams{

		Context: ctx,
	}
}

// NewGetNodesNodeFirmwareStatusParamsWithHTTPClient creates a new GetNodesNodeFirmwareStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNodesNodeFirmwareStatusParamsWithHTTPClient(client *http.Client) *GetNodesNodeFirmwareStatusParams {
	var ()
	return &GetNodesNodeFirmwareStatusParams{
		HTTPClient: client,
	}
}

/*GetNodesNodeFirmwareStatusParams contains all the parameters to send to the API endpoint
for the get nodes node firmware status operation typically these are written to a http.Request
*/
type GetNodesNodeFirmwareStatusParams struct {

	/*Lnn*/
	Lnn int64
	/*Devices
	  Show devices. If false, this returns an empty list. Default is false.

	*/
	Devices *bool
	/*Package
	  Show package. If false, this returns an empty list.Default is false.

	*/
	Package *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get nodes node firmware status params
func (o *GetNodesNodeFirmwareStatusParams) WithTimeout(timeout time.Duration) *GetNodesNodeFirmwareStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nodes node firmware status params
func (o *GetNodesNodeFirmwareStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nodes node firmware status params
func (o *GetNodesNodeFirmwareStatusParams) WithContext(ctx context.Context) *GetNodesNodeFirmwareStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nodes node firmware status params
func (o *GetNodesNodeFirmwareStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nodes node firmware status params
func (o *GetNodesNodeFirmwareStatusParams) WithHTTPClient(client *http.Client) *GetNodesNodeFirmwareStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nodes node firmware status params
func (o *GetNodesNodeFirmwareStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLnn adds the lnn to the get nodes node firmware status params
func (o *GetNodesNodeFirmwareStatusParams) WithLnn(lnn int64) *GetNodesNodeFirmwareStatusParams {
	o.SetLnn(lnn)
	return o
}

// SetLnn adds the lnn to the get nodes node firmware status params
func (o *GetNodesNodeFirmwareStatusParams) SetLnn(lnn int64) {
	o.Lnn = lnn
}

// WithDevices adds the devices to the get nodes node firmware status params
func (o *GetNodesNodeFirmwareStatusParams) WithDevices(devices *bool) *GetNodesNodeFirmwareStatusParams {
	o.SetDevices(devices)
	return o
}

// SetDevices adds the devices to the get nodes node firmware status params
func (o *GetNodesNodeFirmwareStatusParams) SetDevices(devices *bool) {
	o.Devices = devices
}

// WithPackage adds the packageVar to the get nodes node firmware status params
func (o *GetNodesNodeFirmwareStatusParams) WithPackage(packageVar *bool) *GetNodesNodeFirmwareStatusParams {
	o.SetPackage(packageVar)
	return o
}

// SetPackage adds the package to the get nodes node firmware status params
func (o *GetNodesNodeFirmwareStatusParams) SetPackage(packageVar *bool) {
	o.Package = packageVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetNodesNodeFirmwareStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Lnn
	if err := r.SetPathParam("Lnn", swag.FormatInt64(o.Lnn)); err != nil {
		return err
	}

	if o.Devices != nil {

		// query param devices
		var qrDevices bool
		if o.Devices != nil {
			qrDevices = *o.Devices
		}
		qDevices := swag.FormatBool(qrDevices)
		if qDevices != "" {
			if err := r.SetQueryParam("devices", qDevices); err != nil {
				return err
			}
		}

	}

	if o.Package != nil {

		// query param package
		var qrPackage bool
		if o.Package != nil {
			qrPackage = *o.Package
		}
		qPackage := swag.FormatBool(qrPackage)
		if qPackage != "" {
			if err := r.SetQueryParam("package", qPackage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}