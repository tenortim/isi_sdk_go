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

// GetFileMetadataReader is a Reader for the GetFileMetadata structure.
type GetFileMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFileMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFileMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFileMetadataOK creates a GetFileMetadataOK with default headers values
func NewGetFileMetadataOK() *GetFileMetadataOK {
	return &GetFileMetadataOK{}
}

/*GetFileMetadataOK handles this case with default header values.

Get file metadata.
*/
type GetFileMetadataOK struct {
	Payload *models.NamespaceMetadataList
}

func (o *GetFileMetadataOK) Error() string {
	return fmt.Sprintf("[GET /namespace/{FileMetadataPath}][%d] getFileMetadataOK  %+v", 200, o.Payload)
}

func (o *GetFileMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NamespaceMetadataList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}