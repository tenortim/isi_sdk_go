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

// GetNodeDrivesPurposelistReader is a Reader for the GetNodeDrivesPurposelist structure.
type GetNodeDrivesPurposelistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNodeDrivesPurposelistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNodeDrivesPurposelistOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetNodeDrivesPurposelistDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNodeDrivesPurposelistOK creates a GetNodeDrivesPurposelistOK with default headers values
func NewGetNodeDrivesPurposelistOK() *GetNodeDrivesPurposelistOK {
	return &GetNodeDrivesPurposelistOK{}
}

/*GetNodeDrivesPurposelistOK handles this case with default header values.

Lists the available purposes for drives in this node.
*/
type GetNodeDrivesPurposelistOK struct {
	Payload *models.NodeDrivesPurposelist
}

func (o *GetNodeDrivesPurposelistOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/cluster/nodes/{Lnn}/drives-purposelist][%d] getNodeDrivesPurposelistOK  %+v", 200, o.Payload)
}

func (o *GetNodeDrivesPurposelistOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeDrivesPurposelist)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodeDrivesPurposelistDefault creates a GetNodeDrivesPurposelistDefault with default headers values
func NewGetNodeDrivesPurposelistDefault(code int) *GetNodeDrivesPurposelistDefault {
	return &GetNodeDrivesPurposelistDefault{
		_statusCode: code,
	}
}

/*GetNodeDrivesPurposelistDefault handles this case with default header values.

Unexpected error
*/
type GetNodeDrivesPurposelistDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get node drives purposelist default response
func (o *GetNodeDrivesPurposelistDefault) Code() int {
	return o._statusCode
}

func (o *GetNodeDrivesPurposelistDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/cluster/nodes/{Lnn}/drives-purposelist][%d] getNodeDrivesPurposelist default  %+v", o._statusCode, o.Payload)
}

func (o *GetNodeDrivesPurposelistDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
