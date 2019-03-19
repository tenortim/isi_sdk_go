// Code generated by go-swagger; DO NOT EDIT.

package sync

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

// NewCreateSyncReportsRotateItemParams creates a new CreateSyncReportsRotateItemParams object
// with the default values initialized.
func NewCreateSyncReportsRotateItemParams() *CreateSyncReportsRotateItemParams {
	var ()
	return &CreateSyncReportsRotateItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSyncReportsRotateItemParamsWithTimeout creates a new CreateSyncReportsRotateItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSyncReportsRotateItemParamsWithTimeout(timeout time.Duration) *CreateSyncReportsRotateItemParams {
	var ()
	return &CreateSyncReportsRotateItemParams{

		timeout: timeout,
	}
}

// NewCreateSyncReportsRotateItemParamsWithContext creates a new CreateSyncReportsRotateItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSyncReportsRotateItemParamsWithContext(ctx context.Context) *CreateSyncReportsRotateItemParams {
	var ()
	return &CreateSyncReportsRotateItemParams{

		Context: ctx,
	}
}

// NewCreateSyncReportsRotateItemParamsWithHTTPClient creates a new CreateSyncReportsRotateItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSyncReportsRotateItemParamsWithHTTPClient(client *http.Client) *CreateSyncReportsRotateItemParams {
	var ()
	return &CreateSyncReportsRotateItemParams{
		HTTPClient: client,
	}
}

/*CreateSyncReportsRotateItemParams contains all the parameters to send to the API endpoint
for the create sync reports rotate item operation typically these are written to a http.Request
*/
type CreateSyncReportsRotateItemParams struct {

	/*SyncReportsRotateItem*/
	SyncReportsRotateItem models.Empty

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create sync reports rotate item params
func (o *CreateSyncReportsRotateItemParams) WithTimeout(timeout time.Duration) *CreateSyncReportsRotateItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create sync reports rotate item params
func (o *CreateSyncReportsRotateItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create sync reports rotate item params
func (o *CreateSyncReportsRotateItemParams) WithContext(ctx context.Context) *CreateSyncReportsRotateItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create sync reports rotate item params
func (o *CreateSyncReportsRotateItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create sync reports rotate item params
func (o *CreateSyncReportsRotateItemParams) WithHTTPClient(client *http.Client) *CreateSyncReportsRotateItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create sync reports rotate item params
func (o *CreateSyncReportsRotateItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSyncReportsRotateItem adds the syncReportsRotateItem to the create sync reports rotate item params
func (o *CreateSyncReportsRotateItemParams) WithSyncReportsRotateItem(syncReportsRotateItem models.Empty) *CreateSyncReportsRotateItemParams {
	o.SetSyncReportsRotateItem(syncReportsRotateItem)
	return o
}

// SetSyncReportsRotateItem adds the syncReportsRotateItem to the create sync reports rotate item params
func (o *CreateSyncReportsRotateItemParams) SetSyncReportsRotateItem(syncReportsRotateItem models.Empty) {
	o.SyncReportsRotateItem = syncReportsRotateItem
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSyncReportsRotateItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SyncReportsRotateItem != nil {
		if err := r.SetBodyParam(o.SyncReportsRotateItem); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
