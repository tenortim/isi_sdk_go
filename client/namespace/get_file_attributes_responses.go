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

// GetFileAttributesReader is a Reader for the GetFileAttributes structure.
type GetFileAttributesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFileAttributesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFileAttributesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFileAttributesOK creates a GetFileAttributesOK with default headers values
func NewGetFileAttributesOK() *GetFileAttributesOK {
	return &GetFileAttributesOK{}
}

/*GetFileAttributesOK handles this case with default header values.

Get attribute information for file.
*/
type GetFileAttributesOK struct {
	Payload models.Empty
}

func (o *GetFileAttributesOK) Error() string {
	return fmt.Sprintf("[HEAD /namespace/{FilePath}][%d] getFileAttributesOK  %+v", 200, o.Payload)
}

func (o *GetFileAttributesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
