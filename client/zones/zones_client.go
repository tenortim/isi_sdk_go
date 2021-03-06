// Code generated by go-swagger; DO NOT EDIT.

package zones

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new zones API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for zones API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateZone Create a new access zone.
*/
func (a *Client) CreateZone(params *CreateZoneParams, authInfo runtime.ClientAuthInfoWriter) (*CreateZoneOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateZoneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createZone",
		Method:             "POST",
		PathPattern:        "/platform/3/zones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateZoneReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateZoneOK), nil

}

/*
DeleteZone Delete the access zone.
*/
func (a *Client) DeleteZone(params *DeleteZoneParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteZoneNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteZoneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteZone",
		Method:             "DELETE",
		PathPattern:        "/platform/3/zones/{ZoneId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteZoneReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteZoneNoContent), nil

}

/*
GetZone Retrieve the access zone information.
*/
func (a *Client) GetZone(params *GetZoneParams, authInfo runtime.ClientAuthInfoWriter) (*GetZoneOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetZoneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getZone",
		Method:             "GET",
		PathPattern:        "/platform/3/zones/{ZoneId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetZoneReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetZoneOK), nil

}

/*
ListZones List all access zones.
*/
func (a *Client) ListZones(params *ListZonesParams, authInfo runtime.ClientAuthInfoWriter) (*ListZonesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListZonesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listZones",
		Method:             "GET",
		PathPattern:        "/platform/3/zones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListZonesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListZonesOK), nil

}

/*
UpdateZone Modify the access zone. All input fields are optional, but one or more must be supplied.
*/
func (a *Client) UpdateZone(params *UpdateZoneParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateZoneNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateZoneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateZone",
		Method:             "PUT",
		PathPattern:        "/platform/3/zones/{ZoneId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateZoneReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateZoneNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
