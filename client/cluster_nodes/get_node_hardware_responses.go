// Code generated by go-swagger; DO NOT EDIT.

package cluster_nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetNodeHardwareReader is a Reader for the GetNodeHardware structure.
type GetNodeHardwareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNodeHardwareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNodeHardwareOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetNodeHardwareDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNodeHardwareOK creates a GetNodeHardwareOK with default headers values
func NewGetNodeHardwareOK() *GetNodeHardwareOK {
	return &GetNodeHardwareOK{}
}

/*GetNodeHardwareOK handles this case with default header values.

Retrieve node hardware identity information.
*/
type GetNodeHardwareOK struct {
	Payload *models.NodeHardware
}

func (o *GetNodeHardwareOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/cluster/nodes/{Lnn}/hardware][%d] getNodeHardwareOK  %+v", 200, o.Payload)
}

func (o *GetNodeHardwareOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeHardware)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodeHardwareDefault creates a GetNodeHardwareDefault with default headers values
func NewGetNodeHardwareDefault(code int) *GetNodeHardwareDefault {
	return &GetNodeHardwareDefault{
		_statusCode: code,
	}
}

/*GetNodeHardwareDefault handles this case with default header values.

Unexpected error
*/
type GetNodeHardwareDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get node hardware default response
func (o *GetNodeHardwareDefault) Code() int {
	return o._statusCode
}

func (o *GetNodeHardwareDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/cluster/nodes/{Lnn}/hardware][%d] getNodeHardware default  %+v", o._statusCode, o.Payload)
}

func (o *GetNodeHardwareDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}