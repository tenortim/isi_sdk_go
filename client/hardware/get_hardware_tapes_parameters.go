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
)

// NewGetHardwareTapesParams creates a new GetHardwareTapesParams object
// with the default values initialized.
func NewGetHardwareTapesParams() *GetHardwareTapesParams {
	var ()
	return &GetHardwareTapesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHardwareTapesParamsWithTimeout creates a new GetHardwareTapesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHardwareTapesParamsWithTimeout(timeout time.Duration) *GetHardwareTapesParams {
	var ()
	return &GetHardwareTapesParams{

		timeout: timeout,
	}
}

// NewGetHardwareTapesParamsWithContext creates a new GetHardwareTapesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHardwareTapesParamsWithContext(ctx context.Context) *GetHardwareTapesParams {
	var ()
	return &GetHardwareTapesParams{

		Context: ctx,
	}
}

// NewGetHardwareTapesParamsWithHTTPClient creates a new GetHardwareTapesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHardwareTapesParamsWithHTTPClient(client *http.Client) *GetHardwareTapesParams {
	var ()
	return &GetHardwareTapesParams{
		HTTPClient: client,
	}
}

/*GetHardwareTapesParams contains all the parameters to send to the API endpoint
for the get hardware tapes operation typically these are written to a http.Request
*/
type GetHardwareTapesParams struct {

	/*Activepath
	  List devices with only active paths.

	*/
	Activepath *string
	/*Devname
	  List only the named device.

	*/
	Devname *string
	/*Limit
	  Return no more than this many results at once (see resume).

	*/
	Limit *int64
	/*Node
	  List only devices on the specified node.

	*/
	Node *string
	/*Resume
	  Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options).

	*/
	Resume *string
	/*Type
	  Restrict to list on tape or mc device.

	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get hardware tapes params
func (o *GetHardwareTapesParams) WithTimeout(timeout time.Duration) *GetHardwareTapesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get hardware tapes params
func (o *GetHardwareTapesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get hardware tapes params
func (o *GetHardwareTapesParams) WithContext(ctx context.Context) *GetHardwareTapesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get hardware tapes params
func (o *GetHardwareTapesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get hardware tapes params
func (o *GetHardwareTapesParams) WithHTTPClient(client *http.Client) *GetHardwareTapesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get hardware tapes params
func (o *GetHardwareTapesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActivepath adds the activepath to the get hardware tapes params
func (o *GetHardwareTapesParams) WithActivepath(activepath *string) *GetHardwareTapesParams {
	o.SetActivepath(activepath)
	return o
}

// SetActivepath adds the activepath to the get hardware tapes params
func (o *GetHardwareTapesParams) SetActivepath(activepath *string) {
	o.Activepath = activepath
}

// WithDevname adds the devname to the get hardware tapes params
func (o *GetHardwareTapesParams) WithDevname(devname *string) *GetHardwareTapesParams {
	o.SetDevname(devname)
	return o
}

// SetDevname adds the devname to the get hardware tapes params
func (o *GetHardwareTapesParams) SetDevname(devname *string) {
	o.Devname = devname
}

// WithLimit adds the limit to the get hardware tapes params
func (o *GetHardwareTapesParams) WithLimit(limit *int64) *GetHardwareTapesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get hardware tapes params
func (o *GetHardwareTapesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNode adds the node to the get hardware tapes params
func (o *GetHardwareTapesParams) WithNode(node *string) *GetHardwareTapesParams {
	o.SetNode(node)
	return o
}

// SetNode adds the node to the get hardware tapes params
func (o *GetHardwareTapesParams) SetNode(node *string) {
	o.Node = node
}

// WithResume adds the resume to the get hardware tapes params
func (o *GetHardwareTapesParams) WithResume(resume *string) *GetHardwareTapesParams {
	o.SetResume(resume)
	return o
}

// SetResume adds the resume to the get hardware tapes params
func (o *GetHardwareTapesParams) SetResume(resume *string) {
	o.Resume = resume
}

// WithType adds the typeVar to the get hardware tapes params
func (o *GetHardwareTapesParams) WithType(typeVar *string) *GetHardwareTapesParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get hardware tapes params
func (o *GetHardwareTapesParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetHardwareTapesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Activepath != nil {

		// query param activepath
		var qrActivepath string
		if o.Activepath != nil {
			qrActivepath = *o.Activepath
		}
		qActivepath := qrActivepath
		if qActivepath != "" {
			if err := r.SetQueryParam("activepath", qActivepath); err != nil {
				return err
			}
		}

	}

	if o.Devname != nil {

		// query param devname
		var qrDevname string
		if o.Devname != nil {
			qrDevname = *o.Devname
		}
		qDevname := qrDevname
		if qDevname != "" {
			if err := r.SetQueryParam("devname", qDevname); err != nil {
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

	if o.Node != nil {

		// query param node
		var qrNode string
		if o.Node != nil {
			qrNode = *o.Node
		}
		qNode := qrNode
		if qNode != "" {
			if err := r.SetQueryParam("node", qNode); err != nil {
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
