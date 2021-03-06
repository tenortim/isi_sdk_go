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

// GetNetworkPoolsReader is a Reader for the GetNetworkPools structure.
type GetNetworkPoolsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkPoolsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNetworkPoolsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetNetworkPoolsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworkPoolsOK creates a GetNetworkPoolsOK with default headers values
func NewGetNetworkPoolsOK() *GetNetworkPoolsOK {
	return &GetNetworkPoolsOK{}
}

/*GetNetworkPoolsOK handles this case with default header values.

Get a list of flexnet pools.
*/
type GetNetworkPoolsOK struct {
	Payload *models.NetworkPools
}

func (o *GetNetworkPoolsOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/network/pools][%d] getNetworkPoolsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkPoolsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkPools)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkPoolsDefault creates a GetNetworkPoolsDefault with default headers values
func NewGetNetworkPoolsDefault(code int) *GetNetworkPoolsDefault {
	return &GetNetworkPoolsDefault{
		_statusCode: code,
	}
}

/*GetNetworkPoolsDefault handles this case with default header values.

Unexpected error
*/
type GetNetworkPoolsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get network pools default response
func (o *GetNetworkPoolsDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworkPoolsDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/network/pools][%d] getNetworkPools default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetworkPoolsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
