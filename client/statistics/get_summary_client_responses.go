// Code generated by go-swagger; DO NOT EDIT.

package statistics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetSummaryClientReader is a Reader for the GetSummaryClient structure.
type GetSummaryClientReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSummaryClientReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSummaryClientOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetSummaryClientDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSummaryClientOK creates a GetSummaryClientOK with default headers values
func NewGetSummaryClientOK() *GetSummaryClientOK {
	return &GetSummaryClientOK{}
}

/*GetSummaryClientOK handles this case with default header values.

GetSummaryClientOK get summary client o k
*/
type GetSummaryClientOK struct {
	Payload *models.SummaryClient
}

func (o *GetSummaryClientOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/statistics/summary/client][%d] getSummaryClientOK  %+v", 200, o.Payload)
}

func (o *GetSummaryClientOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SummaryClient)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSummaryClientDefault creates a GetSummaryClientDefault with default headers values
func NewGetSummaryClientDefault(code int) *GetSummaryClientDefault {
	return &GetSummaryClientDefault{
		_statusCode: code,
	}
}

/*GetSummaryClientDefault handles this case with default header values.

Unexpected error
*/
type GetSummaryClientDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get summary client default response
func (o *GetSummaryClientDefault) Code() int {
	return o._statusCode
}

func (o *GetSummaryClientDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/statistics/summary/client][%d] getSummaryClient default  %+v", o._statusCode, o.Payload)
}

func (o *GetSummaryClientDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
