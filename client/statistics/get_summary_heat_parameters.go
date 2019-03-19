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

// NewGetSummaryHeatParams creates a new GetSummaryHeatParams object
// with the default values initialized.
func NewGetSummaryHeatParams() *GetSummaryHeatParams {
	var ()
	return &GetSummaryHeatParams{

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewGetSummaryHeatParamsWithTimeout creates a new GetSummaryHeatParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSummaryHeatParamsWithTimeout(timeout time.Duration) *GetSummaryHeatParams {
	var ()
	return &GetSummaryHeatParams{

		requestTimeout: timeout,
	}
}

// NewGetSummaryHeatParamsWithContext creates a new GetSummaryHeatParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSummaryHeatParamsWithContext(ctx context.Context) *GetSummaryHeatParams {
	var ()
	return &GetSummaryHeatParams{

		Context: ctx,
	}
}

// NewGetSummaryHeatParamsWithHTTPClient creates a new GetSummaryHeatParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSummaryHeatParamsWithHTTPClient(client *http.Client) *GetSummaryHeatParams {
	var ()
	return &GetSummaryHeatParams{
		HTTPClient: client,
	}
}

/*GetSummaryHeatParams contains all the parameters to send to the API endpoint
for the get summary heat operation typically these are written to a http.Request
*/
type GetSummaryHeatParams struct {

	/*Classes
	  Specify class(es) for which statistics should be reported. Default is all supported classes. See [...]/platform/3/statistics/summary/filters/classes for a complete list.

	*/
	Classes *string
	/*Convertlin
	  Convert lin to hex. Defaults to true.

	*/
	Convertlin *bool
	/*Degraded
	  Continue to report if some nodes do not respond.

	*/
	Degraded *bool
	/*Events
	  Only report specified event types(s). A comma separated list of events. Defaults to all supported events. See [...]/platform/3/statistics/summary/filters/events for a complete list.

	*/
	Events *string
	/*Maxpath
	  Maximum bytes allocated for looking up a path. An ASCII character is 1 byte (It may be more for other types of encoding). The default is 1024 bytes. Zero (0) means unlimited (Subject to the system limits.)

	*/
	Maxpath *int64
	/*Nodes
	  Specify node(s) for which statistics should be reported. A comma separated set of numbers. Default is 'all'. Zero (0) should be passed in as the sole argument to indicate only the local node.

	*/
	Nodes *string
	/*Numeric
	  Whether to convert hostnames or usernames to their human readable form. False by default.

	*/
	Numeric *bool
	/*Pathdepth
	  Squash paths to this directory depth. Defaults to none, ie. the paths are not limited (Subject to the system limits.)

	*/
	Pathdepth *int64
	/*Sort
	  { drive_id | type | xfers_in | bytes_in | xfer_size_in | xfers_out | bytes_out | xfer_size_out | access_latency | access_slow | iosched_latency | iosched_queue | busy | used_bytes_percent | used_inodes } Sort data by the specified comma-separated field(s). Prepend 'asc:' or 'desc:' to a field to change the sort direction.

	*/
	Sort *string
	/*Timeout
	  Timeout remote commands after NUM seconds, where NUM is the integer passed to the argument.

	*/
	Timeout *int64
	/*Totalby
	  Aggregate per specified fields(s). Defaults to none.

	*/
	Totalby *string

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithRequestTimeout adds the timeout to the get summary heat params
func (o *GetSummaryHeatParams) WithRequestTimeout(timeout time.Duration) *GetSummaryHeatParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the get summary heat params
func (o *GetSummaryHeatParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the get summary heat params
func (o *GetSummaryHeatParams) WithContext(ctx context.Context) *GetSummaryHeatParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get summary heat params
func (o *GetSummaryHeatParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get summary heat params
func (o *GetSummaryHeatParams) WithHTTPClient(client *http.Client) *GetSummaryHeatParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get summary heat params
func (o *GetSummaryHeatParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClasses adds the classes to the get summary heat params
func (o *GetSummaryHeatParams) WithClasses(classes *string) *GetSummaryHeatParams {
	o.SetClasses(classes)
	return o
}

// SetClasses adds the classes to the get summary heat params
func (o *GetSummaryHeatParams) SetClasses(classes *string) {
	o.Classes = classes
}

// WithConvertlin adds the convertlin to the get summary heat params
func (o *GetSummaryHeatParams) WithConvertlin(convertlin *bool) *GetSummaryHeatParams {
	o.SetConvertlin(convertlin)
	return o
}

// SetConvertlin adds the convertlin to the get summary heat params
func (o *GetSummaryHeatParams) SetConvertlin(convertlin *bool) {
	o.Convertlin = convertlin
}

// WithDegraded adds the degraded to the get summary heat params
func (o *GetSummaryHeatParams) WithDegraded(degraded *bool) *GetSummaryHeatParams {
	o.SetDegraded(degraded)
	return o
}

// SetDegraded adds the degraded to the get summary heat params
func (o *GetSummaryHeatParams) SetDegraded(degraded *bool) {
	o.Degraded = degraded
}

// WithEvents adds the events to the get summary heat params
func (o *GetSummaryHeatParams) WithEvents(events *string) *GetSummaryHeatParams {
	o.SetEvents(events)
	return o
}

// SetEvents adds the events to the get summary heat params
func (o *GetSummaryHeatParams) SetEvents(events *string) {
	o.Events = events
}

// WithMaxpath adds the maxpath to the get summary heat params
func (o *GetSummaryHeatParams) WithMaxpath(maxpath *int64) *GetSummaryHeatParams {
	o.SetMaxpath(maxpath)
	return o
}

// SetMaxpath adds the maxpath to the get summary heat params
func (o *GetSummaryHeatParams) SetMaxpath(maxpath *int64) {
	o.Maxpath = maxpath
}

// WithNodes adds the nodes to the get summary heat params
func (o *GetSummaryHeatParams) WithNodes(nodes *string) *GetSummaryHeatParams {
	o.SetNodes(nodes)
	return o
}

// SetNodes adds the nodes to the get summary heat params
func (o *GetSummaryHeatParams) SetNodes(nodes *string) {
	o.Nodes = nodes
}

// WithNumeric adds the numeric to the get summary heat params
func (o *GetSummaryHeatParams) WithNumeric(numeric *bool) *GetSummaryHeatParams {
	o.SetNumeric(numeric)
	return o
}

// SetNumeric adds the numeric to the get summary heat params
func (o *GetSummaryHeatParams) SetNumeric(numeric *bool) {
	o.Numeric = numeric
}

// WithPathdepth adds the pathdepth to the get summary heat params
func (o *GetSummaryHeatParams) WithPathdepth(pathdepth *int64) *GetSummaryHeatParams {
	o.SetPathdepth(pathdepth)
	return o
}

// SetPathdepth adds the pathdepth to the get summary heat params
func (o *GetSummaryHeatParams) SetPathdepth(pathdepth *int64) {
	o.Pathdepth = pathdepth
}

// WithSort adds the sort to the get summary heat params
func (o *GetSummaryHeatParams) WithSort(sort *string) *GetSummaryHeatParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get summary heat params
func (o *GetSummaryHeatParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithTimeout adds the timeout to the get summary heat params
func (o *GetSummaryHeatParams) WithTimeout(timeout *int64) *GetSummaryHeatParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get summary heat params
func (o *GetSummaryHeatParams) SetTimeout(timeout *int64) {
	o.Timeout = timeout
}

// WithTotalby adds the totalby to the get summary heat params
func (o *GetSummaryHeatParams) WithTotalby(totalby *string) *GetSummaryHeatParams {
	o.SetTotalby(totalby)
	return o
}

// SetTotalby adds the totalby to the get summary heat params
func (o *GetSummaryHeatParams) SetTotalby(totalby *string) {
	o.Totalby = totalby
}

// WriteToRequest writes these params to a swagger request
func (o *GetSummaryHeatParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.requestTimeout); err != nil {
		return err
	}
	var res []error

	if o.Classes != nil {

		// query param classes
		var qrClasses string
		if o.Classes != nil {
			qrClasses = *o.Classes
		}
		qClasses := qrClasses
		if qClasses != "" {
			if err := r.SetQueryParam("classes", qClasses); err != nil {
				return err
			}
		}

	}

	if o.Convertlin != nil {

		// query param convertlin
		var qrConvertlin bool
		if o.Convertlin != nil {
			qrConvertlin = *o.Convertlin
		}
		qConvertlin := swag.FormatBool(qrConvertlin)
		if qConvertlin != "" {
			if err := r.SetQueryParam("convertlin", qConvertlin); err != nil {
				return err
			}
		}

	}

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

	if o.Events != nil {

		// query param events
		var qrEvents string
		if o.Events != nil {
			qrEvents = *o.Events
		}
		qEvents := qrEvents
		if qEvents != "" {
			if err := r.SetQueryParam("events", qEvents); err != nil {
				return err
			}
		}

	}

	if o.Maxpath != nil {

		// query param maxpath
		var qrMaxpath int64
		if o.Maxpath != nil {
			qrMaxpath = *o.Maxpath
		}
		qMaxpath := swag.FormatInt64(qrMaxpath)
		if qMaxpath != "" {
			if err := r.SetQueryParam("maxpath", qMaxpath); err != nil {
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

	if o.Numeric != nil {

		// query param numeric
		var qrNumeric bool
		if o.Numeric != nil {
			qrNumeric = *o.Numeric
		}
		qNumeric := swag.FormatBool(qrNumeric)
		if qNumeric != "" {
			if err := r.SetQueryParam("numeric", qNumeric); err != nil {
				return err
			}
		}

	}

	if o.Pathdepth != nil {

		// query param pathdepth
		var qrPathdepth int64
		if o.Pathdepth != nil {
			qrPathdepth = *o.Pathdepth
		}
		qPathdepth := swag.FormatInt64(qrPathdepth)
		if qPathdepth != "" {
			if err := r.SetQueryParam("pathdepth", qPathdepth); err != nil {
				return err
			}
		}

	}

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
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

	if o.Totalby != nil {

		// query param totalby
		var qrTotalby string
		if o.Totalby != nil {
			qrTotalby = *o.Totalby
		}
		qTotalby := qrTotalby
		if qTotalby != "" {
			if err := r.SetQueryParam("totalby", qTotalby); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}