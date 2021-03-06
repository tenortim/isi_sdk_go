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

// GetAuthWellknownsReader is a Reader for the GetAuthWellknowns structure.
type GetAuthWellknownsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthWellknownsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAuthWellknownsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetAuthWellknownsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAuthWellknownsOK creates a GetAuthWellknownsOK with default headers values
func NewGetAuthWellknownsOK() *GetAuthWellknownsOK {
	return &GetAuthWellknownsOK{}
}

/*GetAuthWellknownsOK handles this case with default header values.

List all wellknown personas.
*/
type GetAuthWellknownsOK struct {
	Payload *models.AuthWellknowns
}

func (o *GetAuthWellknownsOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/auth/wellknowns][%d] getAuthWellknownsOK  %+v", 200, o.Payload)
}

func (o *GetAuthWellknownsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthWellknowns)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthWellknownsDefault creates a GetAuthWellknownsDefault with default headers values
func NewGetAuthWellknownsDefault(code int) *GetAuthWellknownsDefault {
	return &GetAuthWellknownsDefault{
		_statusCode: code,
	}
}

/*GetAuthWellknownsDefault handles this case with default header values.

Unexpected error
*/
type GetAuthWellknownsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get auth wellknowns default response
func (o *GetAuthWellknownsDefault) Code() int {
	return o._statusCode
}

func (o *GetAuthWellknownsDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/auth/wellknowns][%d] getAuthWellknowns default  %+v", o._statusCode, o.Payload)
}

func (o *GetAuthWellknownsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
