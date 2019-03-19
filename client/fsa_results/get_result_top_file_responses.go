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

// GetResultTopFileReader is a Reader for the GetResultTopFile structure.
type GetResultTopFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResultTopFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetResultTopFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetResultTopFileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetResultTopFileOK creates a GetResultTopFileOK with default headers values
func NewGetResultTopFileOK() *GetResultTopFileOK {
	return &GetResultTopFileOK{}
}

/*GetResultTopFileOK handles this case with default header values.

This resource retrieves the top files. ID in the resource path is the result set ID.
*/
type GetResultTopFileOK struct {
	Payload *models.ResultTopFiles
}

func (o *GetResultTopFileOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/fsa/results/{Id}/top-files/{ResultTopFileId}][%d] getResultTopFileOK  %+v", 200, o.Payload)
}

func (o *GetResultTopFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResultTopFiles)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResultTopFileDefault creates a GetResultTopFileDefault with default headers values
func NewGetResultTopFileDefault(code int) *GetResultTopFileDefault {
	return &GetResultTopFileDefault{
		_statusCode: code,
	}
}

/*GetResultTopFileDefault handles this case with default header values.

Unexpected error
*/
type GetResultTopFileDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get result top file default response
func (o *GetResultTopFileDefault) Code() int {
	return o._statusCode
}

func (o *GetResultTopFileDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/fsa/results/{Id}/top-files/{ResultTopFileId}][%d] getResultTopFile default  %+v", o._statusCode, o.Payload)
}

func (o *GetResultTopFileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
