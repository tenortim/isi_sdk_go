// Code generated by go-swagger; DO NOT EDIT.

package cloud

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// UpdateCloudJobReader is a Reader for the UpdateCloudJob structure.
type UpdateCloudJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCloudJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateCloudJobNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateCloudJobDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateCloudJobNoContent creates a UpdateCloudJobNoContent with default headers values
func NewUpdateCloudJobNoContent() *UpdateCloudJobNoContent {
	return &UpdateCloudJobNoContent{}
}

/*UpdateCloudJobNoContent handles this case with default header values.

Success.
*/
type UpdateCloudJobNoContent struct {
}

func (o *UpdateCloudJobNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/3/cloud/jobs/{CloudJobId}][%d] updateCloudJobNoContent ", 204)
}

func (o *UpdateCloudJobNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateCloudJobDefault creates a UpdateCloudJobDefault with default headers values
func NewUpdateCloudJobDefault(code int) *UpdateCloudJobDefault {
	return &UpdateCloudJobDefault{
		_statusCode: code,
	}
}

/*UpdateCloudJobDefault handles this case with default header values.

Unexpected error
*/
type UpdateCloudJobDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update cloud job default response
func (o *UpdateCloudJobDefault) Code() int {
	return o._statusCode
}

func (o *UpdateCloudJobDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/3/cloud/jobs/{CloudJobId}][%d] updateCloudJob default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateCloudJobDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
