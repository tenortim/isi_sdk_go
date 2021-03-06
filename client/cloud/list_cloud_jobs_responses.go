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

// ListCloudJobsReader is a Reader for the ListCloudJobs structure.
type ListCloudJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCloudJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListCloudJobsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListCloudJobsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListCloudJobsOK creates a ListCloudJobsOK with default headers values
func NewListCloudJobsOK() *ListCloudJobsOK {
	return &ListCloudJobsOK{}
}

/*ListCloudJobsOK handles this case with default header values.

List all cloudpools jobs.
*/
type ListCloudJobsOK struct {
	Payload *models.CloudJobsExtended
}

func (o *ListCloudJobsOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/cloud/jobs][%d] listCloudJobsOK  %+v", 200, o.Payload)
}

func (o *ListCloudJobsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudJobsExtended)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCloudJobsDefault creates a ListCloudJobsDefault with default headers values
func NewListCloudJobsDefault(code int) *ListCloudJobsDefault {
	return &ListCloudJobsDefault{
		_statusCode: code,
	}
}

/*ListCloudJobsDefault handles this case with default header values.

Unexpected error
*/
type ListCloudJobsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list cloud jobs default response
func (o *ListCloudJobsDefault) Code() int {
	return o._statusCode
}

func (o *ListCloudJobsDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/cloud/jobs][%d] listCloudJobs default  %+v", o._statusCode, o.Payload)
}

func (o *ListCloudJobsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
