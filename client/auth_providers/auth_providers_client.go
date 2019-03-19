// Code generated by go-swagger; DO NOT EDIT.

package auth_providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new auth providers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for auth providers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetAdsProviderControllers List all ADS controllers.
*/
func (a *Client) GetAdsProviderControllers(params *GetAdsProviderControllersParams, authInfo runtime.ClientAuthInfoWriter) (*GetAdsProviderControllersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAdsProviderControllersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAdsProviderControllers",
		Method:             "GET",
		PathPattern:        "/platform/1/auth/providers/ads/{Id}/controllers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAdsProviderControllersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAdsProviderControllersOK), nil

}

/*
GetAdsProviderDomain Retrieve the ADS domain information.
*/
func (a *Client) GetAdsProviderDomain(params *GetAdsProviderDomainParams, authInfo runtime.ClientAuthInfoWriter) (*GetAdsProviderDomainOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAdsProviderDomainParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAdsProviderDomain",
		Method:             "GET",
		PathPattern:        "/platform/3/auth/providers/ads/{Id}/domains/{AdsProviderDomainId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAdsProviderDomainReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAdsProviderDomainOK), nil

}

/*
GetAdsProviderDomains List all ADS domains.
*/
func (a *Client) GetAdsProviderDomains(params *GetAdsProviderDomainsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAdsProviderDomainsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAdsProviderDomainsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAdsProviderDomains",
		Method:             "GET",
		PathPattern:        "/platform/3/auth/providers/ads/{Id}/domains",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAdsProviderDomainsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAdsProviderDomainsOK), nil

}

/*
GetAdsProviderSearch Retrieve search results.
*/
func (a *Client) GetAdsProviderSearch(params *GetAdsProviderSearchParams, authInfo runtime.ClientAuthInfoWriter) (*GetAdsProviderSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAdsProviderSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAdsProviderSearch",
		Method:             "GET",
		PathPattern:        "/platform/1/auth/providers/ads/{Id}/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAdsProviderSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAdsProviderSearchOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
