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

// DeleteNtpServersReader is a Reader for the DeleteNtpServers structure.
type DeleteNtpServersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNtpServersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteNtpServersNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteNtpServersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNtpServersNoContent creates a DeleteNtpServersNoContent with default headers values
func NewDeleteNtpServersNoContent() *DeleteNtpServersNoContent {
	return &DeleteNtpServersNoContent{}
}

/*DeleteNtpServersNoContent handles this case with default header values.

Success.
*/
type DeleteNtpServersNoContent struct {
}

func (o *DeleteNtpServersNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/protocols/ntp/servers][%d] deleteNtpServersNoContent ", 204)
}

func (o *DeleteNtpServersNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNtpServersDefault creates a DeleteNtpServersDefault with default headers values
func NewDeleteNtpServersDefault(code int) *DeleteNtpServersDefault {
	return &DeleteNtpServersDefault{
		_statusCode: code,
	}
}

/*DeleteNtpServersDefault handles this case with default header values.

Unexpected error
*/
type DeleteNtpServersDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete ntp servers default response
func (o *DeleteNtpServersDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNtpServersDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/protocols/ntp/servers][%d] deleteNtpServers default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteNtpServersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}