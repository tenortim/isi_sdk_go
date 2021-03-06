// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// UpdateJobJobReader is a Reader for the UpdateJobJob structure.
type UpdateJobJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateJobJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateJobJobNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateJobJobDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateJobJobNoContent creates a UpdateJobJobNoContent with default headers values
func NewUpdateJobJobNoContent() *UpdateJobJobNoContent {
	return &UpdateJobJobNoContent{}
}

/*UpdateJobJobNoContent handles this case with default header values.

Success.
*/
type UpdateJobJobNoContent struct {
}

func (o *UpdateJobJobNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/3/job/jobs/{JobJobId}][%d] updateJobJobNoContent ", 204)
}

func (o *UpdateJobJobNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateJobJobDefault creates a UpdateJobJobDefault with default headers values
func NewUpdateJobJobDefault(code int) *UpdateJobJobDefault {
	return &UpdateJobJobDefault{
		_statusCode: code,
	}
}

/*UpdateJobJobDefault handles this case with default header values.

Unexpected error
*/
type UpdateJobJobDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update job job default response
func (o *UpdateJobJobDefault) Code() int {
	return o._statusCode
}

func (o *UpdateJobJobDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/3/job/jobs/{JobJobId}][%d] updateJobJob default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateJobJobDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
