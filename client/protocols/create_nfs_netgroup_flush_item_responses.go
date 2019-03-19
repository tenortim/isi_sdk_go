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

// CreateNfsNetgroupFlushItemReader is a Reader for the CreateNfsNetgroupFlushItem structure.
type CreateNfsNetgroupFlushItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNfsNetgroupFlushItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateNfsNetgroupFlushItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateNfsNetgroupFlushItemDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateNfsNetgroupFlushItemOK creates a CreateNfsNetgroupFlushItemOK with default headers values
func NewCreateNfsNetgroupFlushItemOK() *CreateNfsNetgroupFlushItemOK {
	return &CreateNfsNetgroupFlushItemOK{}
}

/*CreateNfsNetgroupFlushItemOK handles this case with default header values.

Flush the NFS netgroups in the cache.
*/
type CreateNfsNetgroupFlushItemOK struct {
	Payload models.Empty
}

func (o *CreateNfsNetgroupFlushItemOK) Error() string {
	return fmt.Sprintf("[POST /platform/3/protocols/nfs/netgroup/flush][%d] createNfsNetgroupFlushItemOK  %+v", 200, o.Payload)
}

func (o *CreateNfsNetgroupFlushItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNfsNetgroupFlushItemDefault creates a CreateNfsNetgroupFlushItemDefault with default headers values
func NewCreateNfsNetgroupFlushItemDefault(code int) *CreateNfsNetgroupFlushItemDefault {
	return &CreateNfsNetgroupFlushItemDefault{
		_statusCode: code,
	}
}

/*CreateNfsNetgroupFlushItemDefault handles this case with default header values.

Unexpected error
*/
type CreateNfsNetgroupFlushItemDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create nfs netgroup flush item default response
func (o *CreateNfsNetgroupFlushItemDefault) Code() int {
	return o._statusCode
}

func (o *CreateNfsNetgroupFlushItemDefault) Error() string {
	return fmt.Sprintf("[POST /platform/3/protocols/nfs/netgroup/flush][%d] createNfsNetgroupFlushItem default  %+v", o._statusCode, o.Payload)
}

func (o *CreateNfsNetgroupFlushItemDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}