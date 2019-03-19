// Code generated by go-swagger; DO NOT EDIT.

package statistics

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

// NewGetSummaryProtocolStatsParams creates a new GetSummaryProtocolStatsParams object
// with the default values initialized.
func NewGetSummaryProtocolStatsParams() *GetSummaryProtocolStatsParams {
	var ()
	return &GetSummaryProtocolStatsParams{

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewGetSummaryProtocolStatsParamsWithTimeout creates a new GetSummaryProtocolStatsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSummaryProtocolStatsParamsWithTimeout(timeout time.Duration) *GetSummaryProtocolStatsParams {
	var ()
	return &GetSummaryProtocolStatsParams{

		requestTimeout: timeout,
	}
}

// NewGetSummaryProtocolStatsParamsWithContext creates a new GetSummaryProtocolStatsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSummaryProtocolStatsParamsWithContext(ctx context.Context) *GetSummaryProtocolStatsParams {
	var ()
	return &GetSummaryProtocolStatsParams{

		Context: ctx,
	}
}

// NewGetSummaryProtocolStatsParamsWithHTTPClient creates a new GetSummaryProtocolStatsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSummaryProtocolStatsParamsWithHTTPClient(client *http.Client) *GetSummaryProtocolStatsParams {
	var ()
	return &GetSummaryProtocolStatsParams{
		HTTPClient: client,
	}
}

/*GetSummaryProtocolStatsParams contains all the parameters to send to the API endpoint
for the get summary protocol stats operation typically these are written to a http.Request
*/
type GetSummaryProtocolStatsParams struct {

	/*Degraded
	  Continue to report if some nodes do not respond.

	*/
	Degraded *bool
	/*Nodes
	  Specify node(s) for which statistics should be reported. A comma separated set of numbers. Default is 'all'. Zero (0) should be passed in as the sole argument to indicate only the local node.

	*/
	Nodes *string
	/*Protocol
	  Specify protocol(s) for which statistics should be reported. Default is all external protocols.

	*/
	Protocol *string
	/*Timeout
	  Timeout remote commands after NUM seconds, where NUM is the integer passed to the argument.

	*/
	Timeout *int64

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithRequestTimeout adds the timeout to the get summary protocol stats params
func (o *GetSummaryProtocolStatsParams) WithRequestTimeout(timeout time.Duration) *GetSummaryProtocolStatsParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the get summary protocol stats params
func (o *GetSummaryProtocolStatsParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the get summary protocol stats params
func (o *GetSummaryProtocolStatsParams) WithContext(ctx context.Context) *GetSummaryProtocolStatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get summary protocol stats params
func (o *GetSummaryProtocolStatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get summary protocol stats params
func (o *GetSummaryProtocolStatsParams) WithHTTPClient(client *http.Client) *GetSummaryProtocolStatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get summary protocol stats params
func (o *GetSummaryProtocolStatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDegraded adds the degraded to the get summary protocol stats params
func (o *GetSummaryProtocolStatsParams) WithDegraded(degraded *bool) *GetSummaryProtocolStatsParams {
	o.SetDegraded(degraded)
	return o
}

// SetDegraded adds the degraded to the get summary protocol stats params
func (o *GetSummaryProtocolStatsParams) SetDegraded(degraded *bool) {
	o.Degraded = degraded
}

// WithNodes adds the nodes to the get summary protocol stats params
func (o *GetSummaryProtocolStatsParams) WithNodes(nodes *string) *GetSummaryProtocolStatsParams {
	o.SetNodes(nodes)
	return o
}

// SetNodes adds the nodes to the get summary protocol stats params
func (o *GetSummaryProtocolStatsParams) SetNodes(nodes *string) {
	o.Nodes = nodes
}

// WithProtocol adds the protocol to the get summary protocol stats params
func (o *GetSummaryProtocolStatsParams) WithProtocol(protocol *string) *GetSummaryProtocolStatsParams {
	o.SetProtocol(protocol)
	return o
}

// SetProtocol adds the protocol to the get summary protocol stats params
func (o *GetSummaryProtocolStatsParams) SetProtocol(protocol *string) {
	o.Protocol = protocol
}

// WithTimeout adds the timeout to the get summary protocol stats params
func (o *GetSummaryProtocolStatsParams) WithTimeout(timeout *int64) *GetSummaryProtocolStatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get summary protocol stats params
func (o *GetSummaryProtocolStatsParams) SetTimeout(timeout *int64) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *GetSummaryProtocolStatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.requestTimeout); err != nil {
		return err
	}
	var res []error

	if o.Degraded != nil {

		// query param degraded
		var qrDegraded bool
		if o.Degraded != nil {
			qrDegraded = *o.Degraded
		}
		qDegraded := swag.FormatBool(qrDegraded)
		if qDegraded != "" {
			if err := r.SetQueryParam("degraded", qDegraded); err != nil {
				return err
			}
		}

	}

	if o.Nodes != nil {

		// query param nodes
		var qrNodes string
		if o.Nodes != nil {
			qrNodes = *o.Nodes
		}
		qNodes := qrNodes
		if qNodes != "" {
			if err := r.SetQueryParam("nodes", qNodes); err != nil {
				return err
			}
		}

	}

	if o.Protocol != nil {

		// query param protocol
		var qrProtocol string
		if o.Protocol != nil {
			qrProtocol = *o.Protocol
		}
		qProtocol := qrProtocol
		if qProtocol != "" {
			if err := r.SetQueryParam("protocol", qProtocol); err != nil {
				return err
			}
		}

	}

	if o.Timeout != nil {

		// query param timeout
		var qrTimeout int64
		if o.Timeout != nil {
			qrTimeout = *o.Timeout
		}
		qTimeout := swag.FormatInt64(qrTimeout)
		if qTimeout != "" {
			if err := r.SetQueryParam("timeout", qTimeout); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
