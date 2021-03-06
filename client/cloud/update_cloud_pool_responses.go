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

// UpdateCloudPoolReader is a Reader for the UpdateCloudPool structure.
type UpdateCloudPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCloudPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateCloudPoolNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateCloudPoolDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateCloudPoolNoContent creates a UpdateCloudPoolNoContent with default headers values
func NewUpdateCloudPoolNoContent() *UpdateCloudPoolNoContent {
	return &UpdateCloudPoolNoContent{}
}

/*UpdateCloudPoolNoContent handles this case with default header values.

Success.
*/
type UpdateCloudPoolNoContent struct {
}

func (o *UpdateCloudPoolNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/3/cloud/pools/{CloudPoolId}][%d] updateCloudPoolNoContent ", 204)
}

func (o *UpdateCloudPoolNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateCloudPoolDefault creates a UpdateCloudPoolDefault with default headers values
func NewUpdateCloudPoolDefault(code int) *UpdateCloudPoolDefault {
	return &UpdateCloudPoolDefault{
		_statusCode: code,
	}
}

/*UpdateCloudPoolDefault handles this case with default header values.

Unexpected error
*/
type UpdateCloudPoolDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update cloud pool default response
func (o *UpdateCloudPoolDefault) Code() int {
	return o._statusCode
}

func (o *UpdateCloudPoolDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/3/cloud/pools/{CloudPoolId}][%d] updateCloudPool default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateCloudPoolDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
