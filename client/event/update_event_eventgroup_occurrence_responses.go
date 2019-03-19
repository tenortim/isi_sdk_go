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

// UpdateEventEventgroupOccurrenceReader is a Reader for the UpdateEventEventgroupOccurrence structure.
type UpdateEventEventgroupOccurrenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateEventEventgroupOccurrenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateEventEventgroupOccurrenceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateEventEventgroupOccurrenceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateEventEventgroupOccurrenceNoContent creates a UpdateEventEventgroupOccurrenceNoContent with default headers values
func NewUpdateEventEventgroupOccurrenceNoContent() *UpdateEventEventgroupOccurrenceNoContent {
	return &UpdateEventEventgroupOccurrenceNoContent{}
}

/*UpdateEventEventgroupOccurrenceNoContent handles this case with default header values.

Success.
*/
type UpdateEventEventgroupOccurrenceNoContent struct {
}

func (o *UpdateEventEventgroupOccurrenceNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/3/event/eventgroup-occurrences/{EventEventgroupOccurrenceId}][%d] updateEventEventgroupOccurrenceNoContent ", 204)
}

func (o *UpdateEventEventgroupOccurrenceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateEventEventgroupOccurrenceDefault creates a UpdateEventEventgroupOccurrenceDefault with default headers values
func NewUpdateEventEventgroupOccurrenceDefault(code int) *UpdateEventEventgroupOccurrenceDefault {
	return &UpdateEventEventgroupOccurrenceDefault{
		_statusCode: code,
	}
}

/*UpdateEventEventgroupOccurrenceDefault handles this case with default header values.

Unexpected error
*/
type UpdateEventEventgroupOccurrenceDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update event eventgroup occurrence default response
func (o *UpdateEventEventgroupOccurrenceDefault) Code() int {
	return o._statusCode
}

func (o *UpdateEventEventgroupOccurrenceDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/3/event/eventgroup-occurrences/{EventEventgroupOccurrenceId}][%d] updateEventEventgroupOccurrence default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateEventEventgroupOccurrenceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}