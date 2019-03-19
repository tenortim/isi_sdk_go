// Code generated by go-swagger; DO NOT EDIT.

package filepool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// CreateFilepoolPolicyReader is a Reader for the CreateFilepoolPolicy structure.
type CreateFilepoolPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateFilepoolPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateFilepoolPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateFilepoolPolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateFilepoolPolicyOK creates a CreateFilepoolPolicyOK with default headers values
func NewCreateFilepoolPolicyOK() *CreateFilepoolPolicyOK {
	return &CreateFilepoolPolicyOK{}
}

/*CreateFilepoolPolicyOK handles this case with default header values.

Create a new policy.
*/
type CreateFilepoolPolicyOK struct {
	Payload *models.CreateFilepoolPolicyResponse
}

func (o *CreateFilepoolPolicyOK) Error() string {
	return fmt.Sprintf("[POST /platform/1/filepool/policies][%d] createFilepoolPolicyOK  %+v", 200, o.Payload)
}

func (o *CreateFilepoolPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateFilepoolPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFilepoolPolicyDefault creates a CreateFilepoolPolicyDefault with default headers values
func NewCreateFilepoolPolicyDefault(code int) *CreateFilepoolPolicyDefault {
	return &CreateFilepoolPolicyDefault{
		_statusCode: code,
	}
}

/*CreateFilepoolPolicyDefault handles this case with default header values.

Unexpected error
*/
type CreateFilepoolPolicyDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create filepool policy default response
func (o *CreateFilepoolPolicyDefault) Code() int {
	return o._statusCode
}

func (o *CreateFilepoolPolicyDefault) Error() string {
	return fmt.Sprintf("[POST /platform/1/filepool/policies][%d] createFilepoolPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *CreateFilepoolPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}