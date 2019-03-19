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

	models "github.com/tenortim/isi_sdk_go/models"
)

// NewCreateJobPolicyParams creates a new CreateJobPolicyParams object
// with the default values initialized.
func NewCreateJobPolicyParams() *CreateJobPolicyParams {
	var ()
	return &CreateJobPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateJobPolicyParamsWithTimeout creates a new CreateJobPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateJobPolicyParamsWithTimeout(timeout time.Duration) *CreateJobPolicyParams {
	var ()
	return &CreateJobPolicyParams{

		timeout: timeout,
	}
}

// NewCreateJobPolicyParamsWithContext creates a new CreateJobPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateJobPolicyParamsWithContext(ctx context.Context) *CreateJobPolicyParams {
	var ()
	return &CreateJobPolicyParams{

		Context: ctx,
	}
}

// NewCreateJobPolicyParamsWithHTTPClient creates a new CreateJobPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateJobPolicyParamsWithHTTPClient(client *http.Client) *CreateJobPolicyParams {
	var ()
	return &CreateJobPolicyParams{
		HTTPClient: client,
	}
}

/*CreateJobPolicyParams contains all the parameters to send to the API endpoint
for the create job policy operation typically these are written to a http.Request
*/
type CreateJobPolicyParams struct {

	/*JobPolicy*/
	JobPolicy *models.JobPolicyCreateParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create job policy params
func (o *CreateJobPolicyParams) WithTimeout(timeout time.Duration) *CreateJobPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create job policy params
func (o *CreateJobPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create job policy params
func (o *CreateJobPolicyParams) WithContext(ctx context.Context) *CreateJobPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create job policy params
func (o *CreateJobPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create job policy params
func (o *CreateJobPolicyParams) WithHTTPClient(client *http.Client) *CreateJobPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create job policy params
func (o *CreateJobPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobPolicy adds the jobPolicy to the create job policy params
func (o *CreateJobPolicyParams) WithJobPolicy(jobPolicy *models.JobPolicyCreateParams) *CreateJobPolicyParams {
	o.SetJobPolicy(jobPolicy)
	return o
}

// SetJobPolicy adds the jobPolicy to the create job policy params
func (o *CreateJobPolicyParams) SetJobPolicy(jobPolicy *models.JobPolicyCreateParams) {
	o.JobPolicy = jobPolicy
}

// WriteToRequest writes these params to a swagger request
func (o *CreateJobPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.JobPolicy != nil {
		if err := r.SetBodyParam(o.JobPolicy); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
