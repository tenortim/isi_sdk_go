// Code generated by go-swagger; DO NOT EDIT.

package filepool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// UpdateFilepoolDefaultPolicyReader is a Reader for the UpdateFilepoolDefaultPolicy structure.
type UpdateFilepoolDefaultPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateFilepoolDefaultPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateFilepoolDefaultPolicyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateFilepoolDefaultPolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateFilepoolDefaultPolicyNoContent creates a UpdateFilepoolDefaultPolicyNoContent with default headers values
func NewUpdateFilepoolDefaultPolicyNoContent() *UpdateFilepoolDefaultPolicyNoContent {
	return &UpdateFilepoolDefaultPolicyNoContent{}
}

/*UpdateFilepoolDefaultPolicyNoContent handles this case with default header values.

Success.
*/
type UpdateFilepoolDefaultPolicyNoContent struct {
}

func (o *UpdateFilepoolDefaultPolicyNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/1/filepool/default-policy][%d] updateFilepoolDefaultPolicyNoContent ", 204)
}

func (o *UpdateFilepoolDefaultPolicyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateFilepoolDefaultPolicyDefault creates a UpdateFilepoolDefaultPolicyDefault with default headers values
func NewUpdateFilepoolDefaultPolicyDefault(code int) *UpdateFilepoolDefaultPolicyDefault {
	return &UpdateFilepoolDefaultPolicyDefault{
		_statusCode: code,
	}
}

/*UpdateFilepoolDefaultPolicyDefault handles this case with default header values.

Unexpected error
*/
type UpdateFilepoolDefaultPolicyDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update filepool default policy default response
func (o *UpdateFilepoolDefaultPolicyDefault) Code() int {
	return o._statusCode
}

func (o *UpdateFilepoolDefaultPolicyDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/1/filepool/default-policy][%d] updateFilepoolDefaultPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateFilepoolDefaultPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
