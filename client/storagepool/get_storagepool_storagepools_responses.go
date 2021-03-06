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

// GetStoragepoolStoragepoolsReader is a Reader for the GetStoragepoolStoragepools structure.
type GetStoragepoolStoragepoolsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStoragepoolStoragepoolsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetStoragepoolStoragepoolsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetStoragepoolStoragepoolsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStoragepoolStoragepoolsOK creates a GetStoragepoolStoragepoolsOK with default headers values
func NewGetStoragepoolStoragepoolsOK() *GetStoragepoolStoragepoolsOK {
	return &GetStoragepoolStoragepoolsOK{}
}

/*GetStoragepoolStoragepoolsOK handles this case with default header values.

List all storage pools.
*/
type GetStoragepoolStoragepoolsOK struct {
	Payload *models.StoragepoolStoragepools
}

func (o *GetStoragepoolStoragepoolsOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/storagepool/storagepools][%d] getStoragepoolStoragepoolsOK  %+v", 200, o.Payload)
}

func (o *GetStoragepoolStoragepoolsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StoragepoolStoragepools)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStoragepoolStoragepoolsDefault creates a GetStoragepoolStoragepoolsDefault with default headers values
func NewGetStoragepoolStoragepoolsDefault(code int) *GetStoragepoolStoragepoolsDefault {
	return &GetStoragepoolStoragepoolsDefault{
		_statusCode: code,
	}
}

/*GetStoragepoolStoragepoolsDefault handles this case with default header values.

Unexpected error
*/
type GetStoragepoolStoragepoolsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get storagepool storagepools default response
func (o *GetStoragepoolStoragepoolsDefault) Code() int {
	return o._statusCode
}

func (o *GetStoragepoolStoragepoolsDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/storagepool/storagepools][%d] getStoragepoolStoragepools default  %+v", o._statusCode, o.Payload)
}

func (o *GetStoragepoolStoragepoolsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
