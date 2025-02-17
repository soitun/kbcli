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

// GetInvoiceMPTemplateReader is a Reader for the GetInvoiceMPTemplate structure.
type GetInvoiceMPTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoiceMPTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInvoiceMPTemplateOK()
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

// NewGetInvoiceMPTemplateOK creates a GetInvoiceMPTemplateOK with default headers values
func NewGetInvoiceMPTemplateOK() *GetInvoiceMPTemplateOK {
	return &GetInvoiceMPTemplateOK{}
}

/*GetInvoiceMPTemplateOK handles this case with default header values.

successful operation
*/
type GetInvoiceMPTemplateOK struct {
	Payload string

	HttpResponse runtime.ClientResponse
}

func (o *GetInvoiceMPTemplateOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/manualPayTemplate/{locale}][%d] getInvoiceMPTemplateOK  %+v", 200, o.Payload)
}

func (o *GetInvoiceMPTemplateOK) GetPayload() string {
	return o.Payload
}

func (o *GetInvoiceMPTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoiceMPTemplateNotFound creates a GetInvoiceMPTemplateNotFound with default headers values
func NewGetInvoiceMPTemplateNotFound() *GetInvoiceMPTemplateNotFound {
	return &GetInvoiceMPTemplateNotFound{}
}

/*GetInvoiceMPTemplateNotFound handles this case with default header values.

Template not found
*/
type GetInvoiceMPTemplateNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetInvoiceMPTemplateNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/manualPayTemplate/{locale}][%d] getInvoiceMPTemplateNotFound ", 404)
}

func (o *GetInvoiceMPTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
