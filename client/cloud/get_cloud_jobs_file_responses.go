// Code generated by go-swagger; DO NOT EDIT.

package cloud

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetCloudJobsFileReader is a Reader for the GetCloudJobsFile structure.
type GetCloudJobsFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudJobsFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCloudJobsFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetCloudJobsFileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCloudJobsFileOK creates a GetCloudJobsFileOK with default headers values
func NewGetCloudJobsFileOK() *GetCloudJobsFileOK {
	return &GetCloudJobsFileOK{}
}

/*GetCloudJobsFileOK handles this case with default header values.

Retrieve files associated with a cloudpool job.
*/
type GetCloudJobsFileOK struct {
	Payload *models.CloudJobsFiles
}

func (o *GetCloudJobsFileOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/cloud/jobs-files/{CloudJobsFileId}][%d] getCloudJobsFileOK  %+v", 200, o.Payload)
}

func (o *GetCloudJobsFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudJobsFiles)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudJobsFileDefault creates a GetCloudJobsFileDefault with default headers values
func NewGetCloudJobsFileDefault(code int) *GetCloudJobsFileDefault {
	return &GetCloudJobsFileDefault{
		_statusCode: code,
	}
}

/*GetCloudJobsFileDefault handles this case with default header values.

Unexpected error
*/
type GetCloudJobsFileDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get cloud jobs file default response
func (o *GetCloudJobsFileDefault) Code() int {
	return o._statusCode
}

func (o *GetCloudJobsFileDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/cloud/jobs-files/{CloudJobsFileId}][%d] getCloudJobsFile default  %+v", o._statusCode, o.Payload)
}

func (o *GetCloudJobsFileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
