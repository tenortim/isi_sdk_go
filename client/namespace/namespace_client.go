// Code generated by go-swagger; DO NOT EDIT.

package namespace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new namespace API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for namespace API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CopyDirectory Recursively copies a directory to a specified destination path. Symbolic links are copied as regular files.
*/
func (a *Client) CopyDirectory(params *CopyDirectoryParams, authInfo runtime.ClientAuthInfoWriter) (*CopyDirectoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCopyDirectoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "copyDirectory",
		Method:             "PUT",
		PathPattern:        "/namespace/{DirectoryCopyTarget}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CopyDirectoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CopyDirectoryOK), nil

}

/*
CopyFile Copies a file to the specified destination path.
*/
func (a *Client) CopyFile(params *CopyFileParams, authInfo runtime.ClientAuthInfoWriter) (*CopyFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCopyFileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "copyFile",
		Method:             "PUT",
		PathPattern:        "/namespace/{FileCopyTarget}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CopyFileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CopyFileOK), nil

}

/*
CreateAccessPoint Creates a namespace access point in the file system. Only root users can create or change namespace access points.
*/
func (a *Client) CreateAccessPoint(params *CreateAccessPointParams, authInfo runtime.ClientAuthInfoWriter) (*CreateAccessPointOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAccessPointParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createAccessPoint",
		Method:             "PUT",
		PathPattern:        "/namespace/{AccessPointName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateAccessPointReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateAccessPointOK), nil

}

/*
CreateDirectory Creates a directory with a specified path.
*/
func (a *Client) CreateDirectory(params *CreateDirectoryParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDirectoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDirectoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createDirectory",
		Method:             "PUT",
		PathPattern:        "/namespace/{DirectoryPath}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDirectoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateDirectoryOK), nil

}

/*
CreateFile Creates a file object with a given path. Note that file streaming is not supported.
*/
func (a *Client) CreateFile(params *CreateFileParams, authInfo runtime.ClientAuthInfoWriter) (*CreateFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateFileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createFile",
		Method:             "PUT",
		PathPattern:        "/namespace/{FilePath}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateFileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateFileOK), nil

}

/*
DeleteAccessPoint Deletes a namespace access point. Only root users can delete namespace access points.
*/
func (a *Client) DeleteAccessPoint(params *DeleteAccessPointParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteAccessPointOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAccessPointParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteAccessPoint",
		Method:             "DELETE",
		PathPattern:        "/namespace/{AccessPointName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAccessPointReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteAccessPointOK), nil

}

/*
DeleteDirectory Deletes the directory at the specified path.
*/
func (a *Client) DeleteDirectory(params *DeleteDirectoryParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteDirectoryNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDirectoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteDirectory",
		Method:             "DELETE",
		PathPattern:        "/namespace/{DirectoryPath}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDirectoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteDirectoryNoContent), nil

}

/*
DeleteFile Deletes the specified file.
*/
func (a *Client) DeleteFile(params *DeleteFileParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteFileNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteFile",
		Method:             "DELETE",
		PathPattern:        "/namespace/{FilePath}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteFileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteFileNoContent), nil

}

/*
GetACL Retrieves the access control list for a namespace object.
*/
func (a *Client) GetACL(params *GetACLParams, authInfo runtime.ClientAuthInfoWriter) (*GetACLOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetACLParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAcl",
		Method:             "GET",
		PathPattern:        "/namespace/{NamespacePath}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetACLReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetACLOK), nil

}

/*
GetDirectoryAttributes Retrieves the attribute information for a specified directory without transferring the contents of the directory.
*/
func (a *Client) GetDirectoryAttributes(params *GetDirectoryAttributesParams, authInfo runtime.ClientAuthInfoWriter) (*GetDirectoryAttributesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDirectoryAttributesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDirectoryAttributes",
		Method:             "HEAD",
		PathPattern:        "/namespace/{DirectoryPath}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDirectoryAttributesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDirectoryAttributesOK), nil

}

/*
GetDirectoryContents Retrieves a list of files and subdirectories from a directory.
*/
func (a *Client) GetDirectoryContents(params *GetDirectoryContentsParams, authInfo runtime.ClientAuthInfoWriter) (*GetDirectoryContentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDirectoryContentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDirectoryContents",
		Method:             "GET",
		PathPattern:        "/namespace/{DirectoryPath}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDirectoryContentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDirectoryContentsOK), nil

}

/*
GetDirectoryMetadata Retrieves the attribute information for a specified directory with the metadata query argument.
*/
func (a *Client) GetDirectoryMetadata(params *GetDirectoryMetadataParams, authInfo runtime.ClientAuthInfoWriter) (*GetDirectoryMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDirectoryMetadataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDirectoryMetadata",
		Method:             "GET",
		PathPattern:        "/namespace/{DirectoryMetadataPath}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDirectoryMetadataReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDirectoryMetadataOK), nil

}

