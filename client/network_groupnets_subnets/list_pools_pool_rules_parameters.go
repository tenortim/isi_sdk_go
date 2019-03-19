// Code generated by go-swagger; DO NOT EDIT.

package network_groupnets_subnets

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

// NewListPoolsPoolRulesParams creates a new ListPoolsPoolRulesParams object
// with the default values initialized.
func NewListPoolsPoolRulesParams() *ListPoolsPoolRulesParams {
	var ()
	return &ListPoolsPoolRulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListPoolsPoolRulesParamsWithTimeout creates a new ListPoolsPoolRulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListPoolsPoolRulesParamsWithTimeout(timeout time.Duration) *ListPoolsPoolRulesParams {
	var ()
	return &ListPoolsPoolRulesParams{

		timeout: timeout,
	}
}

// NewListPoolsPoolRulesParamsWithContext creates a new ListPoolsPoolRulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListPoolsPoolRulesParamsWithContext(ctx context.Context) *ListPoolsPoolRulesParams {
	var ()
	return &ListPoolsPoolRulesParams{

		Context: ctx,
	}
}

// NewListPoolsPoolRulesParamsWithHTTPClient creates a new ListPoolsPoolRulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListPoolsPoolRulesParamsWithHTTPClient(client *http.Client) *ListPoolsPoolRulesParams {
	var ()
	return &ListPoolsPoolRulesParams{
		HTTPClient: client,
	}
}

/*ListPoolsPoolRulesParams contains all the parameters to send to the API endpoint
for the list pools pool rules operation typically these are written to a http.Request
*/
type ListPoolsPoolRulesParams struct {

	/*Groupnet*/
	Groupnet string
	/*Pool*/
	Pool string
	/*Subnet*/
	Subnet string
	/*Dir
	  The direction of the sort.

	*/
	Dir *string
	/*Limit
	  Return no more than this many results at once (see resume).

	*/
	Limit *int64
	/*Resume
	  Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options).

	*/
	Resume *string
	/*Sort
	  The field that will be used for sorting.

	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list pools pool rules params
func (o *ListPoolsPoolRulesParams) WithTimeout(timeout time.Duration) *ListPoolsPoolRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list pools pool rules params
func (o *ListPoolsPoolRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list pools pool rules params
func (o *ListPoolsPoolRulesParams) WithContext(ctx context.Context) *ListPoolsPoolRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list pools pool rules params
func (o *ListPoolsPoolRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list pools pool rules params
func (o *ListPoolsPoolRulesParams) WithHTTPClient(client *http.Client) *ListPoolsPoolRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list pools pool rules params
func (o *ListPoolsPoolRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupnet adds the groupnet to the list pools pool rules params
func (o *ListPoolsPoolRulesParams) WithGroupnet(groupnet string) *ListPoolsPoolRulesParams {
	o.SetGroupnet(groupnet)
	return o
}

// SetGroupnet adds the groupnet to the list pools pool rules params
func (o *ListPoolsPoolRulesParams) SetGroupnet(groupnet string) {
	o.Groupnet = groupnet
}

// WithPool adds the pool to the list pools pool rules params
func (o *ListPoolsPoolRulesParams) WithPool(pool string) *ListPoolsPoolRulesParams {
	o.SetPool(pool)
	return o
}

// SetPool adds the pool to the list pools pool rules params
func (o *ListPoolsPoolRulesParams) SetPool(pool string) {
	o.Pool = pool
}

// WithSubnet adds the subnet to the list pools pool rules params
func (o *ListPoolsPoolRulesParams) WithSubnet(subnet string) *ListPoolsPoolRulesParams {
	o.SetSubnet(subnet)
	return o
}

// SetSubnet adds the subnet to the list pools pool rules params
func (o *ListPoolsPoolRulesParams) SetSubnet(subnet string) {
	o.Subnet = subnet
}

// WithDir adds the dir to the list pools pool rules params
func (o *ListPoolsPoolRulesParams) WithDir(dir *string) *ListPoolsPoolRulesParams {
	o.SetDir(dir)
	return o
}

// SetDir adds the dir to the list pools pool rules params
func (o *ListPoolsPoolRulesParams) SetDir(dir *string) {
	o.Dir = dir
}

// WithLimit adds the limit to the list pools pool rules params
func (o *ListPoolsPoolRulesParams) WithLimit(limit *int64) *ListPoolsPoolRulesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list pools pool rules params
func (o *ListPoolsPoolRulesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithResume adds the resume to the list pools pool rules params
func (o *ListPoolsPoolRulesParams) WithResume(resume *string) *ListPoolsPoolRulesParams {
	o.SetResume(resume)
	return o
}

// SetResume adds the resume to the list pools pool rules params
func (o *ListPoolsPoolRulesParams) SetResume(resume *string) {
	o.Resume = resume
}

// WithSort adds the sort to the list pools pool rules params
func (o *ListPoolsPoolRulesParams) WithSort(sort *string) *ListPoolsPoolRulesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the list pools pool rules params
func (o *ListPoolsPoolRulesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *ListPoolsPoolRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Groupnet
	if err := r.SetPathParam("Groupnet", o.Groupnet); err != nil {
		return err
	}

	// path param Pool
	if err := r.SetPathParam("Pool", o.Pool); err != nil {
		return err
	}

	// path param Subnet
	if err := r.SetPathParam("Subnet", o.Subnet); err != nil {
		return err
	}

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

	if o.Resume != nil {

		// query param resume
		var qrResume string
		if o.Resume != nil {
			qrResume = *o.Resume
		}
		qResume := qrResume
		if qResume != "" {
			if err := r.SetQueryParam("resume", qResume); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
