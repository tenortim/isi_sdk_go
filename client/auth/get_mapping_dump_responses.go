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

// GetMappingDumpReader is a Reader for the GetMappingDump structure.
type GetMappingDumpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMappingDumpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMappingDumpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMappingDumpDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMappingDumpOK creates a GetMappingDumpOK with default headers values
func NewGetMappingDumpOK() *GetMappingDumpOK {
	return &GetMappingDumpOK{}
}

/*GetMappingDumpOK handles this case with default header values.

Retrieve all identity mappings (uid, gid, sid, and on-disk) for the supplied source persona.
*/
type GetMappingDumpOK struct {
	Payload *models.MappingDump
}

func (o *GetMappingDumpOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/auth/mapping/dump][%d] getMappingDumpOK  %+v", 200, o.Payload)
}

func (o *GetMappingDumpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MappingDump)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMappingDumpDefault creates a GetMappingDumpDefault with default headers values
func NewGetMappingDumpDefault(code int) *GetMappingDumpDefault {
	return &GetMappingDumpDefault{
		_statusCode: code,
	}
}

/*GetMappingDumpDefault handles this case with default header values.

Unexpected error
*/
type GetMappingDumpDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get mapping dump default response
func (o *GetMappingDumpDefault) Code() int {
	return o._statusCode
}

func (o *GetMappingDumpDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/auth/mapping/dump][%d] getMappingDump default  %+v", o._statusCode, o.Payload)
}

func (o *GetMappingDumpDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
