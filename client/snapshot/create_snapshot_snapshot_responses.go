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

// CreateSnapshotSnapshotReader is a Reader for the CreateSnapshotSnapshot structure.
type CreateSnapshotSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSnapshotSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateSnapshotSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateSnapshotSnapshotDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSnapshotSnapshotOK creates a CreateSnapshotSnapshotOK with default headers values
func NewCreateSnapshotSnapshotOK() *CreateSnapshotSnapshotOK {
	return &CreateSnapshotSnapshotOK{}
}

/*CreateSnapshotSnapshotOK handles this case with default header values.

Create a new snapshot.
*/
type CreateSnapshotSnapshotOK struct {
	Payload *models.SnapshotSnapshotExtended
}

func (o *CreateSnapshotSnapshotOK) Error() string {
	return fmt.Sprintf("[POST /platform/1/snapshot/snapshots][%d] createSnapshotSnapshotOK  %+v", 200, o.Payload)
}

func (o *CreateSnapshotSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapshotSnapshotExtended)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSnapshotSnapshotDefault creates a CreateSnapshotSnapshotDefault with default headers values
func NewCreateSnapshotSnapshotDefault(code int) *CreateSnapshotSnapshotDefault {
	return &CreateSnapshotSnapshotDefault{
		_statusCode: code,
	}
}

/*CreateSnapshotSnapshotDefault handles this case with default header values.

Unexpected error
*/
type CreateSnapshotSnapshotDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create snapshot snapshot default response
func (o *CreateSnapshotSnapshotDefault) Code() int {
	return o._statusCode
}

func (o *CreateSnapshotSnapshotDefault) Error() string {
	return fmt.Sprintf("[POST /platform/1/snapshot/snapshots][%d] createSnapshotSnapshot default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSnapshotSnapshotDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
