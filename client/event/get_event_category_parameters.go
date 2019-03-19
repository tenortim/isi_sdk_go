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

// NewGetEventCategoryParams creates a new GetEventCategoryParams object
// with the default values initialized.
func NewGetEventCategoryParams() *GetEventCategoryParams {
	var ()
	return &GetEventCategoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEventCategoryParamsWithTimeout creates a new GetEventCategoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEventCategoryParamsWithTimeout(timeout time.Duration) *GetEventCategoryParams {
	var ()
	return &GetEventCategoryParams{

		timeout: timeout,
	}
}

// NewGetEventCategoryParamsWithContext creates a new GetEventCategoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEventCategoryParamsWithContext(ctx context.Context) *GetEventCategoryParams {
	var ()
	return &GetEventCategoryParams{

		Context: ctx,
	}
}

// NewGetEventCategoryParamsWithHTTPClient creates a new GetEventCategoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEventCategoryParamsWithHTTPClient(client *http.Client) *GetEventCategoryParams {
	var ()
	return &GetEventCategoryParams{
		HTTPClient: client,
	}
}

/*GetEventCategoryParams contains all the parameters to send to the API endpoint
for the get event category operation typically these are written to a http.Request
*/
type GetEventCategoryParams struct {

	/*EventCategoryID
	  Retrieve the eventgroup category.

	*/
	EventCategoryID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get event category params
func (o *GetEventCategoryParams) WithTimeout(timeout time.Duration) *GetEventCategoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get event category params
func (o *GetEventCategoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get event category params
func (o *GetEventCategoryParams) WithContext(ctx context.Context) *GetEventCategoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get event category params
func (o *GetEventCategoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get event category params
func (o *GetEventCategoryParams) WithHTTPClient(client *http.Client) *GetEventCategoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get event category params
func (o *GetEventCategoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEventCategoryID adds the eventCategoryID to the get event category params
func (o *GetEventCategoryParams) WithEventCategoryID(eventCategoryID string) *GetEventCategoryParams {
	o.SetEventCategoryID(eventCategoryID)
	return o
}

// SetEventCategoryID adds the eventCategoryId to the get event category params
func (o *GetEventCategoryParams) SetEventCategoryID(eventCategoryID string) {
	o.EventCategoryID = eventCategoryID
}

// WriteToRequest writes these params to a swagger request
func (o *GetEventCategoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param EventCategoryId
	if err := r.SetPathParam("EventCategoryId", o.EventCategoryID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}