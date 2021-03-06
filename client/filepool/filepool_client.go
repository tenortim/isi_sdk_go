// Code generated by go-swagger; DO NOT EDIT.

package filepool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new filepool API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for filepool API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateFilepoolPolicy Create a new policy.
*/
func (a *Client) CreateFilepoolPolicy(params *CreateFilepoolPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*CreateFilepoolPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateFilepoolPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createFilepoolPolicy",
		Method:             "POST",
		PathPattern:        "/platform/1/filepool/policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateFilepoolPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateFilepoolPolicyOK), nil

}

/*
DeleteFilepoolPolicy Delete file pool policy.
*/
func (a *Client) DeleteFilepoolPolicy(params *DeleteFilepoolPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteFilepoolPolicyNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFilepoolPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteFilepoolPolicy",
		Method:             "DELETE",
		PathPattern:        "/platform/1/filepool/policies/{FilepoolPolicyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteFilepoolPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteFilepoolPolicyNoContent), nil

}

/*
GetFilepoolDefaultPolicy List default file pool policy.
*/
func (a *Client) GetFilepoolDefaultPolicy(params *GetFilepoolDefaultPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*GetFilepoolDefaultPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFilepoolDefaultPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFilepoolDefaultPolicy",
		Method:             "GET",
		PathPattern:        "/platform/1/filepool/default-policy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFilepoolDefaultPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFilepoolDefaultPolicyOK), nil

}

/*
GetFilepoolPolicy Retrieve file pool policy information.
*/
func (a *Client) GetFilepoolPolicy(params *GetFilepoolPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*GetFilepoolPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFilepoolPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFilepoolPolicy",
		Method:             "GET",
		PathPattern:        "/platform/1/filepool/policies/{FilepoolPolicyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFilepoolPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFilepoolPolicyOK), nil

}

/*
GetFilepoolTemplate List all templates.
*/
func (a *Client) GetFilepoolTemplate(params *GetFilepoolTemplateParams, authInfo runtime.ClientAuthInfoWriter) (*GetFilepoolTemplateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFilepoolTemplateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFilepoolTemplate",
		Method:             "GET",
		PathPattern:        "/platform/1/filepool/templates/{FilepoolTemplateId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFilepoolTemplateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFilepoolTemplateOK), nil

}

/*
GetFilepoolTemplates List all templates.
*/
func (a *Client) GetFilepoolTemplates(params *GetFilepoolTemplatesParams, authInfo runtime.ClientAuthInfoWriter) (*GetFilepoolTemplatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFilepoolTemplatesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFilepoolTemplates",
		Method:             "GET",
		PathPattern:        "/platform/1/filepool/templates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFilepoolTemplatesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFilepoolTemplatesOK), nil

}

/*
ListFilepoolPolicies List all policies.
*/
func (a *Client) ListFilepoolPolicies(params *ListFilepoolPoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*ListFilepoolPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListFilepoolPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listFilepoolPolicies",
		Method:             "GET",
		PathPattern:        "/platform/1/filepool/policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListFilepoolPoliciesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListFilepoolPoliciesOK), nil

}

/*
UpdateFilepoolDefaultPolicy Set default file pool policy.
*/
func (a *Client) UpdateFilepoolDefaultPolicy(params *UpdateFilepoolDefaultPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateFilepoolDefaultPolicyNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateFilepoolDefaultPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateFilepoolDefaultPolicy",
		Method:             "PUT",
		PathPattern:        "/platform/1/filepool/default-policy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateFilepoolDefaultPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateFilepoolDefaultPolicyNoContent), nil

}

/*
UpdateFilepoolPolicy Modify file pool policy. All input fields are optional, but one or more must be supplied.
*/
func (a *Client) UpdateFilepoolPolicy(params *UpdateFilepoolPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateFilepoolPolicyNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateFilepoolPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateFilepoolPolicy",
		Method:             "PUT",
		PathPattern:        "/platform/1/filepool/policies/{FilepoolPolicyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateFilepoolPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateFilepoolPolicyNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
