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

// ListNfsAliasesReader is a Reader for the ListNfsAliases structure.
type ListNfsAliasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListNfsAliasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListNfsAliasesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListNfsAliasesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListNfsAliasesOK creates a ListNfsAliasesOK with default headers values
func NewListNfsAliasesOK() *ListNfsAliasesOK {
	return &ListNfsAliasesOK{}
}

/*ListNfsAliasesOK handles this case with default header values.

List all NFS aliases.
*/
type ListNfsAliasesOK struct {
	Payload *models.NfsAliasesExtended
}

func (o *ListNfsAliasesOK) Error() string {
	return fmt.Sprintf("[GET /platform/2/protocols/nfs/aliases][%d] listNfsAliasesOK  %+v", 200, o.Payload)
}

func (o *ListNfsAliasesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NfsAliasesExtended)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNfsAliasesDefault creates a ListNfsAliasesDefault with default headers values
func NewListNfsAliasesDefault(code int) *ListNfsAliasesDefault {
	return &ListNfsAliasesDefault{
		_statusCode: code,
	}
}

/*ListNfsAliasesDefault handles this case with default header values.

Unexpected error
*/
type ListNfsAliasesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list nfs aliases default response
func (o *ListNfsAliasesDefault) Code() int {
	return o._statusCode
}

func (o *ListNfsAliasesDefault) Error() string {
	return fmt.Sprintf("[GET /platform/2/protocols/nfs/aliases][%d] listNfsAliases default  %+v", o._statusCode, o.Payload)
}

func (o *ListNfsAliasesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
