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

// DeleteSnapshotSchedulesReader is a Reader for the DeleteSnapshotSchedules structure.
type DeleteSnapshotSchedulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSnapshotSchedulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSnapshotSchedulesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteSnapshotSchedulesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSnapshotSchedulesNoContent creates a DeleteSnapshotSchedulesNoContent with default headers values
func NewDeleteSnapshotSchedulesNoContent() *DeleteSnapshotSchedulesNoContent {
	return &DeleteSnapshotSchedulesNoContent{}
}

/*DeleteSnapshotSchedulesNoContent handles this case with default header values.

Success.
*/
type DeleteSnapshotSchedulesNoContent struct {
}

func (o *DeleteSnapshotSchedulesNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/snapshot/schedules][%d] deleteSnapshotSchedulesNoContent ", 204)
}

func (o *DeleteSnapshotSchedulesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSnapshotSchedulesDefault creates a DeleteSnapshotSchedulesDefault with default headers values
func NewDeleteSnapshotSchedulesDefault(code int) *DeleteSnapshotSchedulesDefault {
	return &DeleteSnapshotSchedulesDefault{
		_statusCode: code,
	}
}

/*DeleteSnapshotSchedulesDefault handles this case with default header values.

Unexpected error
*/
type DeleteSnapshotSchedulesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete snapshot schedules default response
func (o *DeleteSnapshotSchedulesDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSnapshotSchedulesDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/snapshot/schedules][%d] deleteSnapshotSchedules default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSnapshotSchedulesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
