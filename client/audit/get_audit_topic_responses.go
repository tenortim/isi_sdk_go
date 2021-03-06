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

// GetAuditTopicReader is a Reader for the GetAuditTopic structure.
type GetAuditTopicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuditTopicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAuditTopicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetAuditTopicDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAuditTopicOK creates a GetAuditTopicOK with default headers values
func NewGetAuditTopicOK() *GetAuditTopicOK {
	return &GetAuditTopicOK{}
}

/*GetAuditTopicOK handles this case with default header values.

Retrieve the audit topic information.
*/
type GetAuditTopicOK struct {
	Payload *models.AuditTopics
}

func (o *GetAuditTopicOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/audit/topics/{AuditTopicId}][%d] getAuditTopicOK  %+v", 200, o.Payload)
}

func (o *GetAuditTopicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuditTopics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditTopicDefault creates a GetAuditTopicDefault with default headers values
func NewGetAuditTopicDefault(code int) *GetAuditTopicDefault {
	return &GetAuditTopicDefault{
		_statusCode: code,
	}
}

/*GetAuditTopicDefault handles this case with default header values.

Unexpected error
*/
type GetAuditTopicDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get audit topic default response
func (o *GetAuditTopicDefault) Code() int {
	return o._statusCode
}

func (o *GetAuditTopicDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/audit/topics/{AuditTopicId}][%d] getAuditTopic default  %+v", o._statusCode, o.Payload)
}

func (o *GetAuditTopicDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
