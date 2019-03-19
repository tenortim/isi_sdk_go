// Code generated by go-swagger; DO NOT EDIT.

package zones_summary

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetZonesSummaryZoneReader is a Reader for the GetZonesSummaryZone structure.
type GetZonesSummaryZoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetZonesSummaryZoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetZonesSummaryZoneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetZonesSummaryZoneDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetZonesSummaryZoneOK creates a GetZonesSummaryZoneOK with default headers values
func NewGetZonesSummaryZoneOK() *GetZonesSummaryZoneOK {
	return &GetZonesSummaryZoneOK{}
}

/*GetZonesSummaryZoneOK handles this case with default header values.

Retrieve non-privileged access zone information.
*/
type GetZonesSummaryZoneOK struct {
	Payload *models.ZonesSummary
}

func (o *GetZonesSummaryZoneOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/zones-summary/{ZonesSummaryZone}][%d] getZonesSummaryZoneOK  %+v", 200, o.Payload)
}

func (o *GetZonesSummaryZoneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZonesSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetZonesSummaryZoneDefault creates a GetZonesSummaryZoneDefault with default headers values
func NewGetZonesSummaryZoneDefault(code int) *GetZonesSummaryZoneDefault {
	return &GetZonesSummaryZoneDefault{
		_statusCode: code,
	}
}

/*GetZonesSummaryZoneDefault handles this case with default header values.

Unexpected error
*/
type GetZonesSummaryZoneDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get zones summary zone default response
func (o *GetZonesSummaryZoneDefault) Code() int {
	return o._statusCode
}

func (o *GetZonesSummaryZoneDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/zones-summary/{ZonesSummaryZone}][%d] getZonesSummaryZone default  %+v", o._statusCode, o.Payload)
}

func (o *GetZonesSummaryZoneDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}