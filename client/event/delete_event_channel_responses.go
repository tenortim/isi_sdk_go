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

// DeleteEventChannelReader is a Reader for the DeleteEventChannel structure.
type DeleteEventChannelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEventChannelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteEventChannelNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteEventChannelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteEventChannelNoContent creates a DeleteEventChannelNoContent with default headers values
func NewDeleteEventChannelNoContent() *DeleteEventChannelNoContent {
	return &DeleteEventChannelNoContent{}
}

/*DeleteEventChannelNoContent handles this case with default header values.

Success.
*/
type DeleteEventChannelNoContent struct {
}

func (o *DeleteEventChannelNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/event/channels/{EventChannelId}][%d] deleteEventChannelNoContent ", 204)
}

func (o *DeleteEventChannelNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEventChannelDefault creates a DeleteEventChannelDefault with default headers values
func NewDeleteEventChannelDefault(code int) *DeleteEventChannelDefault {
	return &DeleteEventChannelDefault{
		_statusCode: code,
	}
}

/*DeleteEventChannelDefault handles this case with default header values.

Unexpected error
*/
type DeleteEventChannelDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete event channel default response
func (o *DeleteEventChannelDefault) Code() int {
	return o._statusCode
}

func (o *DeleteEventChannelDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/event/channels/{EventChannelId}][%d] deleteEventChannel default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteEventChannelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}