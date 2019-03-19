// Code generated by go-swagger; DO NOT EDIT.

package cluster_nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// ListDrivesDriveFirmwareUpdateReader is a Reader for the ListDrivesDriveFirmwareUpdate structure.
type ListDrivesDriveFirmwareUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDrivesDriveFirmwareUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListDrivesDriveFirmwareUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListDrivesDriveFirmwareUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListDrivesDriveFirmwareUpdateOK creates a ListDrivesDriveFirmwareUpdateOK with default headers values
func NewListDrivesDriveFirmwareUpdateOK() *ListDrivesDriveFirmwareUpdateOK {
	return &ListDrivesDriveFirmwareUpdateOK{}
}

/*ListDrivesDriveFirmwareUpdateOK handles this case with default header values.

Retrieve firmware update information.
*/
type ListDrivesDriveFirmwareUpdateOK struct {
	Payload *models.DrivesDriveFirmwareUpdate
}

func (o *ListDrivesDriveFirmwareUpdateOK) Error() string {
	return fmt.Sprintf("[GET /platform/3/cluster/nodes/{Lnn}/drives/{Driveid}/firmware/update][%d] listDrivesDriveFirmwareUpdateOK  %+v", 200, o.Payload)
}

func (o *ListDrivesDriveFirmwareUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DrivesDriveFirmwareUpdate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDrivesDriveFirmwareUpdateDefault creates a ListDrivesDriveFirmwareUpdateDefault with default headers values
func NewListDrivesDriveFirmwareUpdateDefault(code int) *ListDrivesDriveFirmwareUpdateDefault {
	return &ListDrivesDriveFirmwareUpdateDefault{
		_statusCode: code,
	}
}

/*ListDrivesDriveFirmwareUpdateDefault handles this case with default header values.

Unexpected error
*/
type ListDrivesDriveFirmwareUpdateDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list drives drive firmware update default response
func (o *ListDrivesDriveFirmwareUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ListDrivesDriveFirmwareUpdateDefault) Error() string {
	return fmt.Sprintf("[GET /platform/3/cluster/nodes/{Lnn}/drives/{Driveid}/firmware/update][%d] listDrivesDriveFirmwareUpdate default  %+v", o._statusCode, o.Payload)
}

func (o *ListDrivesDriveFirmwareUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}