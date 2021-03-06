// Code generated by go-swagger; DO NOT EDIT.

package cluster_nodes

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

// NewListDrivesDriveFirmwareUpdateParams creates a new ListDrivesDriveFirmwareUpdateParams object
// with the default values initialized.
func NewListDrivesDriveFirmwareUpdateParams() *ListDrivesDriveFirmwareUpdateParams {
	var ()
	return &ListDrivesDriveFirmwareUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListDrivesDriveFirmwareUpdateParamsWithTimeout creates a new ListDrivesDriveFirmwareUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListDrivesDriveFirmwareUpdateParamsWithTimeout(timeout time.Duration) *ListDrivesDriveFirmwareUpdateParams {
	var ()
	return &ListDrivesDriveFirmwareUpdateParams{

		timeout: timeout,
	}
}

// NewListDrivesDriveFirmwareUpdateParamsWithContext creates a new ListDrivesDriveFirmwareUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewListDrivesDriveFirmwareUpdateParamsWithContext(ctx context.Context) *ListDrivesDriveFirmwareUpdateParams {
	var ()
	return &ListDrivesDriveFirmwareUpdateParams{

		Context: ctx,
	}
}

// NewListDrivesDriveFirmwareUpdateParamsWithHTTPClient creates a new ListDrivesDriveFirmwareUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListDrivesDriveFirmwareUpdateParamsWithHTTPClient(client *http.Client) *ListDrivesDriveFirmwareUpdateParams {
	var ()
	return &ListDrivesDriveFirmwareUpdateParams{
		HTTPClient: client,
	}
}

/*ListDrivesDriveFirmwareUpdateParams contains all the parameters to send to the API endpoint
for the list drives drive firmware update operation typically these are written to a http.Request
*/
type ListDrivesDriveFirmwareUpdateParams struct {

	/*Driveid*/
	Driveid string
	/*Lnn*/
	Lnn int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list drives drive firmware update params
func (o *ListDrivesDriveFirmwareUpdateParams) WithTimeout(timeout time.Duration) *ListDrivesDriveFirmwareUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list drives drive firmware update params
func (o *ListDrivesDriveFirmwareUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list drives drive firmware update params
func (o *ListDrivesDriveFirmwareUpdateParams) WithContext(ctx context.Context) *ListDrivesDriveFirmwareUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list drives drive firmware update params
func (o *ListDrivesDriveFirmwareUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list drives drive firmware update params
func (o *ListDrivesDriveFirmwareUpdateParams) WithHTTPClient(client *http.Client) *ListDrivesDriveFirmwareUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list drives drive firmware update params
func (o *ListDrivesDriveFirmwareUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDriveid adds the driveid to the list drives drive firmware update params
func (o *ListDrivesDriveFirmwareUpdateParams) WithDriveid(driveid string) *ListDrivesDriveFirmwareUpdateParams {
	o.SetDriveid(driveid)
	return o
}

// SetDriveid adds the driveid to the list drives drive firmware update params
func (o *ListDrivesDriveFirmwareUpdateParams) SetDriveid(driveid string) {
	o.Driveid = driveid
}

// WithLnn adds the lnn to the list drives drive firmware update params
func (o *ListDrivesDriveFirmwareUpdateParams) WithLnn(lnn int64) *ListDrivesDriveFirmwareUpdateParams {
	o.SetLnn(lnn)
	return o
}

// SetLnn adds the lnn to the list drives drive firmware update params
func (o *ListDrivesDriveFirmwareUpdateParams) SetLnn(lnn int64) {
	o.Lnn = lnn
}

// WriteToRequest writes these params to a swagger request
func (o *ListDrivesDriveFirmwareUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Driveid
	if err := r.SetPathParam("Driveid", o.Driveid); err != nil {
		return err
	}

	// path param Lnn
	if err := r.SetPathParam("Lnn", swag.FormatInt64(o.Lnn)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
