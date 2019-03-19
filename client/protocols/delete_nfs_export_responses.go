// Code generated by go-swagger; DO NOT EDIT.

package protocols

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// DeleteNfsExportReader is a Reader for the DeleteNfsExport structure.
type DeleteNfsExportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNfsExportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteNfsExportNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteNfsExportDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNfsExportNoContent creates a DeleteNfsExportNoContent with default headers values
func NewDeleteNfsExportNoContent() *DeleteNfsExportNoContent {
	return &DeleteNfsExportNoContent{}
}

/*DeleteNfsExportNoContent handles this case with default header values.

Success.
*/
type DeleteNfsExportNoContent struct {
}

func (o *DeleteNfsExportNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/2/protocols/nfs/exports/{NfsExportId}][%d] deleteNfsExportNoContent ", 204)
}

func (o *DeleteNfsExportNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNfsExportDefault creates a DeleteNfsExportDefault with default headers values
func NewDeleteNfsExportDefault(code int) *DeleteNfsExportDefault {
	return &DeleteNfsExportDefault{
		_statusCode: code,
	}
}

/*DeleteNfsExportDefault handles this case with default header values.

Unexpected error
*/
type DeleteNfsExportDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete nfs export default response
func (o *DeleteNfsExportDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNfsExportDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/2/protocols/nfs/exports/{NfsExportId}][%d] deleteNfsExport default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteNfsExportDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}