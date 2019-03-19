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

// ListEventChannelsReader is a Reader for the ListEventChannels structure.
type ListEventChannelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEventChannelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListEventChannelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListEventChannelsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListEventChannelsOK creates a ListEventChannelsOK with default headers values
func NewListEventChannelsOK() *ListEventChannelsOK {
	return &ListEventChannelsOK{}
}

/*ListEventChannelsOK handles this case with default header values.

List all channels.
*/
type ListEventChannelsOK struct {
	Payload *models.EventChannelsExtended
}

func (o *ListEventChannelsOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/event/channels][%d] listEventChannelsOK  %+v", 200, o.Payload)
}

func (o *ListEventChannelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EventChannelsExtended)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEventChannelsDefault creates a ListEventChannelsDefault with default headers values
func NewListEventChannelsDefault(code int) *ListEventChannelsDefault {
	return &ListEventChannelsDefault{
		_statusCode: code,
	}
}

/*ListEventChannelsDefault handles this case with default header values.

Unexpected error
*/
type ListEventChannelsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list event channels default response
func (o *ListEventChannelsDefault) Code() int {
	return o._statusCode
}

func (o *ListEventChannelsDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/event/channels][%d] listEventChannels default  %+v", o._statusCode, o.Payload)
}

func (o *ListEventChannelsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}