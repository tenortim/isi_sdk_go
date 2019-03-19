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

// NewCreateQuotaQuotaParams creates a new CreateQuotaQuotaParams object
// with the default values initialized.
func NewCreateQuotaQuotaParams() *CreateQuotaQuotaParams {
	var ()
	return &CreateQuotaQuotaParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateQuotaQuotaParamsWithTimeout creates a new CreateQuotaQuotaParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateQuotaQuotaParamsWithTimeout(timeout time.Duration) *CreateQuotaQuotaParams {
	var ()
	return &CreateQuotaQuotaParams{

		timeout: timeout,
	}
}

// NewCreateQuotaQuotaParamsWithContext creates a new CreateQuotaQuotaParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateQuotaQuotaParamsWithContext(ctx context.Context) *CreateQuotaQuotaParams {
	var ()
	return &CreateQuotaQuotaParams{

		Context: ctx,
	}
}

// NewCreateQuotaQuotaParamsWithHTTPClient creates a new CreateQuotaQuotaParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateQuotaQuotaParamsWithHTTPClient(client *http.Client) *CreateQuotaQuotaParams {
	var ()
	return &CreateQuotaQuotaParams{
		HTTPClient: client,
	}
}

/*CreateQuotaQuotaParams contains all the parameters to send to the API endpoint
for the create quota quota operation typically these are written to a http.Request
*/
type CreateQuotaQuotaParams struct {

	/*QuotaQuota*/
	QuotaQuota *models.QuotaQuotaCreateParams
	/*Zone
	  Optional named zone to use for user and group resolution.

	*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create quota quota params
func (o *CreateQuotaQuotaParams) WithTimeout(timeout time.Duration) *CreateQuotaQuotaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create quota quota params
func (o *CreateQuotaQuotaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create quota quota params
func (o *CreateQuotaQuotaParams) WithContext(ctx context.Context) *CreateQuotaQuotaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create quota quota params
func (o *CreateQuotaQuotaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create quota quota params
func (o *CreateQuotaQuotaParams) WithHTTPClient(client *http.Client) *CreateQuotaQuotaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create quota quota params
func (o *CreateQuotaQuotaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuotaQuota adds the quotaQuota to the create quota quota params
func (o *CreateQuotaQuotaParams) WithQuotaQuota(quotaQuota *models.QuotaQuotaCreateParams) *CreateQuotaQuotaParams {
	o.SetQuotaQuota(quotaQuota)
	return o
}

// SetQuotaQuota adds the quotaQuota to the create quota quota params
func (o *CreateQuotaQuotaParams) SetQuotaQuota(quotaQuota *models.QuotaQuotaCreateParams) {
	o.QuotaQuota = quotaQuota
}

// WithZone adds the zone to the create quota quota params
func (o *CreateQuotaQuotaParams) WithZone(zone *string) *CreateQuotaQuotaParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the create quota quota params
func (o *CreateQuotaQuotaParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *CreateQuotaQuotaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.QuotaQuota != nil {
		if err := r.SetBodyParam(o.QuotaQuota); err != nil {
			return err
		}
	}

	if o.Zone != nil {

		// query param zone
		var qrZone string
		if o.Zone != nil {
			qrZone = *o.Zone
		}
		qZone := qrZone
		if qZone != "" {
			if err := r.SetQueryParam("zone", qZone); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}