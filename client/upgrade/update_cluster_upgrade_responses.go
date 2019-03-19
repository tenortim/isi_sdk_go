// Code generated by go-swagger; DO NOT EDIT.

package upgrade

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// UpdateClusterUpgradeReader is a Reader for the UpdateClusterUpgrade structure.
type UpdateClusterUpgradeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateClusterUpgradeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateClusterUpgradeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateClusterUpgradeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateClusterUpgradeNoContent creates a UpdateClusterUpgradeNoContent with default headers values
func NewUpdateClusterUpgradeNoContent() *UpdateClusterUpgradeNoContent {
	return &UpdateClusterUpgradeNoContent{}
}

/*UpdateClusterUpgradeNoContent handles this case with default header values.

Success.
*/
type UpdateClusterUpgradeNoContent struct {
}

func (o *UpdateClusterUpgradeNoContent) Error() string {
	return fmt.Sprintf("[PUT /platform/3/upgrade/cluster/upgrade][%d] updateClusterUpgradeNoContent ", 204)
}

func (o *UpdateClusterUpgradeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateClusterUpgradeDefault creates a UpdateClusterUpgradeDefault with default headers values
func NewUpdateClusterUpgradeDefault(code int) *UpdateClusterUpgradeDefault {
	return &UpdateClusterUpgradeDefault{
		_statusCode: code,
	}
}

/*UpdateClusterUpgradeDefault handles this case with default header values.

Unexpected error
*/
type UpdateClusterUpgradeDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update cluster upgrade default response
func (o *UpdateClusterUpgradeDefault) Code() int {
	return o._statusCode
}

func (o *UpdateClusterUpgradeDefault) Error() string {
	return fmt.Sprintf("[PUT /platform/3/upgrade/cluster/upgrade][%d] updateClusterUpgrade default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateClusterUpgradeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}