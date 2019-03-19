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

// ListSnapshotSchedulesReader is a Reader for the ListSnapshotSchedules structure.
type ListSnapshotSchedulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSnapshotSchedulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListSnapshotSchedulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListSnapshotSchedulesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSnapshotSchedulesOK creates a ListSnapshotSchedulesOK with default headers values
func NewListSnapshotSchedulesOK() *ListSnapshotSchedulesOK {
	return &ListSnapshotSchedulesOK{}
}

/*ListSnapshotSchedulesOK handles this case with default header values.

List all or matching schedules.
*/
type ListSnapshotSchedulesOK struct {
	Payload *models.SnapshotSchedulesExtended
}

func (o *ListSnapshotSchedulesOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/snapshot/schedules][%d] listSnapshotSchedulesOK  %+v", 200, o.Payload)
}

func (o *ListSnapshotSchedulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapshotSchedulesExtended)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSnapshotSchedulesDefault creates a ListSnapshotSchedulesDefault with default headers values
func NewListSnapshotSchedulesDefault(code int) *ListSnapshotSchedulesDefault {
	return &ListSnapshotSchedulesDefault{
		_statusCode: code,
	}
}

/*ListSnapshotSchedulesDefault handles this case with default header values.

Unexpected error
*/
type ListSnapshotSchedulesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list snapshot schedules default response
func (o *ListSnapshotSchedulesDefault) Code() int {
	return o._statusCode
}

func (o *ListSnapshotSchedulesDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/snapshot/schedules][%d] listSnapshotSchedules default  %+v", o._statusCode, o.Payload)
}

func (o *ListSnapshotSchedulesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
