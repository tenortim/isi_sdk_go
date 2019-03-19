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

// CreateHardwareTapeNameReader is a Reader for the CreateHardwareTapeName structure.
type CreateHardwareTapeNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateHardwareTapeNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateHardwareTapeNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateHardwareTapeNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateHardwareTapeNameOK creates a CreateHardwareTapeNameOK with default headers values
func NewCreateHardwareTapeNameOK() *CreateHardwareTapeNameOK {
	return &CreateHardwareTapeNameOK{}
}

/*CreateHardwareTapeNameOK handles this case with default header values.

Tape/Changer devices rescan
*/
type CreateHardwareTapeNameOK struct {
	Payload models.Empty
}

func (o *CreateHardwareTapeNameOK) Error() string {
	return fmt.Sprintf("[POST /platform/3/hardware/tape/{HardwareTapeName}][%d] createHardwareTapeNameOK  %+v", 200, o.Payload)
}

func (o *CreateHardwareTapeNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateHardwareTapeNameDefault creates a CreateHardwareTapeNameDefault with default headers values
func NewCreateHardwareTapeNameDefault(code int) *CreateHardwareTapeNameDefault {
	return &CreateHardwareTapeNameDefault{
		_statusCode: code,
	}
}

/*CreateHardwareTapeNameDefault handles this case with default header values.

Unexpected error
*/
type CreateHardwareTapeNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create hardware tape name default response
func (o *CreateHardwareTapeNameDefault) Code() int {
	return o._statusCode
}

func (o *CreateHardwareTapeNameDefault) Error() string {
	return fmt.Sprintf("[POST /platform/3/hardware/tape/{HardwareTapeName}][%d] createHardwareTapeName default  %+v", o._statusCode, o.Payload)
}

func (o *CreateHardwareTapeNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
