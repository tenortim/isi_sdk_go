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

// GetStoragepoolNodepoolReader is a Reader for the GetStoragepoolNodepool structure.
type GetStoragepoolNodepoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStoragepoolNodepoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetStoragepoolNodepoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetStoragepoolNodepoolDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStoragepoolNodepoolOK creates a GetStoragepoolNodepoolOK with default headers values
func NewGetStoragepoolNodepoolOK() *GetStoragepoolNodepoolOK {
	return &GetStoragepoolNodepoolOK{}
}

/*GetStoragepoolNodepoolOK handles this case with default header values.

Retrieve node pool information.
*/
type GetStoragepoolNodepoolOK struct {
	Payload *models.StoragepoolNodepools
}

func (o *GetStoragepoolNodepoolOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/storagepool/nodepools/{StoragepoolNodepoolId}][%d] getStoragepoolNodepoolOK  %+v", 200, o.Payload)
}

func (o *GetStoragepoolNodepoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoragepoolNodepools)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStoragepoolNodepoolDefault creates a GetStoragepoolNodepoolDefault with default headers values
func NewGetStoragepoolNodepoolDefault(code int) *GetStoragepoolNodepoolDefault {
	return &GetStoragepoolNodepoolDefault{
		_statusCode: code,
	}
}

/*GetStoragepoolNodepoolDefault handles this case with default header values.

Unexpected error
*/
type GetStoragepoolNodepoolDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get storagepool nodepool default response
func (o *GetStoragepoolNodepoolDefault) Code() int {
	return o._statusCode
}

func (o *GetStoragepoolNodepoolDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/storagepool/nodepools/{StoragepoolNodepoolId}][%d] getStoragepoolNodepool default  %+v", o._statusCode, o.Payload)
}

func (o *GetStoragepoolNodepoolDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
