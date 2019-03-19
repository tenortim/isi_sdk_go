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

// GetWormPropertiesReader is a Reader for the GetWormProperties structure.
type GetWormPropertiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWormPropertiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWormPropertiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWormPropertiesOK creates a GetWormPropertiesOK with default headers values
func NewGetWormPropertiesOK() *GetWormPropertiesOK {
	return &GetWormPropertiesOK{}
}

/*GetWormPropertiesOK handles this case with default header values.

WORM file properties.
*/
type GetWormPropertiesOK struct {
	Payload *models.WormProperties
}

func (o *GetWormPropertiesOK) Error() string {
	return fmt.Sprintf("[GET /namespace/{WormFilePath}][%d] getWormPropertiesOK  %+v", 200, o.Payload)
}

func (o *GetWormPropertiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WormProperties)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
