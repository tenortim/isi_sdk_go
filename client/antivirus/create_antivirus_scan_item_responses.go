// Code generated by go-swagger; DO NOT EDIT.

package antivirus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// CreateAntivirusScanItemReader is a Reader for the CreateAntivirusScanItem structure.
type CreateAntivirusScanItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAntivirusScanItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateAntivirusScanItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateAntivirusScanItemDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateAntivirusScanItemOK creates a CreateAntivirusScanItemOK with default headers values
func NewCreateAntivirusScanItemOK() *CreateAntivirusScanItemOK {
	return &CreateAntivirusScanItemOK{}
}

/*CreateAntivirusScanItemOK handles this case with default header values.

Manually scan a file.
*/
type CreateAntivirusScanItemOK struct {
	Payload *models.CreateAntivirusScanItemResponse
}

func (o *CreateAntivirusScanItemOK) Error() string {
	return fmt.Sprintf("[POST /platform/3/antivirus/scan][%d] createAntivirusScanItemOK  %+v", 200, o.Payload)
}

func (o *CreateAntivirusScanItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateAntivirusScanItemResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAntivirusScanItemDefault creates a CreateAntivirusScanItemDefault with default headers values
func NewCreateAntivirusScanItemDefault(code int) *CreateAntivirusScanItemDefault {
	return &CreateAntivirusScanItemDefault{
		_statusCode: code,
	}
}

/*CreateAntivirusScanItemDefault handles this case with default header values.

Unexpected error
*/
type CreateAntivirusScanItemDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create antivirus scan item default response
func (o *CreateAntivirusScanItemDefault) Code() int {
	return o._statusCode
}

func (o *CreateAntivirusScanItemDefault) Error() string {
	return fmt.Sprintf("[POST /platform/3/antivirus/scan][%d] createAntivirusScanItem default  %+v", o._statusCode, o.Payload)
}

func (o *CreateAntivirusScanItemDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}