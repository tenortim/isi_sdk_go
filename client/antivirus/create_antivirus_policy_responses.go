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

// CreateAntivirusPolicyReader is a Reader for the CreateAntivirusPolicy structure.
type CreateAntivirusPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAntivirusPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateAntivirusPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateAntivirusPolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateAntivirusPolicyOK creates a CreateAntivirusPolicyOK with default headers values
func NewCreateAntivirusPolicyOK() *CreateAntivirusPolicyOK {
	return &CreateAntivirusPolicyOK{}
}

/*CreateAntivirusPolicyOK handles this case with default header values.

Create new antivirus scan policies.
*/
type CreateAntivirusPolicyOK struct {
	Payload *models.CreateResponse
}

func (o *CreateAntivirusPolicyOK) Error() string {
	return fmt.Sprintf("[POST /platform/3/antivirus/policies][%d] createAntivirusPolicyOK  %+v", 200, o.Payload)
}

func (o *CreateAntivirusPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAntivirusPolicyDefault creates a CreateAntivirusPolicyDefault with default headers values
func NewCreateAntivirusPolicyDefault(code int) *CreateAntivirusPolicyDefault {
	return &CreateAntivirusPolicyDefault{
		_statusCode: code,
	}
}

/*CreateAntivirusPolicyDefault handles this case with default header values.

Unexpected error
*/
type CreateAntivirusPolicyDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create antivirus policy default response
func (o *CreateAntivirusPolicyDefault) Code() int {
	return o._statusCode
}

func (o *CreateAntivirusPolicyDefault) Error() string {
	return fmt.Sprintf("[POST /platform/3/antivirus/policies][%d] createAntivirusPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *CreateAntivirusPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
