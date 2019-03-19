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

// UpdateNetworkGroupnetReader is a Reader for the UpdateNetworkGroupnet structure.
type UpdateNetworkGroupnetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkGroupnetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateNetworkGroupnetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateNetworkGroupnetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateNetworkGroupnetNoContent creates a UpdateNetworkGroupnetNoContent with default headers values
func NewUpdateNetworkGroupnetNoContent() *UpdateNetworkGroupnetNoContent {
	return &UpdateNetworkGroupnetNoContent{}
}

/*UpdateNetworkGroupnetNoContent handles this case with default header values.

Success.
*/
type UpdateNetworkGroupnetNoContent struct {
}

func (o *UpdateNetworkGroupnetNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/3/network/groupnets/{NetworkGroupnetId}][%d] updateNetworkGroupnetNoContent ", 204)
}

func (o *UpdateNetworkGroupnetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateNetworkGroupnetDefault creates a UpdateNetworkGroupnetDefault with default headers values
func NewUpdateNetworkGroupnetDefault(code int) *UpdateNetworkGroupnetDefault {
	return &UpdateNetworkGroupnetDefault{
		_statusCode: code,
	}
}

/*UpdateNetworkGroupnetDefault handles this case with default header values.

Unexpected error
*/
type UpdateNetworkGroupnetDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update network groupnet default response
func (o *UpdateNetworkGroupnetDefault) Code() int {
	return o._statusCode
}

func (o *UpdateNetworkGroupnetDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/3/network/groupnets/{NetworkGroupnetId}][%d] updateNetworkGroupnet default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateNetworkGroupnetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
