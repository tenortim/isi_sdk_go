// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetSettingsKrb5RealmReader is a Reader for the GetSettingsKrb5Realm structure.
type GetSettingsKrb5RealmReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSettingsKrb5RealmReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSettingsKrb5RealmOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetSettingsKrb5RealmDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSettingsKrb5RealmOK creates a GetSettingsKrb5RealmOK with default headers values
func NewGetSettingsKrb5RealmOK() *GetSettingsKrb5RealmOK {
	return &GetSettingsKrb5RealmOK{}
}

/*GetSettingsKrb5RealmOK handles this case with default header values.

Retrieve the krb5 settings for realms.
*/
type GetSettingsKrb5RealmOK struct {
	Payload *models.SettingsKrb5Realms
}

func (o *GetSettingsKrb5RealmOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/auth/settings/krb5/realms/{SettingsKrb5RealmId}][%d] getSettingsKrb5RealmOK  %+v", 200, o.Payload)
}

func (o *GetSettingsKrb5RealmOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SettingsKrb5Realms)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSettingsKrb5RealmDefault creates a GetSettingsKrb5RealmDefault with default headers values
func NewGetSettingsKrb5RealmDefault(code int) *GetSettingsKrb5RealmDefault {
	return &GetSettingsKrb5RealmDefault{
		_statusCode: code,
	}
}

/*GetSettingsKrb5RealmDefault handles this case with default header values.

Unexpected error
*/
type GetSettingsKrb5RealmDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get settings krb5 realm default response
func (o *GetSettingsKrb5RealmDefault) Code() int {
	return o._statusCode
}

func (o *GetSettingsKrb5RealmDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/auth/settings/krb5/realms/{SettingsKrb5RealmId}][%d] getSettingsKrb5Realm default  %+v", o._statusCode, o.Payload)
}

func (o *GetSettingsKrb5RealmDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
