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

// ListQuotaReportsReader is a Reader for the ListQuotaReports structure.
type ListQuotaReportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListQuotaReportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListQuotaReportsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListQuotaReportsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListQuotaReportsOK creates a ListQuotaReportsOK with default headers values
func NewListQuotaReportsOK() *ListQuotaReportsOK {
	return &ListQuotaReportsOK{}
}

/*ListQuotaReportsOK handles this case with default header values.

List all or matching reports.
*/
type ListQuotaReportsOK struct {
	Payload *models.QuotaReports
}

func (o *ListQuotaReportsOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/quota/reports][%d] listQuotaReportsOK  %+v", 200, o.Payload)
}

func (o *ListQuotaReportsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QuotaReports)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListQuotaReportsDefault creates a ListQuotaReportsDefault with default headers values
func NewListQuotaReportsDefault(code int) *ListQuotaReportsDefault {
	return &ListQuotaReportsDefault{
		_statusCode: code,
	}
}

/*ListQuotaReportsDefault handles this case with default header values.

Unexpected error
*/
type ListQuotaReportsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list quota reports default response
func (o *ListQuotaReportsDefault) Code() int {
	return o._statusCode
}

func (o *ListQuotaReportsDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/quota/reports][%d] listQuotaReports default  %+v", o._statusCode, o.Payload)
}

func (o *ListQuotaReportsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}