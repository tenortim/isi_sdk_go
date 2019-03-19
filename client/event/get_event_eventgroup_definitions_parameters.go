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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetEventEventgroupDefinitionsParams creates a new GetEventEventgroupDefinitionsParams object
// with the default values initialized.
func NewGetEventEventgroupDefinitionsParams() *GetEventEventgroupDefinitionsParams {
	var ()
	return &GetEventEventgroupDefinitionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEventEventgroupDefinitionsParamsWithTimeout creates a new GetEventEventgroupDefinitionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEventEventgroupDefinitionsParamsWithTimeout(timeout time.Duration) *GetEventEventgroupDefinitionsParams {
	var ()
	return &GetEventEventgroupDefinitionsParams{

		timeout: timeout,
	}
}

// NewGetEventEventgroupDefinitionsParamsWithContext creates a new GetEventEventgroupDefinitionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEventEventgroupDefinitionsParamsWithContext(ctx context.Context) *GetEventEventgroupDefinitionsParams {
	var ()
	return &GetEventEventgroupDefinitionsParams{

		Context: ctx,
	}
}

// NewGetEventEventgroupDefinitionsParamsWithHTTPClient creates a new GetEventEventgroupDefinitionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEventEventgroupDefinitionsParamsWithHTTPClient(client *http.Client) *GetEventEventgroupDefinitionsParams {
	var ()
	return &GetEventEventgroupDefinitionsParams{
		HTTPClient: client,
	}
}

/*GetEventEventgroupDefinitionsParams contains all the parameters to send to the API endpoint
for the get event eventgroup definitions operation typically these are written to a http.Request
*/
type GetEventEventgroupDefinitionsParams struct {

	/*Category
	  Return eventgroups in the specified category

	*/
	Category *int64
	/*Limit
	  Return no more than this many results at once (see resume).

	*/
	Limit *int64
	/*Resume
	  Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options).

	*/
	Resume *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get event eventgroup definitions params
func (o *GetEventEventgroupDefinitionsParams) WithTimeout(timeout time.Duration) *GetEventEventgroupDefinitionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get event eventgroup definitions params
func (o *GetEventEventgroupDefinitionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get event eventgroup definitions params
func (o *GetEventEventgroupDefinitionsParams) WithContext(ctx context.Context) *GetEventEventgroupDefinitionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get event eventgroup definitions params
func (o *GetEventEventgroupDefinitionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get event eventgroup definitions params
func (o *GetEventEventgroupDefinitionsParams) WithHTTPClient(client *http.Client) *GetEventEventgroupDefinitionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get event eventgroup definitions params
func (o *GetEventEventgroupDefinitionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategory adds the category to the get event eventgroup definitions params
func (o *GetEventEventgroupDefinitionsParams) WithCategory(category *int64) *GetEventEventgroupDefinitionsParams {
	o.SetCategory(category)
	return o
}

// SetCategory adds the category to the get event eventgroup definitions params
func (o *GetEventEventgroupDefinitionsParams) SetCategory(category *int64) {
	o.Category = category
}

// WithLimit adds the limit to the get event eventgroup definitions params
func (o *GetEventEventgroupDefinitionsParams) WithLimit(limit *int64) *GetEventEventgroupDefinitionsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get event eventgroup definitions params
func (o *GetEventEventgroupDefinitionsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithResume adds the resume to the get event eventgroup definitions params
func (o *GetEventEventgroupDefinitionsParams) WithResume(resume *string) *GetEventEventgroupDefinitionsParams {
	o.SetResume(resume)
	return o
}

// SetResume adds the resume to the get event eventgroup definitions params
func (o *GetEventEventgroupDefinitionsParams) SetResume(resume *string) {
	o.Resume = resume
}

// WriteToRequest writes these params to a swagger request
func (o *GetEventEventgroupDefinitionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Category != nil {

		// query param category
		var qrCategory int64
		if o.Category != nil {
			qrCategory = *o.Category
		}
		qCategory := swag.FormatInt64(qrCategory)
		if qCategory != "" {
			if err := r.SetQueryParam("category", qCategory); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Resume != nil {

		// query param resume
		var qrResume string
		if o.Resume != nil {
			qrResume = *o.Resume
		}
		qResume := qrResume
		if qResume != "" {
			if err := r.SetQueryParam("resume", qResume); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
