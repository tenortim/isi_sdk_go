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

// DeleteClusterPatchPatchReader is a Reader for the DeleteClusterPatchPatch structure.
type DeleteClusterPatchPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClusterPatchPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteClusterPatchPatchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteClusterPatchPatchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteClusterPatchPatchNoContent creates a DeleteClusterPatchPatchNoContent with default headers values
func NewDeleteClusterPatchPatchNoContent() *DeleteClusterPatchPatchNoContent {
	return &DeleteClusterPatchPatchNoContent{}
}

/*DeleteClusterPatchPatchNoContent handles this case with default header values.

Success.
*/
type DeleteClusterPatchPatchNoContent struct {
}

func (o *DeleteClusterPatchPatchNoContent) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/upgrade/cluster/patch/patches/{ClusterPatchPatchId}][%d] deleteClusterPatchPatchNoContent ", 204)
}

func (o *DeleteClusterPatchPatchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClusterPatchPatchDefault creates a DeleteClusterPatchPatchDefault with default headers values
func NewDeleteClusterPatchPatchDefault(code int) *DeleteClusterPatchPatchDefault {
	return &DeleteClusterPatchPatchDefault{
		_statusCode: code,
	}
}

/*DeleteClusterPatchPatchDefault handles this case with default header values.

Unexpected error
*/
type DeleteClusterPatchPatchDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete cluster patch patch default response
func (o *DeleteClusterPatchPatchDefault) Code() int {
	return o._statusCode
}

func (o *DeleteClusterPatchPatchDefault) Error() string {
	return fmt.Sprintf("[DELETE /platform/3/upgrade/cluster/patch/patches/{ClusterPatchPatchId}][%d] deleteClusterPatchPatch default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteClusterPatchPatchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}