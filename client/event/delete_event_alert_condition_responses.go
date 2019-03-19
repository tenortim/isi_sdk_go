// Code generated by go-swagger; DO NOT EDIT.

package event

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// DeleteEventAlertConditionReader is a Reader for the DeleteEventAlertCondition structure.
type DeleteEventAlertConditionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEventAlertConditionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteEventAlertConditionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteEventAlertConditionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteEventAlertConditionNoContent creates a DeleteEventAlertConditionNoContent with default headers values
func NewDeleteEventAlertConditionNoContent() *DeleteEventAlertConditionNoContent {
	return &DeleteEventAlertConditionNoContent{}
}

/*DeleteEventAlertConditionNoContent handles this case with default header values.

Success.
*/
type DeleteEventAlertConditionNoContent struct {
}

func (o *DeleteEventAlertConditionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/event/alert-conditions/{EventAlertConditionId}][%d] deleteEventAlertConditionNoContent ", 204)
}

func (o *DeleteEventAlertConditionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEventAlertConditionDefault creates a DeleteEventAlertConditionDefault with default headers values
func NewDeleteEventAlertConditionDefault(code int) *DeleteEventAlertConditionDefault {
	return &DeleteEventAlertConditionDefault{
		_statusCode: code,
	}
}

/*DeleteEventAlertConditionDefault handles this case with default header values.

Unexpected error
*/
type DeleteEventAlertConditionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete event alert condition default response
func (o *DeleteEventAlertConditionDefault) Code() int {
	return o._statusCode
}

func (o *DeleteEventAlertConditionDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/event/alert-conditions/{EventAlertConditionId}][%d] deleteEventAlertCondition default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteEventAlertConditionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}