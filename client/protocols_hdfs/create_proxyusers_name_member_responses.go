// Code generated by go-swagger; DO NOT EDIT.

package protocols_hdfs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// CreateProxyusersNameMemberReader is a Reader for the CreateProxyusersNameMember structure.
type CreateProxyusersNameMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProxyusersNameMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateProxyusersNameMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateProxyusersNameMemberDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateProxyusersNameMemberOK creates a CreateProxyusersNameMemberOK with default headers values
func NewCreateProxyusersNameMemberOK() *CreateProxyusersNameMemberOK {
	return &CreateProxyusersNameMemberOK{}
}

/*CreateProxyusersNameMemberOK handles this case with default header values.

Add a member to the HDFS proxyuser.
*/
type CreateProxyusersNameMemberOK struct {
	Payload models.Empty
}

func (o *CreateProxyusersNameMemberOK) Error() string {
	return fmt.Sprintf("[POST /platform/1/protocols/hdfs/proxyusers/{Name}/members][%d] createProxyusersNameMemberOK  %+v", 200, o.Payload)
}

func (o *CreateProxyusersNameMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProxyusersNameMemberDefault creates a CreateProxyusersNameMemberDefault with default headers values
func NewCreateProxyusersNameMemberDefault(code int) *CreateProxyusersNameMemberDefault {
	return &CreateProxyusersNameMemberDefault{
		_statusCode: code,
	}
}

/*CreateProxyusersNameMemberDefault handles this case with default header values.

Unexpected error
*/
type CreateProxyusersNameMemberDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create proxyusers name member default response
func (o *CreateProxyusersNameMemberDefault) Code() int {
	return o._statusCode
}

func (o *CreateProxyusersNameMemberDefault) Error() string {
	return fmt.Sprintf("[POST /platform/1/protocols/hdfs/proxyusers/{Name}/members][%d] createProxyusersNameMember default  %+v", o._statusCode, o.Payload)
}

func (o *CreateProxyusersNameMemberDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}