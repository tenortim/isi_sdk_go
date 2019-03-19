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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetFileAttributesParams creates a new GetFileAttributesParams object
// with the default values initialized.
func NewGetFileAttributesParams() *GetFileAttributesParams {
	var ()
	return &GetFileAttributesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFileAttributesParamsWithTimeout creates a new GetFileAttributesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFileAttributesParamsWithTimeout(timeout time.Duration) *GetFileAttributesParams {
	var ()
	return &GetFileAttributesParams{

		timeout: timeout,
	}
}

// NewGetFileAttributesParamsWithContext creates a new GetFileAttributesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFileAttributesParamsWithContext(ctx context.Context) *GetFileAttributesParams {
	var ()
	return &GetFileAttributesParams{

		Context: ctx,
	}
}

// NewGetFileAttributesParamsWithHTTPClient creates a new GetFileAttributesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFileAttributesParamsWithHTTPClient(client *http.Client) *GetFileAttributesParams {
	var ()
	return &GetFileAttributesParams{
		HTTPClient: client,
	}
}

/*GetFileAttributesParams contains all the parameters to send to the API endpoint
for the get file attributes operation typically these are written to a http.Request
*/
type GetFileAttributesParams struct {

	/*FilePath
	  File path relative to /.

	*/
	FilePath string
	/*IfModifiedSince
	  Returns only files that were modified since the specified time. If no files were modified since this time, a 304 message is returned.

	*/
	IfModifiedSince *string
	/*IfUnmodifiedSince
	  Returns only files that were not modified since the specified time. If there are no unmodified files since this time, a 412 message is returned to indicate that the precondition failed.

	*/
	IfUnmodifiedSince *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get file attributes params
func (o *GetFileAttributesParams) WithTimeout(timeout time.Duration) *GetFileAttributesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get file attributes params
func (o *GetFileAttributesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get file attributes params
func (o *GetFileAttributesParams) WithContext(ctx context.Context) *GetFileAttributesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get file attributes params
func (o *GetFileAttributesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get file attributes params
func (o *GetFileAttributesParams) WithHTTPClient(client *http.Client) *GetFileAttributesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get file attributes params
func (o *GetFileAttributesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilePath adds the filePath to the get file attributes params
func (o *GetFileAttributesParams) WithFilePath(filePath string) *GetFileAttributesParams {
	o.SetFilePath(filePath)
	return o
}

// SetFilePath adds the filePath to the get file attributes params
func (o *GetFileAttributesParams) SetFilePath(filePath string) {
	o.FilePath = filePath
}

// WithIfModifiedSince adds the ifModifiedSince to the get file attributes params
func (o *GetFileAttributesParams) WithIfModifiedSince(ifModifiedSince *string) *GetFileAttributesParams {
	o.SetIfModifiedSince(ifModifiedSince)
	return o
}

// SetIfModifiedSince adds the ifModifiedSince to the get file attributes params
func (o *GetFileAttributesParams) SetIfModifiedSince(ifModifiedSince *string) {
	o.IfModifiedSince = ifModifiedSince
}

// WithIfUnmodifiedSince adds the ifUnmodifiedSince to the get file attributes params
func (o *GetFileAttributesParams) WithIfUnmodifiedSince(ifUnmodifiedSince *string) *GetFileAttributesParams {
	o.SetIfUnmodifiedSince(ifUnmodifiedSince)
	return o
}

// SetIfUnmodifiedSince adds the ifUnmodifiedSince to the get file attributes params
func (o *GetFileAttributesParams) SetIfUnmodifiedSince(ifUnmodifiedSince *string) {
	o.IfUnmodifiedSince = ifUnmodifiedSince
}

// WriteToRequest writes these params to a swagger request
func (o *GetFileAttributesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param FilePath
	if err := r.SetPathParam("FilePath", o.FilePath); err != nil {
		return err
	}

	if o.IfModifiedSince != nil {

		// header param If-Modified-Since
		if err := r.SetHeaderParam("If-Modified-Since", *o.IfModifiedSince); err != nil {
			return err
		}

	}

	if o.IfUnmodifiedSince != nil {

		// header param If-Unmodified-Since
		if err := r.SetHeaderParam("If-Unmodified-Since", *o.IfUnmodifiedSince); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}