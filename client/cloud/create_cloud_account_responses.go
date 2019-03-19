// Code generated by go-swagger; DO NOT EDIT.

package cloud

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// CreateCloudAccountReader is a Reader for the CreateCloudAccount structure.
type CreateCloudAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCloudAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateCloudAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateCloudAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateCloudAccountOK creates a CreateCloudAccountOK with default headers values
func NewCreateCloudAccountOK() *CreateCloudAccountOK {
	return &CreateCloudAccountOK{}
}

/*CreateCloudAccountOK handles this case with default header values.

Create a new account.
*/
type CreateCloudAccountOK struct {
	Payload *models.CreateCloudAccountResponse
}

func (o *CreateCloudAccountOK) Error() string {
	return fmt.Sprintf("[POST /platform/3/cloud/accounts][%d] createCloudAccountOK  %+v", 200, o.Payload)
}

func (o *CreateCloudAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateCloudAccountResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCloudAccountDefault creates a CreateCloudAccountDefault with default headers values
func NewCreateCloudAccountDefault(code int) *CreateCloudAccountDefault {
	return &CreateCloudAccountDefault{
		_statusCode: code,
	}
}

/*CreateCloudAccountDefault handles this case with default header values.

Unexpected error
*/
type CreateCloudAccountDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create cloud account default response
func (o *CreateCloudAccountDefault) Code() int {
	return o._statusCode
}

func (o *CreateCloudAccountDefault) Error() string {
	return fmt.Sprintf("[POST /platform/3/cloud/accounts][%d] createCloudAccount default  %+v", o._statusCode, o.Payload)
}

func (o *CreateCloudAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
