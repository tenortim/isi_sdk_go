// Code generated by go-swagger; DO NOT EDIT.

package job

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

// NewGetJobTypeParams creates a new GetJobTypeParams object
// with the default values initialized.
func NewGetJobTypeParams() *GetJobTypeParams {
	var ()
	return &GetJobTypeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetJobTypeParamsWithTimeout creates a new GetJobTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetJobTypeParamsWithTimeout(timeout time.Duration) *GetJobTypeParams {
	var ()
	return &GetJobTypeParams{

		timeout: timeout,
	}
}

// NewGetJobTypeParamsWithContext creates a new GetJobTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetJobTypeParamsWithContext(ctx context.Context) *GetJobTypeParams {
	var ()
	return &GetJobTypeParams{

		Context: ctx,
	}
}

// NewGetJobTypeParamsWithHTTPClient creates a new GetJobTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetJobTypeParamsWithHTTPClient(client *http.Client) *GetJobTypeParams {
	var ()
	return &GetJobTypeParams{
		HTTPClient: client,
	}
}

/*GetJobTypeParams contains all the parameters to send to the API endpoint
for the get job type operation typically these are written to a http.Request
*/
type GetJobTypeParams struct {

	/*JobTypeID
	  Retrieve job type information.

	*/
	JobTypeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get job type params
func (o *GetJobTypeParams) WithTimeout(timeout time.Duration) *GetJobTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get job type params
func (o *GetJobTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get job type params
func (o *GetJobTypeParams) WithContext(ctx context.Context) *GetJobTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get job type params
func (o *GetJobTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get job type params
func (o *GetJobTypeParams) WithHTTPClient(client *http.Client) *GetJobTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get job type params
func (o *GetJobTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobTypeID adds the jobTypeID to the get job type params
func (o *GetJobTypeParams) WithJobTypeID(jobTypeID string) *GetJobTypeParams {
	o.SetJobTypeID(jobTypeID)
	return o
}

// SetJobTypeID adds the jobTypeId to the get job type params
func (o *GetJobTypeParams) SetJobTypeID(jobTypeID string) {
	o.JobTypeID = jobTypeID
}

// WriteToRequest writes these params to a swagger request
func (o *GetJobTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param JobTypeId
	if err := r.SetPathParam("JobTypeId", o.JobTypeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}