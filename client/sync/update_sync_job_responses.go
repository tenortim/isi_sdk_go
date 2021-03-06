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

// UpdateSyncJobReader is a Reader for the UpdateSyncJob structure.
type UpdateSyncJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSyncJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateSyncJobNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateSyncJobDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSyncJobNoContent creates a UpdateSyncJobNoContent with default headers values
func NewUpdateSyncJobNoContent() *UpdateSyncJobNoContent {
	return &UpdateSyncJobNoContent{}
}

/*UpdateSyncJobNoContent handles this case with default header values.

Success.
*/
type UpdateSyncJobNoContent struct {
}

func (o *UpdateSyncJobNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/3/sync/jobs/{SyncJobId}][%d] updateSyncJobNoContent ", 204)
}

func (o *UpdateSyncJobNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSyncJobDefault creates a UpdateSyncJobDefault with default headers values
func NewUpdateSyncJobDefault(code int) *UpdateSyncJobDefault {
	return &UpdateSyncJobDefault{
		_statusCode: code,
	}
}

/*UpdateSyncJobDefault handles this case with default header values.

Unexpected error
*/
type UpdateSyncJobDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update sync job default response
func (o *UpdateSyncJobDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSyncJobDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/3/sync/jobs/{SyncJobId}][%d] updateSyncJob default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSyncJobDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
