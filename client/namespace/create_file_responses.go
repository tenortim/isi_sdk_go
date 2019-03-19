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

// CreateFileReader is a Reader for the CreateFile structure.
type CreateFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateFileOK creates a CreateFileOK with default headers values
func NewCreateFileOK() *CreateFileOK {
	return &CreateFileOK{}
}

/*CreateFileOK handles this case with default header values.

Created file object.
*/
type CreateFileOK struct {
	Payload models.Empty
}

func (o *CreateFileOK) Error() string {
	return fmt.Sprintf("[PUT /namespace/{FilePath}][%d] createFileOK  %+v", 200, o.Payload)
}

func (o *CreateFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
