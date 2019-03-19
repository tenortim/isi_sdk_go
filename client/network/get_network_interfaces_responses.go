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

// GetNetworkInterfacesReader is a Reader for the GetNetworkInterfaces structure.
type GetNetworkInterfacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkInterfacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNetworkInterfacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetNetworkInterfacesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworkInterfacesOK creates a GetNetworkInterfacesOK with default headers values
func NewGetNetworkInterfacesOK() *GetNetworkInterfacesOK {
	return &GetNetworkInterfacesOK{}
}

/*GetNetworkInterfacesOK handles this case with default header values.

Get a list of interfaces.
*/
type GetNetworkInterfacesOK struct {
	Payload *models.PoolsPoolInterfaces
}

func (o *GetNetworkInterfacesOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/network/interfaces][%d] getNetworkInterfacesOK  %+v", 200, o.Payload)
}

func (o *GetNetworkInterfacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PoolsPoolInterfaces)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkInterfacesDefault creates a GetNetworkInterfacesDefault with default headers values
func NewGetNetworkInterfacesDefault(code int) *GetNetworkInterfacesDefault {
	return &GetNetworkInterfacesDefault{
		_statusCode: code,
	}
}

/*GetNetworkInterfacesDefault handles this case with default header values.

Unexpected error
*/
type GetNetworkInterfacesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get network interfaces default response
func (o *GetNetworkInterfacesDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworkInterfacesDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/network/interfaces][%d] getNetworkInterfaces default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetworkInterfacesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
