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

// CreateNfsNlmSessionsCheckItemReader is a Reader for the CreateNfsNlmSessionsCheckItem structure.
type CreateNfsNlmSessionsCheckItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNfsNlmSessionsCheckItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateNfsNlmSessionsCheckItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateNfsNlmSessionsCheckItemDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateNfsNlmSessionsCheckItemOK creates a CreateNfsNlmSessionsCheckItemOK with default headers values
func NewCreateNfsNlmSessionsCheckItemOK() *CreateNfsNlmSessionsCheckItemOK {
	return &CreateNfsNlmSessionsCheckItemOK{}
}

/*CreateNfsNlmSessionsCheckItemOK handles this case with default header values.

Perform an active scan for lost NFSv3 locks.
*/
type CreateNfsNlmSessionsCheckItemOK struct {
	Payload *models.CreateNfsNlmSessionsCheckItemResponse
}

func (o *CreateNfsNlmSessionsCheckItemOK) Error() string {
	return fmt.Sprintf("[POST /platform/3/protocols/nfs/nlm/sessions-check][%d] createNfsNlmSessionsCheckItemOK  %+v", 200, o.Payload)
}

func (o *CreateNfsNlmSessionsCheckItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateNfsNlmSessionsCheckItemResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNfsNlmSessionsCheckItemDefault creates a CreateNfsNlmSessionsCheckItemDefault with default headers values
func NewCreateNfsNlmSessionsCheckItemDefault(code int) *CreateNfsNlmSessionsCheckItemDefault {
	return &CreateNfsNlmSessionsCheckItemDefault{
		_statusCode: code,
	}
}

/*CreateNfsNlmSessionsCheckItemDefault handles this case with default header values.

Unexpected error
*/
type CreateNfsNlmSessionsCheckItemDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create nfs nlm sessions check item default response
func (o *CreateNfsNlmSessionsCheckItemDefault) Code() int {
	return o._statusCode
}

func (o *CreateNfsNlmSessionsCheckItemDefault) Error() string {
	return fmt.Sprintf("[POST /platform/3/protocols/nfs/nlm/sessions-check][%d] createNfsNlmSessionsCheckItem default  %+v", o._statusCode, o.Payload)
}

func (o *CreateNfsNlmSessionsCheckItemDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}