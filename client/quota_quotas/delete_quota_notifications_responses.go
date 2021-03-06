// Code generated by go-swagger; DO NOT EDIT.

package quota_quotas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// DeleteQuotaNotificationsReader is a Reader for the DeleteQuotaNotifications structure.
type DeleteQuotaNotificationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteQuotaNotificationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteQuotaNotificationsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteQuotaNotificationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteQuotaNotificationsNoContent creates a DeleteQuotaNotificationsNoContent with default headers values
func NewDeleteQuotaNotificationsNoContent() *DeleteQuotaNotificationsNoContent {
	return &DeleteQuotaNotificationsNoContent{}
}

/*DeleteQuotaNotificationsNoContent handles this case with default header values.

Success.
*/
type DeleteQuotaNotificationsNoContent struct {
}

func (o *DeleteQuotaNotificationsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/quota/quotas/{Qid}/notifications][%d] deleteQuotaNotificationsNoContent ", 204)
}

func (o *DeleteQuotaNotificationsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteQuotaNotificationsDefault creates a DeleteQuotaNotificationsDefault with default headers values
func NewDeleteQuotaNotificationsDefault(code int) *DeleteQuotaNotificationsDefault {
	return &DeleteQuotaNotificationsDefault{
		_statusCode: code,
	}
}

/*DeleteQuotaNotificationsDefault handles this case with default header values.

Unexpected error
*/
type DeleteQuotaNotificationsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete quota notifications default response
func (o *DeleteQuotaNotificationsDefault) Code() int {
	return o._statusCode
}

func (o *DeleteQuotaNotificationsDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/quota/quotas/{Qid}/notifications][%d] deleteQuotaNotifications default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteQuotaNotificationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
