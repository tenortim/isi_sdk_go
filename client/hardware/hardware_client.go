// Code generated by go-swagger; DO NOT EDIT.

package hardware

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new hardware API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for hardware API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateHardwareTapeName Tape/Changer devices rescan
*/
func (a *Client) CreateHardwareTapeName(params *CreateHardwareTapeNameParams, authInfo runtime.ClientAuthInfoWriter) (*CreateHardwareTapeNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateHardwareTapeNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createHardwareTapeName",
		Method:             "POST",
		PathPattern:        "/platform/3/hardware/tape/{HardwareTapeName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateHardwareTapeNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateHardwareTapeNameOK), nil

}

/*
DeleteHardwareTapeName Tape/Changer devices remove
*/
func (a *Client) DeleteHardwareTapeName(params *DeleteHardwareTapeNameParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteHardwareTapeNameNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteHardwareTapeNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteHardwareTapeName",
		Method:             "DELETE",
		PathPattern:        "/platform/3/hardware/tape/{HardwareTapeName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteHardwareTapeNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteHardwareTapeNameNoContent), nil

}

/*
GetHardwareFcport Get one fibre-channel port
*/
func (a *Client) GetHardwareFcport(params *GetHardwareFcportParams, authInfo runtime.ClientAuthInfoWriter) (*GetHardwareFcportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHardwareFcportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getHardwareFcport",
		Method:             "GET",
		PathPattern:        "/platform/3/hardware/fcports/{HardwareFcportId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHardwareFcportReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetHardwareFcportOK), nil

}

/*
GetHardwareFcports Get list of fibre-channel ports
*/
func (a *Client) GetHardwareFcports(params *GetHardwareFcportsParams, authInfo runtime.ClientAuthInfoWriter) (*GetHardwareFcportsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHardwareFcportsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getHardwareFcports",
		Method:             "GET",
		PathPattern:        "/platform/3/hardware/fcports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHardwareFcportsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetHardwareFcportsOK), nil

}

/*
GetHardwareTapes Get list Tape and Changer devices
*/
func (a *Client) GetHardwareTapes(params *GetHardwareTapesParams, authInfo runtime.ClientAuthInfoWriter) (*GetHardwareTapesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHardwareTapesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getHardwareTapes",
		Method:             "GET",
		PathPattern:        "/platform/3/hardware/tapes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHardwareTapesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetHardwareTapesOK), nil

}

/*
UpdateHardwareFcport Change wwnn, wwpn, state, topology, or rate of a FC port
*/
func (a *Client) UpdateHardwareFcport(params *UpdateHardwareFcportParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateHardwareFcportNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateHardwareFcportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateHardwareFcport",
		Method:             "PUT",
		PathPattern:        "/platform/3/hardware/fcports/{HardwareFcportId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateHardwareFcportReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateHardwareFcportNoContent), nil

}

/*
UpdateHardwareTapeName Tape/Changer device modify
*/
func (a *Client) UpdateHardwareTapeName(params *UpdateHardwareTapeNameParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateHardwareTapeNameNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateHardwareTapeNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateHardwareTapeName",
		Method:             "PUT",
		PathPattern:        "/platform/3/hardware/tape/{HardwareTapeName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateHardwareTapeNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateHardwareTapeNameNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
