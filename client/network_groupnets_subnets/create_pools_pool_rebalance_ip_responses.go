// Code generated by go-swagger; DO NOT EDIT.

package network_groupnets_subnets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// CreatePoolsPoolRebalanceIPReader is a Reader for the CreatePoolsPoolRebalanceIP structure.
type CreatePoolsPoolRebalanceIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePoolsPoolRebalanceIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreatePoolsPoolRebalanceIPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreatePoolsPoolRebalanceIPDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreatePoolsPoolRebalanceIPOK creates a CreatePoolsPoolRebalanceIPOK with default headers values
func NewCreatePoolsPoolRebalanceIPOK() *CreatePoolsPoolRebalanceIPOK {
	return &CreatePoolsPoolRebalanceIPOK{}
}

/*CreatePoolsPoolRebalanceIPOK handles this case with default header values.

Rebalance IP addresses in specified pool.
*/
type CreatePoolsPoolRebalanceIPOK struct {
	Payload models.Empty
}

func (o *CreatePoolsPoolRebalanceIPOK) Error() string {
	return fmt.Sprintf("[POST /platform/3/network/groupnets/{Groupnet}/subnets/{Subnet}/pools/{Pool}/rebalance-ips][%d] createPoolsPoolRebalanceIpOK  %+v", 200, o.Payload)
}

func (o *CreatePoolsPoolRebalanceIPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePoolsPoolRebalanceIPDefault creates a CreatePoolsPoolRebalanceIPDefault with default headers values
func NewCreatePoolsPoolRebalanceIPDefault(code int) *CreatePoolsPoolRebalanceIPDefault {
	return &CreatePoolsPoolRebalanceIPDefault{
		_statusCode: code,
	}
}

/*CreatePoolsPoolRebalanceIPDefault handles this case with default header values.

Unexpected error
*/
type CreatePoolsPoolRebalanceIPDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create pools pool rebalance Ip default response
func (o *CreatePoolsPoolRebalanceIPDefault) Code() int {
	return o._statusCode
}

func (o *CreatePoolsPoolRebalanceIPDefault) Error() string {
	return fmt.Sprintf("[POST /platform/3/network/groupnets/{Groupnet}/subnets/{Subnet}/pools/{Pool}/rebalance-ips][%d] createPoolsPoolRebalanceIp default  %+v", o._statusCode, o.Payload)
}

func (o *CreatePoolsPoolRebalanceIPDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
