// Code generated by go-swagger; DO NOT EDIT.

package fsa

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetFsaResultsReader is a Reader for the GetFsaResults structure.
type GetFsaResultsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFsaResultsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFsaResultsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetFsaResultsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFsaResultsOK creates a GetFsaResultsOK with default headers values
func NewGetFsaResultsOK() *GetFsaResultsOK {
	return &GetFsaResultsOK{}
}

/*GetFsaResultsOK handles this case with default header values.

List all results.
*/
type GetFsaResultsOK struct {
	Payload *models.FsaResultsExtended
}

func (o *GetFsaResultsOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/fsa/results][%d] getFsaResultsOK  %+v", 200, o.Payload)
}

func (o *GetFsaResultsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FsaResultsExtended)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFsaResultsDefault creates a GetFsaResultsDefault with default headers values
func NewGetFsaResultsDefault(code int) *GetFsaResultsDefault {
	return &GetFsaResultsDefault{
		_statusCode: code,
	}
}

/*GetFsaResultsDefault handles this case with default header values.

Unexpected error
*/
type GetFsaResultsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get fsa results default response
func (o *GetFsaResultsDefault) Code() int {
	return o._statusCode
}

func (o *GetFsaResultsDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/fsa/results][%d] getFsaResults default  %+v", o._statusCode, o.Payload)
}

func (o *GetFsaResultsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}