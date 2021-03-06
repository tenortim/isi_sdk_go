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

// DeleteSmbSessionsComputerUserReader is a Reader for the DeleteSmbSessionsComputerUser structure.
type DeleteSmbSessionsComputerUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSmbSessionsComputerUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSmbSessionsComputerUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteSmbSessionsComputerUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSmbSessionsComputerUserNoContent creates a DeleteSmbSessionsComputerUserNoContent with default headers values
func NewDeleteSmbSessionsComputerUserNoContent() *DeleteSmbSessionsComputerUserNoContent {
	return &DeleteSmbSessionsComputerUserNoContent{}
}

/*DeleteSmbSessionsComputerUserNoContent handles this case with default header values.

Success.
*/
type DeleteSmbSessionsComputerUserNoContent struct {
}

func (o *DeleteSmbSessionsComputerUserNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/protocols/smb/sessions/{Computer}/{SmbSessionsComputerUser}][%d] deleteSmbSessionsComputerUserNoContent ", 204)
}

func (o *DeleteSmbSessionsComputerUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSmbSessionsComputerUserDefault creates a DeleteSmbSessionsComputerUserDefault with default headers values
func NewDeleteSmbSessionsComputerUserDefault(code int) *DeleteSmbSessionsComputerUserDefault {
	return &DeleteSmbSessionsComputerUserDefault{
		_statusCode: code,
	}
}

/*DeleteSmbSessionsComputerUserDefault handles this case with default header values.

Unexpected error
*/
type DeleteSmbSessionsComputerUserDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete smb sessions computer user default response
func (o *DeleteSmbSessionsComputerUserDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSmbSessionsComputerUserDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/protocols/smb/sessions/{Computer}/{SmbSessionsComputerUser}][%d] deleteSmbSessionsComputerUser default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSmbSessionsComputerUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
