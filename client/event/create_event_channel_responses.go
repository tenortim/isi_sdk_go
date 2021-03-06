// Code generated by go-swagger; DO NOT EDIT.

package event

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// CreateEventChannelReader is a Reader for the CreateEventChannel structure.
type CreateEventChannelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateEventChannelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateEventChannelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateEventChannelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateEventChannelOK creates a CreateEventChannelOK with default headers values
func NewCreateEventChannelOK() *CreateEventChannelOK {
	return &CreateEventChannelOK{}
}

/*CreateEventChannelOK handles this case with default header values.

Create a new channel.
*/
type CreateEventChannelOK struct {
	Payload *models.CreateResponse
}

func (o *CreateEventChannelOK) Error() string {
	return fmt.Sprintf("[POST /platform/3/event/channels][%d] createEventChannelOK  %+v", 200, o.Payload)
}

func (o *CreateEventChannelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEventChannelDefault creates a CreateEventChannelDefault with default headers values
func NewCreateEventChannelDefault(code int) *CreateEventChannelDefault {
	return &CreateEventChannelDefault{
		_statusCode: code,
	}
}

/*CreateEventChannelDefault handles this case with default header values.

Unexpected error
*/
type CreateEventChannelDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create event channel default response
func (o *CreateEventChannelDefault) Code() int {
	return o._statusCode
}

func (o *CreateEventChannelDefault) Error() string {
	return fmt.Sprintf("[POST /platform/3/event/channels][%d] createEventChannel default  %+v", o._statusCode, o.Payload)
}

func (o *CreateEventChannelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
