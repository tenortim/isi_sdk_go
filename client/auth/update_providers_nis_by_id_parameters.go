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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// NewUpdateProvidersNisByIDParams creates a new UpdateProvidersNisByIDParams object
// with the default values initialized.
func NewUpdateProvidersNisByIDParams() *UpdateProvidersNisByIDParams {
	var ()
	return &UpdateProvidersNisByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateProvidersNisByIDParamsWithTimeout creates a new UpdateProvidersNisByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateProvidersNisByIDParamsWithTimeout(timeout time.Duration) *UpdateProvidersNisByIDParams {
	var ()
	return &UpdateProvidersNisByIDParams{

		timeout: timeout,
	}
}

// NewUpdateProvidersNisByIDParamsWithContext creates a new UpdateProvidersNisByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateProvidersNisByIDParamsWithContext(ctx context.Context) *UpdateProvidersNisByIDParams {
	var ()
	return &UpdateProvidersNisByIDParams{

		Context: ctx,
	}
}

// NewUpdateProvidersNisByIDParamsWithHTTPClient creates a new UpdateProvidersNisByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateProvidersNisByIDParamsWithHTTPClient(client *http.Client) *UpdateProvidersNisByIDParams {
	var ()
	return &UpdateProvidersNisByIDParams{
		HTTPClient: client,
	}
}

/*UpdateProvidersNisByIDParams contains all the parameters to send to the API endpoint
for the update providers nis by Id operation typically these are written to a http.Request
*/
type UpdateProvidersNisByIDParams struct {

	/*ProvidersNisID
	  Modify the NIS provider.

	*/
	ProvidersNisID string
	/*ProvidersNisIDParams*/
	ProvidersNisIDParams *models.ProvidersNisIDParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update providers nis by Id params
func (o *UpdateProvidersNisByIDParams) WithTimeout(timeout time.Duration) *UpdateProvidersNisByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update providers nis by Id params
func (o *UpdateProvidersNisByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update providers nis by Id params
func (o *UpdateProvidersNisByIDParams) WithContext(ctx context.Context) *UpdateProvidersNisByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update providers nis by Id params
func (o *UpdateProvidersNisByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update providers nis by Id params
func (o *UpdateProvidersNisByIDParams) WithHTTPClient(client *http.Client) *UpdateProvidersNisByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update providers nis by Id params
func (o *UpdateProvidersNisByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProvidersNisID adds the providersNisID to the update providers nis by Id params
func (o *UpdateProvidersNisByIDParams) WithProvidersNisID(providersNisID string) *UpdateProvidersNisByIDParams {
	o.SetProvidersNisID(providersNisID)
	return o
}

// SetProvidersNisID adds the providersNisId to the update providers nis by Id params
func (o *UpdateProvidersNisByIDParams) SetProvidersNisID(providersNisID string) {
	o.ProvidersNisID = providersNisID
}

// WithProvidersNisIDParams adds the providersNisIDParams to the update providers nis by Id params
func (o *UpdateProvidersNisByIDParams) WithProvidersNisIDParams(providersNisIDParams *models.ProvidersNisIDParams) *UpdateProvidersNisByIDParams {
	o.SetProvidersNisIDParams(providersNisIDParams)
	return o
}

// SetProvidersNisIDParams adds the providersNisIdParams to the update providers nis by Id params
func (o *UpdateProvidersNisByIDParams) SetProvidersNisIDParams(providersNisIDParams *models.ProvidersNisIDParams) {
	o.ProvidersNisIDParams = providersNisIDParams
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateProvidersNisByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ProvidersNisId
	if err := r.SetPathParam("ProvidersNisId", o.ProvidersNisID); err != nil {
		return err
	}

	if o.ProvidersNisIDParams != nil {
		if err := r.SetBodyParam(o.ProvidersNisIDParams); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}