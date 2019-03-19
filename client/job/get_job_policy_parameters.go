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

// NewGetJobPolicyParams creates a new GetJobPolicyParams object
// with the default values initialized.
func NewGetJobPolicyParams() *GetJobPolicyParams {
	var ()
	return &GetJobPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetJobPolicyParamsWithTimeout creates a new GetJobPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetJobPolicyParamsWithTimeout(timeout time.Duration) *GetJobPolicyParams {
	var ()
	return &GetJobPolicyParams{

		timeout: timeout,
	}
}

// NewGetJobPolicyParamsWithContext creates a new GetJobPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetJobPolicyParamsWithContext(ctx context.Context) *GetJobPolicyParams {
	var ()
	return &GetJobPolicyParams{

		Context: ctx,
	}
}

// NewGetJobPolicyParamsWithHTTPClient creates a new GetJobPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetJobPolicyParamsWithHTTPClient(client *http.Client) *GetJobPolicyParams {
	var ()
	return &GetJobPolicyParams{
		HTTPClient: client,
	}
}

/*GetJobPolicyParams contains all the parameters to send to the API endpoint
for the get job policy operation typically these are written to a http.Request
*/
type GetJobPolicyParams struct {

	/*JobPolicyID
	  View a single job impact policy.

	*/
	JobPolicyID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get job policy params
func (o *GetJobPolicyParams) WithTimeout(timeout time.Duration) *GetJobPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get job policy params
func (o *GetJobPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get job policy params
func (o *GetJobPolicyParams) WithContext(ctx context.Context) *GetJobPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get job policy params
func (o *GetJobPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get job policy params
func (o *GetJobPolicyParams) WithHTTPClient(client *http.Client) *GetJobPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get job policy params
func (o *GetJobPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobPolicyID adds the jobPolicyID to the get job policy params
func (o *GetJobPolicyParams) WithJobPolicyID(jobPolicyID string) *GetJobPolicyParams {
	o.SetJobPolicyID(jobPolicyID)
	return o
}

// SetJobPolicyID adds the jobPolicyId to the get job policy params
func (o *GetJobPolicyParams) SetJobPolicyID(jobPolicyID string) {
	o.JobPolicyID = jobPolicyID
}

// WriteToRequest writes these params to a swagger request
func (o *GetJobPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param JobPolicyId
	if err := r.SetPathParam("JobPolicyId", o.JobPolicyID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}