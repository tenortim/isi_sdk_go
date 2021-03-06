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

// ListStoragepoolNodepoolsReader is a Reader for the ListStoragepoolNodepools structure.
type ListStoragepoolNodepoolsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListStoragepoolNodepoolsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListStoragepoolNodepoolsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListStoragepoolNodepoolsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListStoragepoolNodepoolsOK creates a ListStoragepoolNodepoolsOK with default headers values
func NewListStoragepoolNodepoolsOK() *ListStoragepoolNodepoolsOK {
	return &ListStoragepoolNodepoolsOK{}
}

/*ListStoragepoolNodepoolsOK handles this case with default header values.

List all node pools.
*/
type ListStoragepoolNodepoolsOK struct {
	Payload *models.StoragepoolNodepoolsExtended
}

func (o *ListStoragepoolNodepoolsOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/storagepool/nodepools][%d] listStoragepoolNodepoolsOK  %+v", 200, o.Payload)
}

func (o *ListStoragepoolNodepoolsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoragepoolNodepoolsExtended)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListStoragepoolNodepoolsDefault creates a ListStoragepoolNodepoolsDefault with default headers values
func NewListStoragepoolNodepoolsDefault(code int) *ListStoragepoolNodepoolsDefault {
	return &ListStoragepoolNodepoolsDefault{
		_statusCode: code,
	}
}

/*ListStoragepoolNodepoolsDefault handles this case with default header values.

Unexpected error
*/
type ListStoragepoolNodepoolsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list storagepool nodepools default response
func (o *ListStoragepoolNodepoolsDefault) Code() int {
	return o._statusCode
}

func (o *ListStoragepoolNodepoolsDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/storagepool/nodepools][%d] listStoragepoolNodepools default  %+v", o._statusCode, o.Payload)
}

func (o *ListStoragepoolNodepoolsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
