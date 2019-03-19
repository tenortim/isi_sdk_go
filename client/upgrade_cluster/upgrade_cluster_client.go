// Code generated by go-swagger; DO NOT EDIT.

package upgrade_cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new upgrade cluster API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for upgrade cluster API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetNodesNodeFirmwareStatus The firmware status for the node.
*/
func (a *Client) GetNodesNodeFirmwareStatus(params *GetNodesNodeFirmwareStatusParams, authInfo runtime.ClientAuthInfoWriter) (*GetNodesNodeFirmwareStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodesNodeFirmwareStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNodesNodeFirmwareStatus",
		Method:             "GET",
		PathPattern:        "/platform/3/upgrade/cluster/nodes/{Lnn}/firmware/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNodesNodeFirmwareStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodesNodeFirmwareStatusOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}