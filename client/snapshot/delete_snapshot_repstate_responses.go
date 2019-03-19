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

// DeleteSnapshotRepstateReader is a Reader for the DeleteSnapshotRepstate structure.
type DeleteSnapshotRepstateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSnapshotRepstateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSnapshotRepstateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteSnapshotRepstateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSnapshotRepstateNoContent creates a DeleteSnapshotRepstateNoContent with default headers values
func NewDeleteSnapshotRepstateNoContent() *DeleteSnapshotRepstateNoContent {
	return &DeleteSnapshotRepstateNoContent{}
}

/*DeleteSnapshotRepstateNoContent handles this case with default header values.

Success.
*/
type DeleteSnapshotRepstateNoContent struct {
}

func (o *DeleteSnapshotRepstateNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/snapshot/repstates/{SnapshotRepstateId}][%d] deleteSnapshotRepstateNoContent ", 204)
}

func (o *DeleteSnapshotRepstateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSnapshotRepstateDefault creates a DeleteSnapshotRepstateDefault with default headers values
func NewDeleteSnapshotRepstateDefault(code int) *DeleteSnapshotRepstateDefault {
	return &DeleteSnapshotRepstateDefault{
		_statusCode: code,
	}
}

/*DeleteSnapshotRepstateDefault handles this case with default header values.

Unexpected error
*/
type DeleteSnapshotRepstateDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete snapshot repstate default response
func (o *DeleteSnapshotRepstateDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSnapshotRepstateDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/snapshot/repstates/{SnapshotRepstateId}][%d] deleteSnapshotRepstate default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSnapshotRepstateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}