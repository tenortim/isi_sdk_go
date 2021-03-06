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

// GetHardwareTapesReader is a Reader for the GetHardwareTapes structure.
type GetHardwareTapesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHardwareTapesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetHardwareTapesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetHardwareTapesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHardwareTapesOK creates a GetHardwareTapesOK with default headers values
func NewGetHardwareTapesOK() *GetHardwareTapesOK {
	return &GetHardwareTapesOK{}
}

/*GetHardwareTapesOK handles this case with default header values.

Get list Tape and Changer devices
*/
type GetHardwareTapesOK struct {
	Payload *models.HardwareTapes
}

func (o *GetHardwareTapesOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/hardware/tapes][%d] getHardwareTapesOK  %+v", 200, o.Payload)
}

func (o *GetHardwareTapesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HardwareTapes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHardwareTapesDefault creates a GetHardwareTapesDefault with default headers values
func NewGetHardwareTapesDefault(code int) *GetHardwareTapesDefault {
	return &GetHardwareTapesDefault{
		_statusCode: code,
	}
}

/*GetHardwareTapesDefault handles this case with default header values.

Unexpected error
*/
type GetHardwareTapesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get hardware tapes default response
func (o *GetHardwareTapesDefault) Code() int {
	return o._statusCode
}

func (o *GetHardwareTapesDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/hardware/tapes][%d] getHardwareTapes default  %+v", o._statusCode, o.Payload)
}

func (o *GetHardwareTapesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
