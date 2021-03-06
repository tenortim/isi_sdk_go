// Code generated by go-swagger; DO NOT EDIT.

package upgrade

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// CreateClusterArchiveItemReader is a Reader for the CreateClusterArchiveItem structure.
type CreateClusterArchiveItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateClusterArchiveItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateClusterArchiveItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateClusterArchiveItemDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateClusterArchiveItemOK creates a CreateClusterArchiveItemOK with default headers values
func NewCreateClusterArchiveItemOK() *CreateClusterArchiveItemOK {
	return &CreateClusterArchiveItemOK{}
}

/*CreateClusterArchiveItemOK handles this case with default header values.

Start an archive of an upgrade.
*/
type CreateClusterArchiveItemOK struct {
	Payload models.Empty
}

func (o *CreateClusterArchiveItemOK) Error() string {
	return fmt.Sprintf("[POST /platform/3/upgrade/cluster/archive][%d] createClusterArchiveItemOK  %+v", 200, o.Payload)
}

func (o *CreateClusterArchiveItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterArchiveItemDefault creates a CreateClusterArchiveItemDefault with default headers values
func NewCreateClusterArchiveItemDefault(code int) *CreateClusterArchiveItemDefault {
	return &CreateClusterArchiveItemDefault{
		_statusCode: code,
	}
}

/*CreateClusterArchiveItemDefault handles this case with default header values.

Unexpected error
*/
type CreateClusterArchiveItemDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create cluster archive item default response
func (o *CreateClusterArchiveItemDefault) Code() int {
	return o._statusCode
}

func (o *CreateClusterArchiveItemDefault) Error() string {
	return fmt.Sprintf("[POST /platform/3/upgrade/cluster/archive][%d] createClusterArchiveItem default  %+v", o._statusCode, o.Payload)
}

func (o *CreateClusterArchiveItemDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
