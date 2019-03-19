// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// CreateDnscacheFlushItemReader is a Reader for the CreateDnscacheFlushItem structure.
type CreateDnscacheFlushItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDnscacheFlushItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateDnscacheFlushItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateDnscacheFlushItemDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateDnscacheFlushItemOK creates a CreateDnscacheFlushItemOK with default headers values
func NewCreateDnscacheFlushItemOK() *CreateDnscacheFlushItemOK {
	return &CreateDnscacheFlushItemOK{}
}

/*CreateDnscacheFlushItemOK handles this case with default header values.

Flush the DNSCache.
*/
type CreateDnscacheFlushItemOK struct {
	Payload models.Empty
}

func (o *CreateDnscacheFlushItemOK) Error() string {
	return fmt.Sprintf("[POST /platform/3/network/dnscache/flush][%d] createDnscacheFlushItemOK  %+v", 200, o.Payload)
}

func (o *CreateDnscacheFlushItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDnscacheFlushItemDefault creates a CreateDnscacheFlushItemDefault with default headers values
func NewCreateDnscacheFlushItemDefault(code int) *CreateDnscacheFlushItemDefault {
	return &CreateDnscacheFlushItemDefault{
		_statusCode: code,
	}
}

/*CreateDnscacheFlushItemDefault handles this case with default header values.

Unexpected error
*/
type CreateDnscacheFlushItemDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create dnscache flush item default response
func (o *CreateDnscacheFlushItemDefault) Code() int {
	return o._statusCode
}

func (o *CreateDnscacheFlushItemDefault) Error() string {
	return fmt.Sprintf("[POST /platform/3/network/dnscache/flush][%d] createDnscacheFlushItem default  %+v", o._statusCode, o.Payload)
}

func (o *CreateDnscacheFlushItemDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}