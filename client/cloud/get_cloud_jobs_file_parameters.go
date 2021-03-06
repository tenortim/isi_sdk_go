// Code generated by go-swagger; DO NOT EDIT.

package cloud

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

// NewGetCloudJobsFileParams creates a new GetCloudJobsFileParams object
// with the default values initialized.
func NewGetCloudJobsFileParams() *GetCloudJobsFileParams {
	var ()
	return &GetCloudJobsFileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCloudJobsFileParamsWithTimeout creates a new GetCloudJobsFileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCloudJobsFileParamsWithTimeout(timeout time.Duration) *GetCloudJobsFileParams {
	var ()
	return &GetCloudJobsFileParams{

		timeout: timeout,
	}
}

// NewGetCloudJobsFileParamsWithContext creates a new GetCloudJobsFileParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCloudJobsFileParamsWithContext(ctx context.Context) *GetCloudJobsFileParams {
	var ()
	return &GetCloudJobsFileParams{

		Context: ctx,
	}
}

// NewGetCloudJobsFileParamsWithHTTPClient creates a new GetCloudJobsFileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCloudJobsFileParamsWithHTTPClient(client *http.Client) *GetCloudJobsFileParams {
	var ()
	return &GetCloudJobsFileParams{
		HTTPClient: client,
	}
}

/*GetCloudJobsFileParams contains all the parameters to send to the API endpoint
for the get cloud jobs file operation typically these are written to a http.Request
*/
type GetCloudJobsFileParams struct {

	/*CloudJobsFileID
	  Retrieve files associated with a cloudpool job.

	*/
	CloudJobsFileID string
	/*Dir
	  The direction of the sort.

	*/
	Dir *string
	/*Limit
	  Number of files to display; range from 1 to 100000.

	*/
	Limit *int64
	/*Offset
	  Specifies the starting entry to be displayed.

	*/
	Offset *int64
	/*Page
	  Used with limit option. If exists, specifies the starting page to display where page size is specified by limit. This option will be deprecated; please use offset option instead.

	*/
	Page *int64
	/*Resume
	  Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options).

	*/
	Resume *string
	/*Sort
	  The field that will be used for sorting. Output is limited to maximum of 100000 files.

	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cloud jobs file params
func (o *GetCloudJobsFileParams) WithTimeout(timeout time.Duration) *GetCloudJobsFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cloud jobs file params
func (o *GetCloudJobsFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cloud jobs file params
func (o *GetCloudJobsFileParams) WithContext(ctx context.Context) *GetCloudJobsFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cloud jobs file params
func (o *GetCloudJobsFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cloud jobs file params
func (o *GetCloudJobsFileParams) WithHTTPClient(client *http.Client) *GetCloudJobsFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cloud jobs file params
func (o *GetCloudJobsFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudJobsFileID adds the cloudJobsFileID to the get cloud jobs file params
func (o *GetCloudJobsFileParams) WithCloudJobsFileID(cloudJobsFileID string) *GetCloudJobsFileParams {
	o.SetCloudJobsFileID(cloudJobsFileID)
	return o
}

// SetCloudJobsFileID adds the cloudJobsFileId to the get cloud jobs file params
func (o *GetCloudJobsFileParams) SetCloudJobsFileID(cloudJobsFileID string) {
	o.CloudJobsFileID = cloudJobsFileID
}

// WithDir adds the dir to the get cloud jobs file params
func (o *GetCloudJobsFileParams) WithDir(dir *string) *GetCloudJobsFileParams {
	o.SetDir(dir)
	return o
}

// SetDir adds the dir to the get cloud jobs file params
func (o *GetCloudJobsFileParams) SetDir(dir *string) {
	o.Dir = dir
}

// WithLimit adds the limit to the get cloud jobs file params
func (o *GetCloudJobsFileParams) WithLimit(limit *int64) *GetCloudJobsFileParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get cloud jobs file params
func (o *GetCloudJobsFileParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get cloud jobs file params
func (o *GetCloudJobsFileParams) WithOffset(offset *int64) *GetCloudJobsFileParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get cloud jobs file params
func (o *GetCloudJobsFileParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPage adds the page to the get cloud jobs file params
func (o *GetCloudJobsFileParams) WithPage(page *int64) *GetCloudJobsFileParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get cloud jobs file params
func (o *GetCloudJobsFileParams) SetPage(page *int64) {
	o.Page = page
}

// WithResume adds the resume to the get cloud jobs file params
func (o *GetCloudJobsFileParams) WithResume(resume *string) *GetCloudJobsFileParams {
	o.SetResume(resume)
	return o
}

// SetResume adds the resume to the get cloud jobs file params
func (o *GetCloudJobsFileParams) SetResume(resume *string) {
	o.Resume = resume
}

// WithSort adds the sort to the get cloud jobs file params
func (o *GetCloudJobsFileParams) WithSort(sort *string) *GetCloudJobsFileParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get cloud jobs file params
func (o *GetCloudJobsFileParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetCloudJobsFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param CloudJobsFileId
	if err := r.SetPathParam("CloudJobsFileId", o.CloudJobsFileID); err != nil {
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

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
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
