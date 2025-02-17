// Code generated by go-swagger; DO NOT EDIT.

package invoice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/v3/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// UploadInvoiceTemplateReader is a Reader for the UploadInvoiceTemplate structure.
type UploadInvoiceTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadInvoiceTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewUploadInvoiceTemplateCreated()
		result.HttpResponse = response
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		errorResult := kbcommon.NewKillbillError(response.Code())
		if err := consumer.Consume(response.Body(), &errorResult); err != nil && err != io.EOF {
			return nil, err
		}
		return nil, errorResult
	}
}

// NewUploadInvoiceTemplateCreated creates a UploadInvoiceTemplateCreated with default headers values
func NewUploadInvoiceTemplateCreated() *UploadInvoiceTemplateCreated {
	return &UploadInvoiceTemplateCreated{}
}

/*UploadInvoiceTemplateCreated handles this case with default header values.

Uploaded invoice template Successfully
*/
type UploadInvoiceTemplateCreated struct {
	Payload string

	HttpResponse runtime.ClientResponse
}

func (o *UploadInvoiceTemplateCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices/template][%d] uploadInvoiceTemplateCreated  %+v", 201, o.Payload)
}

func (o *UploadInvoiceTemplateCreated) GetPayload() string {
	return o.Payload
}

func (o *UploadInvoiceTemplateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
