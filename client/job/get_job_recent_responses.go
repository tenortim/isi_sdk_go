// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetJobRecentReader is a Reader for the GetJobRecent structure.
type GetJobRecentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJobRecentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetJobRecentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetJobRecentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetJobRecentOK creates a GetJobRecentOK with default headers values
func NewGetJobRecentOK() *GetJobRecentOK {
	return &GetJobRecentOK{}
}

/*GetJobRecentOK handles this case with default header values.

List recently completed jobs.
*/
type GetJobRecentOK struct {
	Payload *models.JobRecent
}

func (o *GetJobRecentOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/job/recent][%d] getJobRecentOK  %+v", 200, o.Payload)
}

func (o *GetJobRecentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobRecent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobRecentDefault creates a GetJobRecentDefault with default headers values
func NewGetJobRecentDefault(code int) *GetJobRecentDefault {
	return &GetJobRecentDefault{
		_statusCode: code,
	}
}

/*GetJobRecentDefault handles this case with default header values.

Unexpected error
*/
type GetJobRecentDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get job recent default response
func (o *GetJobRecentDefault) Code() int {
	return o._statusCode
}

func (o *GetJobRecentDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/job/recent][%d] getJobRecent default  %+v", o._statusCode, o.Payload)
}

func (o *GetJobRecentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}