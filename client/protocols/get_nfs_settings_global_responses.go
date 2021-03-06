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

// GetNfsSettingsGlobalReader is a Reader for the GetNfsSettingsGlobal structure.
type GetNfsSettingsGlobalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNfsSettingsGlobalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNfsSettingsGlobalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetNfsSettingsGlobalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNfsSettingsGlobalOK creates a GetNfsSettingsGlobalOK with default headers values
func NewGetNfsSettingsGlobalOK() *GetNfsSettingsGlobalOK {
	return &GetNfsSettingsGlobalOK{}
}

/*GetNfsSettingsGlobalOK handles this case with default header values.

Retrieve the NFS configuration.
*/
type GetNfsSettingsGlobalOK struct {
	Payload *models.NfsSettingsGlobal
}

func (o *GetNfsSettingsGlobalOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/protocols/nfs/settings/global][%d] getNfsSettingsGlobalOK  %+v", 200, o.Payload)
}

func (o *GetNfsSettingsGlobalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NfsSettingsGlobal)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNfsSettingsGlobalDefault creates a GetNfsSettingsGlobalDefault with default headers values
func NewGetNfsSettingsGlobalDefault(code int) *GetNfsSettingsGlobalDefault {
	return &GetNfsSettingsGlobalDefault{
		_statusCode: code,
	}
}

/*GetNfsSettingsGlobalDefault handles this case with default header values.

Unexpected error
*/
type GetNfsSettingsGlobalDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get nfs settings global default response
func (o *GetNfsSettingsGlobalDefault) Code() int {
	return o._statusCode
}

func (o *GetNfsSettingsGlobalDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/protocols/nfs/settings/global][%d] getNfsSettingsGlobal default  %+v", o._statusCode, o.Payload)
}

func (o *GetNfsSettingsGlobalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
