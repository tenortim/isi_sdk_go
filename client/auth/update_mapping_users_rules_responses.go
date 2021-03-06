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

// UpdateMappingUsersRulesReader is a Reader for the UpdateMappingUsersRules structure.
type UpdateMappingUsersRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMappingUsersRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateMappingUsersRulesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateMappingUsersRulesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateMappingUsersRulesNoContent creates a UpdateMappingUsersRulesNoContent with default headers values
func NewUpdateMappingUsersRulesNoContent() *UpdateMappingUsersRulesNoContent {
	return &UpdateMappingUsersRulesNoContent{}
}

/*UpdateMappingUsersRulesNoContent handles this case with default header values.

Success.
*/
type UpdateMappingUsersRulesNoContent struct {
}

func (o *UpdateMappingUsersRulesNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/1/auth/mapping/users/rules][%d] updateMappingUsersRulesNoContent ", 204)
}

func (o *UpdateMappingUsersRulesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMappingUsersRulesDefault creates a UpdateMappingUsersRulesDefault with default headers values
func NewUpdateMappingUsersRulesDefault(code int) *UpdateMappingUsersRulesDefault {
	return &UpdateMappingUsersRulesDefault{
		_statusCode: code,
	}
}

/*UpdateMappingUsersRulesDefault handles this case with default header values.

Unexpected error
*/
type UpdateMappingUsersRulesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update mapping users rules default response
func (o *UpdateMappingUsersRulesDefault) Code() int {
	return o._statusCode
}

func (o *UpdateMappingUsersRulesDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/1/auth/mapping/users/rules][%d] updateMappingUsersRules default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateMappingUsersRulesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
