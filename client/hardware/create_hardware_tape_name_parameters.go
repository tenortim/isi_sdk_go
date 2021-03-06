// Code generated by go-swagger; DO NOT EDIT.

package hardware

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

// NewCreateHardwareTapeNameParams creates a new CreateHardwareTapeNameParams object
// with the default values initialized.
func NewCreateHardwareTapeNameParams() *CreateHardwareTapeNameParams {
	var ()
	return &CreateHardwareTapeNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateHardwareTapeNameParamsWithTimeout creates a new CreateHardwareTapeNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateHardwareTapeNameParamsWithTimeout(timeout time.Duration) *CreateHardwareTapeNameParams {
	var ()
	return &CreateHardwareTapeNameParams{

		timeout: timeout,
	}
}

// NewCreateHardwareTapeNameParamsWithContext creates a new CreateHardwareTapeNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateHardwareTapeNameParamsWithContext(ctx context.Context) *CreateHardwareTapeNameParams {
	var ()
	return &CreateHardwareTapeNameParams{

		Context: ctx,
	}
}

// NewCreateHardwareTapeNameParamsWithHTTPClient creates a new CreateHardwareTapeNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateHardwareTapeNameParamsWithHTTPClient(client *http.Client) *CreateHardwareTapeNameParams {
	var ()
	return &CreateHardwareTapeNameParams{
		HTTPClient: client,
	}
}

/*CreateHardwareTapeNameParams contains all the parameters to send to the API endpoint
for the create hardware tape name operation typically these are written to a http.Request
*/
type CreateHardwareTapeNameParams struct {

	/*HardwareTapeName*/
	HardwareTapeName models.Empty
	/*HardwareTapeName
	  Tape/Changer devices rescan

	*/
	PathHardwareTapeName string
	/*Lnn
	  Logical node number.

	*/
	Lnn *string
	/*Port
	  Scan only specified port.

	*/
	Port *int64
	/*Reconcile
	  Remove entries for devices or paths that have become inaccessible.

	*/
	Reconcile *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create hardware tape name params
func (o *CreateHardwareTapeNameParams) WithTimeout(timeout time.Duration) *CreateHardwareTapeNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create hardware tape name params
func (o *CreateHardwareTapeNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create hardware tape name params
func (o *CreateHardwareTapeNameParams) WithContext(ctx context.Context) *CreateHardwareTapeNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create hardware tape name params
func (o *CreateHardwareTapeNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create hardware tape name params
func (o *CreateHardwareTapeNameParams) WithHTTPClient(client *http.Client) *CreateHardwareTapeNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create hardware tape name params
func (o *CreateHardwareTapeNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHardwareTapeName adds the hardwareTapeName to the create hardware tape name params
func (o *CreateHardwareTapeNameParams) WithHardwareTapeName(hardwareTapeName models.Empty) *CreateHardwareTapeNameParams {
	o.SetHardwareTapeName(hardwareTapeName)
	return o
}

// SetHardwareTapeName adds the hardwareTapeName to the create hardware tape name params
func (o *CreateHardwareTapeNameParams) SetHardwareTapeName(hardwareTapeName models.Empty) {
	o.HardwareTapeName = hardwareTapeName
}

// WithPathHardwareTapeName adds the hardwareTapeName to the create hardware tape name params
func (o *CreateHardwareTapeNameParams) WithPathHardwareTapeName(hardwareTapeName string) *CreateHardwareTapeNameParams {
	o.SetPathHardwareTapeName(hardwareTapeName)
	return o
}

// SetPathHardwareTapeName adds the hardwareTapeName to the create hardware tape name params
func (o *CreateHardwareTapeNameParams) SetPathHardwareTapeName(hardwareTapeName string) {
	o.PathHardwareTapeName = hardwareTapeName
}

// WithLnn adds the lnn to the create hardware tape name params
func (o *CreateHardwareTapeNameParams) WithLnn(lnn *string) *CreateHardwareTapeNameParams {
	o.SetLnn(lnn)
	return o
}

// SetLnn adds the lnn to the create hardware tape name params
func (o *CreateHardwareTapeNameParams) SetLnn(lnn *string) {
	o.Lnn = lnn
}

// WithPort adds the port to the create hardware tape name params
func (o *CreateHardwareTapeNameParams) WithPort(port *int64) *CreateHardwareTapeNameParams {
	o.SetPort(port)
	return o
}

// SetPort adds the port to the create hardware tape name params
func (o *CreateHardwareTapeNameParams) SetPort(port *int64) {
	o.Port = port
}

// WithReconcile adds the reconcile to the create hardware tape name params
func (o *CreateHardwareTapeNameParams) WithReconcile(reconcile *bool) *CreateHardwareTapeNameParams {
	o.SetReconcile(reconcile)
	return o
}

// SetReconcile adds the reconcile to the create hardware tape name params
func (o *CreateHardwareTapeNameParams) SetReconcile(reconcile *bool) {
	o.Reconcile = reconcile
}

// WriteToRequest writes these params to a swagger request
func (o *CreateHardwareTapeNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.HardwareTapeName != nil {
		if err := r.SetBodyParam(o.HardwareTapeName); err != nil {
			return err
		}
	}

	// path param HardwareTapeName
	if err := r.SetPathParam("HardwareTapeName", o.PathHardwareTapeName); err != nil {
		return err
	}

	if o.Lnn != nil {

		// query param lnn
		var qrLnn string
		if o.Lnn != nil {
			qrLnn = *o.Lnn
		}
		qLnn := qrLnn
		if qLnn != "" {
			if err := r.SetQueryParam("lnn", qLnn); err != nil {
				return err
			}
		}

	}

	if o.Port != nil {

		// query param port
		var qrPort int64
		if o.Port != nil {
			qrPort = *o.Port
		}
		qPort := swag.FormatInt64(qrPort)
		if qPort != "" {
			if err := r.SetQueryParam("port", qPort); err != nil {
				return err
			}
		}

	}

	if o.Reconcile != nil {

		// query param reconcile
		var qrReconcile bool
		if o.Reconcile != nil {
			qrReconcile = *o.Reconcile
		}
		qReconcile := swag.FormatBool(qrReconcile)
		if qReconcile != "" {
			if err := r.SetQueryParam("reconcile", qReconcile); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
