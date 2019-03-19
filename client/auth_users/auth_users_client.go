// Code generated by go-swagger; DO NOT EDIT.

package auth_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new auth users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for auth users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateUserMemberOfItem Add the user to a group.
*/
func (a *Client) CreateUserMemberOfItem(params *CreateUserMemberOfItemParams, authInfo runtime.ClientAuthInfoWriter) (*CreateUserMemberOfItemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUserMemberOfItemParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createUserMemberOfItem",
		Method:             "POST",
		PathPattern:        "/platform/3/auth/users/{User}/member-of",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateUserMemberOfItemReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateUserMemberOfItemOK), nil

}

/*
DeleteUserMemberOfMemberOf Remove the user from the group.
*/
func (a *Client) DeleteUserMemberOfMemberOf(params *DeleteUserMemberOfMemberOfParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteUserMemberOfMemberOfNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserMemberOfMemberOfParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteUserMemberOfMemberOf",
		Method:             "DELETE",
		PathPattern:        "/platform/3/auth/users/{User}/member-of/{UserMemberOfMemberOf}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUserMemberOfMemberOfReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteUserMemberOfMemberOfNoContent), nil

}

/*
ListUserMemberOf List all groups the user is a member of.
*/
func (a *Client) ListUserMemberOf(params *ListUserMemberOfParams, authInfo runtime.ClientAuthInfoWriter) (*ListUserMemberOfOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListUserMemberOfParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listUserMemberOf",
		Method:             "GET",
		PathPattern:        "/platform/3/auth/users/{User}/member-of",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListUserMemberOfReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListUserMemberOfOK), nil

}

/*
UpdateUserChangePassword Change the user's password.
*/
func (a *Client) UpdateUserChangePassword(params *UpdateUserChangePasswordParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateUserChangePasswordNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserChangePasswordParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateUserChangePassword",
		Method:             "PUT",
		PathPattern:        "/platform/3/auth/users/{User}/change-password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateUserChangePasswordReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateUserChangePasswordNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}