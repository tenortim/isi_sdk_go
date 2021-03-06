// Code generated by go-swagger; DO NOT EDIT.

package protocols

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetNdmpLogsReader is a Reader for the GetNdmpLogs structure.
type GetNdmpLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNdmpLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNdmpLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetNdmpLogsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNdmpLogsOK creates a GetNdmpLogsOK with default headers values
func NewGetNdmpLogsOK() *GetNdmpLogsOK {
	return &GetNdmpLogsOK{}
}

/*GetNdmpLogsOK handles this case with default header values.

Get NDMP logs
*/
type GetNdmpLogsOK struct {
	Payload *models.NdmpLogs
}

func (o *GetNdmpLogsOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/protocols/ndmp/logs][%d] getNdmpLogsOK  %+v", 200, o.Payload)
}

func (o *GetNdmpLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NdmpLogs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNdmpLogsDefault creates a GetNdmpLogsDefault with default headers values
func NewGetNdmpLogsDefault(code int) *GetNdmpLogsDefault {
	return &GetNdmpLogsDefault{
		_statusCode: code,
	}
}

/*GetNdmpLogsDefault handles this case with default header values.

Unexpected error
*/
type GetNdmpLogsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get ndmp logs default response
func (o *GetNdmpLogsDefault) Code() int {
	return o._statusCode
}

func (o *GetNdmpLogsDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/protocols/ndmp/logs][%d] getNdmpLogs default  %+v", o._statusCode, o.Payload)
}

func (o *GetNdmpLogsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
