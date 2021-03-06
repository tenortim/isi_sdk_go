// Code generated by go-swagger; DO NOT EDIT.

package quota

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
)

// NewGetSettingsNotificationParams creates a new GetSettingsNotificationParams object
// with the default values initialized.
func NewGetSettingsNotificationParams() *GetSettingsNotificationParams {
	var ()
	return &GetSettingsNotificationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSettingsNotificationParamsWithTimeout creates a new GetSettingsNotificationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSettingsNotificationParamsWithTimeout(timeout time.Duration) *GetSettingsNotificationParams {
	var ()
	return &GetSettingsNotificationParams{

		timeout: timeout,
	}
}

// NewGetSettingsNotificationParamsWithContext creates a new GetSettingsNotificationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSettingsNotificationParamsWithContext(ctx context.Context) *GetSettingsNotificationParams {
	var ()
	return &GetSettingsNotificationParams{

		Context: ctx,
	}
}

// NewGetSettingsNotificationParamsWithHTTPClient creates a new GetSettingsNotificationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSettingsNotificationParamsWithHTTPClient(client *http.Client) *GetSettingsNotificationParams {
	var ()
	return &GetSettingsNotificationParams{
		HTTPClient: client,
	}
}

/*GetSettingsNotificationParams contains all the parameters to send to the API endpoint
for the get settings notification operation typically these are written to a http.Request
*/
type GetSettingsNotificationParams struct {

	/*SettingsNotificationID
	  Retrieve notification rule information.

	*/
	SettingsNotificationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get settings notification params
func (o *GetSettingsNotificationParams) WithTimeout(timeout time.Duration) *GetSettingsNotificationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get settings notification params
func (o *GetSettingsNotificationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get settings notification params
func (o *GetSettingsNotificationParams) WithContext(ctx context.Context) *GetSettingsNotificationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get settings notification params
func (o *GetSettingsNotificationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get settings notification params
func (o *GetSettingsNotificationParams) WithHTTPClient(client *http.Client) *GetSettingsNotificationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get settings notification params
func (o *GetSettingsNotificationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSettingsNotificationID adds the settingsNotificationID to the get settings notification params
func (o *GetSettingsNotificationParams) WithSettingsNotificationID(settingsNotificationID string) *GetSettingsNotificationParams {
	o.SetSettingsNotificationID(settingsNotificationID)
	return o
}

// SetSettingsNotificationID adds the settingsNotificationId to the get settings notification params
func (o *GetSettingsNotificationParams) SetSettingsNotificationID(settingsNotificationID string) {
	o.SettingsNotificationID = settingsNotificationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSettingsNotificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param SettingsNotificationId
	if err := r.SetPathParam("SettingsNotificationId", o.SettingsNotificationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
