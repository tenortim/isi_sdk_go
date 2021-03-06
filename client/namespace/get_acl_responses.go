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

// GetACLReader is a Reader for the GetACL structure.
type GetACLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetACLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetACLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetACLOK creates a GetACLOK with default headers values
func NewGetACLOK() *GetACLOK {
	return &GetACLOK{}
}

/*GetACLOK handles this case with default header values.

Get namespace ACL.
*/
type GetACLOK struct {
	Payload *models.NamespaceACL
}

func (o *GetACLOK) Error() string {
	return fmt.Sprintf("[GET /namespace/{NamespacePath}][%d] getAclOK  %+v", 200, o.Payload)
}

func (o *GetACLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NamespaceACL)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
