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

// CreateCloudPoolReader is a Reader for the CreateCloudPool structure.
type CreateCloudPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCloudPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateCloudPoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateCloudPoolDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateCloudPoolOK creates a CreateCloudPoolOK with default headers values
func NewCreateCloudPoolOK() *CreateCloudPoolOK {
	return &CreateCloudPoolOK{}
}

/*CreateCloudPoolOK handles this case with default header values.

Create a new pool.
*/
type CreateCloudPoolOK struct {
	Payload *models.CreateCloudPoolResponse
}

func (o *CreateCloudPoolOK) Error() string {
	return fmt.Sprintf("[POST /platform/3/cloud/pools][%d] createCloudPoolOK  %+v", 200, o.Payload)
}

func (o *CreateCloudPoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateCloudPoolResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCloudPoolDefault creates a CreateCloudPoolDefault with default headers values
func NewCreateCloudPoolDefault(code int) *CreateCloudPoolDefault {
	return &CreateCloudPoolDefault{
		_statusCode: code,
	}
}

/*CreateCloudPoolDefault handles this case with default header values.

Unexpected error
*/
type CreateCloudPoolDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create cloud pool default response
func (o *CreateCloudPoolDefault) Code() int {
	return o._statusCode
}

func (o *CreateCloudPoolDefault) Error() string {
	return fmt.Sprintf("[POST /platform/3/cloud/pools][%d] createCloudPool default  %+v", o._statusCode, o.Payload)
}

func (o *CreateCloudPoolDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}