// Code generated by go-swagger; DO NOT EDIT.

package snapshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// UpdateSnapshotScheduleReader is a Reader for the UpdateSnapshotSchedule structure.
type UpdateSnapshotScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSnapshotScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateSnapshotScheduleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateSnapshotScheduleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSnapshotScheduleNoContent creates a UpdateSnapshotScheduleNoContent with default headers values
func NewUpdateSnapshotScheduleNoContent() *UpdateSnapshotScheduleNoContent {
	return &UpdateSnapshotScheduleNoContent{}
}

/*UpdateSnapshotScheduleNoContent handles this case with default header values.

Success.
*/
type UpdateSnapshotScheduleNoContent struct {
}

func (o *UpdateSnapshotScheduleNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/3/snapshot/schedules/{SnapshotScheduleId}][%d] updateSnapshotScheduleNoContent ", 204)
}

func (o *UpdateSnapshotScheduleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSnapshotScheduleDefault creates a UpdateSnapshotScheduleDefault with default headers values
func NewUpdateSnapshotScheduleDefault(code int) *UpdateSnapshotScheduleDefault {
	return &UpdateSnapshotScheduleDefault{
		_statusCode: code,
	}
}

/*UpdateSnapshotScheduleDefault handles this case with default header values.

Unexpected error
*/
type UpdateSnapshotScheduleDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update snapshot schedule default response
func (o *UpdateSnapshotScheduleDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSnapshotScheduleDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/3/snapshot/schedules/{SnapshotScheduleId}][%d] updateSnapshotSchedule default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSnapshotScheduleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}