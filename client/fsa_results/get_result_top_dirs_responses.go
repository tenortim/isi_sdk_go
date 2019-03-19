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

// GetResultTopDirsReader is a Reader for the GetResultTopDirs structure.
type GetResultTopDirsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResultTopDirsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetResultTopDirsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetResultTopDirsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetResultTopDirsOK creates a GetResultTopDirsOK with default headers values
func NewGetResultTopDirsOK() *GetResultTopDirsOK {
	return &GetResultTopDirsOK{}
}

/*GetResultTopDirsOK handles this case with default header values.

This resource retrieves the top directories. ID in the resource path is the result set ID.
*/
type GetResultTopDirsOK struct {
	Payload *models.ResultTopDirs
}

func (o *GetResultTopDirsOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/fsa/results/{Id}/top-dirs][%d] getResultTopDirsOK  %+v", 200, o.Payload)
}

func (o *GetResultTopDirsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResultTopDirs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResultTopDirsDefault creates a GetResultTopDirsDefault with default headers values
func NewGetResultTopDirsDefault(code int) *GetResultTopDirsDefault {
	return &GetResultTopDirsDefault{
		_statusCode: code,
	}
}

/*GetResultTopDirsDefault handles this case with default header values.

Unexpected error
*/
type GetResultTopDirsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get result top dirs default response
func (o *GetResultTopDirsDefault) Code() int {
	return o._statusCode
}

func (o *GetResultTopDirsDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/fsa/results/{Id}/top-dirs][%d] getResultTopDirs default  %+v", o._statusCode, o.Payload)
}

func (o *GetResultTopDirsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
