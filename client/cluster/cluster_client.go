// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new cluster API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cluster API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateClusterAddNodeItem Serial number and arguments of node to add.
*/
func (a *Client) CreateClusterAddNodeItem(params *CreateClusterAddNodeItemParams, authInfo runtime.ClientAuthInfoWriter) (*CreateClusterAddNodeItemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateClusterAddNodeItemParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createClusterAddNodeItem",
		Method:             "POST",
		PathPattern:        "/platform/3/cluster/add-node",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateClusterAddNodeItemReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateClusterAddNodeItemOK), nil

}

/*
GetClusterConfig Retrieve the cluster information.
*/
func (a *Client) GetClusterConfig(params *GetClusterConfigParams, authInfo runtime.ClientAuthInfoWriter) (*GetClusterConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getClusterConfig",
		Method:             "GET",
		PathPattern:        "/platform/3/cluster/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterConfigOK), nil

}

/*
GetClusterEmail Get the cluster email notification settings.
*/
func (a *Client) GetClusterEmail(params *GetClusterEmailParams, authInfo runtime.ClientAuthInfoWriter) (*GetClusterEmailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterEmailParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getClusterEmail",
		Method:             "GET",
		PathPattern:        "/platform/1/cluster/email",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterEmailReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterEmailOK), nil

}

/*
GetClusterExternalIps Retrieve the cluster IP addresses including IPV4 and IPV6.
*/
func (a *Client) GetClusterExternalIps(params *GetClusterExternalIpsParams, authInfo runtime.ClientAuthInfoWriter) (*GetClusterExternalIpsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterExternalIpsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getClusterExternalIps",
		Method:             "GET",
		PathPattern:        "/platform/2/cluster/external-ips",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterExternalIpsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterExternalIpsOK), nil

}

/*
GetClusterIdentity Retrieve the login information.
*/
func (a *Client) GetClusterIdentity(params *GetClusterIdentityParams, authInfo runtime.ClientAuthInfoWriter) (*GetClusterIdentityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterIdentityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getClusterIdentity",
		Method:             "GET",
		PathPattern:        "/platform/3/cluster/identity",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterIdentityReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterIdentityOK), nil

}

/*
GetClusterNodesAvailable List all nodes that are available to add to this cluster.
*/
func (a *Client) GetClusterNodesAvailable(params *GetClusterNodesAvailableParams, authInfo runtime.ClientAuthInfoWriter) (*GetClusterNodesAvailableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterNodesAvailableParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getClusterNodesAvailable",
		Method:             "GET",
		PathPattern:        "/platform/3/cluster/nodes-available",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterNodesAvailableReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterNodesAvailableOK), nil

}

/*
GetClusterOwner Get the cluster contact info settings
*/
func (a *Client) GetClusterOwner(params *GetClusterOwnerParams, authInfo runtime.ClientAuthInfoWriter) (*GetClusterOwnerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterOwnerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getClusterOwner",
		Method:             "GET",
		PathPattern:        "/platform/1/cluster/owner",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterOwnerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterOwnerOK), nil

}

/*
GetClusterStatfs Retrieve the filesystem statistics.
*/
func (a *Client) GetClusterStatfs(params *GetClusterStatfsParams, authInfo runtime.ClientAuthInfoWriter) (*GetClusterStatfsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterStatfsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getClusterStatfs",
		Method:             "GET",
		PathPattern:        "/platform/1/cluster/statfs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterStatfsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterStatfsOK), nil

}

/*
GetClusterTimezone Get the cluster timezone.
*/
func (a *Client) GetClusterTimezone(params *GetClusterTimezoneParams, authInfo runtime.ClientAuthInfoWriter) (*GetClusterTimezoneOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterTimezoneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getClusterTimezone",
		Method:             "GET",
		PathPattern:        "/platform/3/cluster/timezone",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterTimezoneReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterTimezoneOK), nil

}

