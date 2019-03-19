// Code generated by go-swagger; DO NOT EDIT.

package protocols

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// DeleteNtpServerReader is a Reader for the DeleteNtpServer structure.
type DeleteNtpServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNtpServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteNtpServerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteNtpServerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNtpServerNoContent creates a DeleteNtpServerNoContent with default headers values
func NewDeleteNtpServerNoContent() *DeleteNtpServerNoContent {
	return &DeleteNtpServerNoContent{}
}

/*DeleteNtpServerNoContent handles this case with default header values.

Success.
*/
type DeleteNtpServerNoContent struct {
}

func (o *DeleteNtpServerNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/protocols/ntp/servers/{NtpServerId}][%d] deleteNtpServerNoContent ", 204)
}

func (o *DeleteNtpServerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNtpServerDefault creates a DeleteNtpServerDefault with default headers values
func NewDeleteNtpServerDefault(code int) *DeleteNtpServerDefault {
	return &DeleteNtpServerDefault{
		_statusCode: code,
	}
}

/*DeleteNtpServerDefault handles this case with default header values.

Unexpected error
*/
type DeleteNtpServerDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete ntp server default response
func (o *DeleteNtpServerDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNtpServerDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/protocols/ntp/servers/{NtpServerId}][%d] deleteNtpServer default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteNtpServerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}