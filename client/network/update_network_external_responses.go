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

// UpdateNetworkExternalReader is a Reader for the UpdateNetworkExternal structure.
type UpdateNetworkExternalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkExternalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateNetworkExternalNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateNetworkExternalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateNetworkExternalNoContent creates a UpdateNetworkExternalNoContent with default headers values
func NewUpdateNetworkExternalNoContent() *UpdateNetworkExternalNoContent {
	return &UpdateNetworkExternalNoContent{}
}

/*UpdateNetworkExternalNoContent handles this case with default header values.

Success.
*/
type UpdateNetworkExternalNoContent struct {
}

func (o *UpdateNetworkExternalNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/3/network/external][%d] updateNetworkExternalNoContent ", 204)
}

func (o *UpdateNetworkExternalNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateNetworkExternalDefault creates a UpdateNetworkExternalDefault with default headers values
func NewUpdateNetworkExternalDefault(code int) *UpdateNetworkExternalDefault {
	return &UpdateNetworkExternalDefault{
		_statusCode: code,
	}
}

/*UpdateNetworkExternalDefault handles this case with default header values.

Unexpected error
*/
type UpdateNetworkExternalDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update network external default response
func (o *UpdateNetworkExternalDefault) Code() int {
	return o._statusCode
}

func (o *UpdateNetworkExternalDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/3/network/external][%d] updateNetworkExternal default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateNetworkExternalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}