// Code generated by go-swagger; DO NOT EDIT.

package invoice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbcommon"
)

// GetInvoiceTemplateReader is a Reader for the GetInvoiceTemplate structure.
type GetInvoiceTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoiceTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInvoiceTemplateOK()
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

// NewGetInvoiceTemplateOK creates a GetInvoiceTemplateOK with default headers values
func NewGetInvoiceTemplateOK() *GetInvoiceTemplateOK {
	return &GetInvoiceTemplateOK{}
}

/*
GetInvoiceTemplateOK describes a response with status code 200, with default header values.

successful operation
*/
type GetInvoiceTemplateOK struct {
	Payload      string
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get invoice template o k response
func (o *GetInvoiceTemplateOK) Code() int {
	return 200
}

// IsSuccess returns true when this get invoice template o k response has a 2xx status code
func (o *GetInvoiceTemplateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get invoice template o k response has a 3xx status code
func (o *GetInvoiceTemplateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoice template o k response has a 4xx status code
func (o *GetInvoiceTemplateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get invoice template o k response has a 5xx status code
func (o *GetInvoiceTemplateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoice template o k response a status code equal to that given
func (o *GetInvoiceTemplateOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetInvoiceTemplateOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/template][%d] getInvoiceTemplateOK  %+v", 200, o.Payload)
}

func (o *GetInvoiceTemplateOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/template][%d] getInvoiceTemplateOK  %+v", 200, o.Payload)
}

func (o *GetInvoiceTemplateOK) GetPayload() string {
	return o.Payload
}

func (o *GetInvoiceTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoiceTemplateNotFound creates a GetInvoiceTemplateNotFound with default headers values
func NewGetInvoiceTemplateNotFound() *GetInvoiceTemplateNotFound {
	return &GetInvoiceTemplateNotFound{}
}

/*
GetInvoiceTemplateNotFound describes a response with status code 404, with default header values.

Template not found
*/
type GetInvoiceTemplateNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get invoice template not found response
func (o *GetInvoiceTemplateNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this get invoice template not found response has a 2xx status code
func (o *GetInvoiceTemplateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoice template not found response has a 3xx status code
func (o *GetInvoiceTemplateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoice template not found response has a 4xx status code
func (o *GetInvoiceTemplateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoice template not found response has a 5xx status code
func (o *GetInvoiceTemplateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoice template not found response a status code equal to that given
func (o *GetInvoiceTemplateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetInvoiceTemplateNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/template][%d] getInvoiceTemplateNotFound ", 404)
}

func (o *GetInvoiceTemplateNotFound) String() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/template][%d] getInvoiceTemplateNotFound ", 404)
}

func (o *GetInvoiceTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
