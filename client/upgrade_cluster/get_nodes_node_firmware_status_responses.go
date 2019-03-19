// Code generated by go-swagger; DO NOT EDIT.

package upgrade_cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetNodesNodeFirmwareStatusReader is a Reader for the GetNodesNodeFirmwareStatus structure.
type GetNodesNodeFirmwareStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNodesNodeFirmwareStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNodesNodeFirmwareStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetNodesNodeFirmwareStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNodesNodeFirmwareStatusOK creates a GetNodesNodeFirmwareStatusOK with default headers values
func NewGetNodesNodeFirmwareStatusOK() *GetNodesNodeFirmwareStatusOK {
	return &GetNodesNodeFirmwareStatusOK{}
}

/*GetNodesNodeFirmwareStatusOK handles this case with default header values.

The firmware status for the node.
*/
type GetNodesNodeFirmwareStatusOK struct {
	Payload *models.NodesNodeFirmwareStatus
}

func (o *GetNodesNodeFirmwareStatusOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/upgrade/cluster/nodes/{Lnn}/firmware/status][%d] getNodesNodeFirmwareStatusOK  %+v", 200, o.Payload)
}

func (o *GetNodesNodeFirmwareStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodesNodeFirmwareStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodesNodeFirmwareStatusDefault creates a GetNodesNodeFirmwareStatusDefault with default headers values
func NewGetNodesNodeFirmwareStatusDefault(code int) *GetNodesNodeFirmwareStatusDefault {
	return &GetNodesNodeFirmwareStatusDefault{
		_statusCode: code,
	}
}

/*GetNodesNodeFirmwareStatusDefault handles this case with default header values.

Unexpected error
*/
type GetNodesNodeFirmwareStatusDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get nodes node firmware status default response
func (o *GetNodesNodeFirmwareStatusDefault) Code() int {
	return o._statusCode
}

func (o *GetNodesNodeFirmwareStatusDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/upgrade/cluster/nodes/{Lnn}/firmware/status][%d] getNodesNodeFirmwareStatus default  %+v", o._statusCode, o.Payload)
}

func (o *GetNodesNodeFirmwareStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
