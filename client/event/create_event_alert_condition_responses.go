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

// CreateEventAlertConditionReader is a Reader for the CreateEventAlertCondition structure.
type CreateEventAlertConditionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateEventAlertConditionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateEventAlertConditionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateEventAlertConditionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateEventAlertConditionOK creates a CreateEventAlertConditionOK with default headers values
func NewCreateEventAlertConditionOK() *CreateEventAlertConditionOK {
	return &CreateEventAlertConditionOK{}
}

/*CreateEventAlertConditionOK handles this case with default header values.

Create a new alert condition.
*/
type CreateEventAlertConditionOK struct {
	Payload *models.CreateResponse
}

func (o *CreateEventAlertConditionOK) Error() string {
	return fmt.Sprintf("[POST /platform/3/event/alert-conditions][%d] createEventAlertConditionOK  %+v", 200, o.Payload)
}

func (o *CreateEventAlertConditionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEventAlertConditionDefault creates a CreateEventAlertConditionDefault with default headers values
func NewCreateEventAlertConditionDefault(code int) *CreateEventAlertConditionDefault {
	return &CreateEventAlertConditionDefault{
		_statusCode: code,
	}
}

/*CreateEventAlertConditionDefault handles this case with default header values.

Unexpected error
*/
type CreateEventAlertConditionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create event alert condition default response
func (o *CreateEventAlertConditionDefault) Code() int {
	return o._statusCode
}

func (o *CreateEventAlertConditionDefault) Error() string {
	return fmt.Sprintf("[POST /platform/3/event/alert-conditions][%d] createEventAlertCondition default  %+v", o._statusCode, o.Payload)
}

func (o *CreateEventAlertConditionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
