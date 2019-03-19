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

	models "github.com/tenortim/isi_sdk_go/models"
)

// NewCreateDrivesDriveSuspendItemParams creates a new CreateDrivesDriveSuspendItemParams object
// with the default values initialized.
func NewCreateDrivesDriveSuspendItemParams() *CreateDrivesDriveSuspendItemParams {
	var ()
	return &CreateDrivesDriveSuspendItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDrivesDriveSuspendItemParamsWithTimeout creates a new CreateDrivesDriveSuspendItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateDrivesDriveSuspendItemParamsWithTimeout(timeout time.Duration) *CreateDrivesDriveSuspendItemParams {
	var ()
	return &CreateDrivesDriveSuspendItemParams{

		timeout: timeout,
	}
}

// NewCreateDrivesDriveSuspendItemParamsWithContext creates a new CreateDrivesDriveSuspendItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateDrivesDriveSuspendItemParamsWithContext(ctx context.Context) *CreateDrivesDriveSuspendItemParams {
	var ()
	return &CreateDrivesDriveSuspendItemParams{

		Context: ctx,
	}
}

// NewCreateDrivesDriveSuspendItemParamsWithHTTPClient creates a new CreateDrivesDriveSuspendItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateDrivesDriveSuspendItemParamsWithHTTPClient(client *http.Client) *CreateDrivesDriveSuspendItemParams {
	var ()
	return &CreateDrivesDriveSuspendItemParams{
		HTTPClient: client,
	}
}

/*CreateDrivesDriveSuspendItemParams contains all the parameters to send to the API endpoint
for the create drives drive suspend item operation typically these are written to a http.Request
*/
type CreateDrivesDriveSuspendItemParams struct {

	/*Driveid*/
	Driveid string
	/*DrivesDriveSuspendItem*/
	DrivesDriveSuspendItem models.Empty
	/*Lnn*/
	Lnn int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create drives drive suspend item params
func (o *CreateDrivesDriveSuspendItemParams) WithTimeout(timeout time.Duration) *CreateDrivesDriveSuspendItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create drives drive suspend item params
func (o *CreateDrivesDriveSuspendItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create drives drive suspend item params
func (o *CreateDrivesDriveSuspendItemParams) WithContext(ctx context.Context) *CreateDrivesDriveSuspendItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create drives drive suspend item params
func (o *CreateDrivesDriveSuspendItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create drives drive suspend item params
func (o *CreateDrivesDriveSuspendItemParams) WithHTTPClient(client *http.Client) *CreateDrivesDriveSuspendItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create drives drive suspend item params
func (o *CreateDrivesDriveSuspendItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDriveid adds the driveid to the create drives drive suspend item params
func (o *CreateDrivesDriveSuspendItemParams) WithDriveid(driveid string) *CreateDrivesDriveSuspendItemParams {
	o.SetDriveid(driveid)
	return o
}

// SetDriveid adds the driveid to the create drives drive suspend item params
func (o *CreateDrivesDriveSuspendItemParams) SetDriveid(driveid string) {
	o.Driveid = driveid
}

// WithDrivesDriveSuspendItem adds the drivesDriveSuspendItem to the create drives drive suspend item params
func (o *CreateDrivesDriveSuspendItemParams) WithDrivesDriveSuspendItem(drivesDriveSuspendItem models.Empty) *CreateDrivesDriveSuspendItemParams {
	o.SetDrivesDriveSuspendItem(drivesDriveSuspendItem)
	return o
}

// SetDrivesDriveSuspendItem adds the drivesDriveSuspendItem to the create drives drive suspend item params
func (o *CreateDrivesDriveSuspendItemParams) SetDrivesDriveSuspendItem(drivesDriveSuspendItem models.Empty) {
	o.DrivesDriveSuspendItem = drivesDriveSuspendItem
}

// WithLnn adds the lnn to the create drives drive suspend item params
func (o *CreateDrivesDriveSuspendItemParams) WithLnn(lnn int64) *CreateDrivesDriveSuspendItemParams {
	o.SetLnn(lnn)
	return o
}

// SetLnn adds the lnn to the create drives drive suspend item params
func (o *CreateDrivesDriveSuspendItemParams) SetLnn(lnn int64) {
	o.Lnn = lnn
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDrivesDriveSuspendItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Driveid
	if err := r.SetPathParam("Driveid", o.Driveid); err != nil {
		return err
	}

	if o.DrivesDriveSuspendItem != nil {
		if err := r.SetBodyParam(o.DrivesDriveSuspendItem); err != nil {
			return err
		}
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
