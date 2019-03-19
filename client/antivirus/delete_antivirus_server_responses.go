// Code generated by go-swagger; DO NOT EDIT.

package antivirus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// DeleteAntivirusServerReader is a Reader for the DeleteAntivirusServer structure.
type DeleteAntivirusServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAntivirusServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteAntivirusServerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteAntivirusServerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAntivirusServerNoContent creates a DeleteAntivirusServerNoContent with default headers values
func NewDeleteAntivirusServerNoContent() *DeleteAntivirusServerNoContent {
	return &DeleteAntivirusServerNoContent{}
}

/*DeleteAntivirusServerNoContent handles this case with default header values.

Success.
*/
type DeleteAntivirusServerNoContent struct {
}

func (o *DeleteAntivirusServerNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/antivirus/servers/{AntivirusServerId}][%d] deleteAntivirusServerNoContent ", 204)
}

func (o *DeleteAntivirusServerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAntivirusServerDefault creates a DeleteAntivirusServerDefault with default headers values
func NewDeleteAntivirusServerDefault(code int) *DeleteAntivirusServerDefault {
	return &DeleteAntivirusServerDefault{
		_statusCode: code,
	}
}

/*DeleteAntivirusServerDefault handles this case with default header values.

Unexpected error
*/
type DeleteAntivirusServerDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete antivirus server default response
func (o *DeleteAntivirusServerDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAntivirusServerDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/antivirus/servers/{AntivirusServerId}][%d] deleteAntivirusServer default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAntivirusServerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
