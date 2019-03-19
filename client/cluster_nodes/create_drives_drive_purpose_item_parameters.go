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

// NewCreateDrivesDrivePurposeItemParams creates a new CreateDrivesDrivePurposeItemParams object
// with the default values initialized.
func NewCreateDrivesDrivePurposeItemParams() *CreateDrivesDrivePurposeItemParams {
	var ()
	return &CreateDrivesDrivePurposeItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDrivesDrivePurposeItemParamsWithTimeout creates a new CreateDrivesDrivePurposeItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateDrivesDrivePurposeItemParamsWithTimeout(timeout time.Duration) *CreateDrivesDrivePurposeItemParams {
	var ()
	return &CreateDrivesDrivePurposeItemParams{

		timeout: timeout,
	}
}

// NewCreateDrivesDrivePurposeItemParamsWithContext creates a new CreateDrivesDrivePurposeItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateDrivesDrivePurposeItemParamsWithContext(ctx context.Context) *CreateDrivesDrivePurposeItemParams {
	var ()
	return &CreateDrivesDrivePurposeItemParams{

		Context: ctx,
	}
}

// NewCreateDrivesDrivePurposeItemParamsWithHTTPClient creates a new CreateDrivesDrivePurposeItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateDrivesDrivePurposeItemParamsWithHTTPClient(client *http.Client) *CreateDrivesDrivePurposeItemParams {
	var ()
	return &CreateDrivesDrivePurposeItemParams{
		HTTPClient: client,
	}
}

/*CreateDrivesDrivePurposeItemParams contains all the parameters to send to the API endpoint
for the create drives drive purpose item operation typically these are written to a http.Request
*/
type CreateDrivesDrivePurposeItemParams struct {

	/*Driveid*/
	Driveid string
	/*DrivesDrivePurposeItem*/
	DrivesDrivePurposeItem *models.DrivesDrivePurposeItem
	/*Lnn*/
	Lnn int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create drives drive purpose item params
func (o *CreateDrivesDrivePurposeItemParams) WithTimeout(timeout time.Duration) *CreateDrivesDrivePurposeItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create drives drive purpose item params
func (o *CreateDrivesDrivePurposeItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create drives drive purpose item params
func (o *CreateDrivesDrivePurposeItemParams) WithContext(ctx context.Context) *CreateDrivesDrivePurposeItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create drives drive purpose item params
func (o *CreateDrivesDrivePurposeItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create drives drive purpose item params
func (o *CreateDrivesDrivePurposeItemParams) WithHTTPClient(client *http.Client) *CreateDrivesDrivePurposeItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create drives drive purpose item params
func (o *CreateDrivesDrivePurposeItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDriveid adds the driveid to the create drives drive purpose item params
func (o *CreateDrivesDrivePurposeItemParams) WithDriveid(driveid string) *CreateDrivesDrivePurposeItemParams {
	o.SetDriveid(driveid)
	return o
}

// SetDriveid adds the driveid to the create drives drive purpose item params
func (o *CreateDrivesDrivePurposeItemParams) SetDriveid(driveid string) {
	o.Driveid = driveid
}

// WithDrivesDrivePurposeItem adds the drivesDrivePurposeItem to the create drives drive purpose item params
func (o *CreateDrivesDrivePurposeItemParams) WithDrivesDrivePurposeItem(drivesDrivePurposeItem *models.DrivesDrivePurposeItem) *CreateDrivesDrivePurposeItemParams {
	o.SetDrivesDrivePurposeItem(drivesDrivePurposeItem)
	return o
}

// SetDrivesDrivePurposeItem adds the drivesDrivePurposeItem to the create drives drive purpose item params
func (o *CreateDrivesDrivePurposeItemParams) SetDrivesDrivePurposeItem(drivesDrivePurposeItem *models.DrivesDrivePurposeItem) {
	o.DrivesDrivePurposeItem = drivesDrivePurposeItem
}

// WithLnn adds the lnn to the create drives drive purpose item params
func (o *CreateDrivesDrivePurposeItemParams) WithLnn(lnn int64) *CreateDrivesDrivePurposeItemParams {
	o.SetLnn(lnn)
	return o
}

// SetLnn adds the lnn to the create drives drive purpose item params
func (o *CreateDrivesDrivePurposeItemParams) SetLnn(lnn int64) {
	o.Lnn = lnn
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDrivesDrivePurposeItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Driveid
	if err := r.SetPathParam("Driveid", o.Driveid); err != nil {
		return err
	}

	if o.DrivesDrivePurposeItem != nil {
		if err := r.SetBodyParam(o.DrivesDrivePurposeItem); err != nil {
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
