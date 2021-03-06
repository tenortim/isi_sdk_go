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

// NewSetDirectoryMetadataParams creates a new SetDirectoryMetadataParams object
// with the default values initialized.
func NewSetDirectoryMetadataParams() *SetDirectoryMetadataParams {
	var ()
	return &SetDirectoryMetadataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetDirectoryMetadataParamsWithTimeout creates a new SetDirectoryMetadataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetDirectoryMetadataParamsWithTimeout(timeout time.Duration) *SetDirectoryMetadataParams {
	var ()
	return &SetDirectoryMetadataParams{

		timeout: timeout,
	}
}

// NewSetDirectoryMetadataParamsWithContext creates a new SetDirectoryMetadataParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetDirectoryMetadataParamsWithContext(ctx context.Context) *SetDirectoryMetadataParams {
	var ()
	return &SetDirectoryMetadataParams{

		Context: ctx,
	}
}

// NewSetDirectoryMetadataParamsWithHTTPClient creates a new SetDirectoryMetadataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetDirectoryMetadataParamsWithHTTPClient(client *http.Client) *SetDirectoryMetadataParams {
	var ()
	return &SetDirectoryMetadataParams{
		HTTPClient: client,
	}
}

/*SetDirectoryMetadataParams contains all the parameters to send to the API endpoint
for the set directory metadata operation typically these are written to a http.Request
*/
type SetDirectoryMetadataParams struct {

	/*DirectoryMetadata
	  Directory metadata parameters model.

	*/
	DirectoryMetadata *models.NamespaceMetadata
	/*DirectoryMetadataPath
	  Directory path relative to /.

	*/
	DirectoryMetadataPath string
	/*Metadata
	  Set directory metadata.

	*/
	Metadata bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set directory metadata params
func (o *SetDirectoryMetadataParams) WithTimeout(timeout time.Duration) *SetDirectoryMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set directory metadata params
func (o *SetDirectoryMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set directory metadata params
func (o *SetDirectoryMetadataParams) WithContext(ctx context.Context) *SetDirectoryMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set directory metadata params
func (o *SetDirectoryMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set directory metadata params
func (o *SetDirectoryMetadataParams) WithHTTPClient(client *http.Client) *SetDirectoryMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set directory metadata params
func (o *SetDirectoryMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDirectoryMetadata adds the directoryMetadata to the set directory metadata params
func (o *SetDirectoryMetadataParams) WithDirectoryMetadata(directoryMetadata *models.NamespaceMetadata) *SetDirectoryMetadataParams {
	o.SetDirectoryMetadata(directoryMetadata)
	return o
}

// SetDirectoryMetadata adds the directoryMetadata to the set directory metadata params
func (o *SetDirectoryMetadataParams) SetDirectoryMetadata(directoryMetadata *models.NamespaceMetadata) {
	o.DirectoryMetadata = directoryMetadata
}

// WithDirectoryMetadataPath adds the directoryMetadataPath to the set directory metadata params
func (o *SetDirectoryMetadataParams) WithDirectoryMetadataPath(directoryMetadataPath string) *SetDirectoryMetadataParams {
	o.SetDirectoryMetadataPath(directoryMetadataPath)
	return o
}

// SetDirectoryMetadataPath adds the directoryMetadataPath to the set directory metadata params
func (o *SetDirectoryMetadataParams) SetDirectoryMetadataPath(directoryMetadataPath string) {
	o.DirectoryMetadataPath = directoryMetadataPath
}

// WithMetadata adds the metadata to the set directory metadata params
func (o *SetDirectoryMetadataParams) WithMetadata(metadata bool) *SetDirectoryMetadataParams {
	o.SetMetadata(metadata)
	return o
}

// SetMetadata adds the metadata to the set directory metadata params
func (o *SetDirectoryMetadataParams) SetMetadata(metadata bool) {
	o.Metadata = metadata
}

// WriteToRequest writes these params to a swagger request
func (o *SetDirectoryMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DirectoryMetadata != nil {
		if err := r.SetBodyParam(o.DirectoryMetadata); err != nil {
			return err
		}
	}

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
