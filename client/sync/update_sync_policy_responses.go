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

// UpdateSyncPolicyReader is a Reader for the UpdateSyncPolicy structure.
type UpdateSyncPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSyncPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateSyncPolicyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateSyncPolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSyncPolicyNoContent creates a UpdateSyncPolicyNoContent with default headers values
func NewUpdateSyncPolicyNoContent() *UpdateSyncPolicyNoContent {
	return &UpdateSyncPolicyNoContent{}
}

/*UpdateSyncPolicyNoContent handles this case with default header values.

Success.
*/
type UpdateSyncPolicyNoContent struct {
}

func (o *UpdateSyncPolicyNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/3/sync/policies/{SyncPolicyId}][%d] updateSyncPolicyNoContent ", 204)
}

func (o *UpdateSyncPolicyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSyncPolicyDefault creates a UpdateSyncPolicyDefault with default headers values
func NewUpdateSyncPolicyDefault(code int) *UpdateSyncPolicyDefault {
	return &UpdateSyncPolicyDefault{
		_statusCode: code,
	}
}

/*UpdateSyncPolicyDefault handles this case with default header values.

Unexpected error
*/
type UpdateSyncPolicyDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update sync policy default response
func (o *UpdateSyncPolicyDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSyncPolicyDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/3/sync/policies/{SyncPolicyId}][%d] updateSyncPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSyncPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
