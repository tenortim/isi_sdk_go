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

// CreateDrivesDriveAddItemReader is a Reader for the CreateDrivesDriveAddItem structure.
type CreateDrivesDriveAddItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDrivesDriveAddItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateDrivesDriveAddItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateDrivesDriveAddItemDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateDrivesDriveAddItemOK creates a CreateDrivesDriveAddItemOK with default headers values
func NewCreateDrivesDriveAddItemOK() *CreateDrivesDriveAddItemOK {
	return &CreateDrivesDriveAddItemOK{}
}

/*CreateDrivesDriveAddItemOK handles this case with default header values.

Add a drive to a node.
*/
type CreateDrivesDriveAddItemOK struct {
	Payload models.Empty
}

func (o *CreateDrivesDriveAddItemOK) Error() string {
	return fmt.Sprintf("[POST /platform/3/cluster/nodes/{Lnn}/drives/{Driveid}/add][%d] createDrivesDriveAddItemOK  %+v", 200, o.Payload)
}

func (o *CreateDrivesDriveAddItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDrivesDriveAddItemDefault creates a CreateDrivesDriveAddItemDefault with default headers values
func NewCreateDrivesDriveAddItemDefault(code int) *CreateDrivesDriveAddItemDefault {
	return &CreateDrivesDriveAddItemDefault{
		_statusCode: code,
	}
}

/*CreateDrivesDriveAddItemDefault handles this case with default header values.

Unexpected error
*/
type CreateDrivesDriveAddItemDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create drives drive add item default response
func (o *CreateDrivesDriveAddItemDefault) Code() int {
	return o._statusCode
}

func (o *CreateDrivesDriveAddItemDefault) Error() string {
	return fmt.Sprintf("[POST /platform/3/cluster/nodes/{Lnn}/drives/{Driveid}/add][%d] createDrivesDriveAddItem default  %+v", o._statusCode, o.Payload)
}

func (o *CreateDrivesDriveAddItemDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}