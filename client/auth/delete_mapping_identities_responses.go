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

// DeleteMappingIdentitiesReader is a Reader for the DeleteMappingIdentities structure.
type DeleteMappingIdentitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMappingIdentitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteMappingIdentitiesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteMappingIdentitiesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMappingIdentitiesNoContent creates a DeleteMappingIdentitiesNoContent with default headers values
func NewDeleteMappingIdentitiesNoContent() *DeleteMappingIdentitiesNoContent {
	return &DeleteMappingIdentitiesNoContent{}
}

/*DeleteMappingIdentitiesNoContent handles this case with default header values.

Success.
*/
type DeleteMappingIdentitiesNoContent struct {
}

func (o *DeleteMappingIdentitiesNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/auth/mapping/identities][%d] deleteMappingIdentitiesNoContent ", 204)
}

func (o *DeleteMappingIdentitiesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMappingIdentitiesDefault creates a DeleteMappingIdentitiesDefault with default headers values
func NewDeleteMappingIdentitiesDefault(code int) *DeleteMappingIdentitiesDefault {
	return &DeleteMappingIdentitiesDefault{
		_statusCode: code,
	}
}

/*DeleteMappingIdentitiesDefault handles this case with default header values.

Unexpected error
*/
type DeleteMappingIdentitiesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete mapping identities default response
func (o *DeleteMappingIdentitiesDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMappingIdentitiesDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/auth/mapping/identities][%d] deleteMappingIdentities default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteMappingIdentitiesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}