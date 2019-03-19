// Code generated by go-swagger; DO NOT EDIT.

package audit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// DeleteAuditTopicReader is a Reader for the DeleteAuditTopic structure.
type DeleteAuditTopicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAuditTopicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteAuditTopicNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteAuditTopicDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAuditTopicNoContent creates a DeleteAuditTopicNoContent with default headers values
func NewDeleteAuditTopicNoContent() *DeleteAuditTopicNoContent {
	return &DeleteAuditTopicNoContent{}
}

/*DeleteAuditTopicNoContent handles this case with default header values.

Success.
*/
type DeleteAuditTopicNoContent struct {
}

func (o *DeleteAuditTopicNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/audit/topics/{AuditTopicId}][%d] deleteAuditTopicNoContent ", 204)
}

func (o *DeleteAuditTopicNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAuditTopicDefault creates a DeleteAuditTopicDefault with default headers values
func NewDeleteAuditTopicDefault(code int) *DeleteAuditTopicDefault {
	return &DeleteAuditTopicDefault{
		_statusCode: code,
	}
}

/*DeleteAuditTopicDefault handles this case with default header values.

Unexpected error
*/
type DeleteAuditTopicDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete audit topic default response
func (o *DeleteAuditTopicDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAuditTopicDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/audit/topics/{AuditTopicId}][%d] deleteAuditTopic default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAuditTopicDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}