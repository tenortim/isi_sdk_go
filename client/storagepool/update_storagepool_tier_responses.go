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

// UpdateStoragepoolTierReader is a Reader for the UpdateStoragepoolTier structure.
type UpdateStoragepoolTierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateStoragepoolTierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateStoragepoolTierNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateStoragepoolTierDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateStoragepoolTierNoContent creates a UpdateStoragepoolTierNoContent with default headers values
func NewUpdateStoragepoolTierNoContent() *UpdateStoragepoolTierNoContent {
	return &UpdateStoragepoolTierNoContent{}
}

/*UpdateStoragepoolTierNoContent handles this case with default header values.

Success.
*/
type UpdateStoragepoolTierNoContent struct {
}

func (o *UpdateStoragepoolTierNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/1/storagepool/tiers/{StoragepoolTierId}][%d] updateStoragepoolTierNoContent ", 204)
}

func (o *UpdateStoragepoolTierNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateStoragepoolTierDefault creates a UpdateStoragepoolTierDefault with default headers values
func NewUpdateStoragepoolTierDefault(code int) *UpdateStoragepoolTierDefault {
	return &UpdateStoragepoolTierDefault{
		_statusCode: code,
	}
}

/*UpdateStoragepoolTierDefault handles this case with default header values.

Unexpected error
*/
type UpdateStoragepoolTierDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update storagepool tier default response
func (o *UpdateStoragepoolTierDefault) Code() int {
	return o._statusCode
}

func (o *UpdateStoragepoolTierDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/1/storagepool/tiers/{StoragepoolTierId}][%d] updateStoragepoolTier default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateStoragepoolTierDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}