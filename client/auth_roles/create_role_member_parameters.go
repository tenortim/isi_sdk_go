// Code generated by go-swagger; DO NOT EDIT.

package auth_roles

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

// NewCreateRoleMemberParams creates a new CreateRoleMemberParams object
// with the default values initialized.
func NewCreateRoleMemberParams() *CreateRoleMemberParams {
	var ()
	return &CreateRoleMemberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRoleMemberParamsWithTimeout creates a new CreateRoleMemberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRoleMemberParamsWithTimeout(timeout time.Duration) *CreateRoleMemberParams {
	var ()
	return &CreateRoleMemberParams{

		timeout: timeout,
	}
}

// NewCreateRoleMemberParamsWithContext creates a new CreateRoleMemberParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRoleMemberParamsWithContext(ctx context.Context) *CreateRoleMemberParams {
	var ()
	return &CreateRoleMemberParams{

		Context: ctx,
	}
}

// NewCreateRoleMemberParamsWithHTTPClient creates a new CreateRoleMemberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRoleMemberParamsWithHTTPClient(client *http.Client) *CreateRoleMemberParams {
	var ()
	return &CreateRoleMemberParams{
		HTTPClient: client,
	}
}

/*CreateRoleMemberParams contains all the parameters to send to the API endpoint
for the create role member operation typically these are written to a http.Request
*/
type CreateRoleMemberParams struct {

	/*Role*/
	Role string
	/*RoleMember*/
	RoleMember *models.GroupMember

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create role member params
func (o *CreateRoleMemberParams) WithTimeout(timeout time.Duration) *CreateRoleMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create role member params
func (o *CreateRoleMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create role member params
func (o *CreateRoleMemberParams) WithContext(ctx context.Context) *CreateRoleMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create role member params
func (o *CreateRoleMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create role member params
func (o *CreateRoleMemberParams) WithHTTPClient(client *http.Client) *CreateRoleMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create role member params
func (o *CreateRoleMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRole adds the role to the create role member params
func (o *CreateRoleMemberParams) WithRole(role string) *CreateRoleMemberParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the create role member params
func (o *CreateRoleMemberParams) SetRole(role string) {
	o.Role = role
}

// WithRoleMember adds the roleMember to the create role member params
func (o *CreateRoleMemberParams) WithRoleMember(roleMember *models.GroupMember) *CreateRoleMemberParams {
	o.SetRoleMember(roleMember)
	return o
}

// SetRoleMember adds the roleMember to the create role member params
func (o *CreateRoleMemberParams) SetRoleMember(roleMember *models.GroupMember) {
	o.RoleMember = roleMember
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRoleMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Role
	if err := r.SetPathParam("Role", o.Role); err != nil {
		return err
	}

	if o.RoleMember != nil {
		if err := r.SetBodyParam(o.RoleMember); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
