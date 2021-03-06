// Code generated by go-swagger; DO NOT EDIT.

package upgrade

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetClusterNodeReader is a Reader for the GetClusterNode structure.
type GetClusterNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClusterNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetClusterNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterNodeOK creates a GetClusterNodeOK with default headers values
func NewGetClusterNodeOK() *GetClusterNodeOK {
	return &GetClusterNodeOK{}
}

/*GetClusterNodeOK handles this case with default header values.

The node details useful during an upgrade or assessment.
*/
type GetClusterNodeOK struct {
	Payload *models.ClusterNodesExtended
}

func (o *GetClusterNodeOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/upgrade/cluster/nodes/{ClusterNodeId}][%d] getClusterNodeOK  %+v", 200, o.Payload)
}

func (o *GetClusterNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterNodesExtended)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterNodeDefault creates a GetClusterNodeDefault with default headers values
func NewGetClusterNodeDefault(code int) *GetClusterNodeDefault {
	return &GetClusterNodeDefault{
		_statusCode: code,
	}
}

/*GetClusterNodeDefault handles this case with default header values.

Unexpected error
*/
type GetClusterNodeDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get cluster node default response
func (o *GetClusterNodeDefault) Code() int {
	return o._statusCode
}

func (o *GetClusterNodeDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/upgrade/cluster/nodes/{ClusterNodeId}][%d] getClusterNode default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
