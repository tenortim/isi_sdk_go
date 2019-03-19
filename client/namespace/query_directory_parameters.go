// Code generated by go-swagger; DO NOT EDIT.

package namespace

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

	models "github.com/tenortim/isi_sdk_go/models"
)

// NewQueryDirectoryParams creates a new QueryDirectoryParams object
// with the default values initialized.
func NewQueryDirectoryParams() *QueryDirectoryParams {
	var ()
	return &QueryDirectoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQueryDirectoryParamsWithTimeout creates a new QueryDirectoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQueryDirectoryParamsWithTimeout(timeout time.Duration) *QueryDirectoryParams {
	var ()
	return &QueryDirectoryParams{

		timeout: timeout,
	}
}

// NewQueryDirectoryParamsWithContext creates a new QueryDirectoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewQueryDirectoryParamsWithContext(ctx context.Context) *QueryDirectoryParams {
	var ()
	return &QueryDirectoryParams{

		Context: ctx,
	}
}

// NewQueryDirectoryParamsWithHTTPClient creates a new QueryDirectoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQueryDirectoryParamsWithHTTPClient(client *http.Client) *QueryDirectoryParams {
	var ()
	return &QueryDirectoryParams{
		HTTPClient: client,
	}
}

/*QueryDirectoryParams contains all the parameters to send to the API endpoint
for the query directory operation typically these are written to a http.Request
*/
type QueryDirectoryParams struct {

	/*DirectoryQuery
	  Directory query parameters model.

	*/
	DirectoryQuery *models.DirectoryQuery
	/*QueryPath
	  Directory path relative to /.

	*/
	QueryPath string
	/*Detail
	  Specifies which object attributes are displayed. If the detail parameter is excluded, only the name of the object is returned.

	*/
	Detail *string
	/*Dir
	  Specifies the sort direction. This value can be either ascending (ASC) or descending (DESC).

	*/
	Dir *string
	/*Hidden
	  Specifies if hidden objects are returned.

	*/
	Hidden *bool
	/*Limit
	  Specifies the maximum number of objects to send to the client. You can set the value to a negative number to retrieve all objects.

	*/
	Limit *int64
	/*MaxDepth
	  Specifies the maximum directory level depth to search for objects.

	*/
	MaxDepth *int64
	/*Query
	  Enable directory query.

	*/
	Query bool
	/*Resume
	  Specifies a token to return in the JSON result to indicate when there is a next page.

	*/
	Resume *string
	/*Sort
	  Specifies one or more attributes to sort on the directory entries.

	*/
	Sort *string
	/*Type
	  Specifies the object type to return, which can be one of the following values: container, object, pipe, character_device, block_device, symbolic_link, socket, or whiteout_file.

	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the query directory params
func (o *QueryDirectoryParams) WithTimeout(timeout time.Duration) *QueryDirectoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query directory params
func (o *QueryDirectoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query directory params
func (o *QueryDirectoryParams) WithContext(ctx context.Context) *QueryDirectoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query directory params
func (o *QueryDirectoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query directory params
func (o *QueryDirectoryParams) WithHTTPClient(client *http.Client) *QueryDirectoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query directory params
func (o *QueryDirectoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDirectoryQuery adds the directoryQuery to the query directory params
func (o *QueryDirectoryParams) WithDirectoryQuery(directoryQuery *models.DirectoryQuery) *QueryDirectoryParams {
	o.SetDirectoryQuery(directoryQuery)
	return o
}

// SetDirectoryQuery adds the directoryQuery to the query directory params
func (o *QueryDirectoryParams) SetDirectoryQuery(directoryQuery *models.DirectoryQuery) {
	o.DirectoryQuery = directoryQuery
}

// WithQueryPath adds the queryPath to the query directory params
func (o *QueryDirectoryParams) WithQueryPath(queryPath string) *QueryDirectoryParams {
	o.SetQueryPath(queryPath)
	return o
}

// SetQueryPath adds the queryPath to the query directory params
func (o *QueryDirectoryParams) SetQueryPath(queryPath string) {
	o.QueryPath = queryPath
}

// WithDetail adds the detail to the query directory params
func (o *QueryDirectoryParams) WithDetail(detail *string) *QueryDirectoryParams {
	o.SetDetail(detail)
	return o
}

// SetDetail adds the detail to the query directory params
func (o *QueryDirectoryParams) SetDetail(detail *string) {
	o.Detail = detail
}

// WithDir adds the dir to the query directory params
func (o *QueryDirectoryParams) WithDir(dir *string) *QueryDirectoryParams {
	o.SetDir(dir)
	return o
}

// SetDir adds the dir to the query directory params
func (o *QueryDirectoryParams) SetDir(dir *string) {
	o.Dir = dir
}

// WithHidden adds the hidden to the query directory params
func (o *QueryDirectoryParams) WithHidden(hidden *bool) *QueryDirectoryParams {
	o.SetHidden(hidden)
	return o
}

// SetHidden adds the hidden to the query directory params
func (o *QueryDirectoryParams) SetHidden(hidden *bool) {
	o.Hidden = hidden
}

// WithLimit adds the limit to the query directory params
func (o *QueryDirectoryParams) WithLimit(limit *int64) *QueryDirectoryParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the query directory params
func (o *QueryDirectoryParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMaxDepth adds the maxDepth to the query directory params
func (o *QueryDirectoryParams) WithMaxDepth(maxDepth *int64) *QueryDirectoryParams {
	o.SetMaxDepth(maxDepth)
	return o
}

// SetMaxDepth adds the maxDepth to the query directory params
func (o *QueryDirectoryParams) SetMaxDepth(maxDepth *int64) {
	o.MaxDepth = maxDepth
}

// WithQuery adds the query to the query directory params
func (o *QueryDirectoryParams) WithQuery(query bool) *QueryDirectoryParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the query directory params
func (o *QueryDirectoryParams) SetQuery(query bool) {
	o.Query = query
}

// WithResume adds the resume to the query directory params
func (o *QueryDirectoryParams) WithResume(resume *string) *QueryDirectoryParams {
	o.SetResume(resume)
	return o
}

// SetResume adds the resume to the query directory params
func (o *QueryDirectoryParams) SetResume(resume *string) {
	o.Resume = resume
}

// WithSort adds the sort to the query directory params
func (o *QueryDirectoryParams) WithSort(sort *string) *QueryDirectoryParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the query directory params
func (o *QueryDirectoryParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithType adds the typeVar to the query directory params
func (o *QueryDirectoryParams) WithType(typeVar *string) *QueryDirectoryParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the query directory params
func (o *QueryDirectoryParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *QueryDirectoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DirectoryQuery != nil {
		if err := r.SetBodyParam(o.DirectoryQuery); err != nil {
			return err
		}
	}

	// path param QueryPath
	if err := r.SetPathParam("QueryPath", o.QueryPath); err != nil {
		return err
	}

	if o.Detail != nil {

		// query param detail
		var qrDetail string
		if o.Detail != nil {
			qrDetail = *o.Detail
		}
		qDetail := qrDetail
		if qDetail != "" {
			if err := r.SetQueryParam("detail", qDetail); err != nil {
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

	if o.Hidden != nil {

		// query param hidden
		var qrHidden bool
		if o.Hidden != nil {
			qrHidden = *o.Hidden
		}
		qHidden := swag.FormatBool(qrHidden)
		if qHidden != "" {
			if err := r.SetQueryParam("hidden", qHidden); err != nil {
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

	if o.MaxDepth != nil {

		// query param max-depth
		var qrMaxDepth int64
		if o.MaxDepth != nil {
			qrMaxDepth = *o.MaxDepth
		}
		qMaxDepth := swag.FormatInt64(qrMaxDepth)
		if qMaxDepth != "" {
			if err := r.SetQueryParam("max-depth", qMaxDepth); err != nil {
				return err
			}
		}

	}

	// query param query
	qrQuery := o.Query
	qQuery := swag.FormatBool(qrQuery)
	if qQuery != "" {
		if err := r.SetQueryParam("query", qQuery); err != nil {
			return err
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

	if o.Type != nil {

		// query param type
		var qrType string
		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {
			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
