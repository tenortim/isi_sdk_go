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

// DeleteQuotaQuotaReader is a Reader for the DeleteQuotaQuota structure.
type DeleteQuotaQuotaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteQuotaQuotaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteQuotaQuotaNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteQuotaQuotaDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteQuotaQuotaNoContent creates a DeleteQuotaQuotaNoContent with default headers values
func NewDeleteQuotaQuotaNoContent() *DeleteQuotaQuotaNoContent {
	return &DeleteQuotaQuotaNoContent{}
}

/*DeleteQuotaQuotaNoContent handles this case with default header values.

Success.
*/
type DeleteQuotaQuotaNoContent struct {
}

func (o *DeleteQuotaQuotaNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/quota/quotas/{QuotaQuotaId}][%d] deleteQuotaQuotaNoContent ", 204)
}

func (o *DeleteQuotaQuotaNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteQuotaQuotaDefault creates a DeleteQuotaQuotaDefault with default headers values
func NewDeleteQuotaQuotaDefault(code int) *DeleteQuotaQuotaDefault {
	return &DeleteQuotaQuotaDefault{
		_statusCode: code,
	}
}

/*DeleteQuotaQuotaDefault handles this case with default header values.

Unexpected error
*/
type DeleteQuotaQuotaDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete quota quota default response
func (o *DeleteQuotaQuotaDefault) Code() int {
	return o._statusCode
}

func (o *DeleteQuotaQuotaDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/quota/quotas/{QuotaQuotaId}][%d] deleteQuotaQuota default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteQuotaQuotaDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
