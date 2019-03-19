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

// GetNdmpSettingsGlobalReader is a Reader for the GetNdmpSettingsGlobal structure.
type GetNdmpSettingsGlobalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNdmpSettingsGlobalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNdmpSettingsGlobalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetNdmpSettingsGlobalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNdmpSettingsGlobalOK creates a GetNdmpSettingsGlobalOK with default headers values
func NewGetNdmpSettingsGlobalOK() *GetNdmpSettingsGlobalOK {
	return &GetNdmpSettingsGlobalOK{}
}

/*GetNdmpSettingsGlobalOK handles this case with default header values.

List global ndmp settings.
*/
type GetNdmpSettingsGlobalOK struct {
	Payload *models.NdmpSettingsGlobal
}

func (o *GetNdmpSettingsGlobalOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/protocols/ndmp/settings/global][%d] getNdmpSettingsGlobalOK  %+v", 200, o.Payload)
}

func (o *GetNdmpSettingsGlobalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NdmpSettingsGlobal)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNdmpSettingsGlobalDefault creates a GetNdmpSettingsGlobalDefault with default headers values
func NewGetNdmpSettingsGlobalDefault(code int) *GetNdmpSettingsGlobalDefault {
	return &GetNdmpSettingsGlobalDefault{
		_statusCode: code,
	}
}

/*GetNdmpSettingsGlobalDefault handles this case with default header values.

Unexpected error
*/
type GetNdmpSettingsGlobalDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get ndmp settings global default response
func (o *GetNdmpSettingsGlobalDefault) Code() int {
	return o._statusCode
}

func (o *GetNdmpSettingsGlobalDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/protocols/ndmp/settings/global][%d] getNdmpSettingsGlobal default  %+v", o._statusCode, o.Payload)
}

func (o *GetNdmpSettingsGlobalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
