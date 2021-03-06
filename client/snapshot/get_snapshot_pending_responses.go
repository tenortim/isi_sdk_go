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

// GetSnapshotPendingReader is a Reader for the GetSnapshotPending structure.
type GetSnapshotPendingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSnapshotPendingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSnapshotPendingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetSnapshotPendingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSnapshotPendingOK creates a GetSnapshotPendingOK with default headers values
func NewGetSnapshotPendingOK() *GetSnapshotPendingOK {
	return &GetSnapshotPendingOK{}
}

/*GetSnapshotPendingOK handles this case with default header values.

Return list of snapshots to be taken.
*/
type GetSnapshotPendingOK struct {
	Payload *models.SnapshotPending
}

func (o *GetSnapshotPendingOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/snapshot/pending][%d] getSnapshotPendingOK  %+v", 200, o.Payload)
}

func (o *GetSnapshotPendingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapshotPending)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSnapshotPendingDefault creates a GetSnapshotPendingDefault with default headers values
func NewGetSnapshotPendingDefault(code int) *GetSnapshotPendingDefault {
	return &GetSnapshotPendingDefault{
		_statusCode: code,
	}
}

/*GetSnapshotPendingDefault handles this case with default header values.

Unexpected error
*/
type GetSnapshotPendingDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get snapshot pending default response
func (o *GetSnapshotPendingDefault) Code() int {
	return o._statusCode
}

func (o *GetSnapshotPendingDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/snapshot/pending][%d] getSnapshotPending default  %+v", o._statusCode, o.Payload)
}

func (o *GetSnapshotPendingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
