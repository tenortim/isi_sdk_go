// Code generated by go-swagger; DO NOT EDIT.

package hardening

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetHardeningStateReader is a Reader for the GetHardeningState structure.
type GetHardeningStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHardeningStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetHardeningStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetHardeningStateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHardeningStateOK creates a GetHardeningStateOK with default headers values
func NewGetHardeningStateOK() *GetHardeningStateOK {
	return &GetHardeningStateOK{}
}

/*GetHardeningStateOK handles this case with default header values.

Get the state of the current hardening operation, if one is happening.  Note that this is different from the /status resource, which returns the overall hardening status of the cluster.
*/
type GetHardeningStateOK struct {
	Payload *models.HardeningState
}

func (o *GetHardeningStateOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/hardening/state][%d] getHardeningStateOK  %+v", 200, o.Payload)
}

func (o *GetHardeningStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HardeningState)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHardeningStateDefault creates a GetHardeningStateDefault with default headers values
func NewGetHardeningStateDefault(code int) *GetHardeningStateDefault {
	return &GetHardeningStateDefault{
		_statusCode: code,
	}
}

/*GetHardeningStateDefault handles this case with default header values.

Unexpected error
*/
type GetHardeningStateDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get hardening state default response
func (o *GetHardeningStateDefault) Code() int {
	return o._statusCode
}

func (o *GetHardeningStateDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/hardening/state][%d] getHardeningState default  %+v", o._statusCode, o.Payload)
}

func (o *GetHardeningStateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
