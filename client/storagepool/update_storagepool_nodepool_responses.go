// Code generated by go-swagger; DO NOT EDIT.

package storagepool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// UpdateStoragepoolNodepoolReader is a Reader for the UpdateStoragepoolNodepool structure.
type UpdateStoragepoolNodepoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateStoragepoolNodepoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateStoragepoolNodepoolNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateStoragepoolNodepoolDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateStoragepoolNodepoolNoContent creates a UpdateStoragepoolNodepoolNoContent with default headers values
func NewUpdateStoragepoolNodepoolNoContent() *UpdateStoragepoolNodepoolNoContent {
	return &UpdateStoragepoolNodepoolNoContent{}
}

/*UpdateStoragepoolNodepoolNoContent handles this case with default header values.

Success.
*/
type UpdateStoragepoolNodepoolNoContent struct {
}

func (o *UpdateStoragepoolNodepoolNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/3/storagepool/nodepools/{StoragepoolNodepoolId}][%d] updateStoragepoolNodepoolNoContent ", 204)
}

func (o *UpdateStoragepoolNodepoolNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateStoragepoolNodepoolDefault creates a UpdateStoragepoolNodepoolDefault with default headers values
func NewUpdateStoragepoolNodepoolDefault(code int) *UpdateStoragepoolNodepoolDefault {
	return &UpdateStoragepoolNodepoolDefault{
		_statusCode: code,
	}
}

/*UpdateStoragepoolNodepoolDefault handles this case with default header values.

Unexpected error
*/
type UpdateStoragepoolNodepoolDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update storagepool nodepool default response
func (o *UpdateStoragepoolNodepoolDefault) Code() int {
	return o._statusCode
}

func (o *UpdateStoragepoolNodepoolDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/3/storagepool/nodepools/{StoragepoolNodepoolId}][%d] updateStoragepoolNodepool default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateStoragepoolNodepoolDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}