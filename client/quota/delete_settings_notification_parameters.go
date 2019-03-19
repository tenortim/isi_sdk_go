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

// NewDeleteSettingsNotificationParams creates a new DeleteSettingsNotificationParams object
// with the default values initialized.
func NewDeleteSettingsNotificationParams() *DeleteSettingsNotificationParams {
	var ()
	return &DeleteSettingsNotificationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSettingsNotificationParamsWithTimeout creates a new DeleteSettingsNotificationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSettingsNotificationParamsWithTimeout(timeout time.Duration) *DeleteSettingsNotificationParams {
	var ()
	return &DeleteSettingsNotificationParams{

		timeout: timeout,
	}
}

// NewDeleteSettingsNotificationParamsWithContext creates a new DeleteSettingsNotificationParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSettingsNotificationParamsWithContext(ctx context.Context) *DeleteSettingsNotificationParams {
	var ()
	return &DeleteSettingsNotificationParams{

		Context: ctx,
	}
}

// NewDeleteSettingsNotificationParamsWithHTTPClient creates a new DeleteSettingsNotificationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSettingsNotificationParamsWithHTTPClient(client *http.Client) *DeleteSettingsNotificationParams {
	var ()
	return &DeleteSettingsNotificationParams{
		HTTPClient: client,
	}
}

/*DeleteSettingsNotificationParams contains all the parameters to send to the API endpoint
for the delete settings notification operation typically these are written to a http.Request
*/
type DeleteSettingsNotificationParams struct {

	/*SettingsNotificationID
	  Delete the notification rule.

	*/
	SettingsNotificationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete settings notification params
func (o *DeleteSettingsNotificationParams) WithTimeout(timeout time.Duration) *DeleteSettingsNotificationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete settings notification params
func (o *DeleteSettingsNotificationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete settings notification params
func (o *DeleteSettingsNotificationParams) WithContext(ctx context.Context) *DeleteSettingsNotificationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete settings notification params
func (o *DeleteSettingsNotificationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete settings notification params
func (o *DeleteSettingsNotificationParams) WithHTTPClient(client *http.Client) *DeleteSettingsNotificationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete settings notification params
func (o *DeleteSettingsNotificationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSettingsNotificationID adds the settingsNotificationID to the delete settings notification params
func (o *DeleteSettingsNotificationParams) WithSettingsNotificationID(settingsNotificationID string) *DeleteSettingsNotificationParams {
	o.SetSettingsNotificationID(settingsNotificationID)
	return o
}

// SetSettingsNotificationID adds the settingsNotificationId to the delete settings notification params
func (o *DeleteSettingsNotificationParams) SetSettingsNotificationID(settingsNotificationID string) {
	o.SettingsNotificationID = settingsNotificationID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSettingsNotificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
