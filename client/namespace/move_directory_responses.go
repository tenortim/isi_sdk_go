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

// MoveDirectoryReader is a Reader for the MoveDirectory structure.
type MoveDirectoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MoveDirectoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewMoveDirectoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMoveDirectoryOK creates a MoveDirectoryOK with default headers values
func NewMoveDirectoryOK() *MoveDirectoryOK {
	return &MoveDirectoryOK{}
}

/*MoveDirectoryOK handles this case with default header values.

Move directory.
*/
type MoveDirectoryOK struct {
	Payload models.Empty
}

func (o *MoveDirectoryOK) Error() string {
	return fmt.Sprintf("[POST /namespace/{DirectoryPath}][%d] moveDirectoryOK  %+v", 200, o.Payload)
}

func (o *MoveDirectoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
