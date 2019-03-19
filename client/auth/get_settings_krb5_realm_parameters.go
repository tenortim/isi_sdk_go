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
)

// NewGetSettingsKrb5RealmParams creates a new GetSettingsKrb5RealmParams object
// with the default values initialized.
func NewGetSettingsKrb5RealmParams() *GetSettingsKrb5RealmParams {
	var ()
	return &GetSettingsKrb5RealmParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSettingsKrb5RealmParamsWithTimeout creates a new GetSettingsKrb5RealmParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSettingsKrb5RealmParamsWithTimeout(timeout time.Duration) *GetSettingsKrb5RealmParams {
	var ()
	return &GetSettingsKrb5RealmParams{

		timeout: timeout,
	}
}

// NewGetSettingsKrb5RealmParamsWithContext creates a new GetSettingsKrb5RealmParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSettingsKrb5RealmParamsWithContext(ctx context.Context) *GetSettingsKrb5RealmParams {
	var ()
	return &GetSettingsKrb5RealmParams{

		Context: ctx,
	}
}

// NewGetSettingsKrb5RealmParamsWithHTTPClient creates a new GetSettingsKrb5RealmParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSettingsKrb5RealmParamsWithHTTPClient(client *http.Client) *GetSettingsKrb5RealmParams {
	var ()
	return &GetSettingsKrb5RealmParams{
		HTTPClient: client,
	}
}

/*GetSettingsKrb5RealmParams contains all the parameters to send to the API endpoint
for the get settings krb5 realm operation typically these are written to a http.Request
*/
type GetSettingsKrb5RealmParams struct {

	/*SettingsKrb5RealmID
	  Retrieve the krb5 settings for realms.

	*/
	SettingsKrb5RealmID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get settings krb5 realm params
func (o *GetSettingsKrb5RealmParams) WithTimeout(timeout time.Duration) *GetSettingsKrb5RealmParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get settings krb5 realm params
func (o *GetSettingsKrb5RealmParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get settings krb5 realm params
func (o *GetSettingsKrb5RealmParams) WithContext(ctx context.Context) *GetSettingsKrb5RealmParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get settings krb5 realm params
func (o *GetSettingsKrb5RealmParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get settings krb5 realm params
func (o *GetSettingsKrb5RealmParams) WithHTTPClient(client *http.Client) *GetSettingsKrb5RealmParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get settings krb5 realm params
func (o *GetSettingsKrb5RealmParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSettingsKrb5RealmID adds the settingsKrb5RealmID to the get settings krb5 realm params
func (o *GetSettingsKrb5RealmParams) WithSettingsKrb5RealmID(settingsKrb5RealmID string) *GetSettingsKrb5RealmParams {
	o.SetSettingsKrb5RealmID(settingsKrb5RealmID)
	return o
}

// SetSettingsKrb5RealmID adds the settingsKrb5RealmId to the get settings krb5 realm params
func (o *GetSettingsKrb5RealmParams) SetSettingsKrb5RealmID(settingsKrb5RealmID string) {
	o.SettingsKrb5RealmID = settingsKrb5RealmID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSettingsKrb5RealmParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param SettingsKrb5RealmId
	if err := r.SetPathParam("SettingsKrb5RealmId", o.SettingsKrb5RealmID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}