// Code generated by go-swagger; DO NOT EDIT.

package sync

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// CreateSyncRuleReader is a Reader for the CreateSyncRule structure.
type CreateSyncRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSyncRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateSyncRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateSyncRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSyncRuleOK creates a CreateSyncRuleOK with default headers values
func NewCreateSyncRuleOK() *CreateSyncRuleOK {
	return &CreateSyncRuleOK{}
}

/*CreateSyncRuleOK handles this case with default header values.

Create a new SyncIQ performance rule.
*/
type CreateSyncRuleOK struct {
	Payload *models.CreateResponse
}

func (o *CreateSyncRuleOK) Error() string {
	return fmt.Sprintf("[POST /platform/3/sync/rules][%d] createSyncRuleOK  %+v", 200, o.Payload)
}

func (o *CreateSyncRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSyncRuleDefault creates a CreateSyncRuleDefault with default headers values
func NewCreateSyncRuleDefault(code int) *CreateSyncRuleDefault {
	return &CreateSyncRuleDefault{
		_statusCode: code,
	}
}

/*CreateSyncRuleDefault handles this case with default header values.

Unexpected error
*/
type CreateSyncRuleDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create sync rule default response
func (o *CreateSyncRuleDefault) Code() int {
	return o._statusCode
}

func (o *CreateSyncRuleDefault) Error() string {
	return fmt.Sprintf("[POST /platform/3/sync/rules][%d] createSyncRule default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSyncRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
