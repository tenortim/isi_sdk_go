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

// GetNodeStatusBatterystatusReader is a Reader for the GetNodeStatusBatterystatus structure.
type GetNodeStatusBatterystatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNodeStatusBatterystatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNodeStatusBatterystatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetNodeStatusBatterystatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNodeStatusBatterystatusOK creates a GetNodeStatusBatterystatusOK with default headers values
func NewGetNodeStatusBatterystatusOK() *GetNodeStatusBatterystatusOK {
	return &GetNodeStatusBatterystatusOK{}
}

/*GetNodeStatusBatterystatusOK handles this case with default header values.

Retrieve node battery status information.
*/
type GetNodeStatusBatterystatusOK struct {
	Payload *models.NodeStatusBatterystatus
}

func (o *GetNodeStatusBatterystatusOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/cluster/nodes/{Lnn}/status/batterystatus][%d] getNodeStatusBatterystatusOK  %+v", 200, o.Payload)
}

func (o *GetNodeStatusBatterystatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeStatusBatterystatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodeStatusBatterystatusDefault creates a GetNodeStatusBatterystatusDefault with default headers values
func NewGetNodeStatusBatterystatusDefault(code int) *GetNodeStatusBatterystatusDefault {
	return &GetNodeStatusBatterystatusDefault{
		_statusCode: code,
	}
}

/*GetNodeStatusBatterystatusDefault handles this case with default header values.

Unexpected error
*/
type GetNodeStatusBatterystatusDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get node status batterystatus default response
func (o *GetNodeStatusBatterystatusDefault) Code() int {
	return o._statusCode
}

func (o *GetNodeStatusBatterystatusDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/cluster/nodes/{Lnn}/status/batterystatus][%d] getNodeStatusBatterystatus default  %+v", o._statusCode, o.Payload)
}

func (o *GetNodeStatusBatterystatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
