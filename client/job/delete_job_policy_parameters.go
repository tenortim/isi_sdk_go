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

// NewDeleteJobPolicyParams creates a new DeleteJobPolicyParams object
// with the default values initialized.
func NewDeleteJobPolicyParams() *DeleteJobPolicyParams {
	var ()
	return &DeleteJobPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteJobPolicyParamsWithTimeout creates a new DeleteJobPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteJobPolicyParamsWithTimeout(timeout time.Duration) *DeleteJobPolicyParams {
	var ()
	return &DeleteJobPolicyParams{

		timeout: timeout,
	}
}

// NewDeleteJobPolicyParamsWithContext creates a new DeleteJobPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteJobPolicyParamsWithContext(ctx context.Context) *DeleteJobPolicyParams {
	var ()
	return &DeleteJobPolicyParams{

		Context: ctx,
	}
}

// NewDeleteJobPolicyParamsWithHTTPClient creates a new DeleteJobPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteJobPolicyParamsWithHTTPClient(client *http.Client) *DeleteJobPolicyParams {
	var ()
	return &DeleteJobPolicyParams{
		HTTPClient: client,
	}
}

/*DeleteJobPolicyParams contains all the parameters to send to the API endpoint
for the delete job policy operation typically these are written to a http.Request
*/
type DeleteJobPolicyParams struct {

	/*JobPolicyID
	  Delete a job impact policy.  System policies may not be deleted.

	*/
	JobPolicyID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete job policy params
func (o *DeleteJobPolicyParams) WithTimeout(timeout time.Duration) *DeleteJobPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete job policy params
func (o *DeleteJobPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete job policy params
func (o *DeleteJobPolicyParams) WithContext(ctx context.Context) *DeleteJobPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete job policy params
func (o *DeleteJobPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete job policy params
func (o *DeleteJobPolicyParams) WithHTTPClient(client *http.Client) *DeleteJobPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete job policy params
func (o *DeleteJobPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobPolicyID adds the jobPolicyID to the delete job policy params
func (o *DeleteJobPolicyParams) WithJobPolicyID(jobPolicyID string) *DeleteJobPolicyParams {
	o.SetJobPolicyID(jobPolicyID)
	return o
}

// SetJobPolicyID adds the jobPolicyId to the delete job policy params
func (o *DeleteJobPolicyParams) SetJobPolicyID(jobPolicyID string) {
	o.JobPolicyID = jobPolicyID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteJobPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
