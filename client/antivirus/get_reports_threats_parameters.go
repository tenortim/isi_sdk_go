// Code generated by go-swagger; DO NOT EDIT.

package antivirus

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

// NewGetReportsThreatsParams creates a new GetReportsThreatsParams object
// with the default values initialized.
func NewGetReportsThreatsParams() *GetReportsThreatsParams {
	var ()
	return &GetReportsThreatsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReportsThreatsParamsWithTimeout creates a new GetReportsThreatsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReportsThreatsParamsWithTimeout(timeout time.Duration) *GetReportsThreatsParams {
	var ()
	return &GetReportsThreatsParams{

		timeout: timeout,
	}
}

// NewGetReportsThreatsParamsWithContext creates a new GetReportsThreatsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReportsThreatsParamsWithContext(ctx context.Context) *GetReportsThreatsParams {
	var ()
	return &GetReportsThreatsParams{

		Context: ctx,
	}
}

// NewGetReportsThreatsParamsWithHTTPClient creates a new GetReportsThreatsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReportsThreatsParamsWithHTTPClient(client *http.Client) *GetReportsThreatsParams {
	var ()
	return &GetReportsThreatsParams{
		HTTPClient: client,
	}
}

/*GetReportsThreatsParams contains all the parameters to send to the API endpoint
for the get reports threats operation typically these are written to a http.Request
*/
type GetReportsThreatsParams struct {

	/*Dir
	  The direction of the sort.

	*/
	Dir *string
	/*File
	  If present, only returns threat reports for the given file.

	*/
	File *string
	/*Limit
	  Return no more than this many results at once (see resume).

	*/
	Limit *int64
	/*Remediation
	  If present, only returns threat reports with the given remediation.

	*/
	Remediation *string
	/*Resume
	  Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options).

	*/
	Resume *string
	/*ScanID
	  If present, only returns threat reports associated with the given scan report.

	*/
	ScanID *string
	/*Sort
	  The field that will be used for sorting.

	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get reports threats params
func (o *GetReportsThreatsParams) WithTimeout(timeout time.Duration) *GetReportsThreatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get reports threats params
func (o *GetReportsThreatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get reports threats params
func (o *GetReportsThreatsParams) WithContext(ctx context.Context) *GetReportsThreatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get reports threats params
func (o *GetReportsThreatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get reports threats params
func (o *GetReportsThreatsParams) WithHTTPClient(client *http.Client) *GetReportsThreatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get reports threats params
func (o *GetReportsThreatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDir adds the dir to the get reports threats params
func (o *GetReportsThreatsParams) WithDir(dir *string) *GetReportsThreatsParams {
	o.SetDir(dir)
	return o
}

// SetDir adds the dir to the get reports threats params
func (o *GetReportsThreatsParams) SetDir(dir *string) {
	o.Dir = dir
}

// WithFile adds the file to the get reports threats params
func (o *GetReportsThreatsParams) WithFile(file *string) *GetReportsThreatsParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the get reports threats params
func (o *GetReportsThreatsParams) SetFile(file *string) {
	o.File = file
}

// WithLimit adds the limit to the get reports threats params
func (o *GetReportsThreatsParams) WithLimit(limit *int64) *GetReportsThreatsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get reports threats params
func (o *GetReportsThreatsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithRemediation adds the remediation to the get reports threats params
func (o *GetReportsThreatsParams) WithRemediation(remediation *string) *GetReportsThreatsParams {
	o.SetRemediation(remediation)
	return o
}

// SetRemediation adds the remediation to the get reports threats params
func (o *GetReportsThreatsParams) SetRemediation(remediation *string) {
	o.Remediation = remediation
}

// WithResume adds the resume to the get reports threats params
func (o *GetReportsThreatsParams) WithResume(resume *string) *GetReportsThreatsParams {
	o.SetResume(resume)
	return o
}

// SetResume adds the resume to the get reports threats params
func (o *GetReportsThreatsParams) SetResume(resume *string) {
	o.Resume = resume
}

// WithScanID adds the scanID to the get reports threats params
func (o *GetReportsThreatsParams) WithScanID(scanID *string) *GetReportsThreatsParams {
	o.SetScanID(scanID)
	return o
}

// SetScanID adds the scanId to the get reports threats params
func (o *GetReportsThreatsParams) SetScanID(scanID *string) {
	o.ScanID = scanID
}

// WithSort adds the sort to the get reports threats params
func (o *GetReportsThreatsParams) WithSort(sort *string) *GetReportsThreatsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get reports threats params
func (o *GetReportsThreatsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetReportsThreatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Dir != nil {

		// query param dir
		var qrDir string
		if o.Dir != nil {
			qrDir = *o.Dir
		}
		qDir := qrDir
		if qDir != "" {
			if err := r.SetQueryParam("dir", qDir); err != nil {
				return err
			}
		}

	}

	if o.File != nil {

		// query param file
		var qrFile string
		if o.File != nil {
			qrFile = *o.File
		}
		qFile := qrFile
		if qFile != "" {
			if err := r.SetQueryParam("file", qFile); err != nil {
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

	if o.Remediation != nil {

		// query param remediation
		var qrRemediation string
		if o.Remediation != nil {
			qrRemediation = *o.Remediation
		}
		qRemediation := qrRemediation
		if qRemediation != "" {
			if err := r.SetQueryParam("remediation", qRemediation); err != nil {
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

	if o.ScanID != nil {

		// query param scan_id
		var qrScanID string
		if o.ScanID != nil {
			qrScanID = *o.ScanID
		}
		qScanID := qrScanID
		if qScanID != "" {
			if err := r.SetQueryParam("scan_id", qScanID); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
