// Code generated by go-swagger; DO NOT EDIT.

package fsa_results

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetResultHistogramStatReader is a Reader for the GetResultHistogramStat structure.
type GetResultHistogramStatReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResultHistogramStatReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetResultHistogramStatOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetResultHistogramStatDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetResultHistogramStatOK creates a GetResultHistogramStatOK with default headers values
func NewGetResultHistogramStatOK() *GetResultHistogramStatOK {
	return &GetResultHistogramStatOK{}
}

/*GetResultHistogramStatOK handles this case with default header values.

This resource retrieves a histogram of file counts for an individual FSA result set. ID in the resource path is the result set ID.
*/
type GetResultHistogramStatOK struct {
	Payload *models.ResultHistogram
}

func (o *GetResultHistogramStatOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/fsa/results/{Id}/histogram/{ResultHistogramStat}][%d] getResultHistogramStatOK  %+v", 200, o.Payload)
}

func (o *GetResultHistogramStatOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResultHistogram)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResultHistogramStatDefault creates a GetResultHistogramStatDefault with default headers values
func NewGetResultHistogramStatDefault(code int) *GetResultHistogramStatDefault {
	return &GetResultHistogramStatDefault{
		_statusCode: code,
	}
}

/*GetResultHistogramStatDefault handles this case with default header values.

Unexpected error
*/
type GetResultHistogramStatDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get result histogram stat default response
func (o *GetResultHistogramStatDefault) Code() int {
	return o._statusCode
}

func (o *GetResultHistogramStatDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/fsa/results/{Id}/histogram/{ResultHistogramStat}][%d] getResultHistogramStat default  %+v", o._statusCode, o.Payload)
}

func (o *GetResultHistogramStatDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
