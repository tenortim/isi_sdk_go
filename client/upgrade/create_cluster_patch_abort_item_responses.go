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

// CreateClusterPatchAbortItemReader is a Reader for the CreateClusterPatchAbortItem structure.
type CreateClusterPatchAbortItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateClusterPatchAbortItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateClusterPatchAbortItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateClusterPatchAbortItemDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateClusterPatchAbortItemOK creates a CreateClusterPatchAbortItemOK with default headers values
func NewCreateClusterPatchAbortItemOK() *CreateClusterPatchAbortItemOK {
	return &CreateClusterPatchAbortItemOK{}
}

/*CreateClusterPatchAbortItemOK handles this case with default header values.

Abort the previous action performed by the patch system.
*/
type CreateClusterPatchAbortItemOK struct {
	Payload models.Empty
}

func (o *CreateClusterPatchAbortItemOK) Error() string {
	return fmt.Sprintf("[POST /platform/3/upgrade/cluster/patch/abort][%d] createClusterPatchAbortItemOK  %+v", 200, o.Payload)
}

func (o *CreateClusterPatchAbortItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterPatchAbortItemDefault creates a CreateClusterPatchAbortItemDefault with default headers values
func NewCreateClusterPatchAbortItemDefault(code int) *CreateClusterPatchAbortItemDefault {
	return &CreateClusterPatchAbortItemDefault{
		_statusCode: code,
	}
}

/*CreateClusterPatchAbortItemDefault handles this case with default header values.

Unexpected error
*/
type CreateClusterPatchAbortItemDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create cluster patch abort item default response
func (o *CreateClusterPatchAbortItemDefault) Code() int {
	return o._statusCode
}

func (o *CreateClusterPatchAbortItemDefault) Error() string {
	return fmt.Sprintf("[POST /platform/3/upgrade/cluster/patch/abort][%d] createClusterPatchAbortItem default  %+v", o._statusCode, o.Payload)
}

func (o *CreateClusterPatchAbortItemDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
