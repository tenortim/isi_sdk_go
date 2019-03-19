// Code generated by go-swagger; DO NOT EDIT.

package storagepool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// ListCompatibilitiesClassActiveReader is a Reader for the ListCompatibilitiesClassActive structure.
type ListCompatibilitiesClassActiveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCompatibilitiesClassActiveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListCompatibilitiesClassActiveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListCompatibilitiesClassActiveDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListCompatibilitiesClassActiveOK creates a ListCompatibilitiesClassActiveOK with default headers values
func NewListCompatibilitiesClassActiveOK() *ListCompatibilitiesClassActiveOK {
	return &ListCompatibilitiesClassActiveOK{}
}

/*ListCompatibilitiesClassActiveOK handles this case with default header values.

Get a list of active compatibilities
*/
type ListCompatibilitiesClassActiveOK struct {
	Payload *models.CompatibilitiesClassActiveExtended
}

func (o *ListCompatibilitiesClassActiveOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/storagepool/compatibilities/class/active][%d] listCompatibilitiesClassActiveOK  %+v", 200, o.Payload)
}

func (o *ListCompatibilitiesClassActiveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CompatibilitiesClassActiveExtended)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCompatibilitiesClassActiveDefault creates a ListCompatibilitiesClassActiveDefault with default headers values
func NewListCompatibilitiesClassActiveDefault(code int) *ListCompatibilitiesClassActiveDefault {
	return &ListCompatibilitiesClassActiveDefault{
		_statusCode: code,
	}
}

/*ListCompatibilitiesClassActiveDefault handles this case with default header values.

Unexpected error
*/
type ListCompatibilitiesClassActiveDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list compatibilities class active default response
func (o *ListCompatibilitiesClassActiveDefault) Code() int {
	return o._statusCode
}

func (o *ListCompatibilitiesClassActiveDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/storagepool/compatibilities/class/active][%d] listCompatibilitiesClassActive default  %+v", o._statusCode, o.Payload)
}

func (o *ListCompatibilitiesClassActiveDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
