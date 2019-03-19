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

// GetEventChannelReader is a Reader for the GetEventChannel structure.
type GetEventChannelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventChannelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEventChannelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetEventChannelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEventChannelOK creates a GetEventChannelOK with default headers values
func NewGetEventChannelOK() *GetEventChannelOK {
	return &GetEventChannelOK{}
}

/*GetEventChannelOK handles this case with default header values.

Retrieve the alert-condition.
*/
type GetEventChannelOK struct {
	Payload *models.EventChannels
}

func (o *GetEventChannelOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/event/channels/{EventChannelId}][%d] getEventChannelOK  %+v", 200, o.Payload)
}

func (o *GetEventChannelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EventChannels)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventChannelDefault creates a GetEventChannelDefault with default headers values
func NewGetEventChannelDefault(code int) *GetEventChannelDefault {
	return &GetEventChannelDefault{
		_statusCode: code,
	}
}

/*GetEventChannelDefault handles this case with default header values.

Unexpected error
*/
type GetEventChannelDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get event channel default response
func (o *GetEventChannelDefault) Code() int {
	return o._statusCode
}

func (o *GetEventChannelDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/event/channels/{EventChannelId}][%d] getEventChannel default  %+v", o._statusCode, o.Payload)
}

func (o *GetEventChannelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
