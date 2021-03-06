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

// UpdateHdfsLogLevelReader is a Reader for the UpdateHdfsLogLevel structure.
type UpdateHdfsLogLevelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateHdfsLogLevelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateHdfsLogLevelNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateHdfsLogLevelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateHdfsLogLevelNoContent creates a UpdateHdfsLogLevelNoContent with default headers values
func NewUpdateHdfsLogLevelNoContent() *UpdateHdfsLogLevelNoContent {
	return &UpdateHdfsLogLevelNoContent{}
}

/*UpdateHdfsLogLevelNoContent handles this case with default header values.

Success.
*/
type UpdateHdfsLogLevelNoContent struct {
}

func (o *UpdateHdfsLogLevelNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/3/protocols/hdfs/log-level][%d] updateHdfsLogLevelNoContent ", 204)
}

func (o *UpdateHdfsLogLevelNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateHdfsLogLevelDefault creates a UpdateHdfsLogLevelDefault with default headers values
func NewUpdateHdfsLogLevelDefault(code int) *UpdateHdfsLogLevelDefault {
	return &UpdateHdfsLogLevelDefault{
		_statusCode: code,
	}
}

/*UpdateHdfsLogLevelDefault handles this case with default header values.

Unexpected error
*/
type UpdateHdfsLogLevelDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update hdfs log level default response
func (o *UpdateHdfsLogLevelDefault) Code() int {
	return o._statusCode
}

func (o *UpdateHdfsLogLevelDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/3/protocols/hdfs/log-level][%d] updateHdfsLogLevel default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateHdfsLogLevelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
