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

// GetAntivirusServerReader is a Reader for the GetAntivirusServer structure.
type GetAntivirusServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAntivirusServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAntivirusServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetAntivirusServerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAntivirusServerOK creates a GetAntivirusServerOK with default headers values
func NewGetAntivirusServerOK() *GetAntivirusServerOK {
	return &GetAntivirusServerOK{}
}

/*GetAntivirusServerOK handles this case with default header values.

Retrieve one antivirus server entry.
*/
type GetAntivirusServerOK struct {
	Payload *models.AntivirusServers
}

func (o *GetAntivirusServerOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/antivirus/servers/{AntivirusServerId}][%d] getAntivirusServerOK  %+v", 200, o.Payload)
}

func (o *GetAntivirusServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AntivirusServers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAntivirusServerDefault creates a GetAntivirusServerDefault with default headers values
func NewGetAntivirusServerDefault(code int) *GetAntivirusServerDefault {
	return &GetAntivirusServerDefault{
		_statusCode: code,
	}
}

/*GetAntivirusServerDefault handles this case with default header values.

Unexpected error
*/
type GetAntivirusServerDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get antivirus server default response
func (o *GetAntivirusServerDefault) Code() int {
	return o._statusCode
}

func (o *GetAntivirusServerDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/antivirus/servers/{AntivirusServerId}][%d] getAntivirusServer default  %+v", o._statusCode, o.Payload)
}

func (o *GetAntivirusServerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}