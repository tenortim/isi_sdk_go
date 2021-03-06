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

// ListSyncRulesReader is a Reader for the ListSyncRules structure.
type ListSyncRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSyncRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListSyncRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListSyncRulesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSyncRulesOK creates a ListSyncRulesOK with default headers values
func NewListSyncRulesOK() *ListSyncRulesOK {
	return &ListSyncRulesOK{}
}

/*ListSyncRulesOK handles this case with default header values.

List all SyncIQ performance rules.
*/
type ListSyncRulesOK struct {
	Payload *models.SyncRulesExtended
}

func (o *ListSyncRulesOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/sync/rules][%d] listSyncRulesOK  %+v", 200, o.Payload)
}

func (o *ListSyncRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SyncRulesExtended)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSyncRulesDefault creates a ListSyncRulesDefault with default headers values
func NewListSyncRulesDefault(code int) *ListSyncRulesDefault {
	return &ListSyncRulesDefault{
		_statusCode: code,
	}
}

/*ListSyncRulesDefault handles this case with default header values.

Unexpected error
*/
type ListSyncRulesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list sync rules default response
func (o *ListSyncRulesDefault) Code() int {
	return o._statusCode
}

func (o *ListSyncRulesDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/sync/rules][%d] listSyncRules default  %+v", o._statusCode, o.Payload)
}

func (o *ListSyncRulesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
