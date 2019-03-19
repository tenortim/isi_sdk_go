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

// DeleteNdmpSettingsVariableReader is a Reader for the DeleteNdmpSettingsVariable structure.
type DeleteNdmpSettingsVariableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNdmpSettingsVariableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteNdmpSettingsVariableNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteNdmpSettingsVariableDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNdmpSettingsVariableNoContent creates a DeleteNdmpSettingsVariableNoContent with default headers values
func NewDeleteNdmpSettingsVariableNoContent() *DeleteNdmpSettingsVariableNoContent {
	return &DeleteNdmpSettingsVariableNoContent{}
}

/*DeleteNdmpSettingsVariableNoContent handles this case with default header values.

Success.
*/
type DeleteNdmpSettingsVariableNoContent struct {
}

func (o *DeleteNdmpSettingsVariableNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/protocols/ndmp/settings/variables/{NdmpSettingsVariableId}][%d] deleteNdmpSettingsVariableNoContent ", 204)
}

func (o *DeleteNdmpSettingsVariableNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNdmpSettingsVariableDefault creates a DeleteNdmpSettingsVariableDefault with default headers values
func NewDeleteNdmpSettingsVariableDefault(code int) *DeleteNdmpSettingsVariableDefault {
	return &DeleteNdmpSettingsVariableDefault{
		_statusCode: code,
	}
}

/*DeleteNdmpSettingsVariableDefault handles this case with default header values.

Unexpected error
*/
type DeleteNdmpSettingsVariableDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete ndmp settings variable default response
func (o *DeleteNdmpSettingsVariableDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNdmpSettingsVariableDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/protocols/ndmp/settings/variables/{NdmpSettingsVariableId}][%d] deleteNdmpSettingsVariable default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteNdmpSettingsVariableDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}