/*
GetFileAttributes Retrieves the attribute information for a specified file.
*/
func (a *Client) GetFileAttributes(params *GetFileAttributesParams, authInfo runtime.ClientAuthInfoWriter) (*GetFileAttributesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFileAttributesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFileAttributes",
		Method:             "HEAD",
		PathPattern:        "/namespace/{FilePath}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFileAttributesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFileAttributesOK), nil

}

/*
GetFileContents Retrieves the contents of a file from a specified path. Note that file streaming is not supported.
*/
func (a *Client) GetFileContents(params *GetFileContentsParams, authInfo runtime.ClientAuthInfoWriter) (*GetFileContentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFileContentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFileContents",
		Method:             "GET",
		PathPattern:        "/namespace/{FilePath}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFileContentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFileContentsOK), nil

}

/*
GetFileMetadata Retrieves the attribute information for a specified file with the metadata query argument.
*/
func (a *Client) GetFileMetadata(params *GetFileMetadataParams, authInfo runtime.ClientAuthInfoWriter) (*GetFileMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFileMetadataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFileMetadata",
		Method:             "GET",
		PathPattern:        "/namespace/{FileMetadataPath}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFileMetadataReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFileMetadataOK), nil

}

/*
GetWormProperties Retrieves the WORM retention date and committed state of the file.
*/
func (a *Client) GetWormProperties(params *GetWormPropertiesParams, authInfo runtime.ClientAuthInfoWriter) (*GetWormPropertiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWormPropertiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWormProperties",
		Method:             "GET",
		PathPattern:        "/namespace/{WormFilePath}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWormPropertiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWormPropertiesOK), nil

}

/*
ListAccessPoints Retrieves the namespace access points available for the authenticated user.
*/
func (a *Client) ListAccessPoints(params *ListAccessPointsParams, authInfo runtime.ClientAuthInfoWriter) (*ListAccessPointsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAccessPointsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAccessPoints",
		Method:             "GET",
		PathPattern:        "/namespace",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAccessPointsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListAccessPointsOK), nil

}

/*
MoveDirectory Moves a directory from an existing source to a new destination path.
*/
func (a *Client) MoveDirectory(params *MoveDirectoryParams, authInfo runtime.ClientAuthInfoWriter) (*MoveDirectoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMoveDirectoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "moveDirectory",
		Method:             "POST",
		PathPattern:        "/namespace/{DirectoryPath}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MoveDirectoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MoveDirectoryOK), nil

}

/*
MoveFile Moves a file to a destination path that does not yet exist.
*/
func (a *Client) MoveFile(params *MoveFileParams, authInfo runtime.ClientAuthInfoWriter) (*MoveFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMoveFileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "moveFile",
		Method:             "POST",
		PathPattern:        "/namespace/{FilePath}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MoveFileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MoveFileOK), nil

}

/*
QueryDirectory Query objects by system-defined and user-defined attributes in a directory.
*/
func (a *Client) QueryDirectory(params *QueryDirectoryParams, authInfo runtime.ClientAuthInfoWriter) (*QueryDirectoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryDirectoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "queryDirectory",
		Method:             "POST",
		PathPattern:        "/namespace/{QueryPath}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryDirectoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*QueryDirectoryOK), nil

}

/*
SetACL Sets the access control list for a namespace.
*/
func (a *Client) SetACL(params *SetACLParams, authInfo runtime.ClientAuthInfoWriter) (*SetACLOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetACLParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setAcl",
		Method:             "PUT",
		PathPattern:        "/namespace/{NamespacePath}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetACLReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetACLOK), nil

}

/*
SetDirectoryMetadata Sets attributes on a specified directory with the metadata query argument.
*/
func (a *Client) SetDirectoryMetadata(params *SetDirectoryMetadataParams, authInfo runtime.ClientAuthInfoWriter) (*SetDirectoryMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetDirectoryMetadataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setDirectoryMetadata",
		Method:             "PUT",
		PathPattern:        "/namespace/{DirectoryMetadataPath}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetDirectoryMetadataReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetDirectoryMetadataOK), nil

}

/*
SetFileMetadata Sets attributes on a specified file with the metadata query argument through the JSON body.
*/
func (a *Client) SetFileMetadata(params *SetFileMetadataParams, authInfo runtime.ClientAuthInfoWriter) (*SetFileMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetFileMetadataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setFileMetadata",
		Method:             "PUT",
		PathPattern:        "/namespace/{FileMetadataPath}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetFileMetadataReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetFileMetadataOK), nil

}

/*
SetWormProperties Sets the retention period and commits a file in a SmartLock directory.
*/
func (a *Client) SetWormProperties(params *SetWormPropertiesParams, authInfo runtime.ClientAuthInfoWriter) (*SetWormPropertiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetWormPropertiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setWormProperties",
		Method:             "PUT",
		PathPattern:        "/namespace/{WormFilePath}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetWormPropertiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetWormPropertiesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}