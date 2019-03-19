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

// GetJobPolicyReader is a Reader for the GetJobPolicy structure.
type GetJobPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJobPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetJobPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetJobPolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetJobPolicyOK creates a GetJobPolicyOK with default headers values
func NewGetJobPolicyOK() *GetJobPolicyOK {
	return &GetJobPolicyOK{}
}

/*GetJobPolicyOK handles this case with default header values.

View a single job impact policy.
*/
type GetJobPolicyOK struct {
	Payload *models.JobPolicies
}

func (o *GetJobPolicyOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/job/policies/{JobPolicyId}][%d] getJobPolicyOK  %+v", 200, o.Payload)
}

func (o *GetJobPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobPolicies)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobPolicyDefault creates a GetJobPolicyDefault with default headers values
func NewGetJobPolicyDefault(code int) *GetJobPolicyDefault {
	return &GetJobPolicyDefault{
		_statusCode: code,
	}
}

/*GetJobPolicyDefault handles this case with default header values.

Unexpected error
*/
type GetJobPolicyDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get job policy default response
func (o *GetJobPolicyDefault) Code() int {
	return o._statusCode
}

func (o *GetJobPolicyDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/job/policies/{JobPolicyId}][%d] getJobPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *GetJobPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
