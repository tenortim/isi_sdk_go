// Code generated by go-swagger; DO NOT EDIT.

package zones

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// DeleteZoneReader is a Reader for the DeleteZone structure.
type DeleteZoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteZoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteZoneNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteZoneDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteZoneNoContent creates a DeleteZoneNoContent with default headers values
func NewDeleteZoneNoContent() *DeleteZoneNoContent {
	return &DeleteZoneNoContent{}
}

/*DeleteZoneNoContent handles this case with default header values.

Success.
*/
type DeleteZoneNoContent struct {
}

func (o *DeleteZoneNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/zones/{ZoneId}][%d] deleteZoneNoContent ", 204)
}

func (o *DeleteZoneNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteZoneDefault creates a DeleteZoneDefault with default headers values
func NewDeleteZoneDefault(code int) *DeleteZoneDefault {
	return &DeleteZoneDefault{
		_statusCode: code,
	}
}

/*DeleteZoneDefault handles this case with default header values.

Unexpected error
*/
type DeleteZoneDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete zone default response
func (o *DeleteZoneDefault) Code() int {
	return o._statusCode
}

func (o *DeleteZoneDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/zones/{ZoneId}][%d] deleteZone default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteZoneDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}