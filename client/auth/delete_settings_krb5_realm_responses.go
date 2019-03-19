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

// DeleteSettingsKrb5RealmReader is a Reader for the DeleteSettingsKrb5Realm structure.
type DeleteSettingsKrb5RealmReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSettingsKrb5RealmReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSettingsKrb5RealmNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteSettingsKrb5RealmDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSettingsKrb5RealmNoContent creates a DeleteSettingsKrb5RealmNoContent with default headers values
func NewDeleteSettingsKrb5RealmNoContent() *DeleteSettingsKrb5RealmNoContent {
	return &DeleteSettingsKrb5RealmNoContent{}
}

/*DeleteSettingsKrb5RealmNoContent handles this case with default header values.

Success.
*/
type DeleteSettingsKrb5RealmNoContent struct {
}

func (o *DeleteSettingsKrb5RealmNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/auth/settings/krb5/realms/{SettingsKrb5RealmId}][%d] deleteSettingsKrb5RealmNoContent ", 204)
}

func (o *DeleteSettingsKrb5RealmNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSettingsKrb5RealmDefault creates a DeleteSettingsKrb5RealmDefault with default headers values
func NewDeleteSettingsKrb5RealmDefault(code int) *DeleteSettingsKrb5RealmDefault {
	return &DeleteSettingsKrb5RealmDefault{
		_statusCode: code,
	}
}

/*DeleteSettingsKrb5RealmDefault handles this case with default header values.

Unexpected error
*/
type DeleteSettingsKrb5RealmDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete settings krb5 realm default response
func (o *DeleteSettingsKrb5RealmDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSettingsKrb5RealmDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/auth/settings/krb5/realms/{SettingsKrb5RealmId}][%d] deleteSettingsKrb5Realm default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSettingsKrb5RealmDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
