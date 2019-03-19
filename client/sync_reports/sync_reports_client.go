// Code generated by go-swagger; DO NOT EDIT.

package sync_reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new sync reports API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sync reports API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetReportSubreport View a single SyncIQ subreport.
*/
func (a *Client) GetReportSubreport(params *GetReportSubreportParams, authInfo runtime.ClientAuthInfoWriter) (*GetReportSubreportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReportSubreportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getReportSubreport",
		Method:             "GET",
		PathPattern:        "/platform/1/sync/reports/{Rid}/subreports/{ReportSubreportId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetReportSubreportReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetReportSubreportOK), nil

}

/*
GetReportSubreports Get a list of SyncIQ subreports for a report.
*/
func (a *Client) GetReportSubreports(params *GetReportSubreportsParams, authInfo runtime.ClientAuthInfoWriter) (*GetReportSubreportsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReportSubreportsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getReportSubreports",
		Method:             "GET",
		PathPattern:        "/platform/1/sync/reports/{Rid}/subreports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetReportSubreportsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetReportSubreportsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}