// Code generated by go-swagger; DO NOT EDIT.

package protocols

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetNfsCheckReader is a Reader for the GetNfsCheck structure.
type GetNfsCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNfsCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNfsCheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetNfsCheckDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNfsCheckOK creates a GetNfsCheckOK with default headers values
func NewGetNfsCheckOK() *GetNfsCheckOK {
	return &GetNfsCheckOK{}
}

/*GetNfsCheckOK handles this case with default header values.

Retrieve NFS export validation information.
*/
type GetNfsCheckOK struct {
	Payload *models.NfsCheckExtended
}

func (o *GetNfsCheckOK) Error() string {
	return fmt.Sprintf("[GET /platform/2/protocols/nfs/check][%d] getNfsCheckOK  %+v", 200, o.Payload)
}

func (o *GetNfsCheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NfsCheckExtended)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNfsCheckDefault creates a GetNfsCheckDefault with default headers values
func NewGetNfsCheckDefault(code int) *GetNfsCheckDefault {
	return &GetNfsCheckDefault{
		_statusCode: code,
	}
}

/*GetNfsCheckDefault handles this case with default header values.

Unexpected error
*/
type GetNfsCheckDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get nfs check default response
func (o *GetNfsCheckDefault) Code() int {
	return o._statusCode
}

func (o *GetNfsCheckDefault) Error() string {
	return fmt.Sprintf("[GET /platform/2/protocols/nfs/check][%d] getNfsCheck default  %+v", o._statusCode, o.Payload)
}

func (o *GetNfsCheckDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
