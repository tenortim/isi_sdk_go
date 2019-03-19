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

// DeleteSnapshotSnapshotReader is a Reader for the DeleteSnapshotSnapshot structure.
type DeleteSnapshotSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSnapshotSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSnapshotSnapshotNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteSnapshotSnapshotDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSnapshotSnapshotNoContent creates a DeleteSnapshotSnapshotNoContent with default headers values
func NewDeleteSnapshotSnapshotNoContent() *DeleteSnapshotSnapshotNoContent {
	return &DeleteSnapshotSnapshotNoContent{}
}

/*DeleteSnapshotSnapshotNoContent handles this case with default header values.

Success.
*/
type DeleteSnapshotSnapshotNoContent struct {
}

func (o *DeleteSnapshotSnapshotNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/snapshot/snapshots/{SnapshotSnapshotId}][%d] deleteSnapshotSnapshotNoContent ", 204)
}

func (o *DeleteSnapshotSnapshotNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSnapshotSnapshotDefault creates a DeleteSnapshotSnapshotDefault with default headers values
func NewDeleteSnapshotSnapshotDefault(code int) *DeleteSnapshotSnapshotDefault {
	return &DeleteSnapshotSnapshotDefault{
		_statusCode: code,
	}
}

/*DeleteSnapshotSnapshotDefault handles this case with default header values.

Unexpected error
*/
type DeleteSnapshotSnapshotDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete snapshot snapshot default response
func (o *DeleteSnapshotSnapshotDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSnapshotSnapshotDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/snapshot/snapshots/{SnapshotSnapshotId}][%d] deleteSnapshotSnapshot default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSnapshotSnapshotDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
