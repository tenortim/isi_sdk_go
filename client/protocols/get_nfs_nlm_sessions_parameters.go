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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetNfsNlmSessionsParams creates a new GetNfsNlmSessionsParams object
// with the default values initialized.
func NewGetNfsNlmSessionsParams() *GetNfsNlmSessionsParams {
	var ()
	return &GetNfsNlmSessionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNfsNlmSessionsParamsWithTimeout creates a new GetNfsNlmSessionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNfsNlmSessionsParamsWithTimeout(timeout time.Duration) *GetNfsNlmSessionsParams {
	var ()
	return &GetNfsNlmSessionsParams{

		timeout: timeout,
	}
}

// NewGetNfsNlmSessionsParamsWithContext creates a new GetNfsNlmSessionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNfsNlmSessionsParamsWithContext(ctx context.Context) *GetNfsNlmSessionsParams {
	var ()
	return &GetNfsNlmSessionsParams{

		Context: ctx,
	}
}

// NewGetNfsNlmSessionsParamsWithHTTPClient creates a new GetNfsNlmSessionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNfsNlmSessionsParamsWithHTTPClient(client *http.Client) *GetNfsNlmSessionsParams {
	var ()
	return &GetNfsNlmSessionsParams{
		HTTPClient: client,
	}
}

/*GetNfsNlmSessionsParams contains all the parameters to send to the API endpoint
for the get nfs nlm sessions operation typically these are written to a http.Request
*/
type GetNfsNlmSessionsParams struct {

	/*Dir
	  The direction of the sort.

	*/
	Dir *string
	/*IP
	  An IP address for which NSM has client records

	*/
	IP *string
	/*Limit
	  Return no more than this many results at once (see resume).

	*/
	Limit *int64
	/*Sort
	  The field that will be used for sorting.

	*/
	Sort *string
	/*Zone
	  Represents an extant auth zone

	*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get nfs nlm sessions params
func (o *GetNfsNlmSessionsParams) WithTimeout(timeout time.Duration) *GetNfsNlmSessionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nfs nlm sessions params
func (o *GetNfsNlmSessionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nfs nlm sessions params
func (o *GetNfsNlmSessionsParams) WithContext(ctx context.Context) *GetNfsNlmSessionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nfs nlm sessions params
func (o *GetNfsNlmSessionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nfs nlm sessions params
func (o *GetNfsNlmSessionsParams) WithHTTPClient(client *http.Client) *GetNfsNlmSessionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nfs nlm sessions params
func (o *GetNfsNlmSessionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDir adds the dir to the get nfs nlm sessions params
func (o *GetNfsNlmSessionsParams) WithDir(dir *string) *GetNfsNlmSessionsParams {
	o.SetDir(dir)
	return o
}

// SetDir adds the dir to the get nfs nlm sessions params
func (o *GetNfsNlmSessionsParams) SetDir(dir *string) {
	o.Dir = dir
}

// WithIP adds the ip to the get nfs nlm sessions params
func (o *GetNfsNlmSessionsParams) WithIP(ip *string) *GetNfsNlmSessionsParams {
	o.SetIP(ip)
	return o
}

// SetIP adds the ip to the get nfs nlm sessions params
func (o *GetNfsNlmSessionsParams) SetIP(ip *string) {
	o.IP = ip
}

// WithLimit adds the limit to the get nfs nlm sessions params
func (o *GetNfsNlmSessionsParams) WithLimit(limit *int64) *GetNfsNlmSessionsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get nfs nlm sessions params
func (o *GetNfsNlmSessionsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithSort adds the sort to the get nfs nlm sessions params
func (o *GetNfsNlmSessionsParams) WithSort(sort *string) *GetNfsNlmSessionsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get nfs nlm sessions params
func (o *GetNfsNlmSessionsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithZone adds the zone to the get nfs nlm sessions params
func (o *GetNfsNlmSessionsParams) WithZone(zone *string) *GetNfsNlmSessionsParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the get nfs nlm sessions params
func (o *GetNfsNlmSessionsParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *GetNfsNlmSessionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Dir != nil {

		// query param dir
		var qrDir string
		if o.Dir != nil {
			qrDir = *o.Dir
		}
		qDir := qrDir
		if qDir != "" {
			if err := r.SetQueryParam("dir", qDir); err != nil {
				return err
			}
		}

	}

	if o.IP != nil {

		// query param ip
		var qrIP string
		if o.IP != nil {
			qrIP = *o.IP
		}
		qIP := qrIP
		if qIP != "" {
			if err := r.SetQueryParam("ip", qIP); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
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