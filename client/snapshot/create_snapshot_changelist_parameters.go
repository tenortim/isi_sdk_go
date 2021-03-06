// Code generated by go-swagger; DO NOT EDIT.

package snapshot

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

// NewCreateSnapshotChangelistParams creates a new CreateSnapshotChangelistParams object
// with the default values initialized.
func NewCreateSnapshotChangelistParams() *CreateSnapshotChangelistParams {
	var ()
	return &CreateSnapshotChangelistParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSnapshotChangelistParamsWithTimeout creates a new CreateSnapshotChangelistParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSnapshotChangelistParamsWithTimeout(timeout time.Duration) *CreateSnapshotChangelistParams {
	var ()
	return &CreateSnapshotChangelistParams{

		timeout: timeout,
	}
}

// NewCreateSnapshotChangelistParamsWithContext creates a new CreateSnapshotChangelistParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSnapshotChangelistParamsWithContext(ctx context.Context) *CreateSnapshotChangelistParams {
	var ()
	return &CreateSnapshotChangelistParams{

		Context: ctx,
	}
}

// NewCreateSnapshotChangelistParamsWithHTTPClient creates a new CreateSnapshotChangelistParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSnapshotChangelistParamsWithHTTPClient(client *http.Client) *CreateSnapshotChangelistParams {
	var ()
	return &CreateSnapshotChangelistParams{
		HTTPClient: client,
	}
}

/*CreateSnapshotChangelistParams contains all the parameters to send to the API endpoint
for the create snapshot changelist operation typically these are written to a http.Request
*/
type CreateSnapshotChangelistParams struct {

	/*SnapshotChangelist*/
	SnapshotChangelist *models.SnapshotChangelists

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create snapshot changelist params
func (o *CreateSnapshotChangelistParams) WithTimeout(timeout time.Duration) *CreateSnapshotChangelistParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create snapshot changelist params
func (o *CreateSnapshotChangelistParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create snapshot changelist params
func (o *CreateSnapshotChangelistParams) WithContext(ctx context.Context) *CreateSnapshotChangelistParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create snapshot changelist params
func (o *CreateSnapshotChangelistParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create snapshot changelist params
func (o *CreateSnapshotChangelistParams) WithHTTPClient(client *http.Client) *CreateSnapshotChangelistParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create snapshot changelist params
func (o *CreateSnapshotChangelistParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSnapshotChangelist adds the snapshotChangelist to the create snapshot changelist params
func (o *CreateSnapshotChangelistParams) WithSnapshotChangelist(snapshotChangelist *models.SnapshotChangelists) *CreateSnapshotChangelistParams {
	o.SetSnapshotChangelist(snapshotChangelist)
	return o
}

// SetSnapshotChangelist adds the snapshotChangelist to the create snapshot changelist params
func (o *CreateSnapshotChangelistParams) SetSnapshotChangelist(snapshotChangelist *models.SnapshotChangelists) {
	o.SnapshotChangelist = snapshotChangelist
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSnapshotChangelistParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SnapshotChangelist != nil {
		if err := r.SetBodyParam(o.SnapshotChangelist); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
