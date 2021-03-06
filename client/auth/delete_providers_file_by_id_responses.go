// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// DeleteProvidersFileByIDReader is a Reader for the DeleteProvidersFileByID structure.
type DeleteProvidersFileByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProvidersFileByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteProvidersFileByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteProvidersFileByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteProvidersFileByIDNoContent creates a DeleteProvidersFileByIDNoContent with default headers values
func NewDeleteProvidersFileByIDNoContent() *DeleteProvidersFileByIDNoContent {
	return &DeleteProvidersFileByIDNoContent{}
}

/*DeleteProvidersFileByIDNoContent handles this case with default header values.

Success.
*/
type DeleteProvidersFileByIDNoContent struct {
}

func (o *DeleteProvidersFileByIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/auth/providers/file/{ProvidersFileId}][%d] deleteProvidersFileByIdNoContent ", 204)
}

func (o *DeleteProvidersFileByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProvidersFileByIDDefault creates a DeleteProvidersFileByIDDefault with default headers values
func NewDeleteProvidersFileByIDDefault(code int) *DeleteProvidersFileByIDDefault {
	return &DeleteProvidersFileByIDDefault{
		_statusCode: code,
	}
}

/*DeleteProvidersFileByIDDefault handles this case with default header values.

Unexpected error
*/
type DeleteProvidersFileByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete providers file by Id default response
func (o *DeleteProvidersFileByIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteProvidersFileByIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/auth/providers/file/{ProvidersFileId}][%d] deleteProvidersFileById default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteProvidersFileByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
