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

// ListSnapshotRepstatesReader is a Reader for the ListSnapshotRepstates structure.
type ListSnapshotRepstatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSnapshotRepstatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListSnapshotRepstatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListSnapshotRepstatesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSnapshotRepstatesOK creates a ListSnapshotRepstatesOK with default headers values
func NewListSnapshotRepstatesOK() *ListSnapshotRepstatesOK {
	return &ListSnapshotRepstatesOK{}
}

/*ListSnapshotRepstatesOK handles this case with default header values.

List all repstates.
*/
type ListSnapshotRepstatesOK struct {
	Payload *models.SnapshotRepstatesExtended
}

func (o *ListSnapshotRepstatesOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/snapshot/repstates][%d] listSnapshotRepstatesOK  %+v", 200, o.Payload)
}

func (o *ListSnapshotRepstatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapshotRepstatesExtended)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSnapshotRepstatesDefault creates a ListSnapshotRepstatesDefault with default headers values
func NewListSnapshotRepstatesDefault(code int) *ListSnapshotRepstatesDefault {
	return &ListSnapshotRepstatesDefault{
		_statusCode: code,
	}
}

/*ListSnapshotRepstatesDefault handles this case with default header values.

Unexpected error
*/
type ListSnapshotRepstatesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list snapshot repstates default response
func (o *ListSnapshotRepstatesDefault) Code() int {
	return o._statusCode
}

func (o *ListSnapshotRepstatesDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/snapshot/repstates][%d] listSnapshotRepstates default  %+v", o._statusCode, o.Payload)
}

func (o *ListSnapshotRepstatesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
