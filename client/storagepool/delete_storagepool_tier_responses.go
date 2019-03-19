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

// DeleteStoragepoolTierReader is a Reader for the DeleteStoragepoolTier structure.
type DeleteStoragepoolTierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteStoragepoolTierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteStoragepoolTierNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteStoragepoolTierDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteStoragepoolTierNoContent creates a DeleteStoragepoolTierNoContent with default headers values
func NewDeleteStoragepoolTierNoContent() *DeleteStoragepoolTierNoContent {
	return &DeleteStoragepoolTierNoContent{}
}

/*DeleteStoragepoolTierNoContent handles this case with default header values.

Success.
*/
type DeleteStoragepoolTierNoContent struct {
}

func (o *DeleteStoragepoolTierNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/storagepool/tiers/{StoragepoolTierId}][%d] deleteStoragepoolTierNoContent ", 204)
}

func (o *DeleteStoragepoolTierNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteStoragepoolTierDefault creates a DeleteStoragepoolTierDefault with default headers values
func NewDeleteStoragepoolTierDefault(code int) *DeleteStoragepoolTierDefault {
	return &DeleteStoragepoolTierDefault{
		_statusCode: code,
	}
}

/*DeleteStoragepoolTierDefault handles this case with default header values.

Unexpected error
*/
type DeleteStoragepoolTierDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete storagepool tier default response
func (o *DeleteStoragepoolTierDefault) Code() int {
	return o._statusCode
}

func (o *DeleteStoragepoolTierDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/storagepool/tiers/{StoragepoolTierId}][%d] deleteStoragepoolTier default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteStoragepoolTierDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
