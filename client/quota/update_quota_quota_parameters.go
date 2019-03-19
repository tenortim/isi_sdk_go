// Code generated by go-swagger; DO NOT EDIT.

package quota

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

// NewUpdateQuotaQuotaParams creates a new UpdateQuotaQuotaParams object
// with the default values initialized.
func NewUpdateQuotaQuotaParams() *UpdateQuotaQuotaParams {
	var ()
	return &UpdateQuotaQuotaParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateQuotaQuotaParamsWithTimeout creates a new UpdateQuotaQuotaParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateQuotaQuotaParamsWithTimeout(timeout time.Duration) *UpdateQuotaQuotaParams {
	var ()
	return &UpdateQuotaQuotaParams{

		timeout: timeout,
	}
}

// NewUpdateQuotaQuotaParamsWithContext creates a new UpdateQuotaQuotaParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateQuotaQuotaParamsWithContext(ctx context.Context) *UpdateQuotaQuotaParams {
	var ()
	return &UpdateQuotaQuotaParams{

		Context: ctx,
	}
}

// NewUpdateQuotaQuotaParamsWithHTTPClient creates a new UpdateQuotaQuotaParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateQuotaQuotaParamsWithHTTPClient(client *http.Client) *UpdateQuotaQuotaParams {
	var ()
	return &UpdateQuotaQuotaParams{
		HTTPClient: client,
	}
}

/*UpdateQuotaQuotaParams contains all the parameters to send to the API endpoint
for the update quota quota operation typically these are written to a http.Request
*/
type UpdateQuotaQuotaParams struct {

	/*QuotaQuota*/
	QuotaQuota *models.QuotaQuota
	/*QuotaQuotaID
	  Modify quota. All input fields are optional, but one or more must be supplied.

	*/
	QuotaQuotaID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update quota quota params
func (o *UpdateQuotaQuotaParams) WithTimeout(timeout time.Duration) *UpdateQuotaQuotaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update quota quota params
func (o *UpdateQuotaQuotaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update quota quota params
func (o *UpdateQuotaQuotaParams) WithContext(ctx context.Context) *UpdateQuotaQuotaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update quota quota params
func (o *UpdateQuotaQuotaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update quota quota params
func (o *UpdateQuotaQuotaParams) WithHTTPClient(client *http.Client) *UpdateQuotaQuotaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update quota quota params
func (o *UpdateQuotaQuotaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuotaQuota adds the quotaQuota to the update quota quota params
func (o *UpdateQuotaQuotaParams) WithQuotaQuota(quotaQuota *models.QuotaQuota) *UpdateQuotaQuotaParams {
	o.SetQuotaQuota(quotaQuota)
	return o
}

// SetQuotaQuota adds the quotaQuota to the update quota quota params
func (o *UpdateQuotaQuotaParams) SetQuotaQuota(quotaQuota *models.QuotaQuota) {
	o.QuotaQuota = quotaQuota
}

// WithQuotaQuotaID adds the quotaQuotaID to the update quota quota params
func (o *UpdateQuotaQuotaParams) WithQuotaQuotaID(quotaQuotaID string) *UpdateQuotaQuotaParams {
	o.SetQuotaQuotaID(quotaQuotaID)
	return o
}

// SetQuotaQuotaID adds the quotaQuotaId to the update quota quota params
func (o *UpdateQuotaQuotaParams) SetQuotaQuotaID(quotaQuotaID string) {
	o.QuotaQuotaID = quotaQuotaID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateQuotaQuotaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.QuotaQuota != nil {
		if err := r.SetBodyParam(o.QuotaQuota); err != nil {
			return err
		}
	}

	// path param QuotaQuotaId
	if err := r.SetPathParam("QuotaQuotaId", o.QuotaQuotaID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}