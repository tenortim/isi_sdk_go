// Code generated by go-swagger; DO NOT EDIT.

package job

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

// NewGetJobTypesParams creates a new GetJobTypesParams object
// with the default values initialized.
func NewGetJobTypesParams() *GetJobTypesParams {
	var ()
	return &GetJobTypesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetJobTypesParamsWithTimeout creates a new GetJobTypesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetJobTypesParamsWithTimeout(timeout time.Duration) *GetJobTypesParams {
	var ()
	return &GetJobTypesParams{

		timeout: timeout,
	}
}

// NewGetJobTypesParamsWithContext creates a new GetJobTypesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetJobTypesParamsWithContext(ctx context.Context) *GetJobTypesParams {
	var ()
	return &GetJobTypesParams{

		Context: ctx,
	}
}

// NewGetJobTypesParamsWithHTTPClient creates a new GetJobTypesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetJobTypesParamsWithHTTPClient(client *http.Client) *GetJobTypesParams {
	var ()
	return &GetJobTypesParams{
		HTTPClient: client,
	}
}

/*GetJobTypesParams contains all the parameters to send to the API endpoint
for the get job types operation typically these are written to a http.Request
*/
type GetJobTypesParams struct {

	/*Dir
	  The direction of the sort.

	*/
	Dir *string
	/*ShowAll
	  Whether to show all job types, including hidden ones.  Defaults to false.

	*/
	ShowAll *bool
	/*Sort
	  The field that will be used for sorting.

	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get job types params
func (o *GetJobTypesParams) WithTimeout(timeout time.Duration) *GetJobTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get job types params
func (o *GetJobTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get job types params
func (o *GetJobTypesParams) WithContext(ctx context.Context) *GetJobTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get job types params
func (o *GetJobTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get job types params
func (o *GetJobTypesParams) WithHTTPClient(client *http.Client) *GetJobTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get job types params
func (o *GetJobTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDir adds the dir to the get job types params
func (o *GetJobTypesParams) WithDir(dir *string) *GetJobTypesParams {
	o.SetDir(dir)
	return o
}

// SetDir adds the dir to the get job types params
func (o *GetJobTypesParams) SetDir(dir *string) {
	o.Dir = dir
}

// WithShowAll adds the showAll to the get job types params
func (o *GetJobTypesParams) WithShowAll(showAll *bool) *GetJobTypesParams {
	o.SetShowAll(showAll)
	return o
}

// SetShowAll adds the showAll to the get job types params
func (o *GetJobTypesParams) SetShowAll(showAll *bool) {
	o.ShowAll = showAll
}

// WithSort adds the sort to the get job types params
func (o *GetJobTypesParams) WithSort(sort *string) *GetJobTypesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get job types params
func (o *GetJobTypesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetJobTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ShowAll != nil {

		// query param show_all
		var qrShowAll bool
		if o.ShowAll != nil {
			qrShowAll = *o.ShowAll
		}
		qShowAll := swag.FormatBool(qrShowAll)
		if qShowAll != "" {
			if err := r.SetQueryParam("show_all", qShowAll); err != nil {
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