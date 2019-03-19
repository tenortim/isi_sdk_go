// Code generated by go-swagger; DO NOT EDIT.

package auth

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

// NewCreateSettingsKrb5DomainParams creates a new CreateSettingsKrb5DomainParams object
// with the default values initialized.
func NewCreateSettingsKrb5DomainParams() *CreateSettingsKrb5DomainParams {
	var ()
	return &CreateSettingsKrb5DomainParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSettingsKrb5DomainParamsWithTimeout creates a new CreateSettingsKrb5DomainParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSettingsKrb5DomainParamsWithTimeout(timeout time.Duration) *CreateSettingsKrb5DomainParams {
	var ()
	return &CreateSettingsKrb5DomainParams{

		timeout: timeout,
	}
}

// NewCreateSettingsKrb5DomainParamsWithContext creates a new CreateSettingsKrb5DomainParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSettingsKrb5DomainParamsWithContext(ctx context.Context) *CreateSettingsKrb5DomainParams {
	var ()
	return &CreateSettingsKrb5DomainParams{

		Context: ctx,
	}
}

// NewCreateSettingsKrb5DomainParamsWithHTTPClient creates a new CreateSettingsKrb5DomainParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSettingsKrb5DomainParamsWithHTTPClient(client *http.Client) *CreateSettingsKrb5DomainParams {
	var ()
	return &CreateSettingsKrb5DomainParams{
		HTTPClient: client,
	}
}

/*CreateSettingsKrb5DomainParams contains all the parameters to send to the API endpoint
for the create settings krb5 domain operation typically these are written to a http.Request
*/
type CreateSettingsKrb5DomainParams struct {

	/*SettingsKrb5Domain*/
	SettingsKrb5Domain *models.SettingsKrb5DomainCreateParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create settings krb5 domain params
func (o *CreateSettingsKrb5DomainParams) WithTimeout(timeout time.Duration) *CreateSettingsKrb5DomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create settings krb5 domain params
func (o *CreateSettingsKrb5DomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create settings krb5 domain params
func (o *CreateSettingsKrb5DomainParams) WithContext(ctx context.Context) *CreateSettingsKrb5DomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create settings krb5 domain params
func (o *CreateSettingsKrb5DomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create settings krb5 domain params
func (o *CreateSettingsKrb5DomainParams) WithHTTPClient(client *http.Client) *CreateSettingsKrb5DomainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create settings krb5 domain params
func (o *CreateSettingsKrb5DomainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSettingsKrb5Domain adds the settingsKrb5Domain to the create settings krb5 domain params
func (o *CreateSettingsKrb5DomainParams) WithSettingsKrb5Domain(settingsKrb5Domain *models.SettingsKrb5DomainCreateParams) *CreateSettingsKrb5DomainParams {
	o.SetSettingsKrb5Domain(settingsKrb5Domain)
	return o
}

// SetSettingsKrb5Domain adds the settingsKrb5Domain to the create settings krb5 domain params
func (o *CreateSettingsKrb5DomainParams) SetSettingsKrb5Domain(settingsKrb5Domain *models.SettingsKrb5DomainCreateParams) {
	o.SettingsKrb5Domain = settingsKrb5Domain
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSettingsKrb5DomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SettingsKrb5Domain != nil {
		if err := r.SetBodyParam(o.SettingsKrb5Domain); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
