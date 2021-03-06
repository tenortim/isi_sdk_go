// Code generated by go-swagger; DO NOT EDIT.

package network_groupnets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// DeleteGroupnetSubnetReader is a Reader for the DeleteGroupnetSubnet structure.
type DeleteGroupnetSubnetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGroupnetSubnetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteGroupnetSubnetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteGroupnetSubnetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteGroupnetSubnetNoContent creates a DeleteGroupnetSubnetNoContent with default headers values
func NewDeleteGroupnetSubnetNoContent() *DeleteGroupnetSubnetNoContent {
	return &DeleteGroupnetSubnetNoContent{}
}

/*DeleteGroupnetSubnetNoContent handles this case with default header values.

Success.
*/
type DeleteGroupnetSubnetNoContent struct {
}

func (o *DeleteGroupnetSubnetNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/network/groupnets/{Groupnet}/subnets/{GroupnetSubnetId}][%d] deleteGroupnetSubnetNoContent ", 204)
}

func (o *DeleteGroupnetSubnetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteGroupnetSubnetDefault creates a DeleteGroupnetSubnetDefault with default headers values
func NewDeleteGroupnetSubnetDefault(code int) *DeleteGroupnetSubnetDefault {
	return &DeleteGroupnetSubnetDefault{
		_statusCode: code,
	}
}

/*DeleteGroupnetSubnetDefault handles this case with default header values.

Unexpected error
*/
type DeleteGroupnetSubnetDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete groupnet subnet default response
func (o *DeleteGroupnetSubnetDefault) Code() int {
	return o._statusCode
}

func (o *DeleteGroupnetSubnetDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/network/groupnets/{Groupnet}/subnets/{GroupnetSubnetId}][%d] deleteGroupnetSubnet default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteGroupnetSubnetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
