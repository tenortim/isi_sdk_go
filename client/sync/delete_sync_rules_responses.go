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

// DeleteSyncRulesReader is a Reader for the DeleteSyncRules structure.
type DeleteSyncRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSyncRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSyncRulesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteSyncRulesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSyncRulesNoContent creates a DeleteSyncRulesNoContent with default headers values
func NewDeleteSyncRulesNoContent() *DeleteSyncRulesNoContent {
	return &DeleteSyncRulesNoContent{}
}

/*DeleteSyncRulesNoContent handles this case with default header values.

Success.
*/
type DeleteSyncRulesNoContent struct {
}

func (o *DeleteSyncRulesNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/sync/rules][%d] deleteSyncRulesNoContent ", 204)
}

func (o *DeleteSyncRulesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSyncRulesDefault creates a DeleteSyncRulesDefault with default headers values
func NewDeleteSyncRulesDefault(code int) *DeleteSyncRulesDefault {
	return &DeleteSyncRulesDefault{
		_statusCode: code,
	}
}

/*DeleteSyncRulesDefault handles this case with default header values.

Unexpected error
*/
type DeleteSyncRulesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete sync rules default response
func (o *DeleteSyncRulesDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSyncRulesDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/sync/rules][%d] deleteSyncRules default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSyncRulesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
