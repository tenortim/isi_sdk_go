// Code generated by go-swagger; DO NOT EDIT.

package network_groupnets

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

// NewListSubnetsSubnetPoolsParams creates a new ListSubnetsSubnetPoolsParams object
// with the default values initialized.
func NewListSubnetsSubnetPoolsParams() *ListSubnetsSubnetPoolsParams {
	var ()
	return &ListSubnetsSubnetPoolsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSubnetsSubnetPoolsParamsWithTimeout creates a new ListSubnetsSubnetPoolsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSubnetsSubnetPoolsParamsWithTimeout(timeout time.Duration) *ListSubnetsSubnetPoolsParams {
	var ()
	return &ListSubnetsSubnetPoolsParams{

		timeout: timeout,
	}
}

// NewListSubnetsSubnetPoolsParamsWithContext creates a new ListSubnetsSubnetPoolsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSubnetsSubnetPoolsParamsWithContext(ctx context.Context) *ListSubnetsSubnetPoolsParams {
	var ()
	return &ListSubnetsSubnetPoolsParams{

		Context: ctx,
	}
}

// NewListSubnetsSubnetPoolsParamsWithHTTPClient creates a new ListSubnetsSubnetPoolsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSubnetsSubnetPoolsParamsWithHTTPClient(client *http.Client) *ListSubnetsSubnetPoolsParams {
	var ()
	return &ListSubnetsSubnetPoolsParams{
		HTTPClient: client,
	}
}

/*ListSubnetsSubnetPoolsParams contains all the parameters to send to the API endpoint
for the list subnets subnet pools operation typically these are written to a http.Request
*/
type ListSubnetsSubnetPoolsParams struct {

	/*Groupnet*/
	Groupnet string
	/*Subnet*/
	Subnet string
	/*AccessZone
	  If specified, only pools with this zone name will be returned.

	*/
	AccessZone *string
	/*AllocMethod
	  If specified, only pools with this allocation type will be returned.

	*/
	AllocMethod *string
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

// WithTimeout adds the timeout to the list subnets subnet pools params
func (o *ListSubnetsSubnetPoolsParams) WithTimeout(timeout time.Duration) *ListSubnetsSubnetPoolsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list subnets subnet pools params
func (o *ListSubnetsSubnetPoolsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list subnets subnet pools params
func (o *ListSubnetsSubnetPoolsParams) WithContext(ctx context.Context) *ListSubnetsSubnetPoolsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list subnets subnet pools params
func (o *ListSubnetsSubnetPoolsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list subnets subnet pools params
func (o *ListSubnetsSubnetPoolsParams) WithHTTPClient(client *http.Client) *ListSubnetsSubnetPoolsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list subnets subnet pools params
func (o *ListSubnetsSubnetPoolsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupnet adds the groupnet to the list subnets subnet pools params
func (o *ListSubnetsSubnetPoolsParams) WithGroupnet(groupnet string) *ListSubnetsSubnetPoolsParams {
	o.SetGroupnet(groupnet)
	return o
}

// SetGroupnet adds the groupnet to the list subnets subnet pools params
func (o *ListSubnetsSubnetPoolsParams) SetGroupnet(groupnet string) {
	o.Groupnet = groupnet
}

// WithSubnet adds the subnet to the list subnets subnet pools params
func (o *ListSubnetsSubnetPoolsParams) WithSubnet(subnet string) *ListSubnetsSubnetPoolsParams {
	o.SetSubnet(subnet)
	return o
}

// SetSubnet adds the subnet to the list subnets subnet pools params
func (o *ListSubnetsSubnetPoolsParams) SetSubnet(subnet string) {
	o.Subnet = subnet
}

// WithAccessZone adds the accessZone to the list subnets subnet pools params
func (o *ListSubnetsSubnetPoolsParams) WithAccessZone(accessZone *string) *ListSubnetsSubnetPoolsParams {
	o.SetAccessZone(accessZone)
	return o
}

// SetAccessZone adds the accessZone to the list subnets subnet pools params
func (o *ListSubnetsSubnetPoolsParams) SetAccessZone(accessZone *string) {
	o.AccessZone = accessZone
}

// WithAllocMethod adds the allocMethod to the list subnets subnet pools params
func (o *ListSubnetsSubnetPoolsParams) WithAllocMethod(allocMethod *string) *ListSubnetsSubnetPoolsParams {
	o.SetAllocMethod(allocMethod)
	return o
}

// SetAllocMethod adds the allocMethod to the list subnets subnet pools params
func (o *ListSubnetsSubnetPoolsParams) SetAllocMethod(allocMethod *string) {
	o.AllocMethod = allocMethod
}

// WithDir adds the dir to the list subnets subnet pools params
func (o *ListSubnetsSubnetPoolsParams) WithDir(dir *string) *ListSubnetsSubnetPoolsParams {
	o.SetDir(dir)
	return o
}

// SetDir adds the dir to the list subnets subnet pools params
func (o *ListSubnetsSubnetPoolsParams) SetDir(dir *string) {
	o.Dir = dir
}

// WithLimit adds the limit to the list subnets subnet pools params
func (o *ListSubnetsSubnetPoolsParams) WithLimit(limit *int64) *ListSubnetsSubnetPoolsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list subnets subnet pools params
func (o *ListSubnetsSubnetPoolsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithResume adds the resume to the list subnets subnet pools params
func (o *ListSubnetsSubnetPoolsParams) WithResume(resume *string) *ListSubnetsSubnetPoolsParams {
	o.SetResume(resume)
	return o
}

// SetResume adds the resume to the list subnets subnet pools params
func (o *ListSubnetsSubnetPoolsParams) SetResume(resume *string) {
	o.Resume = resume
}

// WithSort adds the sort to the list subnets subnet pools params
func (o *ListSubnetsSubnetPoolsParams) WithSort(sort *string) *ListSubnetsSubnetPoolsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the list subnets subnet pools params
func (o *ListSubnetsSubnetPoolsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *ListSubnetsSubnetPoolsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Groupnet
	if err := r.SetPathParam("Groupnet", o.Groupnet); err != nil {
		return err
	}

	// path param Subnet
	if err := r.SetPathParam("Subnet", o.Subnet); err != nil {
		return err
	}

	if o.AccessZone != nil {

		// query param access_zone
		var qrAccessZone string
		if o.AccessZone != nil {
			qrAccessZone = *o.AccessZone
		}
		qAccessZone := qrAccessZone
		if qAccessZone != "" {
			if err := r.SetQueryParam("access_zone", qAccessZone); err != nil {
				return err
			}
		}

	}

	if o.AllocMethod != nil {

		// query param alloc_method
		var qrAllocMethod string
		if o.AllocMethod != nil {
			qrAllocMethod = *o.AllocMethod
		}
		qAllocMethod := qrAllocMethod
		if qAllocMethod != "" {
			if err := r.SetQueryParam("alloc_method", qAllocMethod); err != nil {
				return err
			}
		}

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
