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
)

// NewListRolePrivilegesParams creates a new ListRolePrivilegesParams object
// with the default values initialized.
func NewListRolePrivilegesParams() *ListRolePrivilegesParams {
	var ()
	return &ListRolePrivilegesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListRolePrivilegesParamsWithTimeout creates a new ListRolePrivilegesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListRolePrivilegesParamsWithTimeout(timeout time.Duration) *ListRolePrivilegesParams {
	var ()
	return &ListRolePrivilegesParams{

		timeout: timeout,
	}
}

// NewListRolePrivilegesParamsWithContext creates a new ListRolePrivilegesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListRolePrivilegesParamsWithContext(ctx context.Context) *ListRolePrivilegesParams {
	var ()
	return &ListRolePrivilegesParams{

		Context: ctx,
	}
}

// NewListRolePrivilegesParamsWithHTTPClient creates a new ListRolePrivilegesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListRolePrivilegesParamsWithHTTPClient(client *http.Client) *ListRolePrivilegesParams {
	var ()
	return &ListRolePrivilegesParams{
		HTTPClient: client,
	}
}

/*ListRolePrivilegesParams contains all the parameters to send to the API endpoint
for the list role privileges operation typically these are written to a http.Request
*/
type ListRolePrivilegesParams struct {

	/*Role*/
	Role string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list role privileges params
func (o *ListRolePrivilegesParams) WithTimeout(timeout time.Duration) *ListRolePrivilegesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list role privileges params
func (o *ListRolePrivilegesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list role privileges params
func (o *ListRolePrivilegesParams) WithContext(ctx context.Context) *ListRolePrivilegesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list role privileges params
func (o *ListRolePrivilegesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list role privileges params
func (o *ListRolePrivilegesParams) WithHTTPClient(client *http.Client) *ListRolePrivilegesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list role privileges params
func (o *ListRolePrivilegesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRole adds the role to the list role privileges params
func (o *ListRolePrivilegesParams) WithRole(role string) *ListRolePrivilegesParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the list role privileges params
func (o *ListRolePrivilegesParams) SetRole(role string) {
	o.Role = role
}

// WriteToRequest writes these params to a swagger request
func (o *ListRolePrivilegesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Role
	if err := r.SetPathParam("Role", o.Role); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
