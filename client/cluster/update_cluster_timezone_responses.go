// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// UpdateClusterTimezoneReader is a Reader for the UpdateClusterTimezone structure.
type UpdateClusterTimezoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateClusterTimezoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateClusterTimezoneNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateClusterTimezoneDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateClusterTimezoneNoContent creates a UpdateClusterTimezoneNoContent with default headers values
func NewUpdateClusterTimezoneNoContent() *UpdateClusterTimezoneNoContent {
	return &UpdateClusterTimezoneNoContent{}
}

/*UpdateClusterTimezoneNoContent handles this case with default header values.

Success.
*/
type UpdateClusterTimezoneNoContent struct {
}

func (o *UpdateClusterTimezoneNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/3/cluster/timezone][%d] updateClusterTimezoneNoContent ", 204)
}

func (o *UpdateClusterTimezoneNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateClusterTimezoneDefault creates a UpdateClusterTimezoneDefault with default headers values
func NewUpdateClusterTimezoneDefault(code int) *UpdateClusterTimezoneDefault {
	return &UpdateClusterTimezoneDefault{
		_statusCode: code,
	}
}

/*UpdateClusterTimezoneDefault handles this case with default header values.

Unexpected error
*/
type UpdateClusterTimezoneDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update cluster timezone default response
func (o *UpdateClusterTimezoneDefault) Code() int {
	return o._statusCode
}

func (o *UpdateClusterTimezoneDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/3/cluster/timezone][%d] updateClusterTimezone default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateClusterTimezoneDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
