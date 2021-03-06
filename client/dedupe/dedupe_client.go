// Code generated by go-swagger; DO NOT EDIT.

package dedupe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new dedupe API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for dedupe API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetDedupeDedupeSummary Return summary information about dedupe.
*/
func (a *Client) GetDedupeDedupeSummary(params *GetDedupeDedupeSummaryParams, authInfo runtime.ClientAuthInfoWriter) (*GetDedupeDedupeSummaryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDedupeDedupeSummaryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDedupeDedupeSummary",
		Method:             "GET",
		PathPattern:        "/platform/1/dedupe/dedupe-summary",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDedupeDedupeSummaryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDedupeDedupeSummaryOK), nil

}

/*
GetDedupeReport Retrieve a report for a single dedupe job.
*/
func (a *Client) GetDedupeReport(params *GetDedupeReportParams, authInfo runtime.ClientAuthInfoWriter) (*GetDedupeReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDedupeReportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDedupeReport",
		Method:             "GET",
		PathPattern:        "/platform/1/dedupe/reports/{DedupeReportId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDedupeReportReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDedupeReportOK), nil

}

/*
GetDedupeReports List dedupe reports.
*/
func (a *Client) GetDedupeReports(params *GetDedupeReportsParams, authInfo runtime.ClientAuthInfoWriter) (*GetDedupeReportsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDedupeReportsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDedupeReports",
		Method:             "GET",
		PathPattern:        "/platform/1/dedupe/reports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDedupeReportsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDedupeReportsOK), nil

}

/*
GetDedupeSettings Retrieve the dedupe settings.
*/
func (a *Client) GetDedupeSettings(params *GetDedupeSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*GetDedupeSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDedupeSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDedupeSettings",
		Method:             "GET",
		PathPattern:        "/platform/1/dedupe/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDedupeSettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDedupeSettingsOK), nil

}

/*
UpdateDedupeSettings Modify the dedupe settings. All input fields are optional, but one or more must be supplied.
*/
func (a *Client) UpdateDedupeSettings(params *UpdateDedupeSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateDedupeSettingsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDedupeSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateDedupeSettings",
		Method:             "PUT",
		PathPattern:        "/platform/1/dedupe/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDedupeSettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateDedupeSettingsNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
