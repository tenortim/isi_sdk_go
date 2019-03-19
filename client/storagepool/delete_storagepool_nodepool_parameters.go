// Code generated by go-swagger; DO NOT EDIT.

package storagepool

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

// NewDeleteStoragepoolNodepoolParams creates a new DeleteStoragepoolNodepoolParams object
// with the default values initialized.
func NewDeleteStoragepoolNodepoolParams() *DeleteStoragepoolNodepoolParams {
	var ()
	return &DeleteStoragepoolNodepoolParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteStoragepoolNodepoolParamsWithTimeout creates a new DeleteStoragepoolNodepoolParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteStoragepoolNodepoolParamsWithTimeout(timeout time.Duration) *DeleteStoragepoolNodepoolParams {
	var ()
	return &DeleteStoragepoolNodepoolParams{

		timeout: timeout,
	}
}

// NewDeleteStoragepoolNodepoolParamsWithContext creates a new DeleteStoragepoolNodepoolParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteStoragepoolNodepoolParamsWithContext(ctx context.Context) *DeleteStoragepoolNodepoolParams {
	var ()
	return &DeleteStoragepoolNodepoolParams{

		Context: ctx,
	}
}

// NewDeleteStoragepoolNodepoolParamsWithHTTPClient creates a new DeleteStoragepoolNodepoolParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteStoragepoolNodepoolParamsWithHTTPClient(client *http.Client) *DeleteStoragepoolNodepoolParams {
	var ()
	return &DeleteStoragepoolNodepoolParams{
		HTTPClient: client,
	}
}

/*DeleteStoragepoolNodepoolParams contains all the parameters to send to the API endpoint
for the delete storagepool nodepool operation typically these are written to a http.Request
*/
type DeleteStoragepoolNodepoolParams struct {

	/*StoragepoolNodepoolID
	  Delete node pool.

	*/
	StoragepoolNodepoolID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete storagepool nodepool params
func (o *DeleteStoragepoolNodepoolParams) WithTimeout(timeout time.Duration) *DeleteStoragepoolNodepoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete storagepool nodepool params
func (o *DeleteStoragepoolNodepoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete storagepool nodepool params
func (o *DeleteStoragepoolNodepoolParams) WithContext(ctx context.Context) *DeleteStoragepoolNodepoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete storagepool nodepool params
func (o *DeleteStoragepoolNodepoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete storagepool nodepool params
func (o *DeleteStoragepoolNodepoolParams) WithHTTPClient(client *http.Client) *DeleteStoragepoolNodepoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete storagepool nodepool params
func (o *DeleteStoragepoolNodepoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStoragepoolNodepoolID adds the storagepoolNodepoolID to the delete storagepool nodepool params
func (o *DeleteStoragepoolNodepoolParams) WithStoragepoolNodepoolID(storagepoolNodepoolID string) *DeleteStoragepoolNodepoolParams {
	o.SetStoragepoolNodepoolID(storagepoolNodepoolID)
	return o
}

// SetStoragepoolNodepoolID adds the storagepoolNodepoolId to the delete storagepool nodepool params
func (o *DeleteStoragepoolNodepoolParams) SetStoragepoolNodepoolID(storagepoolNodepoolID string) {
	o.StoragepoolNodepoolID = storagepoolNodepoolID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteStoragepoolNodepoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param StoragepoolNodepoolId
	if err := r.SetPathParam("StoragepoolNodepoolId", o.StoragepoolNodepoolID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}