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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetMappingUsersLookupParams creates a new GetMappingUsersLookupParams object
// with the default values initialized.
func NewGetMappingUsersLookupParams() *GetMappingUsersLookupParams {
	var ()
	return &GetMappingUsersLookupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMappingUsersLookupParamsWithTimeout creates a new GetMappingUsersLookupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMappingUsersLookupParamsWithTimeout(timeout time.Duration) *GetMappingUsersLookupParams {
	var ()
	return &GetMappingUsersLookupParams{

		timeout: timeout,
	}
}

// NewGetMappingUsersLookupParamsWithContext creates a new GetMappingUsersLookupParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMappingUsersLookupParamsWithContext(ctx context.Context) *GetMappingUsersLookupParams {
	var ()
	return &GetMappingUsersLookupParams{

		Context: ctx,
	}
}

// NewGetMappingUsersLookupParamsWithHTTPClient creates a new GetMappingUsersLookupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMappingUsersLookupParamsWithHTTPClient(client *http.Client) *GetMappingUsersLookupParams {
	var ()
	return &GetMappingUsersLookupParams{
		HTTPClient: client,
	}
}

/*GetMappingUsersLookupParams contains all the parameters to send to the API endpoint
for the get mapping users lookup operation typically these are written to a http.Request
*/
type GetMappingUsersLookupParams struct {

	/*Gid
	  The IDs of the groups that the user belongs to.

	*/
	Gid []int64
	/*KerberosPrincipal
	  The Kerberos principal name, of the form user@realm.

	*/
	KerberosPrincipal *string
	/*PrimaryGid
	  The user's primary group ID.

	*/
	PrimaryGid *int64
	/*UID
	  The user ID.

	*/
	UID *int64
	/*User
	  The user name.

	*/
	User *string
	/*Zone
	  The zone the user belongs to.

	*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get mapping users lookup params
func (o *GetMappingUsersLookupParams) WithTimeout(timeout time.Duration) *GetMappingUsersLookupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get mapping users lookup params
func (o *GetMappingUsersLookupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get mapping users lookup params
func (o *GetMappingUsersLookupParams) WithContext(ctx context.Context) *GetMappingUsersLookupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get mapping users lookup params
func (o *GetMappingUsersLookupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get mapping users lookup params
func (o *GetMappingUsersLookupParams) WithHTTPClient(client *http.Client) *GetMappingUsersLookupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get mapping users lookup params
func (o *GetMappingUsersLookupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGid adds the gid to the get mapping users lookup params
func (o *GetMappingUsersLookupParams) WithGid(gid []int64) *GetMappingUsersLookupParams {
	o.SetGid(gid)
	return o
}

// SetGid adds the gid to the get mapping users lookup params
func (o *GetMappingUsersLookupParams) SetGid(gid []int64) {
	o.Gid = gid
}

// WithKerberosPrincipal adds the kerberosPrincipal to the get mapping users lookup params
func (o *GetMappingUsersLookupParams) WithKerberosPrincipal(kerberosPrincipal *string) *GetMappingUsersLookupParams {
	o.SetKerberosPrincipal(kerberosPrincipal)
	return o
}

// SetKerberosPrincipal adds the kerberosPrincipal to the get mapping users lookup params
func (o *GetMappingUsersLookupParams) SetKerberosPrincipal(kerberosPrincipal *string) {
	o.KerberosPrincipal = kerberosPrincipal
}

// WithPrimaryGid adds the primaryGid to the get mapping users lookup params
func (o *GetMappingUsersLookupParams) WithPrimaryGid(primaryGid *int64) *GetMappingUsersLookupParams {
	o.SetPrimaryGid(primaryGid)
	return o
}

// SetPrimaryGid adds the primaryGid to the get mapping users lookup params
func (o *GetMappingUsersLookupParams) SetPrimaryGid(primaryGid *int64) {
	o.PrimaryGid = primaryGid
}

// WithUID adds the uid to the get mapping users lookup params
func (o *GetMappingUsersLookupParams) WithUID(uid *int64) *GetMappingUsersLookupParams {
	o.SetUID(uid)
	return o
}

// SetUID adds the uid to the get mapping users lookup params
func (o *GetMappingUsersLookupParams) SetUID(uid *int64) {
	o.UID = uid
}

// WithUser adds the user to the get mapping users lookup params
func (o *GetMappingUsersLookupParams) WithUser(user *string) *GetMappingUsersLookupParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the get mapping users lookup params
func (o *GetMappingUsersLookupParams) SetUser(user *string) {
	o.User = user
}

// WithZone adds the zone to the get mapping users lookup params
func (o *GetMappingUsersLookupParams) WithZone(zone *string) *GetMappingUsersLookupParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the get mapping users lookup params
func (o *GetMappingUsersLookupParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *GetMappingUsersLookupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	var valuesGid []string
	for _, v := range o.Gid {
		valuesGid = append(valuesGid, swag.FormatInt64(v))
	}

	joinedGid := swag.JoinByFormat(valuesGid, "")
	// query array param gid
	if err := r.SetQueryParam("gid", joinedGid...); err != nil {
		return err
	}

	if o.KerberosPrincipal != nil {

		// query param kerberos_principal
		var qrKerberosPrincipal string
		if o.KerberosPrincipal != nil {
			qrKerberosPrincipal = *o.KerberosPrincipal
		}
		qKerberosPrincipal := qrKerberosPrincipal
		if qKerberosPrincipal != "" {
			if err := r.SetQueryParam("kerberos_principal", qKerberosPrincipal); err != nil {
				return err
			}
		}

	}

	if o.PrimaryGid != nil {

		// query param primary_gid
		var qrPrimaryGid int64
		if o.PrimaryGid != nil {
			qrPrimaryGid = *o.PrimaryGid
		}
		qPrimaryGid := swag.FormatInt64(qrPrimaryGid)
		if qPrimaryGid != "" {
			if err := r.SetQueryParam("primary_gid", qPrimaryGid); err != nil {
				return err
			}
		}

	}

	if o.UID != nil {

		// query param uid
		var qrUID int64
		if o.UID != nil {
			qrUID = *o.UID
		}
		qUID := swag.FormatInt64(qrUID)
		if qUID != "" {
			if err := r.SetQueryParam("uid", qUID); err != nil {
				return err
			}
		}

	}

	if o.User != nil {

		// query param user
		var qrUser string
		if o.User != nil {
			qrUser = *o.User
		}
		qUser := qrUser
		if qUser != "" {
			if err := r.SetQueryParam("user", qUser); err != nil {
				return err
			}
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
