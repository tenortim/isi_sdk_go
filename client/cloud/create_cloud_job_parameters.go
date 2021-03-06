// Code generated by go-swagger; DO NOT EDIT.

package cloud

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

// NewCreateCloudJobParams creates a new CreateCloudJobParams object
// with the default values initialized.
func NewCreateCloudJobParams() *CreateCloudJobParams {
	var ()
	return &CreateCloudJobParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCloudJobParamsWithTimeout creates a new CreateCloudJobParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateCloudJobParamsWithTimeout(timeout time.Duration) *CreateCloudJobParams {
	var ()
	return &CreateCloudJobParams{

		timeout: timeout,
	}
}

// NewCreateCloudJobParamsWithContext creates a new CreateCloudJobParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateCloudJobParamsWithContext(ctx context.Context) *CreateCloudJobParams {
	var ()
	return &CreateCloudJobParams{

		Context: ctx,
	}
}

// NewCreateCloudJobParamsWithHTTPClient creates a new CreateCloudJobParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateCloudJobParamsWithHTTPClient(client *http.Client) *CreateCloudJobParams {
	var ()
	return &CreateCloudJobParams{
		HTTPClient: client,
	}
}

/*CreateCloudJobParams contains all the parameters to send to the API endpoint
for the create cloud job operation typically these are written to a http.Request
*/
type CreateCloudJobParams struct {

	/*CloudJob*/
	CloudJob *models.CloudJobCreateParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create cloud job params
func (o *CreateCloudJobParams) WithTimeout(timeout time.Duration) *CreateCloudJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cloud job params
func (o *CreateCloudJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cloud job params
func (o *CreateCloudJobParams) WithContext(ctx context.Context) *CreateCloudJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cloud job params
func (o *CreateCloudJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cloud job params
func (o *CreateCloudJobParams) WithHTTPClient(client *http.Client) *CreateCloudJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cloud job params
func (o *CreateCloudJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudJob adds the cloudJob to the create cloud job params
func (o *CreateCloudJobParams) WithCloudJob(cloudJob *models.CloudJobCreateParams) *CreateCloudJobParams {
	o.SetCloudJob(cloudJob)
	return o
}

// SetCloudJob adds the cloudJob to the create cloud job params
func (o *CreateCloudJobParams) SetCloudJob(cloudJob *models.CloudJobCreateParams) {
	o.CloudJob = cloudJob
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCloudJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CloudJob != nil {
		if err := r.SetBodyParam(o.CloudJob); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
