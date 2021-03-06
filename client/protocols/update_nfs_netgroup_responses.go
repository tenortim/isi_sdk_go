// Code generated by go-swagger; DO NOT EDIT.

package protocols

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// UpdateNfsNetgroupReader is a Reader for the UpdateNfsNetgroup structure.
type UpdateNfsNetgroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNfsNetgroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateNfsNetgroupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateNfsNetgroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateNfsNetgroupNoContent creates a UpdateNfsNetgroupNoContent with default headers values
func NewUpdateNfsNetgroupNoContent() *UpdateNfsNetgroupNoContent {
	return &UpdateNfsNetgroupNoContent{}
}

/*UpdateNfsNetgroupNoContent handles this case with default header values.

Success.
*/
type UpdateNfsNetgroupNoContent struct {
}

func (o *UpdateNfsNetgroupNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/3/protocols/nfs/netgroup][%d] updateNfsNetgroupNoContent ", 204)
}

func (o *UpdateNfsNetgroupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateNfsNetgroupDefault creates a UpdateNfsNetgroupDefault with default headers values
func NewUpdateNfsNetgroupDefault(code int) *UpdateNfsNetgroupDefault {
	return &UpdateNfsNetgroupDefault{
		_statusCode: code,
	}
}

/*UpdateNfsNetgroupDefault handles this case with default header values.

Unexpected error
*/
type UpdateNfsNetgroupDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update nfs netgroup default response
func (o *UpdateNfsNetgroupDefault) Code() int {
	return o._statusCode
}

func (o *UpdateNfsNetgroupDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/3/protocols/nfs/netgroup][%d] updateNfsNetgroup default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateNfsNetgroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
