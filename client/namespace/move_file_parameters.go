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

// NewMoveFileParams creates a new MoveFileParams object
// with the default values initialized.
func NewMoveFileParams() *MoveFileParams {
	var ()
	return &MoveFileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMoveFileParamsWithTimeout creates a new MoveFileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMoveFileParamsWithTimeout(timeout time.Duration) *MoveFileParams {
	var ()
	return &MoveFileParams{

		timeout: timeout,
	}
}

// NewMoveFileParamsWithContext creates a new MoveFileParams object
// with the default values initialized, and the ability to set a context for a request
func NewMoveFileParamsWithContext(ctx context.Context) *MoveFileParams {
	var ()
	return &MoveFileParams{

		Context: ctx,
	}
}

// NewMoveFileParamsWithHTTPClient creates a new MoveFileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMoveFileParamsWithHTTPClient(client *http.Client) *MoveFileParams {
	var ()
	return &MoveFileParams{
		HTTPClient: client,
	}
}

/*MoveFileParams contains all the parameters to send to the API endpoint
for the move file operation typically these are written to a http.Request
*/
type MoveFileParams struct {

	/*FilePath
	  File path relative to /.

	*/
	FilePath string
	/*XIsiIfsSetLocation
	  Specifies the full path for the destination file.

	*/
	XIsiIfsSetLocation string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the move file params
func (o *MoveFileParams) WithTimeout(timeout time.Duration) *MoveFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the move file params
func (o *MoveFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the move file params
func (o *MoveFileParams) WithContext(ctx context.Context) *MoveFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the move file params
func (o *MoveFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the move file params
func (o *MoveFileParams) WithHTTPClient(client *http.Client) *MoveFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the move file params
func (o *MoveFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilePath adds the filePath to the move file params
func (o *MoveFileParams) WithFilePath(filePath string) *MoveFileParams {
	o.SetFilePath(filePath)
	return o
}

// SetFilePath adds the filePath to the move file params
func (o *MoveFileParams) SetFilePath(filePath string) {
	o.FilePath = filePath
}

// WithXIsiIfsSetLocation adds the xIsiIfsSetLocation to the move file params
func (o *MoveFileParams) WithXIsiIfsSetLocation(xIsiIfsSetLocation string) *MoveFileParams {
	o.SetXIsiIfsSetLocation(xIsiIfsSetLocation)
	return o
}

// SetXIsiIfsSetLocation adds the xIsiIfsSetLocation to the move file params
func (o *MoveFileParams) SetXIsiIfsSetLocation(xIsiIfsSetLocation string) {
	o.XIsiIfsSetLocation = xIsiIfsSetLocation
}

// WriteToRequest writes these params to a swagger request
func (o *MoveFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param FilePath
	if err := r.SetPathParam("FilePath", o.FilePath); err != nil {
		return err
	}

	// header param x-isi-ifs-set-location
	if err := r.SetHeaderParam("x-isi-ifs-set-location", o.XIsiIfsSetLocation); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
