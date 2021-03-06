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

// NewGetDirectoryMetadataParams creates a new GetDirectoryMetadataParams object
// with the default values initialized.
func NewGetDirectoryMetadataParams() *GetDirectoryMetadataParams {
	var ()
	return &GetDirectoryMetadataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDirectoryMetadataParamsWithTimeout creates a new GetDirectoryMetadataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDirectoryMetadataParamsWithTimeout(timeout time.Duration) *GetDirectoryMetadataParams {
	var ()
	return &GetDirectoryMetadataParams{

		timeout: timeout,
	}
}

// NewGetDirectoryMetadataParamsWithContext creates a new GetDirectoryMetadataParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDirectoryMetadataParamsWithContext(ctx context.Context) *GetDirectoryMetadataParams {
	var ()
	return &GetDirectoryMetadataParams{

		Context: ctx,
	}
}

// NewGetDirectoryMetadataParamsWithHTTPClient creates a new GetDirectoryMetadataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDirectoryMetadataParamsWithHTTPClient(client *http.Client) *GetDirectoryMetadataParams {
	var ()
	return &GetDirectoryMetadataParams{
		HTTPClient: client,
	}
}

/*GetDirectoryMetadataParams contains all the parameters to send to the API endpoint
for the get directory metadata operation typically these are written to a http.Request
*/
type GetDirectoryMetadataParams struct {

	/*DirectoryMetadataPath
	  Directory path relative to /.

	*/
	DirectoryMetadataPath string
	/*Metadata
	  Show directory metadata.

	*/
	Metadata bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get directory metadata params
func (o *GetDirectoryMetadataParams) WithTimeout(timeout time.Duration) *GetDirectoryMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get directory metadata params
func (o *GetDirectoryMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get directory metadata params
func (o *GetDirectoryMetadataParams) WithContext(ctx context.Context) *GetDirectoryMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get directory metadata params
func (o *GetDirectoryMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get directory metadata params
func (o *GetDirectoryMetadataParams) WithHTTPClient(client *http.Client) *GetDirectoryMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get directory metadata params
func (o *GetDirectoryMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDirectoryMetadataPath adds the directoryMetadataPath to the get directory metadata params
func (o *GetDirectoryMetadataParams) WithDirectoryMetadataPath(directoryMetadataPath string) *GetDirectoryMetadataParams {
	o.SetDirectoryMetadataPath(directoryMetadataPath)
	return o
}

// SetDirectoryMetadataPath adds the directoryMetadataPath to the get directory metadata params
func (o *GetDirectoryMetadataParams) SetDirectoryMetadataPath(directoryMetadataPath string) {
	o.DirectoryMetadataPath = directoryMetadataPath
}

// WithMetadata adds the metadata to the get directory metadata params
func (o *GetDirectoryMetadataParams) WithMetadata(metadata bool) *GetDirectoryMetadataParams {
	o.SetMetadata(metadata)
	return o
}

// SetMetadata adds the metadata to the get directory metadata params
func (o *GetDirectoryMetadataParams) SetMetadata(metadata bool) {
	o.Metadata = metadata
}

// WriteToRequest writes these params to a swagger request
func (o *GetDirectoryMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param DirectoryMetadataPath
	if err := r.SetPathParam("DirectoryMetadataPath", o.DirectoryMetadataPath); err != nil {
		return err
	}

	// query param metadata
	qrMetadata := o.Metadata
	qMetadata := swag.FormatBool(qrMetadata)
	if qMetadata != "" {
		if err := r.SetQueryParam("metadata", qMetadata); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
