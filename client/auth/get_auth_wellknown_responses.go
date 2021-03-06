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

// GetAuthWellknownReader is a Reader for the GetAuthWellknown structure.
type GetAuthWellknownReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthWellknownReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAuthWellknownOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetAuthWellknownDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAuthWellknownOK creates a GetAuthWellknownOK with default headers values
func NewGetAuthWellknownOK() *GetAuthWellknownOK {
	return &GetAuthWellknownOK{}
}

/*GetAuthWellknownOK handles this case with default header values.

Retrieve the wellknown persona.
*/
type GetAuthWellknownOK struct {
	Payload *models.AuthWellknowns
}

func (o *GetAuthWellknownOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/auth/wellknowns/{AuthWellknownId}][%d] getAuthWellknownOK  %+v", 200, o.Payload)
}

func (o *GetAuthWellknownOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthWellknowns)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthWellknownDefault creates a GetAuthWellknownDefault with default headers values
func NewGetAuthWellknownDefault(code int) *GetAuthWellknownDefault {
	return &GetAuthWellknownDefault{
		_statusCode: code,
	}
}

/*GetAuthWellknownDefault handles this case with default header values.

Unexpected error
*/
type GetAuthWellknownDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get auth wellknown default response
func (o *GetAuthWellknownDefault) Code() int {
	return o._statusCode
}

func (o *GetAuthWellknownDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/auth/wellknowns/{AuthWellknownId}][%d] getAuthWellknown default  %+v", o._statusCode, o.Payload)
}

func (o *GetAuthWellknownDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
