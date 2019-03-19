// Code generated by go-swagger; DO NOT EDIT.

package cloud

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

// NewCreateSettingsReportingEulaItemParams creates a new CreateSettingsReportingEulaItemParams object
// with the default values initialized.
func NewCreateSettingsReportingEulaItemParams() *CreateSettingsReportingEulaItemParams {
	var ()
	return &CreateSettingsReportingEulaItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSettingsReportingEulaItemParamsWithTimeout creates a new CreateSettingsReportingEulaItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSettingsReportingEulaItemParamsWithTimeout(timeout time.Duration) *CreateSettingsReportingEulaItemParams {
	var ()
	return &CreateSettingsReportingEulaItemParams{

		timeout: timeout,
	}
}

// NewCreateSettingsReportingEulaItemParamsWithContext creates a new CreateSettingsReportingEulaItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSettingsReportingEulaItemParamsWithContext(ctx context.Context) *CreateSettingsReportingEulaItemParams {
	var ()
	return &CreateSettingsReportingEulaItemParams{

		Context: ctx,
	}
}

// NewCreateSettingsReportingEulaItemParamsWithHTTPClient creates a new CreateSettingsReportingEulaItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSettingsReportingEulaItemParamsWithHTTPClient(client *http.Client) *CreateSettingsReportingEulaItemParams {
	var ()
	return &CreateSettingsReportingEulaItemParams{
		HTTPClient: client,
	}
}

/*CreateSettingsReportingEulaItemParams contains all the parameters to send to the API endpoint
for the create settings reporting eula item operation typically these are written to a http.Request
*/
type CreateSettingsReportingEulaItemParams struct {

	/*SettingsReportingEulaItem*/
	SettingsReportingEulaItem *models.SettingsReportingEulaItem

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create settings reporting eula item params
func (o *CreateSettingsReportingEulaItemParams) WithTimeout(timeout time.Duration) *CreateSettingsReportingEulaItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create settings reporting eula item params
func (o *CreateSettingsReportingEulaItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create settings reporting eula item params
func (o *CreateSettingsReportingEulaItemParams) WithContext(ctx context.Context) *CreateSettingsReportingEulaItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create settings reporting eula item params
func (o *CreateSettingsReportingEulaItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create settings reporting eula item params
func (o *CreateSettingsReportingEulaItemParams) WithHTTPClient(client *http.Client) *CreateSettingsReportingEulaItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create settings reporting eula item params
func (o *CreateSettingsReportingEulaItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSettingsReportingEulaItem adds the settingsReportingEulaItem to the create settings reporting eula item params
func (o *CreateSettingsReportingEulaItemParams) WithSettingsReportingEulaItem(settingsReportingEulaItem *models.SettingsReportingEulaItem) *CreateSettingsReportingEulaItemParams {
	o.SetSettingsReportingEulaItem(settingsReportingEulaItem)
	return o
}

// SetSettingsReportingEulaItem adds the settingsReportingEulaItem to the create settings reporting eula item params
func (o *CreateSettingsReportingEulaItemParams) SetSettingsReportingEulaItem(settingsReportingEulaItem *models.SettingsReportingEulaItem) {
	o.SettingsReportingEulaItem = settingsReportingEulaItem
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSettingsReportingEulaItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SettingsReportingEulaItem != nil {
		if err := r.SetBodyParam(o.SettingsReportingEulaItem); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
