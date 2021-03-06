// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetClusterVersionReader is a Reader for the GetClusterVersion structure.
type GetClusterVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClusterVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetClusterVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterVersionOK creates a GetClusterVersionOK with default headers values
func NewGetClusterVersionOK() *GetClusterVersionOK {
	return &GetClusterVersionOK{}
}

/*GetClusterVersionOK handles this case with default header values.

Retrieve the OneFS version as reported by each node.
*/
type GetClusterVersionOK struct {
	Payload *models.ClusterVersion
}

func (o *GetClusterVersionOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/cluster/version][%d] getClusterVersionOK  %+v", 200, o.Payload)
}

func (o *GetClusterVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterVersionDefault creates a GetClusterVersionDefault with default headers values
func NewGetClusterVersionDefault(code int) *GetClusterVersionDefault {
	return &GetClusterVersionDefault{
		_statusCode: code,
	}
}

/*GetClusterVersionDefault handles this case with default header values.

Unexpected error
*/
type GetClusterVersionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get cluster version default response
func (o *GetClusterVersionDefault) Code() int {
	return o._statusCode
}

func (o *GetClusterVersionDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/cluster/version][%d] getClusterVersion default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
