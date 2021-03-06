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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteNdmpDumpdateParams creates a new DeleteNdmpDumpdateParams object
// with the default values initialized.
func NewDeleteNdmpDumpdateParams() *DeleteNdmpDumpdateParams {
	var ()
	return &DeleteNdmpDumpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNdmpDumpdateParamsWithTimeout creates a new DeleteNdmpDumpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteNdmpDumpdateParamsWithTimeout(timeout time.Duration) *DeleteNdmpDumpdateParams {
	var ()
	return &DeleteNdmpDumpdateParams{

		timeout: timeout,
	}
}

// NewDeleteNdmpDumpdateParamsWithContext creates a new DeleteNdmpDumpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteNdmpDumpdateParamsWithContext(ctx context.Context) *DeleteNdmpDumpdateParams {
	var ()
	return &DeleteNdmpDumpdateParams{

		Context: ctx,
	}
}

// NewDeleteNdmpDumpdateParamsWithHTTPClient creates a new DeleteNdmpDumpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteNdmpDumpdateParamsWithHTTPClient(client *http.Client) *DeleteNdmpDumpdateParams {
	var ()
	return &DeleteNdmpDumpdateParams{
		HTTPClient: client,
	}
}

/*DeleteNdmpDumpdateParams contains all the parameters to send to the API endpoint
for the delete ndmp dumpdate operation typically these are written to a http.Request
*/
type DeleteNdmpDumpdateParams struct {

	/*NdmpDumpdateID
	  Delete dumpdates entries of a specified path.

	*/
	NdmpDumpdateID string
	/*Level
	  Level is an input from 0 to 10.

	*/
	Level *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete ndmp dumpdate params
func (o *DeleteNdmpDumpdateParams) WithTimeout(timeout time.Duration) *DeleteNdmpDumpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete ndmp dumpdate params
func (o *DeleteNdmpDumpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete ndmp dumpdate params
func (o *DeleteNdmpDumpdateParams) WithContext(ctx context.Context) *DeleteNdmpDumpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete ndmp dumpdate params
func (o *DeleteNdmpDumpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete ndmp dumpdate params
func (o *DeleteNdmpDumpdateParams) WithHTTPClient(client *http.Client) *DeleteNdmpDumpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete ndmp dumpdate params
func (o *DeleteNdmpDumpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNdmpDumpdateID adds the ndmpDumpdateID to the delete ndmp dumpdate params
func (o *DeleteNdmpDumpdateParams) WithNdmpDumpdateID(ndmpDumpdateID string) *DeleteNdmpDumpdateParams {
	o.SetNdmpDumpdateID(ndmpDumpdateID)
	return o
}

// SetNdmpDumpdateID adds the ndmpDumpdateId to the delete ndmp dumpdate params
func (o *DeleteNdmpDumpdateParams) SetNdmpDumpdateID(ndmpDumpdateID string) {
	o.NdmpDumpdateID = ndmpDumpdateID
}

// WithLevel adds the level to the delete ndmp dumpdate params
func (o *DeleteNdmpDumpdateParams) WithLevel(level *int64) *DeleteNdmpDumpdateParams {
	o.SetLevel(level)
	return o
}

// SetLevel adds the level to the delete ndmp dumpdate params
func (o *DeleteNdmpDumpdateParams) SetLevel(level *int64) {
	o.Level = level
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNdmpDumpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param NdmpDumpdateId
	if err := r.SetPathParam("NdmpDumpdateId", o.NdmpDumpdateID); err != nil {
		return err
	}

	if o.Level != nil {

		// query param level
		var qrLevel int64
		if o.Level != nil {
			qrLevel = *o.Level
		}
		qLevel := swag.FormatInt64(qrLevel)
		if qLevel != "" {
			if err := r.SetQueryParam("level", qLevel); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
