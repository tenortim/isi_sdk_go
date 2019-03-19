// Code generated by go-swagger; DO NOT EDIT.

package remotesupport

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// UpdateRemotesupportConnectemcReader is a Reader for the UpdateRemotesupportConnectemc structure.
type UpdateRemotesupportConnectemcReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRemotesupportConnectemcReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateRemotesupportConnectemcNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateRemotesupportConnectemcDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateRemotesupportConnectemcNoContent creates a UpdateRemotesupportConnectemcNoContent with default headers values
func NewUpdateRemotesupportConnectemcNoContent() *UpdateRemotesupportConnectemcNoContent {
	return &UpdateRemotesupportConnectemcNoContent{}
}

/*UpdateRemotesupportConnectemcNoContent handles this case with default header values.

Success.
*/
type UpdateRemotesupportConnectemcNoContent struct {
}

func (o *UpdateRemotesupportConnectemcNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/1/remotesupport/connectemc][%d] updateRemotesupportConnectemcNoContent ", 204)
}

func (o *UpdateRemotesupportConnectemcNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRemotesupportConnectemcDefault creates a UpdateRemotesupportConnectemcDefault with default headers values
func NewUpdateRemotesupportConnectemcDefault(code int) *UpdateRemotesupportConnectemcDefault {
	return &UpdateRemotesupportConnectemcDefault{
		_statusCode: code,
	}
}

/*UpdateRemotesupportConnectemcDefault handles this case with default header values.

Unexpected error
*/
type UpdateRemotesupportConnectemcDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update remotesupport connectemc default response
func (o *UpdateRemotesupportConnectemcDefault) Code() int {
	return o._statusCode
}

func (o *UpdateRemotesupportConnectemcDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/1/remotesupport/connectemc][%d] updateRemotesupportConnectemc default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateRemotesupportConnectemcDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}