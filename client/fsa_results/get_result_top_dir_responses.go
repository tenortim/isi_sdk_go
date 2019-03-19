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

// GetResultTopDirReader is a Reader for the GetResultTopDir structure.
type GetResultTopDirReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResultTopDirReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetResultTopDirOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetResultTopDirDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetResultTopDirOK creates a GetResultTopDirOK with default headers values
func NewGetResultTopDirOK() *GetResultTopDirOK {
	return &GetResultTopDirOK{}
}

/*GetResultTopDirOK handles this case with default header values.

This resource retrieves the top directories. ID in the resource path is the result set ID.
*/
type GetResultTopDirOK struct {
	Payload *models.ResultTopDirs
}

func (o *GetResultTopDirOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/fsa/results/{Id}/top-dirs/{ResultTopDirId}][%d] getResultTopDirOK  %+v", 200, o.Payload)
}

func (o *GetResultTopDirOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResultTopDirs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResultTopDirDefault creates a GetResultTopDirDefault with default headers values
func NewGetResultTopDirDefault(code int) *GetResultTopDirDefault {
	return &GetResultTopDirDefault{
		_statusCode: code,
	}
}

/*GetResultTopDirDefault handles this case with default header values.

Unexpected error
*/
type GetResultTopDirDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get result top dir default response
func (o *GetResultTopDirDefault) Code() int {
	return o._statusCode
}

func (o *GetResultTopDirDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/fsa/results/{Id}/top-dirs/{ResultTopDirId}][%d] getResultTopDir default  %+v", o._statusCode, o.Payload)
}

func (o *GetResultTopDirDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}