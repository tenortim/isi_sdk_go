// Code generated by go-swagger; DO NOT EDIT.

package namespace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tenortim/isi_sdk_go/models"
)

// SetWormPropertiesReader is a Reader for the SetWormProperties structure.
type SetWormPropertiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetWormPropertiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetWormPropertiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetWormPropertiesOK creates a SetWormPropertiesOK with default headers values
func NewSetWormPropertiesOK() *SetWormPropertiesOK {
	return &SetWormPropertiesOK{}
}

/*SetWormPropertiesOK handles this case with default header values.

WORM file properties.
*/
type SetWormPropertiesOK struct {
	Payload models.Empty
}

func (o *SetWormPropertiesOK) Error() string {
	return fmt.Sprintf("[PUT /namespace/{WormFilePath}][%d] setWormPropertiesOK  %+v", 200, o.Payload)
}

func (o *SetWormPropertiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
