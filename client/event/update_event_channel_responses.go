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

// UpdateEventChannelReader is a Reader for the UpdateEventChannel structure.
type UpdateEventChannelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateEventChannelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateEventChannelNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateEventChannelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateEventChannelNoContent creates a UpdateEventChannelNoContent with default headers values
func NewUpdateEventChannelNoContent() *UpdateEventChannelNoContent {
	return &UpdateEventChannelNoContent{}
}

/*UpdateEventChannelNoContent handles this case with default header values.

Success.
*/
type UpdateEventChannelNoContent struct {
}

func (o *UpdateEventChannelNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/3/event/channels/{EventChannelId}][%d] updateEventChannelNoContent ", 204)
}

func (o *UpdateEventChannelNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateEventChannelDefault creates a UpdateEventChannelDefault with default headers values
func NewUpdateEventChannelDefault(code int) *UpdateEventChannelDefault {
	return &UpdateEventChannelDefault{
		_statusCode: code,
	}
}

/*UpdateEventChannelDefault handles this case with default header values.

Unexpected error
*/
type UpdateEventChannelDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update event channel default response
func (o *UpdateEventChannelDefault) Code() int {
	return o._statusCode
}

func (o *UpdateEventChannelDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/3/event/channels/{EventChannelId}][%d] updateEventChannel default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateEventChannelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
