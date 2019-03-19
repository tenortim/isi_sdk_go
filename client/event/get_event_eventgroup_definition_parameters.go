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

// NewGetEventEventgroupDefinitionParams creates a new GetEventEventgroupDefinitionParams object
// with the default values initialized.
func NewGetEventEventgroupDefinitionParams() *GetEventEventgroupDefinitionParams {
	var ()
	return &GetEventEventgroupDefinitionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEventEventgroupDefinitionParamsWithTimeout creates a new GetEventEventgroupDefinitionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEventEventgroupDefinitionParamsWithTimeout(timeout time.Duration) *GetEventEventgroupDefinitionParams {
	var ()
	return &GetEventEventgroupDefinitionParams{

		timeout: timeout,
	}
}

// NewGetEventEventgroupDefinitionParamsWithContext creates a new GetEventEventgroupDefinitionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEventEventgroupDefinitionParamsWithContext(ctx context.Context) *GetEventEventgroupDefinitionParams {
	var ()
	return &GetEventEventgroupDefinitionParams{

		Context: ctx,
	}
}

// NewGetEventEventgroupDefinitionParamsWithHTTPClient creates a new GetEventEventgroupDefinitionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEventEventgroupDefinitionParamsWithHTTPClient(client *http.Client) *GetEventEventgroupDefinitionParams {
	var ()
	return &GetEventEventgroupDefinitionParams{
		HTTPClient: client,
	}
}

/*GetEventEventgroupDefinitionParams contains all the parameters to send to the API endpoint
for the get event eventgroup definition operation typically these are written to a http.Request
*/
type GetEventEventgroupDefinitionParams struct {

	/*EventEventgroupDefinitionID
	  Retrieve the eventgroup definition.

	*/
	EventEventgroupDefinitionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get event eventgroup definition params
func (o *GetEventEventgroupDefinitionParams) WithTimeout(timeout time.Duration) *GetEventEventgroupDefinitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get event eventgroup definition params
func (o *GetEventEventgroupDefinitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get event eventgroup definition params
func (o *GetEventEventgroupDefinitionParams) WithContext(ctx context.Context) *GetEventEventgroupDefinitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get event eventgroup definition params
func (o *GetEventEventgroupDefinitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get event eventgroup definition params
func (o *GetEventEventgroupDefinitionParams) WithHTTPClient(client *http.Client) *GetEventEventgroupDefinitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get event eventgroup definition params
func (o *GetEventEventgroupDefinitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEventEventgroupDefinitionID adds the eventEventgroupDefinitionID to the get event eventgroup definition params
func (o *GetEventEventgroupDefinitionParams) WithEventEventgroupDefinitionID(eventEventgroupDefinitionID string) *GetEventEventgroupDefinitionParams {
	o.SetEventEventgroupDefinitionID(eventEventgroupDefinitionID)
	return o
}

// SetEventEventgroupDefinitionID adds the eventEventgroupDefinitionId to the get event eventgroup definition params
func (o *GetEventEventgroupDefinitionParams) SetEventEventgroupDefinitionID(eventEventgroupDefinitionID string) {
	o.EventEventgroupDefinitionID = eventEventgroupDefinitionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetEventEventgroupDefinitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param EventEventgroupDefinitionId
	if err := r.SetPathParam("EventEventgroupDefinitionId", o.EventEventgroupDefinitionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}