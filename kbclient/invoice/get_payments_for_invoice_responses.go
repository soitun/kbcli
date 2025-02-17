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

	kbmodel "github.com/killbill/kbcli/v3/kbmodel"
)

// GetPaymentsForInvoiceReader is a Reader for the GetPaymentsForInvoice structure.
type GetPaymentsForInvoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentsForInvoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentsForInvoiceOK()
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

// NewGetPaymentsForInvoiceOK creates a GetPaymentsForInvoiceOK with default headers values
func NewGetPaymentsForInvoiceOK() *GetPaymentsForInvoiceOK {
	return &GetPaymentsForInvoiceOK{}
}

/*GetPaymentsForInvoiceOK handles this case with default header values.

successful operation
*/
type GetPaymentsForInvoiceOK struct {
	Payload []*kbmodel.InvoicePayment

	HttpResponse runtime.ClientResponse
}

func (o *GetPaymentsForInvoiceOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/{invoiceId}/payments][%d] getPaymentsForInvoiceOK  %+v", 200, o.Payload)
}

func (o *GetPaymentsForInvoiceOK) GetPayload() []*kbmodel.InvoicePayment {
	return o.Payload
}

func (o *GetPaymentsForInvoiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentsForInvoiceBadRequest creates a GetPaymentsForInvoiceBadRequest with default headers values
func NewGetPaymentsForInvoiceBadRequest() *GetPaymentsForInvoiceBadRequest {
	return &GetPaymentsForInvoiceBadRequest{}
}

/*GetPaymentsForInvoiceBadRequest handles this case with default header values.

Invalid invoice id supplied
*/
type GetPaymentsForInvoiceBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetPaymentsForInvoiceBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/{invoiceId}/payments][%d] getPaymentsForInvoiceBadRequest ", 400)
}

func (o *GetPaymentsForInvoiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPaymentsForInvoiceNotFound creates a GetPaymentsForInvoiceNotFound with default headers values
func NewGetPaymentsForInvoiceNotFound() *GetPaymentsForInvoiceNotFound {
	return &GetPaymentsForInvoiceNotFound{}
}

/*GetPaymentsForInvoiceNotFound handles this case with default header values.

Invoice not found
*/
type GetPaymentsForInvoiceNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetPaymentsForInvoiceNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/{invoiceId}/payments][%d] getPaymentsForInvoiceNotFound ", 404)
}

func (o *GetPaymentsForInvoiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