/*
GetClusterVersion Retrieve the OneFS version as reported by each node.
*/
func (a *Client) GetClusterVersion(params *GetClusterVersionParams, authInfo runtime.ClientAuthInfoWriter) (*GetClusterVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getClusterVersion",
		Method:             "GET",
		PathPattern:        "/platform/3/cluster/version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterVersionOK), nil

}

/*
GetTimezoneRegion List timezone regions.
*/
func (a *Client) GetTimezoneRegion(params *GetTimezoneRegionParams, authInfo runtime.ClientAuthInfoWriter) (*GetTimezoneRegionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTimezoneRegionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTimezoneRegion",
		Method:             "GET",
		PathPattern:        "/platform/3/cluster/timezone/regions/{TimezoneRegionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTimezoneRegionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTimezoneRegionOK), nil

}

/*
GetTimezoneSettings Retrieve the cluster timezone.
*/
func (a *Client) GetTimezoneSettings(params *GetTimezoneSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*GetTimezoneSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTimezoneSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTimezoneSettings",
		Method:             "GET",
		PathPattern:        "/platform/3/cluster/timezone/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTimezoneSettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTimezoneSettingsOK), nil

}

/*
UpdateClusterEmail Modify the cluster email notification settings.  All input fields are optional, but one or more must be supplied.
*/
func (a *Client) UpdateClusterEmail(params *UpdateClusterEmailParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateClusterEmailNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateClusterEmailParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateClusterEmail",
		Method:             "PUT",
		PathPattern:        "/platform/1/cluster/email",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateClusterEmailReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateClusterEmailNoContent), nil

}

/*
UpdateClusterIdentity Modify the login information.
*/
func (a *Client) UpdateClusterIdentity(params *UpdateClusterIdentityParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateClusterIdentityNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateClusterIdentityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateClusterIdentity",
		Method:             "PUT",
		PathPattern:        "/platform/3/cluster/identity",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateClusterIdentityReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateClusterIdentityNoContent), nil

}

/*
UpdateClusterNode Modify one or more node settings.
*/
func (a *Client) UpdateClusterNode(params *UpdateClusterNodeParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateClusterNodeNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateClusterNodeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateClusterNode",
		Method:             "PUT",
		PathPattern:        "/platform/3/cluster/nodes/{ClusterNodeId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateClusterNodeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateClusterNodeNoContent), nil

}

/*
UpdateClusterOwner Modify the cluster contact info settings.  All input fields are optional, but one or more must be supplied.
*/
func (a *Client) UpdateClusterOwner(params *UpdateClusterOwnerParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateClusterOwnerNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateClusterOwnerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateClusterOwner",
		Method:             "PUT",
		PathPattern:        "/platform/1/cluster/owner",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateClusterOwnerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateClusterOwnerNoContent), nil

}

/*
UpdateClusterTime Set cluster time.  Time will mostly be synchronized across nodes, but there may be slight drift.
*/
func (a *Client) UpdateClusterTime(params *UpdateClusterTimeParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateClusterTimeNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateClusterTimeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateClusterTime",
		Method:             "PUT",
		PathPattern:        "/platform/3/cluster/time",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateClusterTimeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateClusterTimeNoContent), nil

}

/*
UpdateClusterTimezone Set a new timezone for the cluster.
*/
func (a *Client) UpdateClusterTimezone(params *UpdateClusterTimezoneParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateClusterTimezoneNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateClusterTimezoneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateClusterTimezone",
		Method:             "PUT",
		PathPattern:        "/platform/3/cluster/timezone",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateClusterTimezoneReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateClusterTimezoneNoContent), nil

}

/*
UpdateTimezoneSettings Modify the cluster timezone.
*/
func (a *Client) UpdateTimezoneSettings(params *UpdateTimezoneSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateTimezoneSettingsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateTimezoneSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateTimezoneSettings",
		Method:             "PUT",
		PathPattern:        "/platform/3/cluster/timezone/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateTimezoneSettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateTimezoneSettingsNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
