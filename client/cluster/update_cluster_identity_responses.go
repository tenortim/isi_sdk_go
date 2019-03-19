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

// UpdateClusterIdentityReader is a Reader for the UpdateClusterIdentity structure.
type UpdateClusterIdentityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateClusterIdentityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateClusterIdentityNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateClusterIdentityDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateClusterIdentityNoContent creates a UpdateClusterIdentityNoContent with default headers values
func NewUpdateClusterIdentityNoContent() *UpdateClusterIdentityNoContent {
	return &UpdateClusterIdentityNoContent{}
}

/*UpdateClusterIdentityNoContent handles this case with default header values.

Success.
*/
type UpdateClusterIdentityNoContent struct {
}

func (o *UpdateClusterIdentityNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/3/cluster/identity][%d] updateClusterIdentityNoContent ", 204)
}

func (o *UpdateClusterIdentityNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateClusterIdentityDefault creates a UpdateClusterIdentityDefault with default headers values
func NewUpdateClusterIdentityDefault(code int) *UpdateClusterIdentityDefault {
	return &UpdateClusterIdentityDefault{
		_statusCode: code,
	}
}

/*UpdateClusterIdentityDefault handles this case with default header values.

Unexpected error
*/
type UpdateClusterIdentityDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update cluster identity default response
func (o *UpdateClusterIdentityDefault) Code() int {
	return o._statusCode
}

func (o *UpdateClusterIdentityDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/3/cluster/identity][%d] updateClusterIdentity default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateClusterIdentityDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
