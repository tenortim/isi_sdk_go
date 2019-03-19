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

// QueryDirectoryReader is a Reader for the QueryDirectory structure.
type QueryDirectoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryDirectoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewQueryDirectoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewQueryDirectoryOK creates a QueryDirectoryOK with default headers values
func NewQueryDirectoryOK() *QueryDirectoryOK {
	return &QueryDirectoryOK{}
}

/*QueryDirectoryOK handles this case with default header values.

Directory query results.
*/
type QueryDirectoryOK struct {
	Payload *models.NamespaceObjects
}

func (o *QueryDirectoryOK) Error() string {
	return fmt.Sprintf("[POST /namespace/{QueryPath}][%d] queryDirectoryOK  %+v", 200, o.Payload)
}

func (o *QueryDirectoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NamespaceObjects)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
