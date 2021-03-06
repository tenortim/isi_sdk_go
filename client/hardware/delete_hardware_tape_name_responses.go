// Code generated by go-swagger; DO NOT EDIT.

package hardware

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// DeleteHardwareTapeNameReader is a Reader for the DeleteHardwareTapeName structure.
type DeleteHardwareTapeNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteHardwareTapeNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteHardwareTapeNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteHardwareTapeNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteHardwareTapeNameNoContent creates a DeleteHardwareTapeNameNoContent with default headers values
func NewDeleteHardwareTapeNameNoContent() *DeleteHardwareTapeNameNoContent {
	return &DeleteHardwareTapeNameNoContent{}
}

/*DeleteHardwareTapeNameNoContent handles this case with default header values.

Success.
*/
type DeleteHardwareTapeNameNoContent struct {
}

func (o *DeleteHardwareTapeNameNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/hardware/tape/{HardwareTapeName}][%d] deleteHardwareTapeNameNoContent ", 204)
}

func (o *DeleteHardwareTapeNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteHardwareTapeNameDefault creates a DeleteHardwareTapeNameDefault with default headers values
func NewDeleteHardwareTapeNameDefault(code int) *DeleteHardwareTapeNameDefault {
	return &DeleteHardwareTapeNameDefault{
		_statusCode: code,
	}
}

/*DeleteHardwareTapeNameDefault handles this case with default header values.

Unexpected error
*/
type DeleteHardwareTapeNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete hardware tape name default response
func (o *DeleteHardwareTapeNameDefault) Code() int {
	return o._statusCode
}

func (o *DeleteHardwareTapeNameDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/hardware/tape/{HardwareTapeName}][%d] deleteHardwareTapeName default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteHardwareTapeNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
