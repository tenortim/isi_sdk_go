// Code generated by go-swagger; DO NOT EDIT.

package sync

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// CreateSyncJobReader is a Reader for the CreateSyncJob structure.
type CreateSyncJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSyncJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateSyncJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateSyncJobDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSyncJobOK creates a CreateSyncJobOK with default headers values
func NewCreateSyncJobOK() *CreateSyncJobOK {
	return &CreateSyncJobOK{}
}

/*CreateSyncJobOK handles this case with default header values.

Start a SyncIQ job.
*/
type CreateSyncJobOK struct {
	Payload *models.CreateResponse
}

func (o *CreateSyncJobOK) Error() string {
	return fmt.Sprintf("[POST /platform/3/sync/jobs][%d] createSyncJobOK  %+v", 200, o.Payload)
}

func (o *CreateSyncJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSyncJobDefault creates a CreateSyncJobDefault with default headers values
func NewCreateSyncJobDefault(code int) *CreateSyncJobDefault {
	return &CreateSyncJobDefault{
		_statusCode: code,
	}
}

/*CreateSyncJobDefault handles this case with default header values.

Unexpected error
*/
type CreateSyncJobDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create sync job default response
func (o *CreateSyncJobDefault) Code() int {
	return o._statusCode
}

func (o *CreateSyncJobDefault) Error() string {
	return fmt.Sprintf("[POST /platform/3/sync/jobs][%d] createSyncJob default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSyncJobDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}