// Code generated by go-swagger; DO NOT EDIT.

package sync

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetHistoryCPUReader is a Reader for the GetHistoryCPU structure.
type GetHistoryCPUReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHistoryCPUReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetHistoryCPUOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetHistoryCPUDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHistoryCPUOK creates a GetHistoryCPUOK with default headers values
func NewGetHistoryCPUOK() *GetHistoryCPUOK {
	return &GetHistoryCPUOK{}
}

/*GetHistoryCPUOK handles this case with default header values.

List cpu performance data.
*/
type GetHistoryCPUOK struct {
	Payload *models.HistoryFile
}

func (o *GetHistoryCPUOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/sync/history/cpu][%d] getHistoryCpuOK  %+v", 200, o.Payload)
}

func (o *GetHistoryCPUOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HistoryFile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHistoryCPUDefault creates a GetHistoryCPUDefault with default headers values
func NewGetHistoryCPUDefault(code int) *GetHistoryCPUDefault {
	return &GetHistoryCPUDefault{
		_statusCode: code,
	}
}

/*GetHistoryCPUDefault handles this case with default header values.

Unexpected error
*/
type GetHistoryCPUDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get history Cpu default response
func (o *GetHistoryCPUDefault) Code() int {
	return o._statusCode
}

func (o *GetHistoryCPUDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/sync/history/cpu][%d] getHistoryCpu default  %+v", o._statusCode, o.Payload)
}

func (o *GetHistoryCPUDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
