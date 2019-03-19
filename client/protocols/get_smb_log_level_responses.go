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

// GetSmbLogLevelReader is a Reader for the GetSmbLogLevel structure.
type GetSmbLogLevelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSmbLogLevelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSmbLogLevelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetSmbLogLevelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSmbLogLevelOK creates a GetSmbLogLevelOK with default headers values
func NewGetSmbLogLevelOK() *GetSmbLogLevelOK {
	return &GetSmbLogLevelOK{}
}

/*GetSmbLogLevelOK handles this case with default header values.

Get the current SMB logging level.
*/
type GetSmbLogLevelOK struct {
	Payload *models.SmbLogLevel
}

func (o *GetSmbLogLevelOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/protocols/smb/log-level][%d] getSmbLogLevelOK  %+v", 200, o.Payload)
}

func (o *GetSmbLogLevelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SmbLogLevel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSmbLogLevelDefault creates a GetSmbLogLevelDefault with default headers values
func NewGetSmbLogLevelDefault(code int) *GetSmbLogLevelDefault {
	return &GetSmbLogLevelDefault{
		_statusCode: code,
	}
}

/*GetSmbLogLevelDefault handles this case with default header values.

Unexpected error
*/
type GetSmbLogLevelDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get smb log level default response
func (o *GetSmbLogLevelDefault) Code() int {
	return o._statusCode
}

func (o *GetSmbLogLevelDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/protocols/smb/log-level][%d] getSmbLogLevel default  %+v", o._statusCode, o.Payload)
}

func (o *GetSmbLogLevelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
