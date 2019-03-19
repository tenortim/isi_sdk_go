// Code generated by go-swagger; DO NOT EDIT.

package event

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetEventEventgroupOccurrencesReader is a Reader for the GetEventEventgroupOccurrences structure.
type GetEventEventgroupOccurrencesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventEventgroupOccurrencesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEventEventgroupOccurrencesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetEventEventgroupOccurrencesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEventEventgroupOccurrencesOK creates a GetEventEventgroupOccurrencesOK with default headers values
func NewGetEventEventgroupOccurrencesOK() *GetEventEventgroupOccurrencesOK {
	return &GetEventEventgroupOccurrencesOK{}
}

/*GetEventEventgroupOccurrencesOK handles this case with default header values.

List all eventgroup occurrences.
*/
type GetEventEventgroupOccurrencesOK struct {
	Payload *models.EventEventgroupOccurrencesExtended
}

func (o *GetEventEventgroupOccurrencesOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/event/eventgroup-occurrences][%d] getEventEventgroupOccurrencesOK  %+v", 200, o.Payload)
}

func (o *GetEventEventgroupOccurrencesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EventEventgroupOccurrencesExtended)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventEventgroupOccurrencesDefault creates a GetEventEventgroupOccurrencesDefault with default headers values
func NewGetEventEventgroupOccurrencesDefault(code int) *GetEventEventgroupOccurrencesDefault {
	return &GetEventEventgroupOccurrencesDefault{
		_statusCode: code,
	}
}

/*GetEventEventgroupOccurrencesDefault handles this case with default header values.

Unexpected error
*/
type GetEventEventgroupOccurrencesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get event eventgroup occurrences default response
func (o *GetEventEventgroupOccurrencesDefault) Code() int {
	return o._statusCode
}

func (o *GetEventEventgroupOccurrencesDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/event/eventgroup-occurrences][%d] getEventEventgroupOccurrences default  %+v", o._statusCode, o.Payload)
}

func (o *GetEventEventgroupOccurrencesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
