// Code generated by go-swagger; DO NOT EDIT.

package filepool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// GetFilepoolTemplateReader is a Reader for the GetFilepoolTemplate structure.
type GetFilepoolTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFilepoolTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFilepoolTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetFilepoolTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFilepoolTemplateOK creates a GetFilepoolTemplateOK with default headers values
func NewGetFilepoolTemplateOK() *GetFilepoolTemplateOK {
	return &GetFilepoolTemplateOK{}
}

/*GetFilepoolTemplateOK handles this case with default header values.

List all templates.
*/
type GetFilepoolTemplateOK struct {
	Payload *models.FilepoolTemplates
}

func (o *GetFilepoolTemplateOK) Error() string {
	return fmt.Sprintf("[GET /platform/1/filepool/templates/{FilepoolTemplateId}][%d] getFilepoolTemplateOK  %+v", 200, o.Payload)
}

func (o *GetFilepoolTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FilepoolTemplates)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilepoolTemplateDefault creates a GetFilepoolTemplateDefault with default headers values
func NewGetFilepoolTemplateDefault(code int) *GetFilepoolTemplateDefault {
	return &GetFilepoolTemplateDefault{
		_statusCode: code,
	}
}

/*GetFilepoolTemplateDefault handles this case with default header values.

Unexpected error
*/
type GetFilepoolTemplateDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get filepool template default response
func (o *GetFilepoolTemplateDefault) Code() int {
	return o._statusCode
}

func (o *GetFilepoolTemplateDefault) Error() string {
	return fmt.Sprintf("[GET /platform/1/filepool/templates/{FilepoolTemplateId}][%d] getFilepoolTemplate default  %+v", o._statusCode, o.Payload)
}

func (o *GetFilepoolTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}