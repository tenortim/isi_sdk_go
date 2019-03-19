// Code generated by go-swagger; DO NOT EDIT.

package cloud

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// ListCloudAccountsReader is a Reader for the ListCloudAccounts structure.
type ListCloudAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCloudAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListCloudAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListCloudAccountsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListCloudAccountsOK creates a ListCloudAccountsOK with default headers values
func NewListCloudAccountsOK() *ListCloudAccountsOK {
	return &ListCloudAccountsOK{}
}

/*ListCloudAccountsOK handles this case with default header values.

List all accounts.
*/
type ListCloudAccountsOK struct {
	Payload *models.CloudAccountsExtended
}

func (o *ListCloudAccountsOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/cloud/accounts][%d] listCloudAccountsOK  %+v", 200, o.Payload)
}

func (o *ListCloudAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudAccountsExtended)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCloudAccountsDefault creates a ListCloudAccountsDefault with default headers values
func NewListCloudAccountsDefault(code int) *ListCloudAccountsDefault {
	return &ListCloudAccountsDefault{
		_statusCode: code,
	}
}

/*ListCloudAccountsDefault handles this case with default header values.

Unexpected error
*/
type ListCloudAccountsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list cloud accounts default response
func (o *ListCloudAccountsDefault) Code() int {
	return o._statusCode
}

func (o *ListCloudAccountsDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/cloud/accounts][%d] listCloudAccounts default  %+v", o._statusCode, o.Payload)
}

func (o *ListCloudAccountsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
