// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// ListNetworkGroupnetsReader is a Reader for the ListNetworkGroupnets structure.
type ListNetworkGroupnetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListNetworkGroupnetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListNetworkGroupnetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListNetworkGroupnetsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListNetworkGroupnetsOK creates a ListNetworkGroupnetsOK with default headers values
func NewListNetworkGroupnetsOK() *ListNetworkGroupnetsOK {
	return &ListNetworkGroupnetsOK{}
}

/*ListNetworkGroupnetsOK handles this case with default header values.

Get a list of groupnets.
*/
type ListNetworkGroupnetsOK struct {
	Payload *models.NetworkGroupnetsExtended
}

func (o *ListNetworkGroupnetsOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/network/groupnets][%d] listNetworkGroupnetsOK  %+v", 200, o.Payload)
}

func (o *ListNetworkGroupnetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkGroupnetsExtended)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNetworkGroupnetsDefault creates a ListNetworkGroupnetsDefault with default headers values
func NewListNetworkGroupnetsDefault(code int) *ListNetworkGroupnetsDefault {
	return &ListNetworkGroupnetsDefault{
		_statusCode: code,
	}
}

/*ListNetworkGroupnetsDefault handles this case with default header values.

Unexpected error
*/
type ListNetworkGroupnetsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list network groupnets default response
func (o *ListNetworkGroupnetsDefault) Code() int {
	return o._statusCode
}

func (o *ListNetworkGroupnetsDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/network/groupnets][%d] listNetworkGroupnets default  %+v", o._statusCode, o.Payload)
}

func (o *ListNetworkGroupnetsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
