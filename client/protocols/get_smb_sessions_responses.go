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

// GetSmbSessionsReader is a Reader for the GetSmbSessions structure.
type GetSmbSessionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSmbSessionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSmbSessionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetSmbSessionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSmbSessionsOK creates a GetSmbSessionsOK with default headers values
func NewGetSmbSessionsOK() *GetSmbSessionsOK {
	return &GetSmbSessionsOK{}
}

/*GetSmbSessionsOK handles this case with default header values.

List open sessions.
*/
type GetSmbSessionsOK struct {
	Payload *models.SmbSessions
}

func (o *GetSmbSessionsOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/protocols/smb/sessions][%d] getSmbSessionsOK  %+v", 200, o.Payload)
}

func (o *GetSmbSessionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SmbSessions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSmbSessionsDefault creates a GetSmbSessionsDefault with default headers values
func NewGetSmbSessionsDefault(code int) *GetSmbSessionsDefault {
	return &GetSmbSessionsDefault{
		_statusCode: code,
	}
}

/*GetSmbSessionsDefault handles this case with default header values.

Unexpected error
*/
type GetSmbSessionsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get smb sessions default response
func (o *GetSmbSessionsDefault) Code() int {
	return o._statusCode
}

func (o *GetSmbSessionsDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/protocols/smb/sessions][%d] getSmbSessions default  %+v", o._statusCode, o.Payload)
}

func (o *GetSmbSessionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
