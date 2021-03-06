// Code generated by go-swagger; DO NOT EDIT.

package filepool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetFilepoolDefaultPolicyReader is a Reader for the GetFilepoolDefaultPolicy structure.
type GetFilepoolDefaultPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFilepoolDefaultPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFilepoolDefaultPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetFilepoolDefaultPolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFilepoolDefaultPolicyOK creates a GetFilepoolDefaultPolicyOK with default headers values
func NewGetFilepoolDefaultPolicyOK() *GetFilepoolDefaultPolicyOK {
	return &GetFilepoolDefaultPolicyOK{}
}

/*GetFilepoolDefaultPolicyOK handles this case with default header values.

List default file pool policy.
*/
type GetFilepoolDefaultPolicyOK struct {
	Payload *models.FilepoolDefaultPolicy
}

func (o *GetFilepoolDefaultPolicyOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/filepool/default-policy][%d] getFilepoolDefaultPolicyOK  %+v", 200, o.Payload)
}

func (o *GetFilepoolDefaultPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FilepoolDefaultPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilepoolDefaultPolicyDefault creates a GetFilepoolDefaultPolicyDefault with default headers values
func NewGetFilepoolDefaultPolicyDefault(code int) *GetFilepoolDefaultPolicyDefault {
	return &GetFilepoolDefaultPolicyDefault{
		_statusCode: code,
	}
}

/*GetFilepoolDefaultPolicyDefault handles this case with default header values.

Unexpected error
*/
type GetFilepoolDefaultPolicyDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get filepool default policy default response
func (o *GetFilepoolDefaultPolicyDefault) Code() int {
	return o._statusCode
}

func (o *GetFilepoolDefaultPolicyDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/filepool/default-policy][%d] getFilepoolDefaultPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *GetFilepoolDefaultPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
