// Code generated by go-swagger; DO NOT EDIT.

package protocols

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

// NewUpdateNdmpUserParams creates a new UpdateNdmpUserParams object
// with the default values initialized.
func NewUpdateNdmpUserParams() *UpdateNdmpUserParams {
	var ()
	return &UpdateNdmpUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNdmpUserParamsWithTimeout creates a new UpdateNdmpUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateNdmpUserParamsWithTimeout(timeout time.Duration) *UpdateNdmpUserParams {
	var ()
	return &UpdateNdmpUserParams{

		timeout: timeout,
	}
}

// NewUpdateNdmpUserParamsWithContext creates a new UpdateNdmpUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateNdmpUserParamsWithContext(ctx context.Context) *UpdateNdmpUserParams {
	var ()
	return &UpdateNdmpUserParams{

		Context: ctx,
	}
}

// NewUpdateNdmpUserParamsWithHTTPClient creates a new UpdateNdmpUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateNdmpUserParamsWithHTTPClient(client *http.Client) *UpdateNdmpUserParams {
	var ()
	return &UpdateNdmpUserParams{
		HTTPClient: client,
	}
}

/*UpdateNdmpUserParams contains all the parameters to send to the API endpoint
for the update ndmp user operation typically these are written to a http.Request
*/
type UpdateNdmpUserParams struct {

	/*NdmpUser*/
	NdmpUser *models.NdmpUser
	/*NdmpUserID
	  Modify the user

	*/
	NdmpUserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update ndmp user params
func (o *UpdateNdmpUserParams) WithTimeout(timeout time.Duration) *UpdateNdmpUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update ndmp user params
func (o *UpdateNdmpUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update ndmp user params
func (o *UpdateNdmpUserParams) WithContext(ctx context.Context) *UpdateNdmpUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update ndmp user params
func (o *UpdateNdmpUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update ndmp user params
func (o *UpdateNdmpUserParams) WithHTTPClient(client *http.Client) *UpdateNdmpUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update ndmp user params
func (o *UpdateNdmpUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNdmpUser adds the ndmpUser to the update ndmp user params
func (o *UpdateNdmpUserParams) WithNdmpUser(ndmpUser *models.NdmpUser) *UpdateNdmpUserParams {
	o.SetNdmpUser(ndmpUser)
	return o
}

// SetNdmpUser adds the ndmpUser to the update ndmp user params
func (o *UpdateNdmpUserParams) SetNdmpUser(ndmpUser *models.NdmpUser) {
	o.NdmpUser = ndmpUser
}

// WithNdmpUserID adds the ndmpUserID to the update ndmp user params
func (o *UpdateNdmpUserParams) WithNdmpUserID(ndmpUserID string) *UpdateNdmpUserParams {
	o.SetNdmpUserID(ndmpUserID)
	return o
}

// SetNdmpUserID adds the ndmpUserId to the update ndmp user params
func (o *UpdateNdmpUserParams) SetNdmpUserID(ndmpUserID string) {
	o.NdmpUserID = ndmpUserID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNdmpUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.NdmpUser != nil {
		if err := r.SetBodyParam(o.NdmpUser); err != nil {
			return err
		}
	}

	// path param NdmpUserId
	if err := r.SetPathParam("NdmpUserId", o.NdmpUserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
