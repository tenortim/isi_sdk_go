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
)

// NewListSmbLogLevelFiltersParams creates a new ListSmbLogLevelFiltersParams object
// with the default values initialized.
func NewListSmbLogLevelFiltersParams() *ListSmbLogLevelFiltersParams {
	var ()
	return &ListSmbLogLevelFiltersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSmbLogLevelFiltersParamsWithTimeout creates a new ListSmbLogLevelFiltersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSmbLogLevelFiltersParamsWithTimeout(timeout time.Duration) *ListSmbLogLevelFiltersParams {
	var ()
	return &ListSmbLogLevelFiltersParams{

		timeout: timeout,
	}
}

// NewListSmbLogLevelFiltersParamsWithContext creates a new ListSmbLogLevelFiltersParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSmbLogLevelFiltersParamsWithContext(ctx context.Context) *ListSmbLogLevelFiltersParams {
	var ()
	return &ListSmbLogLevelFiltersParams{

		Context: ctx,
	}
}

// NewListSmbLogLevelFiltersParamsWithHTTPClient creates a new ListSmbLogLevelFiltersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSmbLogLevelFiltersParamsWithHTTPClient(client *http.Client) *ListSmbLogLevelFiltersParams {
	var ()
	return &ListSmbLogLevelFiltersParams{
		HTTPClient: client,
	}
}

/*ListSmbLogLevelFiltersParams contains all the parameters to send to the API endpoint
for the list smb log level filters operation typically these are written to a http.Request
*/
type ListSmbLogLevelFiltersParams struct {

	/*Dir
	  The direction of the sort.

	*/
	Dir *string
	/*Level
	  Only return results with a given level.

	*/
	Level *string
	/*Sort
	  The field that will be used for sorting.

	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list smb log level filters params
func (o *ListSmbLogLevelFiltersParams) WithTimeout(timeout time.Duration) *ListSmbLogLevelFiltersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list smb log level filters params
func (o *ListSmbLogLevelFiltersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list smb log level filters params
func (o *ListSmbLogLevelFiltersParams) WithContext(ctx context.Context) *ListSmbLogLevelFiltersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list smb log level filters params
func (o *ListSmbLogLevelFiltersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list smb log level filters params
func (o *ListSmbLogLevelFiltersParams) WithHTTPClient(client *http.Client) *ListSmbLogLevelFiltersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list smb log level filters params
func (o *ListSmbLogLevelFiltersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDir adds the dir to the list smb log level filters params
func (o *ListSmbLogLevelFiltersParams) WithDir(dir *string) *ListSmbLogLevelFiltersParams {
	o.SetDir(dir)
	return o
}

// SetDir adds the dir to the list smb log level filters params
func (o *ListSmbLogLevelFiltersParams) SetDir(dir *string) {
	o.Dir = dir
}

// WithLevel adds the level to the list smb log level filters params
func (o *ListSmbLogLevelFiltersParams) WithLevel(level *string) *ListSmbLogLevelFiltersParams {
	o.SetLevel(level)
	return o
}

// SetLevel adds the level to the list smb log level filters params
func (o *ListSmbLogLevelFiltersParams) SetLevel(level *string) {
	o.Level = level
}

// WithSort adds the sort to the list smb log level filters params
func (o *ListSmbLogLevelFiltersParams) WithSort(sort *string) *ListSmbLogLevelFiltersParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the list smb log level filters params
func (o *ListSmbLogLevelFiltersParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *ListSmbLogLevelFiltersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Level != nil {

		// query param level
		var qrLevel string
		if o.Level != nil {
			qrLevel = *o.Level
		}
		qLevel := qrLevel
		if qLevel != "" {
			if err := r.SetQueryParam("level", qLevel); err != nil {
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