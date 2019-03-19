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

// ListHdfsRacksReader is a Reader for the ListHdfsRacks structure.
type ListHdfsRacksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListHdfsRacksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListHdfsRacksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListHdfsRacksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListHdfsRacksOK creates a ListHdfsRacksOK with default headers values
func NewListHdfsRacksOK() *ListHdfsRacksOK {
	return &ListHdfsRacksOK{}
}

/*ListHdfsRacksOK handles this case with default header values.

List all racks.
*/
type ListHdfsRacksOK struct {
	Payload *models.HdfsRacksExtended
}

func (o *ListHdfsRacksOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/protocols/hdfs/racks][%d] listHdfsRacksOK  %+v", 200, o.Payload)
}

func (o *ListHdfsRacksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HdfsRacksExtended)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListHdfsRacksDefault creates a ListHdfsRacksDefault with default headers values
func NewListHdfsRacksDefault(code int) *ListHdfsRacksDefault {
	return &ListHdfsRacksDefault{
		_statusCode: code,
	}
}

/*ListHdfsRacksDefault handles this case with default header values.

Unexpected error
*/
type ListHdfsRacksDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list hdfs racks default response
func (o *ListHdfsRacksDefault) Code() int {
	return o._statusCode
}

func (o *ListHdfsRacksDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/protocols/hdfs/racks][%d] listHdfsRacks default  %+v", o._statusCode, o.Payload)
}

func (o *ListHdfsRacksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}