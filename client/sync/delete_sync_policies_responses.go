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

// DeleteSyncPoliciesReader is a Reader for the DeleteSyncPolicies structure.
type DeleteSyncPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSyncPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSyncPoliciesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteSyncPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSyncPoliciesNoContent creates a DeleteSyncPoliciesNoContent with default headers values
func NewDeleteSyncPoliciesNoContent() *DeleteSyncPoliciesNoContent {
	return &DeleteSyncPoliciesNoContent{}
}

/*DeleteSyncPoliciesNoContent handles this case with default header values.

Success.
*/
type DeleteSyncPoliciesNoContent struct {
}

func (o *DeleteSyncPoliciesNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/sync/policies][%d] deleteSyncPoliciesNoContent ", 204)
}

func (o *DeleteSyncPoliciesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSyncPoliciesDefault creates a DeleteSyncPoliciesDefault with default headers values
func NewDeleteSyncPoliciesDefault(code int) *DeleteSyncPoliciesDefault {
	return &DeleteSyncPoliciesDefault{
		_statusCode: code,
	}
}

/*DeleteSyncPoliciesDefault handles this case with default header values.

Unexpected error
*/
type DeleteSyncPoliciesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete sync policies default response
func (o *DeleteSyncPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSyncPoliciesDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/sync/policies][%d] deleteSyncPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSyncPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
