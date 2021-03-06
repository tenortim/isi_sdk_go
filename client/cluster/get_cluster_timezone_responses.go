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

// GetClusterTimezoneReader is a Reader for the GetClusterTimezone structure.
type GetClusterTimezoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterTimezoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClusterTimezoneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetClusterTimezoneDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterTimezoneOK creates a GetClusterTimezoneOK with default headers values
func NewGetClusterTimezoneOK() *GetClusterTimezoneOK {
	return &GetClusterTimezoneOK{}
}

/*GetClusterTimezoneOK handles this case with default header values.

Get the cluster timezone.
*/
type GetClusterTimezoneOK struct {
	Payload *models.ClusterTimezone
}

func (o *GetClusterTimezoneOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/cluster/timezone][%d] getClusterTimezoneOK  %+v", 200, o.Payload)
}

func (o *GetClusterTimezoneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterTimezone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterTimezoneDefault creates a GetClusterTimezoneDefault with default headers values
func NewGetClusterTimezoneDefault(code int) *GetClusterTimezoneDefault {
	return &GetClusterTimezoneDefault{
		_statusCode: code,
	}
}

/*GetClusterTimezoneDefault handles this case with default header values.

Unexpected error
*/
type GetClusterTimezoneDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get cluster timezone default response
func (o *GetClusterTimezoneDefault) Code() int {
	return o._statusCode
}

func (o *GetClusterTimezoneDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/cluster/timezone][%d] getClusterTimezone default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterTimezoneDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
