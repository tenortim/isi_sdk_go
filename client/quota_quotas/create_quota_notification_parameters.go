// Code generated by go-swagger; DO NOT EDIT.

package quota_quotas

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

// NewCreateQuotaNotificationParams creates a new CreateQuotaNotificationParams object
// with the default values initialized.
func NewCreateQuotaNotificationParams() *CreateQuotaNotificationParams {
	var ()
	return &CreateQuotaNotificationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateQuotaNotificationParamsWithTimeout creates a new CreateQuotaNotificationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateQuotaNotificationParamsWithTimeout(timeout time.Duration) *CreateQuotaNotificationParams {
	var ()
	return &CreateQuotaNotificationParams{

		timeout: timeout,
	}
}

// NewCreateQuotaNotificationParamsWithContext creates a new CreateQuotaNotificationParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateQuotaNotificationParamsWithContext(ctx context.Context) *CreateQuotaNotificationParams {
	var ()
	return &CreateQuotaNotificationParams{

		Context: ctx,
	}
}

// NewCreateQuotaNotificationParamsWithHTTPClient creates a new CreateQuotaNotificationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateQuotaNotificationParamsWithHTTPClient(client *http.Client) *CreateQuotaNotificationParams {
	var ()
	return &CreateQuotaNotificationParams{
		HTTPClient: client,
	}
}

/*CreateQuotaNotificationParams contains all the parameters to send to the API endpoint
for the create quota notification operation typically these are written to a http.Request
*/
type CreateQuotaNotificationParams struct {

	/*Qid*/
	Qid string
	/*QuotaNotification*/
	QuotaNotification *models.QuotaNotificationCreateParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create quota notification params
func (o *CreateQuotaNotificationParams) WithTimeout(timeout time.Duration) *CreateQuotaNotificationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create quota notification params
func (o *CreateQuotaNotificationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create quota notification params
func (o *CreateQuotaNotificationParams) WithContext(ctx context.Context) *CreateQuotaNotificationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create quota notification params
func (o *CreateQuotaNotificationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create quota notification params
func (o *CreateQuotaNotificationParams) WithHTTPClient(client *http.Client) *CreateQuotaNotificationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create quota notification params
func (o *CreateQuotaNotificationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQid adds the qid to the create quota notification params
func (o *CreateQuotaNotificationParams) WithQid(qid string) *CreateQuotaNotificationParams {
	o.SetQid(qid)
	return o
}

// SetQid adds the qid to the create quota notification params
func (o *CreateQuotaNotificationParams) SetQid(qid string) {
	o.Qid = qid
}

// WithQuotaNotification adds the quotaNotification to the create quota notification params
func (o *CreateQuotaNotificationParams) WithQuotaNotification(quotaNotification *models.QuotaNotificationCreateParams) *CreateQuotaNotificationParams {
	o.SetQuotaNotification(quotaNotification)
	return o
}

// SetQuotaNotification adds the quotaNotification to the create quota notification params
func (o *CreateQuotaNotificationParams) SetQuotaNotification(quotaNotification *models.QuotaNotificationCreateParams) {
	o.QuotaNotification = quotaNotification
}

// WriteToRequest writes these params to a swagger request
func (o *CreateQuotaNotificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Qid
	if err := r.SetPathParam("Qid", o.Qid); err != nil {
		return err
	}

	if o.QuotaNotification != nil {
		if err := r.SetBodyParam(o.QuotaNotification); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}