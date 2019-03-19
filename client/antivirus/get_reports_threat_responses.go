// Code generated by go-swagger; DO NOT EDIT.

package antivirus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetReportsThreatReader is a Reader for the GetReportsThreat structure.
type GetReportsThreatReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReportsThreatReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetReportsThreatOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetReportsThreatDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetReportsThreatOK creates a GetReportsThreatOK with default headers values
func NewGetReportsThreatOK() *GetReportsThreatOK {
	return &GetReportsThreatOK{}
}

/*GetReportsThreatOK handles this case with default header values.

Retrieve one antivirus threat report.
*/
type GetReportsThreatOK struct {
	Payload *models.ReportsThreats
}

func (o *GetReportsThreatOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/antivirus/reports/threats/{ReportsThreatId}][%d] getReportsThreatOK  %+v", 200, o.Payload)
}

func (o *GetReportsThreatOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReportsThreats)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportsThreatDefault creates a GetReportsThreatDefault with default headers values
func NewGetReportsThreatDefault(code int) *GetReportsThreatDefault {
	return &GetReportsThreatDefault{
		_statusCode: code,
	}
}

/*GetReportsThreatDefault handles this case with default header values.

Unexpected error
*/
type GetReportsThreatDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get reports threat default response
func (o *GetReportsThreatDefault) Code() int {
	return o._statusCode
}

func (o *GetReportsThreatDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/antivirus/reports/threats/{ReportsThreatId}][%d] getReportsThreat default  %+v", o._statusCode, o.Payload)
}

func (o *GetReportsThreatDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}