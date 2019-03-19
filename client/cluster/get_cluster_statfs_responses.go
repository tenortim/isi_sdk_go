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

// GetClusterStatfsReader is a Reader for the GetClusterStatfs structure.
type GetClusterStatfsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterStatfsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClusterStatfsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetClusterStatfsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterStatfsOK creates a GetClusterStatfsOK with default headers values
func NewGetClusterStatfsOK() *GetClusterStatfsOK {
	return &GetClusterStatfsOK{}
}

/*GetClusterStatfsOK handles this case with default header values.

Retrieve the filesystem statistics.
*/
type GetClusterStatfsOK struct {
	Payload *models.ClusterStatfs
}

func (o *GetClusterStatfsOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/cluster/statfs][%d] getClusterStatfsOK  %+v", 200, o.Payload)
}

func (o *GetClusterStatfsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterStatfs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterStatfsDefault creates a GetClusterStatfsDefault with default headers values
func NewGetClusterStatfsDefault(code int) *GetClusterStatfsDefault {
	return &GetClusterStatfsDefault{
		_statusCode: code,
	}
}

/*GetClusterStatfsDefault handles this case with default header values.

Unexpected error
*/
type GetClusterStatfsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get cluster statfs default response
func (o *GetClusterStatfsDefault) Code() int {
	return o._statusCode
}

func (o *GetClusterStatfsDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/cluster/statfs][%d] getClusterStatfs default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterStatfsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}