// Code generated by go-swagger; DO NOT EDIT.

package event

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

// NewDeleteEventChannelParams creates a new DeleteEventChannelParams object
// with the default values initialized.
func NewDeleteEventChannelParams() *DeleteEventChannelParams {
	var ()
	return &DeleteEventChannelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEventChannelParamsWithTimeout creates a new DeleteEventChannelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteEventChannelParamsWithTimeout(timeout time.Duration) *DeleteEventChannelParams {
	var ()
	return &DeleteEventChannelParams{

		timeout: timeout,
	}
}

// NewDeleteEventChannelParamsWithContext creates a new DeleteEventChannelParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteEventChannelParamsWithContext(ctx context.Context) *DeleteEventChannelParams {
	var ()
	return &DeleteEventChannelParams{

		Context: ctx,
	}
}

// NewDeleteEventChannelParamsWithHTTPClient creates a new DeleteEventChannelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteEventChannelParamsWithHTTPClient(client *http.Client) *DeleteEventChannelParams {
	var ()
	return &DeleteEventChannelParams{
		HTTPClient: client,
	}
}

/*DeleteEventChannelParams contains all the parameters to send to the API endpoint
for the delete event channel operation typically these are written to a http.Request
*/
type DeleteEventChannelParams struct {

	/*EventChannelID
	  Delete the alert-condition.

	*/
	EventChannelID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete event channel params
func (o *DeleteEventChannelParams) WithTimeout(timeout time.Duration) *DeleteEventChannelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete event channel params
func (o *DeleteEventChannelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete event channel params
func (o *DeleteEventChannelParams) WithContext(ctx context.Context) *DeleteEventChannelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete event channel params
func (o *DeleteEventChannelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete event channel params
func (o *DeleteEventChannelParams) WithHTTPClient(client *http.Client) *DeleteEventChannelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete event channel params
func (o *DeleteEventChannelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEventChannelID adds the eventChannelID to the delete event channel params
func (o *DeleteEventChannelParams) WithEventChannelID(eventChannelID string) *DeleteEventChannelParams {
	o.SetEventChannelID(eventChannelID)
	return o
}

// SetEventChannelID adds the eventChannelId to the delete event channel params
func (o *DeleteEventChannelParams) SetEventChannelID(eventChannelID string) {
	o.EventChannelID = eventChannelID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEventChannelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param EventChannelId
	if err := r.SetPathParam("EventChannelId", o.EventChannelID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
