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

// UpdateNdmpDiagnosticsReader is a Reader for the UpdateNdmpDiagnostics structure.
type UpdateNdmpDiagnosticsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNdmpDiagnosticsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateNdmpDiagnosticsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateNdmpDiagnosticsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateNdmpDiagnosticsNoContent creates a UpdateNdmpDiagnosticsNoContent with default headers values
func NewUpdateNdmpDiagnosticsNoContent() *UpdateNdmpDiagnosticsNoContent {
	return &UpdateNdmpDiagnosticsNoContent{}
}

/*UpdateNdmpDiagnosticsNoContent handles this case with default header values.

Success.
*/
type UpdateNdmpDiagnosticsNoContent struct {
}

func (o *UpdateNdmpDiagnosticsNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/3/protocols/ndmp/diagnostics][%d] updateNdmpDiagnosticsNoContent ", 204)
}

func (o *UpdateNdmpDiagnosticsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateNdmpDiagnosticsDefault creates a UpdateNdmpDiagnosticsDefault with default headers values
func NewUpdateNdmpDiagnosticsDefault(code int) *UpdateNdmpDiagnosticsDefault {
	return &UpdateNdmpDiagnosticsDefault{
		_statusCode: code,
	}
}

/*UpdateNdmpDiagnosticsDefault handles this case with default header values.

Unexpected error
*/
type UpdateNdmpDiagnosticsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update ndmp diagnostics default response
func (o *UpdateNdmpDiagnosticsDefault) Code() int {
	return o._statusCode
}

func (o *UpdateNdmpDiagnosticsDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/3/protocols/ndmp/diagnostics][%d] updateNdmpDiagnostics default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateNdmpDiagnosticsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}