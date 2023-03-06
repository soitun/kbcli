// Code generated by go-swagger; DO NOT EDIT.

package invoice_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbcommon"

	"github.com/killbill/kbcli/v3/kbmodel"
)

// GetInvoiceItemCustomFieldsReader is a Reader for the GetInvoiceItemCustomFields structure.
type GetInvoiceItemCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoiceItemCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInvoiceItemCustomFieldsOK()
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

// NewGetInvoiceItemCustomFieldsOK creates a GetInvoiceItemCustomFieldsOK with default headers values
func NewGetInvoiceItemCustomFieldsOK() *GetInvoiceItemCustomFieldsOK {
	return &GetInvoiceItemCustomFieldsOK{}
}

/*
GetInvoiceItemCustomFieldsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetInvoiceItemCustomFieldsOK struct {
	Payload      []*kbmodel.CustomField
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get invoice item custom fields o k response
func (o *GetInvoiceItemCustomFieldsOK) Code() int {
	return 200
}

// IsSuccess returns true when this get invoice item custom fields o k response has a 2xx status code
func (o *GetInvoiceItemCustomFieldsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get invoice item custom fields o k response has a 3xx status code
func (o *GetInvoiceItemCustomFieldsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoice item custom fields o k response has a 4xx status code
func (o *GetInvoiceItemCustomFieldsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get invoice item custom fields o k response has a 5xx status code
func (o *GetInvoiceItemCustomFieldsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoice item custom fields o k response a status code equal to that given
func (o *GetInvoiceItemCustomFieldsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetInvoiceItemCustomFieldsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoiceItems/{invoiceItemId}/customFields][%d] getInvoiceItemCustomFieldsOK  %+v", 200, o.Payload)
}

func (o *GetInvoiceItemCustomFieldsOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/invoiceItems/{invoiceItemId}/customFields][%d] getInvoiceItemCustomFieldsOK  %+v", 200, o.Payload)
}

func (o *GetInvoiceItemCustomFieldsOK) GetPayload() []*kbmodel.CustomField {
	return o.Payload
}

func (o *GetInvoiceItemCustomFieldsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoiceItemCustomFieldsBadRequest creates a GetInvoiceItemCustomFieldsBadRequest with default headers values
func NewGetInvoiceItemCustomFieldsBadRequest() *GetInvoiceItemCustomFieldsBadRequest {
	return &GetInvoiceItemCustomFieldsBadRequest{}
}

/*
GetInvoiceItemCustomFieldsBadRequest describes a response with status code 400, with default header values.

Invalid invoice item id supplied
*/
type GetInvoiceItemCustomFieldsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get invoice item custom fields bad request response
func (o *GetInvoiceItemCustomFieldsBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this get invoice item custom fields bad request response has a 2xx status code
func (o *GetInvoiceItemCustomFieldsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoice item custom fields bad request response has a 3xx status code
func (o *GetInvoiceItemCustomFieldsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoice item custom fields bad request response has a 4xx status code
func (o *GetInvoiceItemCustomFieldsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoice item custom fields bad request response has a 5xx status code
func (o *GetInvoiceItemCustomFieldsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoice item custom fields bad request response a status code equal to that given
func (o *GetInvoiceItemCustomFieldsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetInvoiceItemCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoiceItems/{invoiceItemId}/customFields][%d] getInvoiceItemCustomFieldsBadRequest ", 400)
}

func (o *GetInvoiceItemCustomFieldsBadRequest) String() string {
	return fmt.Sprintf("[GET /1.0/kb/invoiceItems/{invoiceItemId}/customFields][%d] getInvoiceItemCustomFieldsBadRequest ", 400)
}

func (o *GetInvoiceItemCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
