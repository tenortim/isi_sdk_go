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

// GetNodeDrivesReader is a Reader for the GetNodeDrives structure.
type GetNodeDrivesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNodeDrivesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNodeDrivesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetNodeDrivesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNodeDrivesOK creates a GetNodeDrivesOK with default headers values
func NewGetNodeDrivesOK() *GetNodeDrivesOK {
	return &GetNodeDrivesOK{}
}

/*GetNodeDrivesOK handles this case with default header values.

List the drives on this node.
*/
type GetNodeDrivesOK struct {
	Payload *models.NodeDrives
}

func (o *GetNodeDrivesOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/cluster/nodes/{Lnn}/drives][%d] getNodeDrivesOK  %+v", 200, o.Payload)
}

func (o *GetNodeDrivesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeDrives)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodeDrivesDefault creates a GetNodeDrivesDefault with default headers values
func NewGetNodeDrivesDefault(code int) *GetNodeDrivesDefault {
	return &GetNodeDrivesDefault{
		_statusCode: code,
	}
}

/*GetNodeDrivesDefault handles this case with default header values.

Unexpected error
*/
type GetNodeDrivesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get node drives default response
func (o *GetNodeDrivesDefault) Code() int {
	return o._statusCode
}

func (o *GetNodeDrivesDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/cluster/nodes/{Lnn}/drives][%d] getNodeDrives default  %+v", o._statusCode, o.Payload)
}

func (o *GetNodeDrivesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
