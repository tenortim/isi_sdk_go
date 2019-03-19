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

// DeleteSyncRuleReader is a Reader for the DeleteSyncRule structure.
type DeleteSyncRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSyncRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSyncRuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteSyncRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSyncRuleNoContent creates a DeleteSyncRuleNoContent with default headers values
func NewDeleteSyncRuleNoContent() *DeleteSyncRuleNoContent {
	return &DeleteSyncRuleNoContent{}
}

/*DeleteSyncRuleNoContent handles this case with default header values.

Success.
*/
type DeleteSyncRuleNoContent struct {
}

func (o *DeleteSyncRuleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/sync/rules/{SyncRuleId}][%d] deleteSyncRuleNoContent ", 204)
}

func (o *DeleteSyncRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSyncRuleDefault creates a DeleteSyncRuleDefault with default headers values
func NewDeleteSyncRuleDefault(code int) *DeleteSyncRuleDefault {
	return &DeleteSyncRuleDefault{
		_statusCode: code,
	}
}

/*DeleteSyncRuleDefault handles this case with default header values.

Unexpected error
*/
type DeleteSyncRuleDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete sync rule default response
func (o *DeleteSyncRuleDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSyncRuleDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/sync/rules/{SyncRuleId}][%d] deleteSyncRule default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSyncRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
