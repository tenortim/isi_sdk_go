// Code generated by go-swagger; DO NOT EDIT.

package worm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetWormDomainReader is a Reader for the GetWormDomain structure.
type GetWormDomainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWormDomainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWormDomainOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetWormDomainDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWormDomainOK creates a GetWormDomainOK with default headers values
func NewGetWormDomainOK() *GetWormDomainOK {
	return &GetWormDomainOK{}
}

/*GetWormDomainOK handles this case with default header values.

View a single WORM domain.
*/
type GetWormDomainOK struct {
	Payload *models.WormDomains
}

func (o *GetWormDomainOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/worm/domains/{WormDomainId}][%d] getWormDomainOK  %+v", 200, o.Payload)
}

func (o *GetWormDomainOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WormDomains)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWormDomainDefault creates a GetWormDomainDefault with default headers values
func NewGetWormDomainDefault(code int) *GetWormDomainDefault {
	return &GetWormDomainDefault{
		_statusCode: code,
	}
}

/*GetWormDomainDefault handles this case with default header values.

Unexpected error
*/
type GetWormDomainDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get worm domain default response
func (o *GetWormDomainDefault) Code() int {
	return o._statusCode
}

func (o *GetWormDomainDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/worm/domains/{WormDomainId}][%d] getWormDomain default  %+v", o._statusCode, o.Payload)
}

func (o *GetWormDomainDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
