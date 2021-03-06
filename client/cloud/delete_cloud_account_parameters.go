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
)

// NewDeleteCloudAccountParams creates a new DeleteCloudAccountParams object
// with the default values initialized.
func NewDeleteCloudAccountParams() *DeleteCloudAccountParams {
	var ()
	return &DeleteCloudAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCloudAccountParamsWithTimeout creates a new DeleteCloudAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCloudAccountParamsWithTimeout(timeout time.Duration) *DeleteCloudAccountParams {
	var ()
	return &DeleteCloudAccountParams{

		timeout: timeout,
	}
}

// NewDeleteCloudAccountParamsWithContext creates a new DeleteCloudAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCloudAccountParamsWithContext(ctx context.Context) *DeleteCloudAccountParams {
	var ()
	return &DeleteCloudAccountParams{

		Context: ctx,
	}
}

// NewDeleteCloudAccountParamsWithHTTPClient creates a new DeleteCloudAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCloudAccountParamsWithHTTPClient(client *http.Client) *DeleteCloudAccountParams {
	var ()
	return &DeleteCloudAccountParams{
		HTTPClient: client,
	}
}

/*DeleteCloudAccountParams contains all the parameters to send to the API endpoint
for the delete cloud account operation typically these are written to a http.Request
*/
type DeleteCloudAccountParams struct {

	/*CloudAccountID
	  Delete cloud account.

	*/
	CloudAccountID string
	/*AcknowledgeForceDelete
	  A value of 1 acknowledges that the user is deleting data.

	*/
	AcknowledgeForceDelete *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete cloud account params
func (o *DeleteCloudAccountParams) WithTimeout(timeout time.Duration) *DeleteCloudAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cloud account params
func (o *DeleteCloudAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cloud account params
func (o *DeleteCloudAccountParams) WithContext(ctx context.Context) *DeleteCloudAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cloud account params
func (o *DeleteCloudAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cloud account params
func (o *DeleteCloudAccountParams) WithHTTPClient(client *http.Client) *DeleteCloudAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cloud account params
func (o *DeleteCloudAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudAccountID adds the cloudAccountID to the delete cloud account params
func (o *DeleteCloudAccountParams) WithCloudAccountID(cloudAccountID string) *DeleteCloudAccountParams {
	o.SetCloudAccountID(cloudAccountID)
	return o
}

// SetCloudAccountID adds the cloudAccountId to the delete cloud account params
func (o *DeleteCloudAccountParams) SetCloudAccountID(cloudAccountID string) {
	o.CloudAccountID = cloudAccountID
}

// WithAcknowledgeForceDelete adds the acknowledgeForceDelete to the delete cloud account params
func (o *DeleteCloudAccountParams) WithAcknowledgeForceDelete(acknowledgeForceDelete *string) *DeleteCloudAccountParams {
	o.SetAcknowledgeForceDelete(acknowledgeForceDelete)
	return o
}

// SetAcknowledgeForceDelete adds the acknowledgeForceDelete to the delete cloud account params
func (o *DeleteCloudAccountParams) SetAcknowledgeForceDelete(acknowledgeForceDelete *string) {
	o.AcknowledgeForceDelete = acknowledgeForceDelete
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCloudAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param CloudAccountId
	if err := r.SetPathParam("CloudAccountId", o.CloudAccountID); err != nil {
		return err
	}

	if o.AcknowledgeForceDelete != nil {

		// query param acknowledge_force_delete
		var qrAcknowledgeForceDelete string
		if o.AcknowledgeForceDelete != nil {
			qrAcknowledgeForceDelete = *o.AcknowledgeForceDelete
		}
		qAcknowledgeForceDelete := qrAcknowledgeForceDelete
		if qAcknowledgeForceDelete != "" {
			if err := r.SetQueryParam("acknowledge_force_delete", qAcknowledgeForceDelete); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
