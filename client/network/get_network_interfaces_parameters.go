// Code generated by go-swagger; DO NOT EDIT.

package network

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

// NewGetNetworkInterfacesParams creates a new GetNetworkInterfacesParams object
// with the default values initialized.
func NewGetNetworkInterfacesParams() *GetNetworkInterfacesParams {
	var ()
	return &GetNetworkInterfacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkInterfacesParamsWithTimeout creates a new GetNetworkInterfacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkInterfacesParamsWithTimeout(timeout time.Duration) *GetNetworkInterfacesParams {
	var ()
	return &GetNetworkInterfacesParams{

		timeout: timeout,
	}
}

// NewGetNetworkInterfacesParamsWithContext creates a new GetNetworkInterfacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkInterfacesParamsWithContext(ctx context.Context) *GetNetworkInterfacesParams {
	var ()
	return &GetNetworkInterfacesParams{

		Context: ctx,
	}
}

// NewGetNetworkInterfacesParamsWithHTTPClient creates a new GetNetworkInterfacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkInterfacesParamsWithHTTPClient(client *http.Client) *GetNetworkInterfacesParams {
	var ()
	return &GetNetworkInterfacesParams{
		HTTPClient: client,
	}
}

/*GetNetworkInterfacesParams contains all the parameters to send to the API endpoint
for the get network interfaces operation typically these are written to a http.Request
*/
type GetNetworkInterfacesParams struct {

	/*AllocMethod
	  Filter addresses and owners by pool address allocation method.

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
	/*Lnns
	  Get a list of interfaces for the specified lnn.

	*/
	Lnns *string
	/*Network
	  Show interfaces associated with external and/or internal networks. Default is 'external'

	*/
	Network *string
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

// WithTimeout adds the timeout to the get network interfaces params
func (o *GetNetworkInterfacesParams) WithTimeout(timeout time.Duration) *GetNetworkInterfacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network interfaces params
func (o *GetNetworkInterfacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network interfaces params
func (o *GetNetworkInterfacesParams) WithContext(ctx context.Context) *GetNetworkInterfacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network interfaces params
func (o *GetNetworkInterfacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network interfaces params
func (o *GetNetworkInterfacesParams) WithHTTPClient(client *http.Client) *GetNetworkInterfacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network interfaces params
func (o *GetNetworkInterfacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllocMethod adds the allocMethod to the get network interfaces params
func (o *GetNetworkInterfacesParams) WithAllocMethod(allocMethod *string) *GetNetworkInterfacesParams {
	o.SetAllocMethod(allocMethod)
	return o
}

// SetAllocMethod adds the allocMethod to the get network interfaces params
func (o *GetNetworkInterfacesParams) SetAllocMethod(allocMethod *string) {
	o.AllocMethod = allocMethod
}

// WithDir adds the dir to the get network interfaces params
func (o *GetNetworkInterfacesParams) WithDir(dir *string) *GetNetworkInterfacesParams {
	o.SetDir(dir)
	return o
}

// SetDir adds the dir to the get network interfaces params
func (o *GetNetworkInterfacesParams) SetDir(dir *string) {
	o.Dir = dir
}

// WithLimit adds the limit to the get network interfaces params
func (o *GetNetworkInterfacesParams) WithLimit(limit *int64) *GetNetworkInterfacesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get network interfaces params
func (o *GetNetworkInterfacesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithLnns adds the lnns to the get network interfaces params
func (o *GetNetworkInterfacesParams) WithLnns(lnns *string) *GetNetworkInterfacesParams {
	o.SetLnns(lnns)
	return o
}

// SetLnns adds the lnns to the get network interfaces params
func (o *GetNetworkInterfacesParams) SetLnns(lnns *string) {
	o.Lnns = lnns
}

// WithNetwork adds the network to the get network interfaces params
func (o *GetNetworkInterfacesParams) WithNetwork(network *string) *GetNetworkInterfacesParams {
	o.SetNetwork(network)
	return o
}

// SetNetwork adds the network to the get network interfaces params
func (o *GetNetworkInterfacesParams) SetNetwork(network *string) {
	o.Network = network
}

// WithResume adds the resume to the get network interfaces params
func (o *GetNetworkInterfacesParams) WithResume(resume *string) *GetNetworkInterfacesParams {
	o.SetResume(resume)
	return o
}

// SetResume adds the resume to the get network interfaces params
func (o *GetNetworkInterfacesParams) SetResume(resume *string) {
	o.Resume = resume
}

// WithSort adds the sort to the get network interfaces params
func (o *GetNetworkInterfacesParams) WithSort(sort *string) *GetNetworkInterfacesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get network interfaces params
func (o *GetNetworkInterfacesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkInterfacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.Lnns != nil {

		// query param lnns
		var qrLnns string
		if o.Lnns != nil {
			qrLnns = *o.Lnns
		}
		qLnns := qrLnns
		if qLnns != "" {
			if err := r.SetQueryParam("lnns", qLnns); err != nil {
				return err
			}
		}

	}

	if o.Network != nil {

		// query param network
		var qrNetwork string
		if o.Network != nil {
			qrNetwork = *o.Network
		}
		qNetwork := qrNetwork
		if qNetwork != "" {
			if err := r.SetQueryParam("network", qNetwork); err != nil {
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
