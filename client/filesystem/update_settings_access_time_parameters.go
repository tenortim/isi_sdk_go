// Code generated by go-swagger; DO NOT EDIT.

package filesystem

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

// NewUpdateSettingsAccessTimeParams creates a new UpdateSettingsAccessTimeParams object
// with the default values initialized.
func NewUpdateSettingsAccessTimeParams() *UpdateSettingsAccessTimeParams {
	var ()
	return &UpdateSettingsAccessTimeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSettingsAccessTimeParamsWithTimeout creates a new UpdateSettingsAccessTimeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSettingsAccessTimeParamsWithTimeout(timeout time.Duration) *UpdateSettingsAccessTimeParams {
	var ()
	return &UpdateSettingsAccessTimeParams{

		timeout: timeout,
	}
}

// NewUpdateSettingsAccessTimeParamsWithContext creates a new UpdateSettingsAccessTimeParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSettingsAccessTimeParamsWithContext(ctx context.Context) *UpdateSettingsAccessTimeParams {
	var ()
	return &UpdateSettingsAccessTimeParams{

		Context: ctx,
	}
}

// NewUpdateSettingsAccessTimeParamsWithHTTPClient creates a new UpdateSettingsAccessTimeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSettingsAccessTimeParamsWithHTTPClient(client *http.Client) *UpdateSettingsAccessTimeParams {
	var ()
	return &UpdateSettingsAccessTimeParams{
		HTTPClient: client,
	}
}

/*UpdateSettingsAccessTimeParams contains all the parameters to send to the API endpoint
for the update settings access time operation typically these are written to a http.Request
*/
type UpdateSettingsAccessTimeParams struct {

	/*SettingsAccessTime*/
	SettingsAccessTime *models.SettingsAccessTimeExtended

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update settings access time params
func (o *UpdateSettingsAccessTimeParams) WithTimeout(timeout time.Duration) *UpdateSettingsAccessTimeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update settings access time params
func (o *UpdateSettingsAccessTimeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update settings access time params
func (o *UpdateSettingsAccessTimeParams) WithContext(ctx context.Context) *UpdateSettingsAccessTimeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update settings access time params
func (o *UpdateSettingsAccessTimeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update settings access time params
func (o *UpdateSettingsAccessTimeParams) WithHTTPClient(client *http.Client) *UpdateSettingsAccessTimeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update settings access time params
func (o *UpdateSettingsAccessTimeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSettingsAccessTime adds the settingsAccessTime to the update settings access time params
func (o *UpdateSettingsAccessTimeParams) WithSettingsAccessTime(settingsAccessTime *models.SettingsAccessTimeExtended) *UpdateSettingsAccessTimeParams {
	o.SetSettingsAccessTime(settingsAccessTime)
	return o
}

// SetSettingsAccessTime adds the settingsAccessTime to the update settings access time params
func (o *UpdateSettingsAccessTimeParams) SetSettingsAccessTime(settingsAccessTime *models.SettingsAccessTimeExtended) {
	o.SettingsAccessTime = settingsAccessTime
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSettingsAccessTimeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SettingsAccessTime != nil {
		if err := r.SetBodyParam(o.SettingsAccessTime); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}