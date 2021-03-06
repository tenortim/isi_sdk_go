// Code generated by go-swagger; DO NOT EDIT.

package dedupe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetDedupeReportsReader is a Reader for the GetDedupeReports structure.
type GetDedupeReportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDedupeReportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDedupeReportsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetDedupeReportsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDedupeReportsOK creates a GetDedupeReportsOK with default headers values
func NewGetDedupeReportsOK() *GetDedupeReportsOK {
	return &GetDedupeReportsOK{}
}

/*GetDedupeReportsOK handles this case with default header values.

List dedupe reports.
*/
type GetDedupeReportsOK struct {
	Payload *models.DedupeReportsExtended
}

func (o *GetDedupeReportsOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/dedupe/reports][%d] getDedupeReportsOK  %+v", 200, o.Payload)
}

func (o *GetDedupeReportsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DedupeReportsExtended)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDedupeReportsDefault creates a GetDedupeReportsDefault with default headers values
func NewGetDedupeReportsDefault(code int) *GetDedupeReportsDefault {
	return &GetDedupeReportsDefault{
		_statusCode: code,
	}
}

/*GetDedupeReportsDefault handles this case with default header values.

Unexpected error
*/
type GetDedupeReportsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get dedupe reports default response
func (o *GetDedupeReportsDefault) Code() int {
	return o._statusCode
}

func (o *GetDedupeReportsDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/dedupe/reports][%d] getDedupeReports default  %+v", o._statusCode, o.Payload)
}

func (o *GetDedupeReportsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
