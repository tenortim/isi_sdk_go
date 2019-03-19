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
)

// NewDeleteDirectoryParams creates a new DeleteDirectoryParams object
// with the default values initialized.
func NewDeleteDirectoryParams() *DeleteDirectoryParams {
	var ()
	return &DeleteDirectoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDirectoryParamsWithTimeout creates a new DeleteDirectoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDirectoryParamsWithTimeout(timeout time.Duration) *DeleteDirectoryParams {
	var ()
	return &DeleteDirectoryParams{

		timeout: timeout,
	}
}

// NewDeleteDirectoryParamsWithContext creates a new DeleteDirectoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDirectoryParamsWithContext(ctx context.Context) *DeleteDirectoryParams {
	var ()
	return &DeleteDirectoryParams{

		Context: ctx,
	}
}

// NewDeleteDirectoryParamsWithHTTPClient creates a new DeleteDirectoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDirectoryParamsWithHTTPClient(client *http.Client) *DeleteDirectoryParams {
	var ()
	return &DeleteDirectoryParams{
		HTTPClient: client,
	}
}

/*DeleteDirectoryParams contains all the parameters to send to the API endpoint
for the delete directory operation typically these are written to a http.Request
*/
type DeleteDirectoryParams struct {

	/*DirectoryPath
	  Directory path relative to /.

	*/
	DirectoryPath string
	/*Recursive
	  Deletes directories recursively, when set to true.

	*/
	Recursive *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete directory params
func (o *DeleteDirectoryParams) WithTimeout(timeout time.Duration) *DeleteDirectoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete directory params
func (o *DeleteDirectoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete directory params
func (o *DeleteDirectoryParams) WithContext(ctx context.Context) *DeleteDirectoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete directory params
func (o *DeleteDirectoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete directory params
func (o *DeleteDirectoryParams) WithHTTPClient(client *http.Client) *DeleteDirectoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete directory params
func (o *DeleteDirectoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDirectoryPath adds the directoryPath to the delete directory params
func (o *DeleteDirectoryParams) WithDirectoryPath(directoryPath string) *DeleteDirectoryParams {
	o.SetDirectoryPath(directoryPath)
	return o
}

// SetDirectoryPath adds the directoryPath to the delete directory params
func (o *DeleteDirectoryParams) SetDirectoryPath(directoryPath string) {
	o.DirectoryPath = directoryPath
}

// WithRecursive adds the recursive to the delete directory params
func (o *DeleteDirectoryParams) WithRecursive(recursive *bool) *DeleteDirectoryParams {
	o.SetRecursive(recursive)
	return o
}

// SetRecursive adds the recursive to the delete directory params
func (o *DeleteDirectoryParams) SetRecursive(recursive *bool) {
	o.Recursive = recursive
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDirectoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param DirectoryPath
	if err := r.SetPathParam("DirectoryPath", o.DirectoryPath); err != nil {
		return err
	}

	if o.Recursive != nil {

		// query param recursive
		var qrRecursive bool
		if o.Recursive != nil {
			qrRecursive = *o.Recursive
		}
		qRecursive := swag.FormatBool(qrRecursive)
		if qRecursive != "" {
			if err := r.SetQueryParam("recursive", qRecursive); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
