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

// GetNfsAliasReader is a Reader for the GetNfsAlias structure.
type GetNfsAliasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNfsAliasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNfsAliasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetNfsAliasDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNfsAliasOK creates a GetNfsAliasOK with default headers values
func NewGetNfsAliasOK() *GetNfsAliasOK {
	return &GetNfsAliasOK{}
}

/*GetNfsAliasOK handles this case with default header values.

Retrieve export information.
*/
type GetNfsAliasOK struct {
	Payload *models.NfsAliases
}

func (o *GetNfsAliasOK) Error() string {
	return fmt.Sprintf("[GET /platform/2/protocols/nfs/aliases/{NfsAliasId}][%d] getNfsAliasOK  %+v", 200, o.Payload)
}

func (o *GetNfsAliasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NfsAliases)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNfsAliasDefault creates a GetNfsAliasDefault with default headers values
func NewGetNfsAliasDefault(code int) *GetNfsAliasDefault {
	return &GetNfsAliasDefault{
		_statusCode: code,
	}
}

/*GetNfsAliasDefault handles this case with default header values.

Unexpected error
*/
type GetNfsAliasDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get nfs alias default response
func (o *GetNfsAliasDefault) Code() int {
	return o._statusCode
}

func (o *GetNfsAliasDefault) Error() string {
	return fmt.Sprintf("[GET /platform/2/protocols/nfs/aliases/{NfsAliasId}][%d] getNfsAlias default  %+v", o._statusCode, o.Payload)
}

func (o *GetNfsAliasDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
