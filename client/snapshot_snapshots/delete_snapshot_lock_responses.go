// Code generated by go-swagger; DO NOT EDIT.

package snapshot_snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// DeleteSnapshotLockReader is a Reader for the DeleteSnapshotLock structure.
type DeleteSnapshotLockReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSnapshotLockReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSnapshotLockNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteSnapshotLockDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSnapshotLockNoContent creates a DeleteSnapshotLockNoContent with default headers values
func NewDeleteSnapshotLockNoContent() *DeleteSnapshotLockNoContent {
	return &DeleteSnapshotLockNoContent{}
}

/*DeleteSnapshotLockNoContent handles this case with default header values.

Success.
*/
type DeleteSnapshotLockNoContent struct {
}

func (o *DeleteSnapshotLockNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/snapshot/snapshots/{Sid}/locks/{SnapshotLockId}][%d] deleteSnapshotLockNoContent ", 204)
}

func (o *DeleteSnapshotLockNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSnapshotLockDefault creates a DeleteSnapshotLockDefault with default headers values
func NewDeleteSnapshotLockDefault(code int) *DeleteSnapshotLockDefault {
	return &DeleteSnapshotLockDefault{
		_statusCode: code,
	}
}

/*DeleteSnapshotLockDefault handles this case with default header values.

Unexpected error
*/
type DeleteSnapshotLockDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete snapshot lock default response
func (o *DeleteSnapshotLockDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSnapshotLockDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/snapshot/snapshots/{Sid}/locks/{SnapshotLockId}][%d] deleteSnapshotLock default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSnapshotLockDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
