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

// NewGetStatisticsHistoryParams creates a new GetStatisticsHistoryParams object
// with the default values initialized.
func NewGetStatisticsHistoryParams() *GetStatisticsHistoryParams {
	var ()
	return &GetStatisticsHistoryParams{

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewGetStatisticsHistoryParamsWithTimeout creates a new GetStatisticsHistoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStatisticsHistoryParamsWithTimeout(timeout time.Duration) *GetStatisticsHistoryParams {
	var ()
	return &GetStatisticsHistoryParams{

		requestTimeout: timeout,
	}
}

// NewGetStatisticsHistoryParamsWithContext creates a new GetStatisticsHistoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStatisticsHistoryParamsWithContext(ctx context.Context) *GetStatisticsHistoryParams {
	var ()
	return &GetStatisticsHistoryParams{

		Context: ctx,
	}
}

// NewGetStatisticsHistoryParamsWithHTTPClient creates a new GetStatisticsHistoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStatisticsHistoryParamsWithHTTPClient(client *http.Client) *GetStatisticsHistoryParams {
	var ()
	return &GetStatisticsHistoryParams{
		HTTPClient: client,
	}
}

/*GetStatisticsHistoryParams contains all the parameters to send to the API endpoint
for the get statistics history operation typically these are written to a http.Request
*/
type GetStatisticsHistoryParams struct {

	/*Begin
	  Earliest time (Unix epoch seconds) of interest. Negative times are interpreted as relative (before) now.

	*/
	Begin *int64
	/*Degraded
	  If true, try to continue even if some stats are unavailable. In this case, errors will be present in the per-key returned data.

	*/
	Degraded *bool
	/*Devid
	  Node devid to query. Either an <integer> or "all". Can be used more than one time to query multiple nodes. "all" queries all up nodes. 0 means query the local node. For "cluster" scoped keys, in any devid including 0 can be used.

	*/
	Devid []string
	/*End
	  Latest time (Unix epoch seconds) of interest. Negative times are interpreted as relative (before) now. If not supplied, use now as the end time.

	*/
	End *int64
	/*ExpandClientid
	  If true, use name resolution to expand client addresses and other IDs.

	*/
	ExpandClientid *bool
	/*Interval
	  Minimum sampling interval time in seconds. If native statistics are higher resolution, data will be down-sampled.

	*/
	Interval *int64
	/*Key
	  One key name. Can be used more than one time to query multiple keys. Also works with 'keys' arguments.

	*/
	Key []string
	/*Keys
	  Multiple key names. May request matching keys or request 'all' keys. Can be comma separated list or can be used more than one time to make queries for multiple keys. May be used in conjunction with 'substr'. Also works with 'key' arguments.

	*/
	Keys []string
	/*MemoryOnly
	  Only use statistics sources that reside in memory (faster, but with less retention).

	*/
	MemoryOnly *bool
	/*Resolution
	  Synonymous with 'interval', if supplied supersedes interval.

	*/
	Resolution *int64
	/*Substr
	  Used in conjunction with the 'keys' argument, alters the behavior of keys. Makes the 'keys' arg perform a partial match. Defaults to false.

	*/
	Substr *bool
	/*Timeout
	  Time in seconds to wait for results from remote nodes.

	*/
	Timeout *int64

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithRequestTimeout adds the timeout to the get statistics history params
func (o *GetStatisticsHistoryParams) WithRequestTimeout(timeout time.Duration) *GetStatisticsHistoryParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the get statistics history params
func (o *GetStatisticsHistoryParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the get statistics history params
func (o *GetStatisticsHistoryParams) WithContext(ctx context.Context) *GetStatisticsHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get statistics history params
func (o *GetStatisticsHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get statistics history params
func (o *GetStatisticsHistoryParams) WithHTTPClient(client *http.Client) *GetStatisticsHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get statistics history params
func (o *GetStatisticsHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBegin adds the begin to the get statistics history params
func (o *GetStatisticsHistoryParams) WithBegin(begin *int64) *GetStatisticsHistoryParams {
	o.SetBegin(begin)
	return o
}

// SetBegin adds the begin to the get statistics history params
func (o *GetStatisticsHistoryParams) SetBegin(begin *int64) {
	o.Begin = begin
}

// WithDegraded adds the degraded to the get statistics history params
func (o *GetStatisticsHistoryParams) WithDegraded(degraded *bool) *GetStatisticsHistoryParams {
	o.SetDegraded(degraded)
	return o
}

// SetDegraded adds the degraded to the get statistics history params
func (o *GetStatisticsHistoryParams) SetDegraded(degraded *bool) {
	o.Degraded = degraded
}

// WithDevid adds the devid to the get statistics history params
func (o *GetStatisticsHistoryParams) WithDevid(devid []string) *GetStatisticsHistoryParams {
	o.SetDevid(devid)
	return o
}

// SetDevid adds the devid to the get statistics history params
func (o *GetStatisticsHistoryParams) SetDevid(devid []string) {
	o.Devid = devid
}

// WithEnd adds the end to the get statistics history params
func (o *GetStatisticsHistoryParams) WithEnd(end *int64) *GetStatisticsHistoryParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get statistics history params
func (o *GetStatisticsHistoryParams) SetEnd(end *int64) {
	o.End = end
}

// WithExpandClientid adds the expandClientid to the get statistics history params
func (o *GetStatisticsHistoryParams) WithExpandClientid(expandClientid *bool) *GetStatisticsHistoryParams {
	o.SetExpandClientid(expandClientid)
	return o
}

// SetExpandClientid adds the expandClientid to the get statistics history params
func (o *GetStatisticsHistoryParams) SetExpandClientid(expandClientid *bool) {
	o.ExpandClientid = expandClientid
}

// WithInterval adds the interval to the get statistics history params
func (o *GetStatisticsHistoryParams) WithInterval(interval *int64) *GetStatisticsHistoryParams {
	o.SetInterval(interval)
	return o
}

// SetInterval adds the interval to the get statistics history params
func (o *GetStatisticsHistoryParams) SetInterval(interval *int64) {
	o.Interval = interval
}

// WithKey adds the key to the get statistics history params
func (o *GetStatisticsHistoryParams) WithKey(key []string) *GetStatisticsHistoryParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the get statistics history params
func (o *GetStatisticsHistoryParams) SetKey(key []string) {
	o.Key = key
}

// WithKeys adds the keys to the get statistics history params
func (o *GetStatisticsHistoryParams) WithKeys(keys []string) *GetStatisticsHistoryParams {
	o.SetKeys(keys)
	return o
}

// SetKeys adds the keys to the get statistics history params
func (o *GetStatisticsHistoryParams) SetKeys(keys []string) {
	o.Keys = keys
}

// WithMemoryOnly adds the memoryOnly to the get statistics history params
func (o *GetStatisticsHistoryParams) WithMemoryOnly(memoryOnly *bool) *GetStatisticsHistoryParams {
	o.SetMemoryOnly(memoryOnly)
	return o
}

// SetMemoryOnly adds the memoryOnly to the get statistics history params
func (o *GetStatisticsHistoryParams) SetMemoryOnly(memoryOnly *bool) {
	o.MemoryOnly = memoryOnly
}

// WithResolution adds the resolution to the get statistics history params
func (o *GetStatisticsHistoryParams) WithResolution(resolution *int64) *GetStatisticsHistoryParams {
	o.SetResolution(resolution)
	return o
}

// SetResolution adds the resolution to the get statistics history params
func (o *GetStatisticsHistoryParams) SetResolution(resolution *int64) {
	o.Resolution = resolution
}

// WithSubstr adds the substr to the get statistics history params
func (o *GetStatisticsHistoryParams) WithSubstr(substr *bool) *GetStatisticsHistoryParams {
	o.SetSubstr(substr)
	return o
}

// SetSubstr adds the substr to the get statistics history params
func (o *GetStatisticsHistoryParams) SetSubstr(substr *bool) {
	o.Substr = substr
}

// WithTimeout adds the timeout to the get statistics history params
func (o *GetStatisticsHistoryParams) WithTimeout(timeout *int64) *GetStatisticsHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get statistics history params
func (o *GetStatisticsHistoryParams) SetTimeout(timeout *int64) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *GetStatisticsHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.requestTimeout); err != nil {
		return err
	}
	var res []error

	if o.Begin != nil {

		// query param begin
		var qrBegin int64
		if o.Begin != nil {
			qrBegin = *o.Begin
		}
		qBegin := swag.FormatInt64(qrBegin)
		if qBegin != "" {
			if err := r.SetQueryParam("begin", qBegin); err != nil {
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

	valuesDevid := o.Devid

	joinedDevid := swag.JoinByFormat(valuesDevid, "")
	// query array param devid
	if err := r.SetQueryParam("devid", joinedDevid...); err != nil {
		return err
	}

	if o.End != nil {

		// query param end
		var qrEnd int64
		if o.End != nil {
			qrEnd = *o.End
		}
		qEnd := swag.FormatInt64(qrEnd)
		if qEnd != "" {
			if err := r.SetQueryParam("end", qEnd); err != nil {
				return err
			}
		}

	}

	if o.ExpandClientid != nil {

		// query param expand_clientid
		var qrExpandClientid bool
		if o.ExpandClientid != nil {
			qrExpandClientid = *o.ExpandClientid
		}
		qExpandClientid := swag.FormatBool(qrExpandClientid)
		if qExpandClientid != "" {
			if err := r.SetQueryParam("expand_clientid", qExpandClientid); err != nil {
				return err
			}
		}

	}

	if o.Interval != nil {

		// query param interval
		var qrInterval int64
		if o.Interval != nil {
			qrInterval = *o.Interval
		}
		qInterval := swag.FormatInt64(qrInterval)
		if qInterval != "" {
			if err := r.SetQueryParam("interval", qInterval); err != nil {
				return err
			}
		}

	}

	valuesKey := o.Key

	joinedKey := swag.JoinByFormat(valuesKey, "")
	// query array param key
	if err := r.SetQueryParam("key", joinedKey...); err != nil {
		return err
	}

	valuesKeys := o.Keys

	joinedKeys := swag.JoinByFormat(valuesKeys, "")
	// query array param keys
	if err := r.SetQueryParam("keys", joinedKeys...); err != nil {
		return err
	}

	if o.MemoryOnly != nil {

		// query param memory_only
		var qrMemoryOnly bool
		if o.MemoryOnly != nil {
			qrMemoryOnly = *o.MemoryOnly
		}
		qMemoryOnly := swag.FormatBool(qrMemoryOnly)
		if qMemoryOnly != "" {
			if err := r.SetQueryParam("memory_only", qMemoryOnly); err != nil {
				return err
			}
		}

	}

	if o.Resolution != nil {

		// query param resolution
		var qrResolution int64
		if o.Resolution != nil {
			qrResolution = *o.Resolution
		}
		qResolution := swag.FormatInt64(qrResolution)
		if qResolution != "" {
			if err := r.SetQueryParam("resolution", qResolution); err != nil {
				return err
			}
		}

	}

	if o.Substr != nil {

		// query param substr
		var qrSubstr bool
		if o.Substr != nil {
			qrSubstr = *o.Substr
		}
		qSubstr := swag.FormatBool(qrSubstr)
		if qSubstr != "" {
			if err := r.SetQueryParam("substr", qSubstr); err != nil {
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
