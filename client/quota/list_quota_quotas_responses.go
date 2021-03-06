// Code generated by go-swagger; DO NOT EDIT.

package quota

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// ListQuotaQuotasReader is a Reader for the ListQuotaQuotas structure.
type ListQuotaQuotasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListQuotaQuotasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListQuotaQuotasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListQuotaQuotasDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListQuotaQuotasOK creates a ListQuotaQuotasOK with default headers values
func NewListQuotaQuotasOK() *ListQuotaQuotasOK {
	return &ListQuotaQuotasOK{}
}

/*ListQuotaQuotasOK handles this case with default header values.

List all or matching quotas. Can also be used to retrieve quota state from existing reports. For any query argument not supplied, the default behavior is return all.
*/
type ListQuotaQuotasOK struct {
	Payload *models.QuotaQuotasExtended
}

func (o *ListQuotaQuotasOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/quota/quotas][%d] listQuotaQuotasOK  %+v", 200, o.Payload)
}

func (o *ListQuotaQuotasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QuotaQuotasExtended)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListQuotaQuotasDefault creates a ListQuotaQuotasDefault with default headers values
func NewListQuotaQuotasDefault(code int) *ListQuotaQuotasDefault {
	return &ListQuotaQuotasDefault{
		_statusCode: code,
	}
}

/*ListQuotaQuotasDefault handles this case with default header values.

Unexpected error
*/
type ListQuotaQuotasDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list quota quotas default response
func (o *ListQuotaQuotasDefault) Code() int {
	return o._statusCode
}

func (o *ListQuotaQuotasDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/quota/quotas][%d] listQuotaQuotas default  %+v", o._statusCode, o.Payload)
}

func (o *ListQuotaQuotasDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
