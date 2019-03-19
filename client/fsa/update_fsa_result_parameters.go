// Code generated by go-swagger; DO NOT EDIT.

package fsa

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

// NewUpdateFsaResultParams creates a new UpdateFsaResultParams object
// with the default values initialized.
func NewUpdateFsaResultParams() *UpdateFsaResultParams {
	var ()
	return &UpdateFsaResultParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateFsaResultParamsWithTimeout creates a new UpdateFsaResultParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateFsaResultParamsWithTimeout(timeout time.Duration) *UpdateFsaResultParams {
	var ()
	return &UpdateFsaResultParams{

		timeout: timeout,
	}
}

// NewUpdateFsaResultParamsWithContext creates a new UpdateFsaResultParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateFsaResultParamsWithContext(ctx context.Context) *UpdateFsaResultParams {
	var ()
	return &UpdateFsaResultParams{

		Context: ctx,
	}
}

// NewUpdateFsaResultParamsWithHTTPClient creates a new UpdateFsaResultParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateFsaResultParamsWithHTTPClient(client *http.Client) *UpdateFsaResultParams {
	var ()
	return &UpdateFsaResultParams{
		HTTPClient: client,
	}
}

/*UpdateFsaResultParams contains all the parameters to send to the API endpoint
for the update fsa result operation typically these are written to a http.Request
*/
type UpdateFsaResultParams struct {

	/*FsaResult*/
	FsaResult *models.FsaResult
	/*FsaResultID
	  Modify result set. Only the pinned property can be changed at this time.

	*/
	FsaResultID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update fsa result params
func (o *UpdateFsaResultParams) WithTimeout(timeout time.Duration) *UpdateFsaResultParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update fsa result params
func (o *UpdateFsaResultParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update fsa result params
func (o *UpdateFsaResultParams) WithContext(ctx context.Context) *UpdateFsaResultParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update fsa result params
func (o *UpdateFsaResultParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update fsa result params
func (o *UpdateFsaResultParams) WithHTTPClient(client *http.Client) *UpdateFsaResultParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update fsa result params
func (o *UpdateFsaResultParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFsaResult adds the fsaResult to the update fsa result params
func (o *UpdateFsaResultParams) WithFsaResult(fsaResult *models.FsaResult) *UpdateFsaResultParams {
	o.SetFsaResult(fsaResult)
	return o
}

// SetFsaResult adds the fsaResult to the update fsa result params
func (o *UpdateFsaResultParams) SetFsaResult(fsaResult *models.FsaResult) {
	o.FsaResult = fsaResult
}

// WithFsaResultID adds the fsaResultID to the update fsa result params
func (o *UpdateFsaResultParams) WithFsaResultID(fsaResultID string) *UpdateFsaResultParams {
	o.SetFsaResultID(fsaResultID)
	return o
}

// SetFsaResultID adds the fsaResultId to the update fsa result params
func (o *UpdateFsaResultParams) SetFsaResultID(fsaResultID string) {
	o.FsaResultID = fsaResultID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateFsaResultParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FsaResult != nil {
		if err := r.SetBodyParam(o.FsaResult); err != nil {
			return err
		}
	}

	// path param FsaResultId
	if err := r.SetPathParam("FsaResultId", o.FsaResultID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}