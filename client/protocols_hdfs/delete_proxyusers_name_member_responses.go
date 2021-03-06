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

// DeleteProxyusersNameMemberReader is a Reader for the DeleteProxyusersNameMember structure.
type DeleteProxyusersNameMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProxyusersNameMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteProxyusersNameMemberNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteProxyusersNameMemberDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteProxyusersNameMemberNoContent creates a DeleteProxyusersNameMemberNoContent with default headers values
func NewDeleteProxyusersNameMemberNoContent() *DeleteProxyusersNameMemberNoContent {
	return &DeleteProxyusersNameMemberNoContent{}
}

/*DeleteProxyusersNameMemberNoContent handles this case with default header values.

Success.
*/
type DeleteProxyusersNameMemberNoContent struct {
}

func (o *DeleteProxyusersNameMemberNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/protocols/hdfs/proxyusers/{Name}/members/{ProxyusersNameMemberId}][%d] deleteProxyusersNameMemberNoContent ", 204)
}

func (o *DeleteProxyusersNameMemberNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProxyusersNameMemberDefault creates a DeleteProxyusersNameMemberDefault with default headers values
func NewDeleteProxyusersNameMemberDefault(code int) *DeleteProxyusersNameMemberDefault {
	return &DeleteProxyusersNameMemberDefault{
		_statusCode: code,
	}
}

/*DeleteProxyusersNameMemberDefault handles this case with default header values.

Unexpected error
*/
type DeleteProxyusersNameMemberDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete proxyusers name member default response
func (o *DeleteProxyusersNameMemberDefault) Code() int {
	return o._statusCode
}

func (o *DeleteProxyusersNameMemberDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/1/protocols/hdfs/proxyusers/{Name}/members/{ProxyusersNameMemberId}][%d] deleteProxyusersNameMember default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteProxyusersNameMemberDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
