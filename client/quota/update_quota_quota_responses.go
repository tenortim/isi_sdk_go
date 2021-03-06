// Code generated by go-swagger; DO NOT EDIT.

package quota

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// UpdateQuotaQuotaReader is a Reader for the UpdateQuotaQuota structure.
type UpdateQuotaQuotaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateQuotaQuotaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateQuotaQuotaNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateQuotaQuotaDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateQuotaQuotaNoContent creates a UpdateQuotaQuotaNoContent with default headers values
func NewUpdateQuotaQuotaNoContent() *UpdateQuotaQuotaNoContent {
	return &UpdateQuotaQuotaNoContent{}
}

/*UpdateQuotaQuotaNoContent handles this case with default header values.

Success.
*/
type UpdateQuotaQuotaNoContent struct {
}

func (o *UpdateQuotaQuotaNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/1/quota/quotas/{QuotaQuotaId}][%d] updateQuotaQuotaNoContent ", 204)
}

func (o *UpdateQuotaQuotaNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateQuotaQuotaDefault creates a UpdateQuotaQuotaDefault with default headers values
func NewUpdateQuotaQuotaDefault(code int) *UpdateQuotaQuotaDefault {
	return &UpdateQuotaQuotaDefault{
		_statusCode: code,
	}
}

/*UpdateQuotaQuotaDefault handles this case with default header values.

Unexpected error
*/
type UpdateQuotaQuotaDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update quota quota default response
func (o *UpdateQuotaQuotaDefault) Code() int {
	return o._statusCode
}

func (o *UpdateQuotaQuotaDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/1/quota/quotas/{QuotaQuotaId}][%d] updateQuotaQuota default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateQuotaQuotaDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
