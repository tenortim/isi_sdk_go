// Code generated by go-swagger; DO NOT EDIT.

package snapshot_changelists

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetChangelistLinReader is a Reader for the GetChangelistLin structure.
type GetChangelistLinReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetChangelistLinReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetChangelistLinOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetChangelistLinDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetChangelistLinOK creates a GetChangelistLinOK with default headers values
func NewGetChangelistLinOK() *GetChangelistLinOK {
	return &GetChangelistLinOK{}
}

/*GetChangelistLinOK handles this case with default header values.

Get a single entry from the changelist.
*/
type GetChangelistLinOK struct {
	Payload *models.ChangelistLins
}

func (o *GetChangelistLinOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/snapshot/changelists/{Changelist}/lins/{ChangelistLinId}][%d] getChangelistLinOK  %+v", 200, o.Payload)
}

func (o *GetChangelistLinOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChangelistLins)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChangelistLinDefault creates a GetChangelistLinDefault with default headers values
func NewGetChangelistLinDefault(code int) *GetChangelistLinDefault {
	return &GetChangelistLinDefault{
		_statusCode: code,
	}
}

/*GetChangelistLinDefault handles this case with default header values.

Unexpected error
*/
type GetChangelistLinDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get changelist lin default response
func (o *GetChangelistLinDefault) Code() int {
	return o._statusCode
}

func (o *GetChangelistLinDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/snapshot/changelists/{Changelist}/lins/{ChangelistLinId}][%d] getChangelistLin default  %+v", o._statusCode, o.Payload)
}

func (o *GetChangelistLinDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
