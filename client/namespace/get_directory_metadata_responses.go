// Code generated by go-swagger; DO NOT EDIT.

package namespace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetDirectoryMetadataReader is a Reader for the GetDirectoryMetadata structure.
type GetDirectoryMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDirectoryMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDirectoryMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDirectoryMetadataOK creates a GetDirectoryMetadataOK with default headers values
func NewGetDirectoryMetadataOK() *GetDirectoryMetadataOK {
	return &GetDirectoryMetadataOK{}
}

/*GetDirectoryMetadataOK handles this case with default header values.

Get directory metadata.
*/
type GetDirectoryMetadataOK struct {
	Payload *models.NamespaceMetadataList
}

func (o *GetDirectoryMetadataOK) Error() string {
	return fmt.Sprintf("[GET /namespace/{DirectoryMetadataPath}][%d] getDirectoryMetadataOK  %+v", 200, o.Payload)
}

func (o *GetDirectoryMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NamespaceMetadataList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
