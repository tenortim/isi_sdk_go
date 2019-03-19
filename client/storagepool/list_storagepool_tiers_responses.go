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

// ListStoragepoolTiersReader is a Reader for the ListStoragepoolTiers structure.
type ListStoragepoolTiersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListStoragepoolTiersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListStoragepoolTiersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListStoragepoolTiersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListStoragepoolTiersOK creates a ListStoragepoolTiersOK with default headers values
func NewListStoragepoolTiersOK() *ListStoragepoolTiersOK {
	return &ListStoragepoolTiersOK{}
}

/*ListStoragepoolTiersOK handles this case with default header values.

List all tiers.
*/
type ListStoragepoolTiersOK struct {
	Payload *models.StoragepoolTiersExtended
}

func (o *ListStoragepoolTiersOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/storagepool/tiers][%d] listStoragepoolTiersOK  %+v", 200, o.Payload)
}

func (o *ListStoragepoolTiersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoragepoolTiersExtended)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListStoragepoolTiersDefault creates a ListStoragepoolTiersDefault with default headers values
func NewListStoragepoolTiersDefault(code int) *ListStoragepoolTiersDefault {
	return &ListStoragepoolTiersDefault{
		_statusCode: code,
	}
}

/*ListStoragepoolTiersDefault handles this case with default header values.

Unexpected error
*/
type ListStoragepoolTiersDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list storagepool tiers default response
func (o *ListStoragepoolTiersDefault) Code() int {
	return o._statusCode
}

func (o *ListStoragepoolTiersDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/storagepool/tiers][%d] listStoragepoolTiers default  %+v", o._statusCode, o.Payload)
}

func (o *ListStoragepoolTiersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}