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

// NewCreateSnapshotRepstateParams creates a new CreateSnapshotRepstateParams object
// with the default values initialized.
func NewCreateSnapshotRepstateParams() *CreateSnapshotRepstateParams {
	var ()
	return &CreateSnapshotRepstateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSnapshotRepstateParamsWithTimeout creates a new CreateSnapshotRepstateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSnapshotRepstateParamsWithTimeout(timeout time.Duration) *CreateSnapshotRepstateParams {
	var ()
	return &CreateSnapshotRepstateParams{

		timeout: timeout,
	}
}

// NewCreateSnapshotRepstateParamsWithContext creates a new CreateSnapshotRepstateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSnapshotRepstateParamsWithContext(ctx context.Context) *CreateSnapshotRepstateParams {
	var ()
	return &CreateSnapshotRepstateParams{

		Context: ctx,
	}
}

// NewCreateSnapshotRepstateParamsWithHTTPClient creates a new CreateSnapshotRepstateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSnapshotRepstateParamsWithHTTPClient(client *http.Client) *CreateSnapshotRepstateParams {
	var ()
	return &CreateSnapshotRepstateParams{
		HTTPClient: client,
	}
}

/*CreateSnapshotRepstateParams contains all the parameters to send to the API endpoint
for the create snapshot repstate operation typically these are written to a http.Request
*/
type CreateSnapshotRepstateParams struct {

	/*SnapshotRepstate*/
	SnapshotRepstate *models.SnapshotRepstates

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create snapshot repstate params
func (o *CreateSnapshotRepstateParams) WithTimeout(timeout time.Duration) *CreateSnapshotRepstateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create snapshot repstate params
func (o *CreateSnapshotRepstateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create snapshot repstate params
func (o *CreateSnapshotRepstateParams) WithContext(ctx context.Context) *CreateSnapshotRepstateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create snapshot repstate params
func (o *CreateSnapshotRepstateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create snapshot repstate params
func (o *CreateSnapshotRepstateParams) WithHTTPClient(client *http.Client) *CreateSnapshotRepstateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create snapshot repstate params
func (o *CreateSnapshotRepstateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSnapshotRepstate adds the snapshotRepstate to the create snapshot repstate params
func (o *CreateSnapshotRepstateParams) WithSnapshotRepstate(snapshotRepstate *models.SnapshotRepstates) *CreateSnapshotRepstateParams {
	o.SetSnapshotRepstate(snapshotRepstate)
	return o
}

// SetSnapshotRepstate adds the snapshotRepstate to the create snapshot repstate params
func (o *CreateSnapshotRepstateParams) SetSnapshotRepstate(snapshotRepstate *models.SnapshotRepstates) {
	o.SnapshotRepstate = snapshotRepstate
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSnapshotRepstateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SnapshotRepstate != nil {
		if err := r.SetBodyParam(o.SnapshotRepstate); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
