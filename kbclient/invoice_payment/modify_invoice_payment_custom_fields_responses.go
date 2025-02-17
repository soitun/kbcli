// Code generated by go-swagger; DO NOT EDIT.

package invoice_payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/v3/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// ModifyInvoicePaymentCustomFieldsReader is a Reader for the ModifyInvoicePaymentCustomFields structure.
type ModifyInvoicePaymentCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyInvoicePaymentCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewModifyInvoicePaymentCustomFieldsNoContent()
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

// NewModifyInvoicePaymentCustomFieldsNoContent creates a ModifyInvoicePaymentCustomFieldsNoContent with default headers values
func NewModifyInvoicePaymentCustomFieldsNoContent() *ModifyInvoicePaymentCustomFieldsNoContent {
	return &ModifyInvoicePaymentCustomFieldsNoContent{}
}

/*ModifyInvoicePaymentCustomFieldsNoContent handles this case with default header values.

Successful operation
*/
type ModifyInvoicePaymentCustomFieldsNoContent struct {
	HttpResponse runtime.ClientResponse
}

func (o *ModifyInvoicePaymentCustomFieldsNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/invoicePayments/{paymentId}/customFields][%d] modifyInvoicePaymentCustomFieldsNoContent ", 204)
}

func (o *ModifyInvoicePaymentCustomFieldsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModifyInvoicePaymentCustomFieldsBadRequest creates a ModifyInvoicePaymentCustomFieldsBadRequest with default headers values
func NewModifyInvoicePaymentCustomFieldsBadRequest() *ModifyInvoicePaymentCustomFieldsBadRequest {
	return &ModifyInvoicePaymentCustomFieldsBadRequest{}
}

/*ModifyInvoicePaymentCustomFieldsBadRequest handles this case with default header values.

Invalid payment id supplied
*/
type ModifyInvoicePaymentCustomFieldsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *ModifyInvoicePaymentCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/invoicePayments/{paymentId}/customFields][%d] modifyInvoicePaymentCustomFieldsBadRequest ", 400)
}

func (o *ModifyInvoicePaymentCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
