// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// CreateProvidersAdsItemReader is a Reader for the CreateProvidersAdsItem structure.
type CreateProvidersAdsItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProvidersAdsItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateProvidersAdsItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateProvidersAdsItemDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateProvidersAdsItemOK creates a CreateProvidersAdsItemOK with default headers values
func NewCreateProvidersAdsItemOK() *CreateProvidersAdsItemOK {
	return &CreateProvidersAdsItemOK{}
}

/*CreateProvidersAdsItemOK handles this case with default header values.

Create a new ADS provider.
*/
type CreateProvidersAdsItemOK struct {
	Payload *models.CreateResponse
}

func (o *CreateProvidersAdsItemOK) Error() string {
	return fmt.Sprintf("[POST /platform/3/auth/providers/ads][%d] createProvidersAdsItemOK  %+v", 200, o.Payload)
}

func (o *CreateProvidersAdsItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProvidersAdsItemDefault creates a CreateProvidersAdsItemDefault with default headers values
func NewCreateProvidersAdsItemDefault(code int) *CreateProvidersAdsItemDefault {
	return &CreateProvidersAdsItemDefault{
		_statusCode: code,
	}
}

/*CreateProvidersAdsItemDefault handles this case with default header values.

Unexpected error
*/
type CreateProvidersAdsItemDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create providers ads item default response
func (o *CreateProvidersAdsItemDefault) Code() int {
	return o._statusCode
}

func (o *CreateProvidersAdsItemDefault) Error() string {
	return fmt.Sprintf("[POST /platform/3/auth/providers/ads][%d] createProvidersAdsItem default  %+v", o._statusCode, o.Payload)
}

func (o *CreateProvidersAdsItemDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
